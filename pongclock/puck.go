package pongclock

import (
	"math/rand"
	"time"
)

type Puck struct {
	SpeedX int
	SpeedY int

	GameObject
}

const PuckSize int = 1

func NewPuck(startX int, startY int) *Puck {
	rand.Seed(time.Now().UnixNano())
	speedX := rand.Intn(1-(-1)) + (-1)
	speedY := rand.Intn(1-(-1)) + (-1)

	return &Puck{
		SpeedX: speedX,
		SpeedY: speedY,
		GameObject: GameObject{
			PositionX: startX,
			PositionY: startY,
			Height:    PuckSize,
			Width:     PuckSize,
		},
	}
}

func (p *Puck) UpdatePuckPosition() {
	p.PositionX += p.SpeedX
	p.PositionY += p.SpeedY
}

func (p *Puck) Display() rune {
	return '\u25CF'
}
