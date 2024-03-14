//https://leetcode.com/problems/reverse-vowels-of-a-string/


/*
The approach used here is the Two-Pointer Approach:

- Initialize two pointers, low and high, pointing to the beginning and end of the string, respectively.
- Initialize a string vowels containing all vowel characters.
- Convert the input string s to a runeSlice to work with individual runes (characters).
- Iterate using the two-pointer technique:
- Move the low pointer to the right until it points to a vowel.
- Move the high pointer to the left until it points to a vowel.
- If low < high, swap the vowels at the positions pointed by low and high.
*/
func reverseVowels(s string) string {
    low, high := 0, len(s) -1
    vowels := "aeiouAEIOU"

    runeSlice := []rune(s) // Convert strin to rune[int32]

    for low < high{
        // Move the low pointer till vowel is encountered.
        for low < high && !strings.Contains(vowels,string(runeSlice[low])){
            low++
        }

        // Move the high pointer to the right till vowel is encountered.
        for low < high && !strings.Contains(vowels,string(runeSlice[high])){
            high--
        }

        if low < high {
            //Swap values.
            runeSlice[low],runeSlice[high] = runeSlice[high],runeSlice[low]
            low++
            high--
        }
    }
    return string(runeSlice)
}
