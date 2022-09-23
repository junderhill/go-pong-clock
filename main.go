package main

import (
	"log"
	"os"
	"pong-clock/pongclock"

	"github.com/gdamore/tcell/v2"
)

func main() {
	screen, err := tcell.NewScreen()

	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorBlack)
	screen.SetStyle(defStyle)

	width, height := screen.Size()

	game := pongclock.NewGame(width, height, &screen)

	go pongclock.GameLoop(*game)

	for {
		screen.Size()
		switch event := screen.PollEvent().(type) {
		case *tcell.EventResize:
			screen.Sync()
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				screen.Fini()
				os.Exit(0)
			}
		}
	}
}
