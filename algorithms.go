package main

import (
	"fmt"
	"math/rand"
	"slices"
	"time"

	"github.com/spakin/disjoint"
)

type Wall struct {
	Exists bool // Flag indicating whether the wall exists or not
	Cell1  *Cell
	Cell2  *Cell
}

type Cell struct {
	Row           int
	Column        int
	Visited       bool
	Adjacent      []*Cell
	WalkedThrough [4]bool // 0: top, 1: right, 2: bottom, 3: left
	IsSearched    bool
	KruskalSet    *disjoint.Element
}

type Maze struct {
	Width  int
	Height int
	Cells  [][]*Cell
	Walks  []Walk // List of walls that have been walked through by the maze generation algorithm
}

// Walk is a struct that represents a wall that has been walked through by the maze generation algorithm
type Walk struct {
	From *Cell
	To   *Cell
}

// InitMaze creates a maze with the given width and height
func InitMaze(width, height int) *Maze {

	// Create a maze with width x height cells and initialized Cells and Walks
	maze := &Maze{
		Width:  width,
		Height: height,
		Cells:  make([][]*Cell, height),
		Walks:  make([]Walk, 0),
	}

	// Create cells
	for i := range maze.Cells {
		maze.Cells[i] = make([]*Cell, width)
		for j := range maze.Cells[i] {
			maze.Cells[i][j] = &Cell{
				Row:        i,
				Column:     j,
				Visited:    false,
				Adjacent:   make([]*Cell, 0),
				IsSearched: false,
			}
		}
	}

	// Populate adjacent cells for each cell
	for i := range maze.Cells {
		for j := range maze.Cells[i] {
			if i > 0 {
				maze.Cells[i][j].Adjacent = append(maze.Cells[i][j].Adjacent, maze.Cells[i-1][j])
			}
			if i < height-1 {
				maze.Cells[i][j].Adjacent = append(maze.Cells[i][j].Adjacent, maze.Cells[i+1][j])
			}
			if j > 0 {
				maze.Cells[i][j].Adjacent = append(maze.Cells[i][j].Adjacent, maze.Cells[i][j-1])
			}
			if j < width-1 {
				maze.Cells[i][j].Adjacent = append(maze.Cells[i][j].Adjacent, maze.Cells[i][j+1])
			}
		}
	}

	return maze
}

// Prim generates a maze using the Prim algorithm
func (maze *Maze) Prim() {
	// Choose the starting cell [0,0] and mark it as visited
	startCell := maze.Cells[0][0]
	startCell.Visited = true

	// A slice of walks representing the currently possible walks from a cell in the Prim tree to a cell not in the tree
	possibleWalks := make([]Walk, 0)
	for _, cell := range startCell.Adjacent { // Add the adjacent walls of the starting cell to the list of possible walks
		possibleWalks = append(possibleWalks, Walk{From: startCell, To: cell})
	}

	// Visited cells represent the cells that have been visited by the Prim algorithm
	visitedCells := make([]*Cell, 0)
	visitedCells = append(visitedCells, startCell)

	// While there are still walks in the list of possible walks
	for len(possibleWalks) > 0 {
		time.Sleep(time.Second / 5000) // remove this line to not see visual generation

		// Choose a random walk from the list of possible walks
		randIndex := rand.Intn(len(possibleWalks))
		walk := possibleWalks[randIndex]
		possibleWalks = append(possibleWalks[:randIndex], possibleWalks[randIndex+1:]...) // Remove the walk from the list of possible walks

		// If the cell on the other side of the walk has not been visited
		if !walk.To.Visited {
			walk.To.Visited = true

			// Add the walk to the list of walks in the maze and add the walked-to cell to the tree
			maze.Walks = append(maze.Walks, walk)
			visitedCells = append(visitedCells, walk.To)

			// Add the adjacent walls of the cell to the list of walls
			for _, cell := range walk.To.Adjacent {
				if !slices.Contains(visitedCells, cell) {
					possibleWalks = append(possibleWalks, Walk{From: walk.To, To: cell})
				}
			}
		}
	}

}

// Kruskal generates a maze using the Kruskal algorithm
func (maze *Maze) Kruskal() {
	possibleWalks := make([]Walk, 0)
	for i := range maze.Cells {
		for j := range maze.Cells[i] {
			maze.Cells[i][j].KruskalSet = disjoint.NewElement()
			for _, walkable := range maze.Cells[i][j].Adjacent {
				possibleWalks = append(possibleWalks, Walk{From: maze.Cells[i][j], To: walkable})
			}
		}
	}
	for len(possibleWalks) > 0 {
		time.Sleep(time.Second / 100) // remove this line to not see visual generation
		randIndex := rand.Intn(len(possibleWalks))
		walk := possibleWalks[randIndex]
		possibleWalks = append(possibleWalks[:randIndex], possibleWalks[randIndex+1:]...)
		if walk.From.KruskalSet.Find() != walk.To.KruskalSet.Find() {
			walk.To.Visited = true
			walk.From.Visited = true
			maze.Walks = append(maze.Walks, walk)
			maze.Walks = append(maze.Walks, Walk{From: walk.To, To: walk.From})
			disjoint.Union(walk.From.KruskalSet, walk.To.KruskalSet)

		}
	}
}

// DFS maze solver algorithm
func (maze *Maze) DFS() {
	// Choose the starting cell [0,0] and mark it as visite
	startCell := maze.Cells[0][0]

	// Solve maze recursively (depth-first search)
	var b bool = recursiveSolve(maze, startCell)

	// Print if the maze was solved
	if b {
		fmt.Println("DFS SOLVED")
		return
	}
	fmt.Println("Not solved")
}

// Recursive function to solve the maze with dfs
func recursiveSolve(maze *Maze, cell *Cell) bool {
	time.Sleep(time.Second / 1000) // remove this line to not see visual generation
	// If the cell is the exit cell, return true - the maze has been solved
	if cell.Column == maze.Width-1 && cell.Row == maze.Height-1 {
		cell.IsSearched = true
		return true
	}

	// If the cell has been searched skip this cell
	if cell.IsSearched {
		return false
	}

	// Mark the cell as searched
	cell.IsSearched = true

	// If the cell is not rightmost column and the maze-gen algorithm has walked through the wall between these two cells, recursively solve the maze at the cell to the right
	if cell.Column < maze.Width-1 && slices.Contains(maze.Walks, Walk{From: cell, To: maze.Cells[cell.Row][cell.Column+1]}) {
		if recursiveSolve(maze, maze.Cells[cell.Row][cell.Column+1]) {
			return true
		}
	}

	// If the cell is not leftmost column and the maze-gen algorithm has walked through the wall between these two cells, recursively solve the maze at the cell to the left
	if cell.Column > 0 && slices.Contains(maze.Walks, Walk{From: cell, To: maze.Cells[cell.Row][cell.Column-1]}) {
		if recursiveSolve(maze, maze.Cells[cell.Row][cell.Column-1]) {
			return true
		}
	}

	// If the cell is not top row and the maze-gen algorithm has walked through the wall between these two cells, recursively solve the maze at the cell above
	if cell.Row > 0 && slices.Contains(maze.Walks, Walk{From: cell, To: maze.Cells[cell.Row-1][cell.Column]}) {
		if recursiveSolve(maze, maze.Cells[cell.Row-1][cell.Column]) {
			return true
		}
	}

	// If the cell is not bottom row and the maze-gen algorithm has walked through the wall between these two cells, recursively solve the maze at the cell below
	if cell.Row < maze.Height-1 && slices.Contains(maze.Walks, Walk{From: cell, To: maze.Cells[cell.Row+1][cell.Column]}) {
		if recursiveSolve(maze, maze.Cells[cell.Row+1][cell.Column]) {
			return true
		}
	}

	return false

}
