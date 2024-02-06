// https://www.youtube.com/watch?v=6kTZYvNNyps

// Given a sorted dictionary (array of words of lower case alphabet) of an alien language, find the order of characters in the language. If there is no solution, return "". If there are multiple possible solutions, return any of them.

// Example:
// Input: words[] = {"baa", "abcd", "abca", "cab", "cad"}
// Output: Order of characters is ‘bdac’
// Explanation: Note that words are sorted and in the given language “baa” comes before “abcd”, therefore ‘b’ is before ‘a’ in output. Similarly we can find other orders.

// Input: words[] = {"caa", "aaa", "aab"}
// Output: Order of characters is ‘cab’

// rule of lexicographically sorting:
// 1. Strings are sorted based on: return firstDifferentLetter[s] < firstDifferentLetter[t]
// 2. if a word 's' is a prefix of a word 't', then s < t

// non-feasible cases:
// 1. contradict the first rule (eg, app < ape & bpp > bpe)
// 2. contradict the second rule (eg, dag > dager)

// Solution:
// build a graph by comparing adjacent words pairs in `words`, from left to right, according to the rules (eg, if we get `e < a`, create an edge e->a). Then find the topological sorted order of the graph: postorder DFS. Return "" if the graph contains a cycle. (ref 210.course-schedule-ii)

// implementation:
// adjacency list: use a hashmap map[char]map[char], to eliminate possible duplicates
// visited, trace := map[byte]bool

func alienOrder(words []string) string {
	// initial graph as adjacency list, create nodes from distinct letters in words
	graph := make(map[byte]map[byte]bool)
	for _, word := range words {
		for i, _ := range word {
			if _, exist := graph[word[i]]; !exist {
				graph[word[i]] = make(map[byte]bool)
			}
		}
	}

	// initial edges, compare neighbor pairs of words
	for i := 0; i < len(words)-1; i++ {
		// compare the first different letter of words[i] and words[i+1].
		l, r := 0, 0
		for l < len(words[i]) && r < len(words[i+1]) {
			// find first different letter, create edge from letter(l) -> letter(r)
			if words[i][l] != words[i+1][r] {
				graph[words[i][l]][words[i+1][r]] = true
				break
			}
			l++
			r++
		}

		// if words[i+1] is prefix, this is contradictory to the rule
		if l < len(words[i]) && r == len(words[i+1]) {
			return ""
		}
	}

	// DFS
	visited := make(map[byte]bool)
	trace := make(map[byte]bool)
	postorder := []byte{}
	for node, _ := range graph {
		if detectCycleAndOrder(graph, node, visited, trace, &postorder) {
			return ""
		}
	}
	slices.Reverse(postorder)
	return string(postorder)
}

func detectCycleAndOrder(graph map[byte]map[byte]bool, node byte, visited, trace map[byte]bool, postorder *[]byte) bool {
	// contain cycle
	if val, exist := trace[node]; exist && val {
		return true
	}
	// node is visited without cycle
	if val, exist := visited[node]; exist && val {
		return false
	}
	// preorder
	trace[node] = true
	visited[node] = true
	for adjNode, _ := range graph[node] {
		if detectCycleAndOrder(graph, adjNode, visited, trace, postorder) {
			return true
		}
	}

	// postorder
	delete(trace, node)
	*postorder = append(*postorder, node)
	return false
}

func main() {
	fmt.Println(alienOrder([]string{"baa", "abcd", "abca", "cab", "cad"})) // bdac
	fmt.Println(alienOrder([]string{"caa", "aaa", "aab"}))                 // cab
	fmt.Println(alienOrder([]string{"kaa", "akcd", "akca", "cak", "cad"})) // kdac
	fmt.Println(alienOrder([]string{"b", "a"}))                            // ba
	fmt.Println(alienOrder([]string{"ab", "a", "b"}))                      // ""
}
