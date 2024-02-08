package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Wall struct {
	Exists bool // Flag indicating whether the wall exists or not
	Cell1  *Cell
	Cell2  *Cell
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
	Walls  map[[2]*Cell]*Wall
}

type Coords struct {
	X int
	Y int
}

func InitMaze(width, height int) *Maze {
	maze := &Maze{
		Width:  width,
		Height: height,
		Cells:  make([][]*Cell, height),
		Walls:  make(map[[2]*Cell]*Wall),
	}
	for i := range maze.Cells {
		maze.Cells[i] = make([]*Cell, width)
		for j := range maze.Cells[i] {
			maze.Cells[i][j] = &Cell{
				Row:      i,
				Column:   j,
				Visited:  false,
				Adjacent: make([]*Cell, 0),
			}
		}
	}
	// Populate adjacent cells and create corresponding walls
	for i := range maze.Cells {
		for j := range maze.Cells[i] {
			cell := maze.Cells[i][j]

			// Populate upper adjacent cell and create wall
			if i > 0 {
				upperCell := maze.Cells[i-1][j]
				upperCell.Adjacent = append(upperCell.Adjacent, cell)
				cell.Adjacent = append(cell.Adjacent, upperCell)
				wall := &Wall{Exists: true, Cell1: upperCell, Cell2: cell}
				maze.Walls[[2]*Cell{upperCell, cell}] = wall
				maze.Walls[[2]*Cell{cell, upperCell}] = wall
			}

			// Populate left adjacent cell and create wall
			if j > 0 {
				leftCell := maze.Cells[i][j-1]
				leftCell.Adjacent = append(leftCell.Adjacent, cell)
				cell.Adjacent = append(cell.Adjacent, leftCell)
				wall := &Wall{Exists: true, Cell1: leftCell, Cell2: cell}
				maze.Walls[[2]*Cell{leftCell, cell}] = wall
				maze.Walls[[2]*Cell{cell, leftCell}] = wall
			}

			// Populate lower adjacent cell and create wall
			if i < maze.Height-1 {
				lowerCell := maze.Cells[i+1][j]
				lowerCell.Adjacent = append(lowerCell.Adjacent, cell)
				cell.Adjacent = append(cell.Adjacent, lowerCell)
				wall := &Wall{Exists: true, Cell1: lowerCell, Cell2: cell}
				maze.Walls[[2]*Cell{lowerCell, cell}] = wall
				maze.Walls[[2]*Cell{cell, lowerCell}] = wall
			}

			// Populate right adjacent cell and create wall
			if j < maze.Width-1 {
				rightCell := maze.Cells[i][j+1]
				rightCell.Adjacent = append(rightCell.Adjacent, cell)
				cell.Adjacent = append(cell.Adjacent, rightCell)
				wall := &Wall{Exists: true, Cell1: rightCell, Cell2: cell}
				maze.Walls[[2]*Cell{rightCell, cell}] = wall
				maze.Walls[[2]*Cell{cell, rightCell}] = wall
			}
		}
	}

	return maze
}
func (maze *Maze) Prim() {
	// Choose the starting cell [0,0] and mark it as visited
	startRow, startCol := 13, 13
	currentCell := maze.Cells[startRow][startCol]
	currentCell.Visited = true
	unvisitedWalls := maze.getUnvisitedWalls(currentCell)
	fmt.Println("Unvisited walls: ", len(unvisitedWalls))
	for len(unvisitedWalls) > 0 {
		time.Sleep(time.Second)
		// Choose a random unvisited wall
		shuffleWalls(unvisitedWalls)
		randomWall := unvisitedWalls[0]
		// Get the neighboring cell of the wall
		neighbor := unvisitedWalls[0].getNeighbor(currentCell)
		neighbor.Visited = true

		// Mark the wall as non-existent
		randomWall.Exists = false
		// Because in the map we have the same wall for both cells, we need to mark the wall as non-existent for the neighbor cell as well
		maze.markWallAsNonExistent(currentCell, neighbor)

		// Add the neighboring cell's unvisited walls to the list
		unvisitedWalls = append(unvisitedWalls, maze.getUnvisitedWalls(neighbor)...)

		// Remove the current wall from the list
		unvisitedWalls = unvisitedWalls[1:]
	}
}

// getUnvisitedWalls returns a list of unvisited walls from the current cell
func (maze *Maze) getUnvisitedWalls(cell *Cell) []*Wall {
	unvisitedWalls := make([]*Wall, 0)
	for _, neighbor := range cell.Adjacent {
		// Check if the neighbor cell is unvisited
		if !neighbor.Visited {
			// Check if the wall between the current cell and its neighbor exists and is unvisited
			wall := maze.Walls[[2]*Cell{cell, neighbor}]
			// get coords of the wall

			if wall != nil && wall.Exists && maze.Walls[[2]*Cell{neighbor, cell}].Exists {
				unvisitedWalls = append(unvisitedWalls, wall)
			}
		}
	}
	return unvisitedWalls
}

// Get the walls associated with a cell
func (maze *Maze) getWalls(cell *Cell) []*Wall {
	walls := make([]*Wall, 0)
if cell.Row > 0 {
		walls = append(walls, maze.Walls[[2]*Cell{cell, maze.Cells[cell.Row-1][cell.Column]}])
	}
	if cell.Column > 0 {
		walls = append(walls, maze.Walls[[2]*Cell{cell, maze.Cells[cell.Row][cell.Column-1]}])
	}
	if cell.Row < maze.Height-1 {
		walls = append(walls, maze.Walls[[2]*Cell{cell, maze.Cells[cell.Row+1][cell.Column]}])
	}
	if cell.Column < maze.Width-1 {
		walls = append(walls, maze.Walls[[2]*Cell{cell, maze.Cells[cell.Row][cell.Column+1]}])
	}

	return walls
}

// getNeighbor returns the neighboring cell of a wall
func (wall *Wall) getNeighbor(cell *Cell) *Cell {
	if wall.Cell1 == cell {
		return wall.Cell2
	}
	if wall.Cell2 == cell {
		return wall.Cell1
	}
	fmt.Println("Error: the cell is not a neighbor of the wall")
	return nil
}

// getUnvisitedNeighbors returns the unvisited neighbors of a cell
func (maze *Maze) getUnvisitedNeighbors(cell *Cell) []*Cell {
	unvisitedNeighbors := make([]*Cell, 0)
	for _, neighbor := range cell.Adjacent {
		if !neighbor.Visited {
			unvisitedNeighbors = append(unvisitedNeighbors, neighbor)
		}
	}
	return unvisitedNeighbors
}

// markWallAsNonExistent marks the wall between two cells as non-existent
func (maze *Maze) markWallAsNonExistent(cell1, cell2 *Cell) {
	for _, neighbor := range cell1.Adjacent {
		if neighbor == cell2 {
			// Found the neighbor cell, mark the corresponding wall as non-existent
			for _, wall := range maze.getWalls(cell1) {
				if (wall.Cell1 == cell1 && wall.Cell2 == cell2) || (wall.Cell1 == cell2 && wall.Cell2 == cell1) {
					wall.Exists = false
					return
				}
			}
		}
	}
}

// Shuffle the slice of cells to randomize traversal
func shuffleNeighbors(neighbors []*Cell) {
	rand.Shuffle(len(neighbors), func(i, j int) {
		neighbors[i], neighbors[j] = neighbors[j], neighbors[i]
	})
}

// Shuffle the slice of walls to randomize traversal
func shuffleWalls(walls []*Wall) {
	rand.Shuffle(len(walls), func(i, j int) {
		walls[i], walls[j] = walls[j], walls[i]
	})
}
