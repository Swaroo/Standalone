package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	temp := x
	rev := 0
	rem := 0
	for temp !=0 {
		rem = temp % 10
		temp = temp/10
		rev = rev *10 + rem
	}
	if rev != x{
		return false
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(12321))
    fmt.Println(isPalindrome(1005))
    fmt.Println(isPalindrome(-123))
}
