func replace(nums []int, i,j int){
    fmt.Println(i,j)
    temp := nums[i]
    nums[i] = nums[j]
    nums[j] = temp
}
func removeDuplicates(nums []int) int {
    n := len(nums)
    currpos := 0
    currval := nums[0]
    for i:=1 ; i< n; i++{
        if(nums[i]!=currval){
            replace(nums,i,currpos)
            if(i!=n-1){
            currpos = currpos + 1
            }
            currval = nums[currpos]
        }
    }    
    fmt.Println(nums)
    return currpos
}
