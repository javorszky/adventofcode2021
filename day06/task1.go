package day06

const (
	targetDay    = 80
	defaultCycle = 7
)

func task1(input string) int {
	return calculateAllSpawns(getSpawnDays(parseFishSplitAtoi(input)), targetDay)
}

func spawnsOn(day, current, until, cycle int) []int {
	spawnDays := make([]int, 0)

	for i := day + current; i < until; i += cycle {
		spawnDays = append(spawnDays, i+1)
	}

	return spawnDays
}

func getSpawnDays(in []int) []int {
	out := make([]int, len(in))
	for i, age := range in {
		out[i] = age - (defaultCycle + 1)
	}

	return out
}

func calculateAllSpawns(in []int, target int) int {
	if len(in) == 0 {
		return 0
	}

	acc := len(in)
	spawns := make([]int, 0)

	for _, d := range in {
		newSpawns := spawnsOn(d, 8, target, defaultCycle)
		spawns = append(spawns, newSpawns...)
	}

	return acc + calculateAllSpawns(spawns, target)
}
