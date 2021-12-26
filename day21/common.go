package day21

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/javorszky/adventofcode2021/util"
)

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs(fn string) (int, int) {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), util.NewLine)

	player1String := lines[0][len(lines[0])-1]
	player2String := lines[1][len(lines[1])-1]

	player1, err := strconv.Atoi(string(player1String))
	if err != nil {
		log.Fatalf("player 1 could not be parsed into an int")
	}

	player2, err := strconv.Atoi(string(player2String))
	if err != nil {
		log.Fatalf("player 2 could not be parsed into an int")
	}

	return player1, player2
}
