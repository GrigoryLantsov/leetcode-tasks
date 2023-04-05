package main
import (
	"fmt"
	"strings"
)
func englishformatinit() {
	result := numberToWords(100000)
	fmt.Println(result)
}

var (
	ones  = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	teens = []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	tens  = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
)

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	// thousands, millions, billions
	groups := []string{"", "Thousand", "Million", "Billion"}
	res := ""
	for i, g := 0, 0; num > 0; i, g, num = i+1, g+1, num/1000 {
		part := num % 1000
		if part > 0 {
			words := groupToWords(part)
			if g > 0 {
				words += " " + groups[g]
			}
			if len(res) > 0 {
				res = words + " " + res
			} else {
				res = words
			}
		}
	}
	return strings.TrimSpace(res)
}

func groupToWords(n int) string {
	if n == 0 {
		return ""
	} else if n < 10 {
		return ones[n]
	} else if n < 20 {
		return teens[n-10]
	} else if n < 100 {
		if n%10 == 0 {
			return tens[n/10]
		}
		return tens[n/10] + " " + ones[n%10]
	} else {
		if n%100 == 0 {
			return ones[n/100] + " Hundred"
		}
		return ones[n/100] + " Hundred " + groupToWords(n%100)
	}
}
