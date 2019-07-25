func reverse(x int) int {
    rev := 0
    neg := false
    if x < 0{
        neg = true
        x = x * -1
    }
    for x!=0{
        rem:= x%10;
        x = x/10
        if rem != 0 && rev == 0{
            rev = rev *10 + rem
        }
    }
    if neg {
        rev = rev * -1
    }
    return rev
}
