package pongclock

import "math"

type Game struct {
	Puck        Puck
	LeftPaddle  Paddle
	RightPaddle Paddle
}

const PuckSize int = 1

func NewGame(screenWidth int, screenHeight int) *Game {
	return &Game{
		Puck: Puck{
			SpeedX: 0,
			SpeedY: 0,
			GameObject: GameObject{
				PositionX: findMiddle(screenWidth),
				PositionY: findMiddle(screenHeight),
				Height:    PuckSize,
				Width:     PuckSize,
			},
		},
		LeftPaddle: Paddle{
			SpeedY: 0,
			GameObject: GameObject{
				PositionX: 0,
				PositionY: findMiddle(screenHeight),
				Height:    0,
				Width:     0,
			},
		},
		RightPaddle: Paddle{
			SpeedY: 0,
			GameObject: GameObject{
				PositionX: 0,
				PositionY: findMiddle(screenHeight),
				Height:    0,
				Width:     0,
			},
		},
	}
}

func findMiddle(dimension int) int {
	return int(math.Round(float64(dimension / 2)))
}
