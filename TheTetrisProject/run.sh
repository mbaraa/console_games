#!/bin/sh


# 1 to tetris 2 for rotation demo
# 3 to build tetris

if [ $1 == 1 ]
then
	go run main.go Constants.go Shared.go TetrisFuncs.go Tetromino.go 

elif [ $1 == 2 ]
then
	go run demo.go Constants.go Shared.go TetrisFuncs.go Tetromino.go

elif [ $1 == 3 ]
then 
	go build main.go Constants.go Shared.go TetrisFuncs.go Tetromino.go
fi
