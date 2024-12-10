package main

import (
	"fmt"
	"os"
	"strings"
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
	Bord [][]Tile
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
		m.Bord = append(m.Bord, lineTiles)
	}

	return err
}

func (m *Map) getGuardPosition() (int, int, Direction, error) {
	for i, row := range m.Bord {
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
	for _, row := range m.Bord {
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

func (m *Map) MoveToNextPosition() bool {
	i, j, d, err := m.getGuardPosition()
	check(err)

	m.Bord[i][j].Type = Visited

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

	if nextI < 0 || nextI >= len(m.Bord) || nextJ < 0 || nextJ >= len(m.Bord[0]) {
		return false
	}

	if m.Bord[nextI][nextJ].Type == Obstacle {
		nextDirection := getNextDirection(d)
		m.Bord[i][j] = Tile{Type: Guard, Direction: nextDirection}
		return true
	}

	m.Bord[nextI][nextJ] = Tile{Type: Guard, Direction: d}
	return true
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
		canMoveToNextPosition := m.MoveToNextPosition()
		if !canMoveToNextPosition {
			hasLeaveTheMap = true
		}

	}
	return m.countVisited()
}

func main() {
	test(sol1("./test-input.txt"), 41)
	test(sol1("./input.txt"), 4903)
}
