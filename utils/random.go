package utils

import (
	"math/rand"
	"time"
)

func GetRandom(leng int, key string) int {
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	var keyNum []int
	for _, v := range key {
		if v >= 48 && v <= 57 {
			keyNum = append(keyNum, int(v)-48)
		}
	}
	var keyInt int
	for _, v := range keyNum {
		keyInt = keyInt*10 + v
	}
	r2 := rand.New(rand.NewSource(int64(keyInt)))
	rand := r1.Intn(leng) + r2.Intn(leng) - rand.Intn(leng)
	if rand < 0 {
		rand = -rand
	}
	return rand
}
