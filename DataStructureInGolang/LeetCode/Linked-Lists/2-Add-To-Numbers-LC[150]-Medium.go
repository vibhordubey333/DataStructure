//Algorithm
/*
 1. Traverse two linked lists.
 2. In each iteration, add the numbers in the nodes of the linked lists.
 3. If the lists are unequal then the smaller one will end before the longer. In this case, we will put only the remaining nodes of the longer list in the resultant list.
 4. If the sum of two digits is greater than 9, then we will have to find out the "carry" to be added in the next iteration.
*/
/*
This is nothing but a simple elementary addition problem. The only difference is that the numbers to be added are represented by linked list where each digit is represented by the nodes of that linked list.

If we see the example then we will see that the digits are in the reverse order i.e.,

First node => ones place
Second node => tens place
Third node => hundreds place
... and so on.
Thus 2 -> 4 -> 3 will actually make 342 and 5 -> 6 -> 4 will actually make 465.

We will have to return a new linked list whose nodes will represent the digits of the sum of the numbers represented by the given two linked list.
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 //Link: https://leetcode.com/problems/add-two-numbers/
 /*
  Time Complexity: O(m+n)
  Space Complexity: O(1)
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var listNodeHead *ListNode
	var listNodeTemp *ListNode

	carry := 0

	for l1 != nil || l2 != nil {
		sum := carry
		//Performing sum operation begins.
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		//Performing sum operation ends.

		//Assigning remainder to new node.
		nodeObject := &ListNode{Val: sum % 10}
		//Computing if any quotient is there.
		carry = sum / 10
		
		//Validating if this is the first iteration then create a new linkedlist. Else continue with created one.
		//Assining nodeObject value to temp node in both the cases
		if listNodeTemp == nil {
			listNodeTemp = nodeObject
			listNodeHead = nodeObject
		} else {
			listNodeTemp.Next = nodeObject
			listNodeTemp = listNodeTemp.Next
		}
		
		//Validating if any carry value is therr. If gt 0 then assign to next node of temp node.
		if carry > 0 {
			listNodeTemp.Next = &ListNode{Val: carry}
		}

	}
	return listNodeHead
}

