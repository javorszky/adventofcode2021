package day06

func task1(input string) interface{} {
	return input
}

const (
	targetDay      = 80
	defaultCycle   = 7
	firstCyclePlus = 2
)

func spawnsOn(day, current, until, cycle int) []int {
	spawnDays := make([]int, 0)

	for i := day + current; i < until; i += cycle {
		spawnDays = append(spawnDays, i+1)
	}

	return spawnDays
}
