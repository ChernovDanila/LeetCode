/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    
    p1, p2 := list1, list2
    dummy := &ListNode{}
    result := dummy
    for true {
        if p1 == nil || p2 == nil{
            break 
        }
        
        if comparisonLists(p1, p2){
            result.Next = &ListNode{Val: p1.Val}
            p1 = p1.Next
        }else{
            result.Next = &ListNode{Val: p2.Val}
            p2 = p2.Next
        }
        
        result = result.Next
    }
    if p1 != nil {
        result.Next = p1
    }else{
        result.Next = p2
    }
    return dummy.Next     

}

func comparisonLists(list1 *ListNode, list2 *ListNode) bool{

    if list1.Val < list2.Val {
        return true
    }
    return false

}