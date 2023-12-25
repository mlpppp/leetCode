/*
 * @lc app=leetcode id=297 lang=golang
 *
 * [297] Serialize and Deserialize Binary Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Encodes a tree to a single string. Implement serialize and deserialize algorithms.

// Preorder Serialization (with null), DFS, append node values or null_string to string builder in preorder place. Deserizalization: The first token in string is the root. DFS and process root node token in preorder place. Then removing the first token before recursive call.
// Serialization: O(n)
// Deserialization: O(n)
import (
	"strconv"
	"strings"
	"log"
)

var SEP byte = ','
var NULL byte = '#'

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var builder strings.Builder
	this.serializeHelper(root, &builder)
	originalStr := builder.String()
	// fmt.Println(originalStr[:len(originalStr)-1])
	// fmt.Println("serialzation END <<<<<<")
	return originalStr[:len(originalStr)-1] // remove the last comma
}

// preorder DFS and build the string
func (this *Codec) serializeHelper(root *TreeNode, builder *strings.Builder) {

	if root == nil {
		builder.WriteByte(NULL)
		builder.WriteByte(SEP)
		return
	}

	// append current node to builder
	builder.WriteString(strconv.Itoa(root.Val))
	builder.WriteByte(SEP)

	this.serializeHelper(root.Left, builder)
	this.serializeHelper(root.Right, builder)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	// turn a string separated by "," to a slice of characters (excluding the ",")
	dataSlice := strings.Split(data, string(SEP))
	// fmt.Println(dataSlice)
	return deserializeHelper(&dataSlice)
}

func deserializeHelper(dataSlice *[]string) *TreeNode {
	if len(*dataSlice) == 0 {
		return nil
	}
	// fmt.Println("current node: ", (*dataSlice)[0])
	if (*dataSlice)[0] == string(NULL) {
		*dataSlice = (*dataSlice)[1:] // remove first element after finishing a node
		// fmt.Println(*dataSlice)
		// fmt.Println("node process END <<<<")
		return nil
	}
	val, err := strconv.Atoi((*dataSlice)[0])
	if err != nil {
		log.Fatal("Error: %v", err)
	}
	root := &TreeNode{Val: val}
	*dataSlice = (*dataSlice)[1:] // remove first element after finishing a node
	// fmt.Println(*dataSlice)
	// fmt.Println("node process END <<<<")

	root.Left = deserializeHelper(dataSlice)
	root.Right = deserializeHelper(dataSlice)

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
// @lc code=end

