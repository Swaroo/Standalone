package main

import (
    "fmt"
    "math"
	"strconv"
    "strings"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverseWithString(x int) int {

	neg := false

	str := strconv.Itoa(x)
	if strings.HasPrefix(str,"-"){
		neg = true
		str = strings.Replace(str,"-","",1)
	}
	for strings.HasSuffix(str,"0"){
		str = strings.TrimSuffix(str,"0")
	}
	str = Reverse(str)

	rev, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	if neg {
		rev = rev * -1
        if rev < math.MinInt32{
            return 0
        }
	} else{
	    if rev > math.MaxInt32{
	        return 0
        }
    }
	return rev
}

func main() {
	fmt.Println(reverseWithString(10220))
    fmt.Println(reverseWithString(123))
    fmt.Println(reverseWithString(-123))
    fmt.Println(reverseWithString(12500))
	fmt.Println(reverseWithString(1534236469))
}
