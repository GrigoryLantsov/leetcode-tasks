package main
import (
	"math"
	"fmt"
)

func AtoiInit() {
	result := myAtoi("         42")
	fmt.Println(result)
}

func myAtoi(s string) int {
	num := 0
	sign := 1
	started := false

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' && !started {
			continue
		} else if s[i] == '-' && !started {
			sign = -1
			started = true
		} else if s[i] == '+' && !started {
			sign = 1
			started = true
		} else if s[i] >= '0' && s[i] <= '9' {
			num = num*10 + int(s[i]-'0')
			started = true
			if sign*num < math.MinInt32 {
				return math.MinInt32
			} else if sign*num > math.MaxInt32 {
				return math.MaxInt32
			}
		} else {
			break
		}
	}

	return num * sign
}
