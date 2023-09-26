//Algorithm
/*
 1. Traverse two linked lists.
 2. In each iteration, add the numbers in the nodes of the linked lists.
 3. If the lists are unequal then the smaller one will end before the longer. In this case, we will put only the remaining nodes of the longer list in the resultant list.
 4. If the sum of two digits is greater than 9, then we will have to find out the "carry" to be added in the next iteration.
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 //Link: https://leetcode.com/problems/add-two-numbers/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var listNodeHead *ListNode
	var listNodeTemp *ListNode

	carry := 0

	for l1 != nil || l2 != nil {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		nodeObject := &ListNode{Val: sum % 10}

		carry = sum / 10

		if listNodeTemp == nil {
			listNodeTemp = nodeObject
			listNodeHead = nodeObject
		} else {
			listNodeTemp.Next = nodeObject
			listNodeTemp = listNodeTemp.Next
		}

		if carry > 0 {
			listNodeTemp.Next = &ListNode{Val: carry}
		}

	}
	return listNodeHead
}

