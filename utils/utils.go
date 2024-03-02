package utils

import (
	"strconv"

	"golang.org/x/term"
)

func termSize() [2]int {
	x, y, err := term.GetSize(0)
	if err != nil {
		panic(err)
	}
	return [2]int{x, y}
}

func TermHeight() int {
	return termSize()[1]
}

func TermWidth() int {
	return termSize()[0]
}

func NewLine() string {
	return "\n"
}

func Atoi(num string) int {
	val, _ := strconv.Atoi(num)
	return val
}
