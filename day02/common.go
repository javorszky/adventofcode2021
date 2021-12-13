package day02

import (
	"io/ioutil"
	"strings"

	"github.com/javorszky/adventofcode2021/util"
)

const (
	filename  = "day02/input.txt"
	byteZero  = 0x30
	byteOne   = 0x31
	byteTwo   = 0x32
	byteThree = 0x33
	byteFour  = 0x34
	byteFive  = 0x35
	byteSix   = 0x36
	byteSeven = 0x37
	byteEight = 0x38
	byteNine  = 0x39
	byteU     = 0x75
	byteD     = 0x64
	byteF     = 0x66

	intZero  = 0
	intOne   = 1
	intTwo   = 2
	intThree = 3
	intFour  = 4
	intFive  = 5
	intSix   = 6
	intSeven = 7
	intEight = 8
	intNine  = 9
)

var byteInt = map[uint8]int{
	byteZero:  intZero,
	byteOne:   intOne,
	byteTwo:   intTwo,
	byteThree: intThree,
	byteFour:  intFour,
	byteFive:  intFive,
	byteSix:   intSix,
	byteSeven: intSeven,
	byteEight: intEight,
	byteNine:  intNine,
}

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs() []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimRight(string(data), util.NewLine), util.NewLine)
}
