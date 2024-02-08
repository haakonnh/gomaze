package main

import (
	"fmt"
	"image/color"
	"log"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct{}

func (g *Game) Update() error {

	//maze.Prim()
	return nil
}

// Init maze with width x height cells
var maze *Maze = InitMaze(63, 63)

func (g *Game) Draw(screen *ebiten.Image) {
	fmt.Println((ebiten.ActualFPS())) // fps tracker

	// Size of each cell in pixels
	cellSize := 5

	currentWallColor := color.RGBA{255, 0, 0, 0} // used for drawing correct colour for walls

	// Loop through all cells and draw cells and walls
	for i := 0; i < len(maze.Cells); i++ {
		for j := 0; j < len(maze.Cells[i]); j++ {

			// Draw cells
			if maze.Cells[i][j].IsSearched { // searched cells are green - for maze exploration
				vector.DrawFilledRect(screen, float32(j*cellSize-1), float32(i*cellSize), float32(cellSize), float32(cellSize), color.RGBA{0, 255, 0, 0}, false)
			} else if maze.Cells[i][j].Visited { // visited cells are red - made for maze generation
				vector.DrawFilledRect(screen, float32(j*cellSize-1), float32(i*cellSize), float32(cellSize), float32(cellSize), color.RGBA{255, 0, 0, 0}, false)
			}
			if i == maze.Width-1 && j == maze.Height-1 { // end cell is blue
				vector.DrawFilledRect(screen, float32(j*cellSize-1), float32(i*cellSize), float32(cellSize), float32(cellSize), color.RGBA{0, 0, 255, 0}, false)
			}
			thickness := float32(1.0) // thickness of walls

			// Draw walls
			if i > 0 { // if not the first row
				// If the maze-gen algorithm has walked through wall between these two cells - draw it green, else red
				if slices.Contains(maze.Walks, Walk{From: maze.Cells[i][j], To: maze.Cells[i-1][j]}) ||
					slices.Contains(maze.Walks, Walk{From: maze.Cells[i-1][j], To: maze.Cells[i][j]}) {
					if maze.Cells[i][j].IsSearched && maze.Cells[i-1][j].IsSearched {
						currentWallColor = color.RGBA{0, 255, 0, 0}
					} else {
						currentWallColor = color.RGBA{255, 0, 0, 0}
					}
					vector.StrokeLine(screen, float32(j*cellSize), float32(i*cellSize), float32(j*cellSize+cellSize), float32(i*cellSize), thickness, currentWallColor, false)
				} else {
					vector.StrokeLine(screen, float32(j*cellSize), float32(i*cellSize), float32(j*cellSize+cellSize), float32(i*cellSize), thickness, color.White, false)
				}
			} else if i == 0 { // draw top wall
				vector.StrokeLine(screen, float32(j*cellSize), float32(i*cellSize), float32(j*cellSize+cellSize), float32(i*cellSize), thickness, color.White, false)

			}

			if i < len(maze.Cells)-1 { // if not the last row
				// If the maze-gen algorithm has walked through wall between these two cells - draw it green, else red
				if slices.Contains(maze.Walks, Walk{From: maze.Cells[i][j], To: maze.Cells[i+1][j]}) ||
					slices.Contains(maze.Walks, Walk{From: maze.Cells[i+1][j], To: maze.Cells[i][j]}) {
					if maze.Cells[i][j].IsSearched && maze.Cells[i+1][j].IsSearched {
						currentWallColor = color.RGBA{0, 255, 0, 0}
					} else {
						currentWallColor = color.RGBA{255, 0, 0, 0}
					}
					vector.StrokeLine(screen, float32(j*cellSize), float32(i*cellSize+cellSize), float32(j*cellSize+cellSize), float32(i*cellSize+cellSize), thickness, currentWallColor, false)
				} else {
					vector.StrokeLine(screen, float32(j*cellSize), float32(i*cellSize+cellSize), float32(j*cellSize+cellSize), float32(i*cellSize+cellSize), thickness, color.White, false)
				}

			} else if i == len(maze.Cells)-1 { // draw bottom wall
				vector.StrokeLine(screen, float32(j*cellSize), float32(i*cellSize+cellSize), float32(j*cellSize+cellSize), float32(i*cellSize+cellSize), thickness, color.White, false)
			}

			if j > 0 { // if not the first column
				if slices.Contains(maze.Walks, Walk{From: maze.Cells[i][j], To: maze.Cells[i][j-1]}) ||
					slices.Contains(maze.Walks, Walk{From: maze.Cells[i][j-1], To: maze.Cells[i][j]}) {
					if maze.Cells[i][j].IsSearched && maze.Cells[i][j-1].IsSearched {
						currentWallColor = color.RGBA{0, 255, 0, 0}
					} else {
						currentWallColor = color.RGBA{255, 0, 0, 0}
					}
					vector.StrokeLine(screen, float32(j*cellSize), float32(i*cellSize), float32(j*cellSize), float32(i*cellSize+cellSize), thickness, currentWallColor, false)
				} else {
					vector.StrokeLine(screen, float32(j*cellSize), float32(i*cellSize), float32(j*cellSize), float32(i*cellSize+cellSize), thickness, color.White, false)
				}

			} else if j == 0 { // draw left wall
				vector.StrokeLine(screen, float32(j*cellSize), float32(i*cellSize), float32(j*cellSize), float32(i*cellSize+cellSize), thickness, color.White, false)
			}

			if j < len(maze.Cells[i])-1 { // if not the last column
				if slices.Contains(maze.Walks, Walk{From: maze.Cells[i][j], To: maze.Cells[i][j+1]}) ||
					slices.Contains(maze.Walks, Walk{From: maze.Cells[i][j+1], To: maze.Cells[i][j]}) {
					if maze.Cells[i][j].IsSearched && maze.Cells[i][j+1].IsSearched {
						currentWallColor = color.RGBA{0, 255, 0, 0}
					} else {
						currentWallColor = color.RGBA{255, 0, 0, 0}
					}
					vector.StrokeLine(screen, float32(j*cellSize+cellSize), float32(i*cellSize), float32(j*cellSize+cellSize), float32(i*cellSize+cellSize), thickness, currentWallColor, false)
				} else {
					vector.StrokeLine(screen, float32(j*cellSize+cellSize), float32(i*cellSize), float32(j*cellSize+cellSize), float32(i*cellSize+cellSize), thickness, color.White, false)
				}
			} else if j == len(maze.Cells[i])-1 { // draw right wall
				vector.StrokeLine(screen, float32(j*cellSize+cellSize), float32(i*cellSize), float32(j*cellSize+cellSize), float32(i*cellSize+cellSize), thickness, color.White, false)
			}

		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 320
}

func main() {
	//maze.Prim()
	ebiten.SetWindowSize(960, 960)
	ebiten.SetWindowTitle("Hello, World!")

	// Create a channel to signal when Prim's algorithm has finished
	done := make(chan struct{})

	// Run Prim's algorithm in a separate goroutine
	go func() {
		maze.Prim()
		maze.DFS()
		// Signal that Prim's algorithm has finished
		close(done)
	}()
	// Run the main loop to continuously draw the maze
	for !isDone(done) {
		// Advance the game state
		if err := ebiten.RunGame(&Game{}); err != nil {
			log.Fatal(err)
		}
	}
}

// isDone checks if Prim's algorithm has finished
func isDone(done <-chan struct{}) bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
