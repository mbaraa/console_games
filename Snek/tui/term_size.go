package tui

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func GetTermSize() (Point2, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin

	out, err := cmd.Output()
	if err != nil {
		return Point2{}, err
	}

	outSpl := strings.Split(string(out), " ")

	y, err := strconv.Atoi(outSpl[0])
	if err != nil {
		return Point2{}, err
	}

	x, err := strconv.Atoi(outSpl[1][:strings.Index(outSpl[1], "\n")])
	if err != nil {
		return Point2{}, err
	}

	return Point2{X: x, Y: y}, nil
}
