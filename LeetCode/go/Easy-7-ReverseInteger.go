package main

import (
    "fmt"
    "math"
)

func reverse(x int) int {

	rev := 0
	neg := false
	still0 := true
	if x < 0 {
		neg = true
		x = x * -1
	}
	for x != 0 {
		rem := x % 10
		x = x / 10
		if !still0  || rem != 0 {
            rev = rev*10 + rem
            still0 = false
        }
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
	fmt.Println(reverse(10220))
    fmt.Println(reverse(123))
    fmt.Println(reverse(-123))
    fmt.Println(reverse(12500))
	fmt.Println(reverse(1534236469))
}
