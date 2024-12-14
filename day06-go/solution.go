package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type TileType int

const (
	Nothing TileType = iota
	Obstacle
	Guard
	Visited
)

type Tile struct {
	Type      TileType
	Direction Direction
}

type Map struct {
	Board        [][]Tile
	visitedTiles map[string]int
}

func createTile(t string) (Tile, error) {
	switch t {
	case ".":
		return Tile{Type: Nothing}, nil
	case "#":
		return Tile{Type: Obstacle}, nil
	case "^":
		return Tile{Type: Guard, Direction: North}, nil
	default:
		return Tile{}, fmt.Errorf("invalid tile type: %s", t)
	}
}

func (m *Map) Init(file string) error {
	data, err := os.ReadFile(file)
	check(err)
	for _, line := range strings.Split(string(data), "\n") {
		lineTiles := []Tile{}
		for _, tile := range strings.Split(line, "") {
			t, err := createTile(tile)
			check(err)
			lineTiles = append(lineTiles, t)
		}
		m.Board = append(m.Board, lineTiles)
		m.visitedTiles = make(map[string]int)
	}
	return err
}

func (m *Map) getGuardPosition() (int, int, Direction, error) {
	for i, row := range m.Board {
		for j, tile := range row {
			if tile.Type == Guard {
				return i, j, tile.Direction, nil
			}
		}
	}
	return -1, -1, North, fmt.Errorf("guard not found")
}

func (m *Map) countVisited() int {
	count := 0
	for _, row := range m.Board {
		for _, tile := range row {
			if tile.Type == Visited {
				count++
			}
		}
	}
	return count
}

func getNextDirection(d Direction) Direction {
	switch d {
	case North:
		return East
	case East:
		return South
	case South:
		return West
	case West:
		return North
	}
	panic("invalid direction")
}

func (m *Map) MoveToNextPosition() (bool, int, int) {
	i, j, d, err := m.getGuardPosition()
	check(err)

	m.Board[i][j].Type = Visited

	nextI := i
	nextJ := j

	switch d {
	case North:
		nextI--
	case East:
		nextJ++
	case South:
		nextI++
	case West:
		nextJ--
	}

	if nextI < 0 || nextI >= len(m.Board) || nextJ < 0 || nextJ >= len(m.Board[0]) {
		return false, -1, -1
	}

	if m.Board[nextI][nextJ].Type == Obstacle {
		nextDirection := getNextDirection(d)
		m.Board[i][j] = Tile{Type: Guard, Direction: nextDirection}
		return true, nextI, nextJ
	}

	m.Board[nextI][nextJ] = Tile{Type: Guard, Direction: d}
	return true, nextI, nextJ
}

func (m *Map) testIfObstruction(i, j int) bool {
	m.Board[i][j].Type = Obstacle

	stateHistory := make(map[string]int)

	guardI, guardJ, guardDir, err := m.getGuardPosition()
	if err != nil {
		return false
	}

	maxMoves := len(m.Board) * len(m.Board[0]) * 4
	moveCount := 0

	for moveCount < maxMoves {
		state := fmt.Sprintf("%d,%d,%d", guardI, guardJ, guardDir)

		if prevMove, exists := stateHistory[state]; exists {
			if moveCount-prevMove > 1 {
				return true
			}
		}
		stateHistory[state] = moveCount

		canMove, nextI, nextJ := m.MoveToNextPosition()
		if !canMove {
			return false
		}

		guardI, guardJ = nextI, nextJ
		_, _, guardDir, err = m.getGuardPosition()
		check(err)
		moveCount++
	}

	return false
}

func (m *Map) DeepCopy() *Map {
	newMap := &Map{
		Board:        make([][]Tile, len(m.Board)),
		visitedTiles: make(map[string]int),
	}

	for i := range m.Board {
		newMap.Board[i] = make([]Tile, len(m.Board[i]))
		copy(newMap.Board[i], m.Board[i])
	}

	for k, v := range m.visitedTiles {
		newMap.visitedTiles[k] = v
	}

	return newMap
}

func processChunk(originalMap *Map, start, end int, width int, guardI, guardJ int, results chan<- int, wg *sync.WaitGroup, progress *int64, totalTests int, startTime time.Time) {
	defer wg.Done()
	count := 0

	for idx := start; idx < end; idx++ {
		i := idx / width
		j := idx % width

		if (i == guardI && j == guardJ) || originalMap.Board[i][j].Type != Nothing {
			continue
		}

		m := originalMap.DeepCopy()
		if m.testIfObstruction(i, j) {
			count++
		}

		current := atomic.AddInt64(progress, 1)
		elapsed := time.Since(startTime)
		if current > 0 {
			testsPerSecond := float64(current) / elapsed.Seconds()
			remainingTests := totalTests - int(current)
			estimatedRemainingSeconds := float64(remainingTests) / testsPerSecond
			remainingDuration := time.Duration(estimatedRemainingSeconds * float64(time.Second))

			fmt.Printf("Test # %d / %d (%.1f%%) - Estimated remaining time: %s\n",
				current,
				totalTests,
				(float64(current)/float64(totalTests))*100,
				remainingDuration.Round(time.Second),
			)
		}
	}

	results <- count
}

func countNumberOfObstructions(file string) int {
	originalMap := NewMap(file)
	height := len(originalMap.Board)
	width := len(originalMap.Board[0])
	totalTiles := height * width

	guardI, guardJ, _, _ := originalMap.getGuardPosition()

	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	chunkSize := (totalTiles + numCPU - 1) / numCPU

	results := make(chan int, numCPU)
	var wg sync.WaitGroup

	var progress int64 = 0
	startTime := time.Now()

	for i := 0; i < numCPU; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > totalTiles {
			end = totalTiles
		}

		wg.Add(1)
		go processChunk(originalMap, start, end, width, guardI, guardJ, results, &wg, &progress, totalTiles, startTime)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	numberOfObstructions := 0
	for result := range results {
		numberOfObstructions += result
	}

	return numberOfObstructions
}

func NewMap(file string) *Map {
	m := &Map{}
	err := m.Init(file)
	check(err)
	return m
}

func test(actual int, expected int) {
	if actual == expected {
		fmt.Println("✅ Test pass")
	} else {
		fmt.Printf("❌ Test fail, Expected: %d, Actual: %d\n", expected, actual)
	}
}

func sol1(file string) int {
	m := NewMap(file)
	hasLeaveTheMap := false

	for !hasLeaveTheMap {
		canMoveToNextPosition, _, _ := m.MoveToNextPosition()
		if !canMoveToNextPosition {
			hasLeaveTheMap = true
		}
	}
	return m.countVisited()
}

func sol2(file string) int {
	numberOfPossibleObstructions := countNumberOfObstructions(file)
	return numberOfPossibleObstructions
}

func main() {
	test(sol1("./test-input.txt"), 41)
	test(sol1("./input.txt"), 4903)

	test(sol2("./test-input.txt"), 6)
	test(sol2("./input.txt"), 10)
}
