/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
We can solve this problem with two approaches.
    - Floyd cycle detection algorithm AKA two pointer[BEST]
    - Hash table approach.
https://leetcode.com/problems/linked-list-cycle
*/

//Floyd's Cycle Finding Algorithm [Two-Pointer Approach]
/*
Floyd's Cycle Finding Algorithm AKA "hare and tortoise" algorithm, this method uses two pointers that traverse the list at different speeds. The slow pointer moves one step at a time, while the fast pointer moves two steps. If there is a cycle, the fast pointer will eventually catch up to the slow pointer.
Time Complexity: O(n)
Space Complexity: O(1)
*/

func hasCycle(head *ListNode) bool{
    
    slow,fast := head,head
    
    for fast != nil && fast.Next != nil{
        slow = slow.Next
        fast = fast.Next.Next

        if slow == fast{
            return true
        }
    }

    return false
}


/*
Hash Table Approach
->Initialization:
   Create an empty set, visited_nodes, to keep track of the nodes that have been visited.
   Traversal and Verification:

-> Traverse the list starting from the head node.
    At each node, check whether it already exists in visited_nodes.
    If it does, return True as a cycle is detected.
    Otherwise, add the node to visited_nodes.

Time Complexity: O(n)
Space Complexity:O(n)
*/
/*
func hasCycle(head *ListNode) bool{
    //Create hashmap
    visitedNode := make(map[*ListNode]bool)
    
    currentNode := head

    for currentNode != nil{
        if visitedNode[currentNode]{
            return true
        }
        visitedNode[currentNode] = true
        currentNode = currentNode.Next
    }
    return false
}
*/
