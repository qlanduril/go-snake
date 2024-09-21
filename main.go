package main

import (
	"fmt"
	"image/color"
	"log"

	snake "github.com/qlanduril/go-snake/snake"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

// Game implements ebiten.Game interface.
type Game struct {
	boardWidth  int
	boardHeight int
	snakeBody   *snake.Snake
}

func NewGame(sizeX, sizeY int) *Game {

	return &Game{
		boardWidth:  sizeX,
		boardHeight: sizeY,
		snakeBody:   snake.NewSnake(2, nil),
	}
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	// Define the color for the rectangle
	rectColor := color.RGBA{255, 0, 0, 255}

	// Create an image to represent the rectangle
	rect := ebiten.NewImage(1, 500)
	rect.Fill(rectColor)

	// Draw the rectangle onto the screen
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0) // Position the rectangle at (50, 50)
	screen.DrawImage(rect, op)

	// Optionally, draw debug information
	//ebitenutil.DebugPrint(screen, "Hello, World!")
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	fmt.Println("hello world")

	game := &Game{}
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("go-snake")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
