package main

import (
	"image/color"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Wall struct {
	Exists bool // Flag indicating whether the wall exists or not
	Cell1  *Cell
	Cell2  *Cell
	Weight int
}

type Cell struct {
	Row      int
	Column   int
	Visited  bool
	Adjacent []*Cell
}

type Maze struct {
	Width  int
	Height int
	Cells  [][]*Cell
}

func InitMaze(width, height int) *Maze {
	// Initialize the maze grid
	cells := make([][]*Cell, height)
	for i := range cells {
		cells[i] = make([]*Cell, width)
		for j := range cells[i] {
			cells[i][j] = &Cell{
				Row:      i,
				Column:   j,
				Visited:  false,
				Adjacent: make([]*Cell, 0), // Initialize the slice for adjacent cells
			}
		}
	}

	// Create shared walls between adjacent cells
	for i := range cells {
		for j := range cells[i] {
			cell := cells[i][j]

			// Add adjacent cells to the current cell
			if i > 0 {
				cell.Adjacent = append(cell.Adjacent, cells[i-1][j])
			}
			if i < height-1 {
				cell.Adjacent = append(cell.Adjacent, cells[i+1][j])
			}
			if j > 0 {
				cell.Adjacent = append(cell.Adjacent, cells[i][j-1])
			}
			if j < width-1 {
				cell.Adjacent = append(cell.Adjacent, cells[i][j+1])
			}

		}

	}
	return &Maze{
		Width:  width,
		Height: height,
		Cells:  cells,
	}
}

func (m *Maze) DrawMaze(screen *ebiten.Image) {
	for i := range m.Cells {
		for j := range m.Cells[i] {
			if j+1 <= len(m.Cells[i])-1 && slices.Contains(m.Cells[i][j].Adjacent, m.Cells[i][j+1]) {
				j := float32(j)
				i := float32(i)
				vector.StrokeLine(screen, (j+1)*10, i*10, (j+1)*10, (i+1)*10, 1, color.RGBA{0x80, 0x80, 0x80, 0xff}, false)
			}
			if i+1 <= len(m.Cells)-1 && slices.Contains(m.Cells[i][j].Adjacent, m.Cells[i+1][j]) {
				j := float32(j)
				i := float32(i)
				vector.StrokeLine(screen, j*10, (i+1)*10, (j+1)*10, (i+1)*10, 1, color.RGBA{0x80, 0x80, 0x80, 0xff}, false)
			}
			if j-1 >= 0 && slices.Contains(m.Cells[i][j].Adjacent, m.Cells[i][j-1]) {
				j := float32(j)
				i := float32(i)
				vector.StrokeLine(screen, j*10, i*10, j*10, (i+1)*10, 1, color.RGBA{0x80, 0x80, 0x80, 0xff}, false)
			}
			if i-1 >= 0 && slices.Contains(m.Cells[i][j].Adjacent, m.Cells[i-1][j]) {
				j := float32(j)
				i := float32(i)
				vector.StrokeLine(screen, j*10, i*10, (j+1)*10, i*10, 1, color.RGBA{0x80, 0x80, 0x80, 0xff}, false)
			}
		}
	}
}

type Graph struct {
	vertices int
	edges    []Edge
}

type Edge struct {
	src, dest, weight int
}

func Prim(maze *Maze) {

}

func DFS(cell *Cell, goal *Cell) bool {
	if cell == goal {
		return true
	}
	cell.Visited = true
	for _, neighbor := range cell.Adjacent {
		if !neighbor.Visited {
			if DFS(neighbor, goal) {
				return true
			}
		}
	}
	return false
}
