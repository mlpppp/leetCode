/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start

// Given an array of strings strs, group the anagrams together. You can return the answer in any order. An Anagram is a word rearranged with the same characters from another word. strs[i] consists of lowercase English letters.

// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

// (hashmap and encoding anagrams): iterate strs and encode each string, then save the string into hashmap [encode]:{strings...}. Then iterate the hashmap and build the answer:

// encode anagram, 0 <= strs[i].length <= 100, means each char occurs at most 100 times. We want a hashable encoding that has 26 positions, each of possible value from (0, 100). We can do it with [26]byte, since byte is int8 and is hashable when turned into string.

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)
	for _, str := range strs {
		encode := encodeWord(str)
		anagramMap[encode] = append(anagramMap[encode], str)
	}
	ans := [][]string{}

	for _, value := range anagramMap {
		ans = append(ans, value)
	}
	return ans
}

func encodeWord(word string) string {
	encode := make([]byte, 26)
	for _, char := range word {
		loc := char - 'a'
		encode[loc] += 1
	}
	return string(encode)
}

// @lc code=end

