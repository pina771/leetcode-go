package main

func main() {
}

func minimumRounds(tasks []int) int {
	retval := 0
	numOfTasks := make(map[int]int)
	for _, val := range tasks {
		numOfTasks[val]++
	}

	for _, val := range numOfTasks {
		if val == 1 {
			return -1
		}
		if val%3 == 0 {
			retval += val / 3
		}

		if val%3 == 2 {
			retval = retval + val/3 + 1
		}

		// let's say it's 7
		// it's 7/3 = 2 -> 2 - 1 = 1 round 3 tasks
		// and 2 rounds of 2 tasks
		if val%3 == 1 {
			retval = retval + (val/3 - 1) + 2
		}

	}

	return retval
}
