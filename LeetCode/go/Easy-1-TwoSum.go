package main

import "fmt"

func twoSum(nums []int, target int) []int {
    var out [2]int
    var m map[int]int
    m = make(map[int]int)
    for i:=0 ; i<len(nums);i++{
        left := target - nums[i]
        if m[left]!=0{
            out[0] = m[left]-1
            out[1] = i
        }else{
            m[nums[i]] = i+1
        }
    }
    return out[:]
}

func main(){
    nums := []int{1,2,3,4}
    target := 7
    fmt.Println(twoSum(nums,target))
}
