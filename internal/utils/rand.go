package utils

import "math/rand"

func RandInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func RandFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
