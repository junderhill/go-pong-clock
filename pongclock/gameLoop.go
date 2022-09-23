package pongclock

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

func GameLoop(g Game) {
	defStyle := tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.ColorDefault)
	paddleStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorWhite)

	screen := *g.Screen

	screen.SetStyle(defStyle)

	for {
		g.UpdateGame()
		screen.Clear()
		screen.SetContent(g.Puck.PositionX, g.Puck.PositionY, g.Puck.Display(), nil, tcell.StyleDefault)

		drawSprite(screen,
			g.LeftPaddle.PositionX,
			g.LeftPaddle.PositionY,
			g.LeftPaddle.PositionX+g.LeftPaddle.Width,
			g.LeftPaddle.PositionY+g.LeftPaddle.Height,
			paddleStyle,
			g.LeftPaddle.Display())

		drawSprite(screen,
			g.RightPaddle.PositionX,
			g.RightPaddle.PositionY,
			g.RightPaddle.PositionX+g.RightPaddle.Width,
			g.RightPaddle.PositionY+g.RightPaddle.Height,
			paddleStyle,
			g.RightPaddle.Display())

		screen.Show()
		time.Sleep(40 * time.Millisecond)
	}
}

func drawSprite(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	row := y1
	col := x1

	for _, r := range []rune(text) {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}
