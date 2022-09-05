package pongclock

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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

	assert.Equal(t, game.LeftPaddle.PositionY, expectedY, "Left paddle should start in middle")
	assert.Equalf(t, game.RightPaddle.PositionY, expectedY, "Right paddle should start in middle")
}
