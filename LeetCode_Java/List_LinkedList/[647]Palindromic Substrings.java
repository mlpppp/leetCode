package LeetCode_Java.List_LinkedList;

//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
    // similar to 5.
    // O(n^2): (There are two kinds of palindrome: odd length with single center, and even length with two center.) Iterate through the string, and for each character in middle(either as a single center or the first character of a dual center), expand it until it is no more a palindrome, while add to the palindrome counter.
    int subStringCounter = 0;
    public int countSubstrings(String s) {
        subStringCounter = 0;
        char[] charArray = s.toCharArray();
        for (int i = 0; i < charArray.length; i++) {
            countSubstringHelper(charArray, i, i);
            if (i + 1 < charArray.length) {
                countSubstringHelper(charArray, i, i+1);
            }
        }
        return subStringCounter;
    }

    public void countSubstringHelper(char[] charArray, int center1, int center2) {
        for (int i = center1, j = center2;
             i >= 0 && j < charArray.length; i--, j++) {
            if (charArray[i] == charArray[j]) {
                subStringCounter++;
            } else {
                break;
            }
        }
    }
}
//leetcode submit region end(Prohibit modification and deletion)
