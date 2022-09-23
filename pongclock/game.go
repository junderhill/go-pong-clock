package pongclock

import (
	"math"

	"github.com/gdamore/tcell/v2"
)

type Game struct {
	Puck        Puck
	LeftPaddle  Paddle
	RightPaddle Paddle
	Screen      *tcell.Screen
}

func NewGame(screenWidth int, screenHeight int, screen *tcell.Screen) *Game {
	return &Game{
		Puck: *NewPuck(findMiddle(screenWidth), findMiddle(screenHeight)),
		LeftPaddle: Paddle{
			SpeedY: -1,
			GameObject: GameObject{
				PositionX: 5,
				PositionY: findMiddle(screenHeight),
				Height:    5,
				Width:     1,
			},
		},
		RightPaddle: Paddle{
			SpeedY: 1,
			GameObject: GameObject{
				PositionX: screenWidth - 5,
				PositionY: findMiddle(screenHeight),
				Height:    5,
				Width:     1,
			},
		},
		Screen: screen,
	}
}

func findMiddle(dimension int) int {
	return int(math.Round(float64(dimension / 2)))
}

func (g *Game) UpdateGame() {
	g.Puck.UpdatePuckPosition()
	g.LeftPaddle.UpdatePaddlePosition()
	g.RightPaddle.UpdatePaddlePosition()
}
