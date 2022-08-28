package pongclock

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
				PositionX: screenWidth / 2,
				PositionY: screenHeight / 2,
				Height:    PuckSize,
				Width:     PuckSize,
			},
		},
		LeftPaddle: Paddle{
			SpeedY: 0,
			GameObject: GameObject{
				PositionX: 0,
				PositionY: 0,
				Height:    0,
				Width:     0,
			},
		},
		RightPaddle: Paddle{
			SpeedY: 0,
			GameObject: GameObject{
				PositionX: 0,
				PositionY: 0,
				Height:    0,
				Width:     0,
			},
		},
	}
}
