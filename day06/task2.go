package day06

const task2TargetDay = 256

func task2(in string) int {
	return tickArray(parseIntoSlice(parseFishSplitAtoi(in)), task2TargetDay)
}
