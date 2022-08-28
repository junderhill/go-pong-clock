package pongclock

type GameObject struct {
	PositionX int
	PositionY int

	Height int
	Width  int
}

type Puck struct {
	SpeedX int
	SpeedY int

	GameObject
}

type Paddle struct {
	SpeedY int

	GameObject
}
