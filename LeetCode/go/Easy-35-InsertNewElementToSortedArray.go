func searchInsert(nums []int, target int) int {
    if(len(nums)==0){
        return 0
    }
    low := 0
    high := len(nums)-1
    var mid int
    for low<=high{
        mid = (low+high)/2
        if(nums[mid]<target){
            low = mid+1
        }else if(nums[mid]>target){
            high = mid-1
        }else{
            return mid
        }
    }
    if(low>high){
        return low
    }
    if(nums[low]>target){
        return low
    }else{
        return low+1
    }
    return low
}
