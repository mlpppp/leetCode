package LeetCode_Java.List_LinkedList;

//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
    // for each location, check i-centered palindrome length and (i,i+1)-centered palindrome length
    int longestLength = 0;
    int longestLeft = -1;// TODO how to initialize?
    int longestRight = -1;// TODO how to initialize?

    public String longestPalindrome(String s) {
        // init the member states
        longestLength = 0;
        longestLeft = -1;
        longestRight = -1;
        // string to char array
        char[] charArray = s.toCharArray();

        // for each location, check i-centered palindrome length and (i,i+1)-centered palindrome length
        for (int i = 0; i < s.length(); i++) {
            findLongestPalindrome(charArray, i, i);
            findLongestPalindrome(charArray, i, i+1);
        }
        // get the longest palindrome string by slicing the original string
        return s.substring(longestLeft, longestRight+1);


    }

    public void findLongestPalindrome(char[] charArray, int center1, int center2) {
        int left = center1;
        int right = center2;
        while (left <= right && left >= 0
                && right < charArray.length
                && charArray[left] == charArray[right]) {

            int length = right - left + 1;
            if (length > longestLength) {
                longestLength = length;
                longestLeft = left;
                longestRight = right;
            }
            left--;
            right++;
        }
    }
}
//leetcode submit region end(Prohibit modification and deletion)
