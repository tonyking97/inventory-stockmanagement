package utils

import "math"

func round(num float32) int {
	return int(float64(num) + math.Copysign(0.5, float64(num)))
}

func RoundOff(num float32, precision int) float32 {
	output := math.Pow(10, float64(precision))
	return float32(round(num*float32(output))) / float32(output)
}
