package utils

import (
	"fmt"
	"strings"
)

func PrintProbabilityMap(m map[rune]float64) {
	mapStrings := TransformMap(m, func(k rune, v float64) string {
		return fmt.Sprintf("%s: %.3f", string(k), v)
	})
	println(strings.Join(mapStrings, ", "))
}
