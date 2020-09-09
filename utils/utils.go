package utils

import (
	"math/rand"
	"time"
)

func RandomNoGen(min int, max int) int {
	if min < max {
		rand.Seed(time.Now().UnixNano())
		randomNo := rand.Intn(max-min+1) + min
		return randomNo
	} else {
		return 0
	}
}
