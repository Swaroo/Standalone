package main

import (
	"fmt"
)

func getValue (s string) uint16{
	switch s{
	case "I": return 1
	case "V": return 5
	case "X": return 10
	case "L": return 50
	case "C": return 100
	case "D": return 500
	case "M": return 1000
	}
	return 0
}
func romanToInt(s string) int {
	var i int
	var finalVal uint16

	i = 0
	finalVal = 0
	nextVal := uint16(0)
	for i < len(s){
		currVal := getValue(string(s[i]))
		if i+1 < len(s) {
			nextVal = getValue(string(s[i+1]))
		}else{
			nextVal = 0
		}
		if nextVal > currVal {
			i = i +2
			finalVal += nextVal-currVal
		}else{
			i++
			finalVal += currVal
		}
	}
	return int(finalVal)
}

func main() {
	fmt.Println(romanToInt("III"))
    fmt.Println(romanToInt("IV"))
    fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
