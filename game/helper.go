package game

import (
	"math/rand"
	"time"
)

func getRandomNumber(max int, exclude ...int) int {
	if len(exclude) == 0 {
		return getRandom(max)
	}
	num := getRandom(max)
	for inExclusion(exclude, num) {
		num = getRandom(max)
	}
	return num
}

func inExclusion(list []int, num int) bool {
	for _, item :=  range list {
		if item == num {
			return true
		}
	}
	return false
}

func getRandom(max int) int {
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	return randomizer.Intn(max)
}

