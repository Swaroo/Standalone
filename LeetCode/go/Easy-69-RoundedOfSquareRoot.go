//@Todo: To be improved by using modified binary search or bit manipulation
func mySqrt(x int) int {
    for i:=1;i<=x;i++{
        if i*i>x{
            return i-1
        }
    }
    if x==0||x==1 {
        return x
    }
    return 0
}
