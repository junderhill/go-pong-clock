package pongclock

import "strings"

type Paddle struct {
	SpeedY int

	GameObject
}

const PaddleSpeed int = 1

func (p *Paddle) UpdatePaddlePosition() {
	p.PositionY += p.SpeedY
}

func (p *Paddle) Display() string {
	return strings.Repeat(" ", p.Height)
}

func (p *Paddle) MoveUp() {
	p.SpeedY = PaddleSpeed
}

func (p *Paddle) MoveDown() {
	p.SpeedY = PaddleSpeed * -1
}

func (p *Paddle) CheckForScreenCollisions(height int) bool {
	if p.PositionY <= 1 {
		return true
	}

	if (p.PositionY + p.Height) >= height-1 {
		return true
	}

	return false
}

func (p *Paddle) ReverseDirection() {
	p.SpeedY = p.SpeedY * -1
}
