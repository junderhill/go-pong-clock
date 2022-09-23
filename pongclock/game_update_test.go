package pongclock

import "testing"

func TestPuckPositionIsUpdated(t *testing.T) {
	game := NewGame(500, 500, nil)

	startX, startY := game.Puck.PositionY, game.Puck.PositionY

	game.UpdateGame()

	expectedX := startX + game.Puck.SpeedX
	expectedY := startY + game.Puck.SpeedY

	if game.Puck.PositionX != expectedX {
		t.Errorf("Puck X Position is %d Should be %d", game.Puck.PositionX, expectedX)
	}
	if game.Puck.PositionY != expectedY {
		t.Errorf("Puck Y Position is %d Should be %d", game.Puck.PositionY, expectedY)
	}
}
