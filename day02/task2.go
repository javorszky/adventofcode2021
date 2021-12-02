package day02

import "log"

func task2(input []string) int {
	aim := 0
	depth := 0
	forward := 0

	for _, instruction := range input {
		cmd, val := fetchCmdAndValue(instruction)
		switch cmd {
		case byteU:
			aim -= byteInt[val]
		case byteD:
			aim += byteInt[val]
		case byteF:
			forward += byteInt[val]
			depth += aim * byteInt[val]
		default:
			log.Fatalf("unrecognisable input: %s", instruction)
		}
	}

	return depth * forward
}
