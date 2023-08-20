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

func TestLeftPaddlePositionIsUpdated(t *testing.T) {
	game := NewGame(500, 500, nil)

	startY := game.LeftPaddle.PositionY

	game.UpdateGame()

	expectedY := startY + game.LeftPaddle.SpeedY

	if game.LeftPaddle.PositionY != expectedY {
		t.Errorf("Left paddle Y position is %d Should be %d", game.LeftPaddle.PositionY, expectedY)
	}
}

func TestRightPaddlePositionIsUpdated(t *testing.T) {
	game := NewGame(500, 500, nil)

	startY := game.RightPaddle.PositionY

	game.UpdateGame()

	expectedY := startY + game.RightPaddle.SpeedY

	if game.RightPaddle.PositionY != expectedY {
		t.Errorf("Left paddle Y position is %d Should be %d", game.LeftPaddle.PositionY, expectedY)
	}
}

func TestPuckCollisionWithTopReversesYDirection(t *testing.T) {
	game := NewGame(500, 500, nil)

	game.Puck.PositionY = 0
	game.Puck.SpeedY = -1

	game.UpdateGame()

	if game.Puck.SpeedY == -1 {
		t.Errorf("Expected Puck SpeedY to be reversed but is %d", game.Puck.SpeedY)
	}
}
