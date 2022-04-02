package main

import (
	"os"
	"os/signal"
	"snek/gamestuff"
	"snek/snakes"
)

var sigint chan os.Signal

func main() {
	sigint = make(chan os.Signal)
	signal.Notify(sigint, os.Interrupt)

	go gamestuff.Start(snakes.LevelNormal)
	for range sigint {
		os.Exit(0)
	}
}
