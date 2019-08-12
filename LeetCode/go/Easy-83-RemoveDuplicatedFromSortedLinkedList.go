/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil{
        return head
    }
    returnPtr := head
    prev := head
    head = head.Next
    for head != nil{
        if head.Val != prev.Val{
            prev = head
        }else{
            prev.Next = head.Next
        }
        head = head.Next
    }
    return returnPtr
}
