package main
import "fmt"
func zigzagconversioninit() {
	result := convert("PAYPALISHIRING", 3)
	fmt.Println(result)
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([]string, numRows)
	row := 0
	direction := 1
	for _, c := range s {
		rows[row] += string(c)
		if row == 0 {
			direction = 1
		} else if row == numRows-1 {
			direction = -1
		}
		row += direction
	}
	result := ""
	for _, r := range rows {
		result += r
	}
	return result
}