/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    out := new(ListNode)
    out.Val=0
    ref := out
    for {        
        if(l1 == nil){
            out.Next=l2
            return ref.Next
        }else if(l2 == nil){
            out.Next=l1
            return ref.Next
        }else if(l1.Val > l2.Val){
            out.Next = l2
            out = out.Next
            l2 = l2.Next
        }else{
            out.Next = l1
            out = out.Next
            l1 = l1.Next
        }
    }
    
    return ref.Next
}
