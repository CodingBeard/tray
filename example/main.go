package main

import (
	"fmt"
	"time"
	"github.com/codingbeard/tray"
)

func main() {
	trayIcon := tray.ClickableIcon{
		IconData: iconData,
		ClickHandler: func(x, y int, rightClick bool) {
			fmt.Println("x", x, "y", y, "isRightClick", rightClick)
		},
	}
	trayIcon.Initialise()

	fmt.Println("Clickable Icon initialised")

	for {
		time.Sleep(time.Second)
	}
}