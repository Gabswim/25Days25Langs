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
	South
	West
	East
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
			//fmt.Println(tile)
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

func (m *Map) String() string {
	str := ""
	for _, row := range m.Bord {
		for _, tile := range row {
			switch tile.Type {
			case Nothing:
				str += "."
			case Obstacle:
				str += "#"
			case Guard:
				str += "^"
			case Visited:
				str += "X"
			}
		}
		str += "\n"
	}
	return str
}

func (m *Map) MoveToNextPosition() bool {
	canMoveToNextPosition := false
	distance := 0
	i, j, d, err := m.getGuardPosition()
	check(err)

	m.Bord[i][j] = Tile{Type: Visited} // reinitialize tile
	nextI := i
	nextJ := j
	if d == East {
		for j+1 < len(m.Bord[i]) && m.Bord[i][j+1].Type != Obstacle {
			j++
			nextJ = j
			m.Bord[i][j] = Tile{Type: Visited}
			canMoveToNextPosition = true
			distance++
		}
	}
	if d == South {
		for i+1 < len(m.Bord) && m.Bord[i+1][j].Type != Obstacle {
			i++
			nextI = i
			canMoveToNextPosition = true
			m.Bord[i][j] = Tile{Type: Visited}
			distance++
		}
	}
	if d == West {
		for j-1 >= 0 && m.Bord[i][j-1].Type != Obstacle {
			j--
			nextJ = j
			canMoveToNextPosition = true
			m.Bord[i][j] = Tile{Type: Visited}
			distance++
		}
	}
	if d == North {
		for i-1 >= 0 && m.Bord[i-1][j].Type != Obstacle {
			i--
			nextI = i
			canMoveToNextPosition = true
			m.Bord[i][j] = Tile{Type: Visited}
			distance++
		}
	}

	fmt.Printf("Guard moved %d steps to the %s\n", distance, [...]string{"North", "South", "West", "East"}[d])

	getNextDirection := func(d Direction) Direction {
		switch d {
		case North:
			return East // iterate to the right: i, j+1
		case East:
			return South // iterate down: i+1, j
		case South:
			return West // iterate to the left: i, j-1
		case West:
			return North // iterate up: i-1, j
		}
		panic("invalid direction") // TODO: handle this error
	}
	nextDirection := getNextDirection(d)
	m.Bord[nextI][nextJ] = Tile{Type: Guard, Direction: nextDirection}

	return canMoveToNextPosition
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
	//fmt.Println(m)
	hasLeaveTheMap := false

	for !hasLeaveTheMap {
		canMoveToNextPosition := m.MoveToNextPosition()
		if !canMoveToNextPosition {
			hasLeaveTheMap = true
		}

	}
	fmt.Println(m)
	return m.countVisited() + 1 // add the end position
}

func main() {

	test(sol1("./test-input.txt"), 41)
	//test(sol1("./input.txt"), 41)
}
