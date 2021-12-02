package day02

import "log"

var byteInt = map[uint8]int{
	0x30: 0,
	0x31: 1,
	0x32: 2,
	0x33: 3,
	0x34: 4,
	0x35: 5,
	0x36: 6,
	0x37: 7,
	0x38: 8,
	0x39: 9,
}

func task1(input []string) int {
	depth := 0
	forward := 0

	for _, instruction := range input {
		cmd, val := fetchCmdAndValue(instruction)
		switch cmd {
		case 0x75:
			depth -= byteInt[val]
		case 0x64:
			depth += byteInt[val]
		case 0x66:
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
