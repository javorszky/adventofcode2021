package day02

import "log"

func task1(input []string) int {
	depth := 0
	forward := 0

	for _, instruction := range input {
		cmd, val := fetchCmdAndValue(instruction)
		switch cmd {
		case byteU:
			depth -= byteInt[val]
		case byteD:
			depth += byteInt[val]
		case byteF:
			forward += byteInt[val]
		default:
			log.Fatalf("unrecognisable input: %s", instruction)
		}
	}

	return depth * forward
}

func fetchCmdAndValue(line string) (uint8, uint8) {
	return line[0], line[len(line)-1]
}
