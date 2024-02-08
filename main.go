package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct{}

func (g *Game) Update() error {

	//maze.Prim()
	return nil
}

// Init maze with 100x100 cells
var maze *Maze = InitMaze(25, 25)

func (g *Game) Draw(screen *ebiten.Image) {
	cellSize := 5
	for i := 0; i < len(maze.Cells); i++ {
		for j := 0; j < len(maze.Cells[i]); j++ {
			if maze.Cells[i][j].Visited {
				vector.DrawFilledRect(screen, float32(j*cellSize-1), float32(i*cellSize), float32(cellSize), float32(cellSize), color.White, false)
			}
		}
	}
	for _, wall := range maze.Walls {
		if wall.Exists {
			cell1 := wall.Cell1
			cell2 := wall.Cell2
			x1 := cell1.Column * cellSize
			y1 := cell1.Row * cellSize
			x2 := cell2.Column * cellSize
			y2 := cell2.Row * cellSize

			// Determine the direction of the wall
			// 0: horizontal, 1: vertical
			var direction int
			if cell1.Row == cell2.Row {
				direction = 0 // horizontal wall
			} else {
				direction = 1 // vertical wall
			}

			// Draw the wall around the cells
			if direction == 0 {
				// Horizontal wall
				if cell1.Column < cell2.Column {
					// Wall to the right of cell1
					vector.StrokeLine(screen, float32(x1+cellSize), float32(y1), float32(x1+cellSize), float32(y1+cellSize), 1, color.Black, false)
				} else {
					// Wall to the right of cell2
					vector.StrokeLine(screen, float32(x2+cellSize), float32(y2), float32(x2+cellSize), float32(y2+cellSize), 1, color.Black, false)
				}
			} else {
				// Vertical wall
				if cell1.Row < cell2.Row {
					// Wall below cell1
					vector.StrokeLine(screen, float32(x1), float32(y1+cellSize), float32(x1+cellSize), float32(y1+cellSize), 1, color.Black, false)
				} else {
					// Wall below cell2
					vector.StrokeLine(screen, float32(x2), float32(y2+cellSize), float32(x2+cellSize), float32(y2+cellSize), 1, color.Black, false)
				}
			}
		}
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	//maze.Prim()
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	// Create a channel to signal when Prim's algorithm has finished
	done := make(chan struct{})

	// Run Prim's algorithm in a separate goroutine
	go func() {
		maze.Prim()
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
