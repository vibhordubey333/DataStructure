/*
Approach

  1.  Traverse the string character by character.
  2.   To keep track of:
        -Open parentheses that still need to be matched.
        -Unmatched closing parentheses.
  3. For every open parenthesis '(', increase the count of open parentheses.
  4. For every closing parenthesis ')', check if there's an unmatched open parenthesis available to pair it with:

        - If yes, decrease the open parenthesis count.
        - If not, increase the unmatched closing parenthesis count.
  5. At the end of the traversal, the result is the sum of unmatched open and unmatched closing parentheses, which indicates how many parentheses are needed to make the string valid.

Complexity:
	Time: O(N)
	Space: O(1)

Link: https://leetcode.com/problems/minimum-add-to-make-parentheses-valid/
*/
package main

func minAddToMakeValid(s string) int {
	open := 0
	closingCount := 0
	for _, c := range s {
		if c == '(' {
			open++
		} else if c == ')' {
			if open > 0 {
				open--
			} else {
				closingCount++
			}
		}
	}
	return open + closingCount
}

/*
func minAddToMakeValid(s string) int {
	open, close := 0, 0
	for _, c := range s {
		if c == '(' {
			open++
		} else if c == ')' && open > 0 {
			open--
		} else {
			close++
		}
	}
	return open + close
}
*/
