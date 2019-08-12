//@Todo : Can be improved to space complexity of O(n)
//Just parse through the array normally while also calculating the min till that position and
//Calculate the profit at each position by subtracting with min value
//The max of these profits will be the answer.
func maxProfit(prices []int) int {
    n := len(prices)
    if n == 0 || n == 1{
        return 0
    }
    diff := make([]int,n-1)    
    for i:=1;i<n;i++{
        diff[i-1]=prices[i]-prices[i-1]
    }
    currSum := 0
    bestSum := 0
    for i:=0;i<n-1;i++{
        currSum += diff[i]
        if currSum < 0 {
            currSum = 0
        }else if currSum > bestSum{
            bestSum = currSum
        }
    }
    return bestSum
}
