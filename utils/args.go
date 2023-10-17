package utils

import (
	"os"
	"strconv"
)

func ReadArguments() (string, int, int) {
	var imageName string
	var quantX int
	var quantY int
	if len(os.Args) > 1 {
		imageName = os.Args[1]
		if len(os.Args) > 2 {
			x, err := strconv.ParseInt(os.Args[2], 10, 0)
			if err != nil {
				panic(err)
			}
			quantX = int(x)
			quantY = quantX
		}
		if len(os.Args) > 3 {
			y, err := strconv.ParseInt(os.Args[3], 10, 0)
			if err != nil {
				panic(err)
			}
			quantY = int(y)
		}
	}
	return imageName, quantX, quantY
}
