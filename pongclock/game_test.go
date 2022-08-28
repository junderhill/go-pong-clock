package pongclock

import "testing"

type location struct {
	X int
	Y int
}

func TestNewGamePlacesPuckInCenter(t *testing.T) {
	screenWidth := 100
	screenHeight := 50

	expectedLocation := location{
		X: 50,
		Y: 25,
	}

	game := NewGame(screenWidth, screenHeight)

	if game.Puck.PositionX != expectedLocation.X || game.Puck.PositionY != expectedLocation.Y {
		t.Errorf("Puck Position %+v is not as expected: %+v", game.Puck, expectedLocation)
	}
}

func TestNewGamePuckSizeIs1(t *testing.T) {

	expectedSize := 1

	game := NewGame(100, 100)

	if game.Puck.Height != expectedSize || game.Puck.Width != expectedSize {
		t.Errorf("New Game Puck Size unexpected. Expecting %d, Puck %+v", expectedSize, game.Puck)
	}

}

func TestNewGamePaddlesStartInMiddle(t *testing.T) {

	screenWidth := 100
	screenHeight := 50

	expectedY := 25

	game := NewGame(screenWidth, screenHeight)

	if game.LeftPaddle.PositionY != expectedY {
		t.Errorf("Expecting LeftPaddle Y Position to be %d, found %d", expectedY, game.LeftPaddle.PositionY)
	}
	if game.RightPaddle.PositionY != expectedY {
		t.Errorf("Expecting RightPaddle Y Position to be %d, found %d", expectedY, game.RightPaddle.PositionY)
	}

}