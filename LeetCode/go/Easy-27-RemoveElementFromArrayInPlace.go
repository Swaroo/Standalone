func replace(nums []int, i,j int){
    temp := nums[i]
    nums[i] = nums[j]
    nums[j] = temp
}
func findLast(nums []int, i,val int) int{
    println(i,val)
    for i>=0 {
        if (nums[i]==val){
            i = i-1
        }else{
            return i
        }
    }
    return 0
}
func removeElement(nums []int, val int) int {
    found := false
    last := findLast(nums,len(nums)-1,val)
    for i:=0;i < last;i++{
        if(nums[i]==val){
            found = true
            replace(nums,i,last)
            last = findLast(nums,last-1,val)
        }
    }
    if last == 0 && found{
        return 0
    }else{
    return last+1
    }
}
