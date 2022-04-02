package snakes

import (
	"math/rand"
	"snek/tui"
	"time"
)

type Apple struct {
	Position    tui.Point2
	Color       tui.TermColor
	GainedScore int
}

type Level uint

const (
	LevelNormal Level = iota + 1
	LevelCrazy
	LevelInsane
	LevelSnek
	LevelBerzerk
)

func GetApples(mxPt tui.Point2, level Level, color tui.TermColor) []Apple {
	var (
		count      = getApplesCount(mxPt, level)
		start, end = getMinMaxPoints(mxPt, level)
		apples     = make([]Apple, count)
		score      = getScore(level)
	)

	for i := uint(0); i < count; i++ {
		rand.Seed(time.Now().UnixMicro())
		apples[i] = Apple{
			Color: color,
			Position: tui.Point2{
				X: rand.Intn(end.X-start.X) + start.X,
				Y: rand.Intn(end.Y-start.Y) + start.Y,
			},
			GainedScore: score,
		}
	}

	return apples
}

func getApplesCount(mxPt tui.Point2, level Level) uint {
	numCells := mxPt.X * mxPt.Y
	rand.Seed(time.Now().UnixMicro())
	switch level {
	default:
		fallthrough
	case LevelNormal:
		return uint(0.01 * float64(numCells))
	case LevelCrazy:
		return uint(0.15 * float64(numCells))
	case LevelInsane:
		return uint(0.25 * float64(numCells))
	case LevelSnek:
		return uint(0.80 * float64(numCells))
	case LevelBerzerk:
		return uint(0.95 * float64(numCells))
	}
}

func getMinMaxPoints(mxPt tui.Point2, level Level) (min, max tui.Point2) {
	excludeRate := 0.05
	switch level {
	case LevelNormal:
		excludeRate = 0.05
	case LevelCrazy:
		excludeRate = 0.03
	case LevelInsane:
		excludeRate = 0.02
	case LevelSnek:
		excludeRate = 0.01
	case LevelBerzerk:
		excludeRate = 0
	}

	min = tui.Point2{
		X: int(excludeRate * float64(mxPt.X)),
		Y: int(excludeRate * float64(mxPt.Y)),
	}

	max = tui.Point2{
		X: mxPt.X - int(excludeRate*float64(mxPt.X)),
		Y: mxPt.Y - int(excludeRate*float64(mxPt.Y)),
	}

	return
}

func getScore(level Level) int {
	return 5 - int(level)
}
