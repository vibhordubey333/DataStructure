/*
Link:https://leetcode.com/problems/longest-substring-without-repeating-characters/description
Given a string s, find the length of the longest
substring
without repeating characters.



Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.



Constraints:

    0 <= s.length <= 5 * 104
    s consists of English letters, digits, symbols and spaces.


*/
package main

func lengthOfLongestSubstring(s string) int {
	// Initialize a map to keep track of the characters in the current window.
	charMap := make(map[byte]int)

	// Initialize variables to keep track of the longest substring length and the current window.
	longest := 0
	start := 0

	// Iterate through the string.
	for i, char := range s {
		// Convert the rune to a byte before using it as a key in the map.
		byteChar := byte(char)

		// If the character is already in the window, update the start index to the last occurrence of the character.
		if lastOccurrence, ok := charMap[byteChar]; ok {
			start = max(start, lastOccurrence+1)
		}

		// Update the length of the longest substring if necessary.
		longest = max(longest, i-start+1)

		// Update the character map with the current character.
		charMap[byteChar] = i
	}

	// Return the length of the longest substring.
	return longest
}

// Helper function to return the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
