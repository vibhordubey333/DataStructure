/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 /*
1. First  - choose the head
2. Second - choose the node with less Val, append it to the new list being built and move the corresponding pointer one step further, e.g. list1 = list1.Next
3. Third - in the end of the cycle you have reached the end of one of the given lists, so what is left is to append the remaining tail to the resulting merged list

Time Complexity: O(m+n)
Space Complexity: O(1)
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := new(ListNode)//&ListNode{} 
    current := dummy
    
    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            current.Next = list1
            list1 = list1.Next
        } else {
            current.Next = list2
            list2 = list2.Next
        }
        current = current.Next
    }
    
    // Append the remaining part of the non-empty list
    if list1 != nil {
        current.Next = list1
    } else {
        current.Next = list2
    }
    
    return dummy.Next
}
