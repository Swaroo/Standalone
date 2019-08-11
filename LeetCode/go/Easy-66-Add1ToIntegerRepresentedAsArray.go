func plusOne(digits []int) []int {
    n := len(digits)-1
    for n>=0{
        val := digits[n]
        if val==9{
            digits[n]=0
            if n ==0{
                a := [1]int{1}
                return append(a[:],digits...)
            }
            n--
        }else{
            digits[n]=val+1
            return digits
        }
    }
    return digits
}
