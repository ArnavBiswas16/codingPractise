package main

import "fmt"

type Node struct {
	item  int
	left  *Node
	right *Node
}

func main() {

	in := []int{2, 1, 3, 0, 3, 1, 2}
	pre := []int{0, 1, 2, 3, 1, 3, 2}
	root := consInPre(in, pre)
	// var post []int
	// post = postOrder(root, post)
	// fmt.Println(post)

	// root1 := consInPost(in, post)
	// fmt.Println(root1.right.right)

	// fmt.Println(checkSymmetry(root))
	flattenBtree(root)
	printTree(root)
}

func consInPre(in []int, pre []int) *Node {

	if len(in) == 0 {
		return nil
	}

	root := &Node{
		item: pre[0],
	}

	if len(in) == 1 {
		return root
	}

	index := indexOf(root.item, in)

	root.left = consInPre(in[:index], pre[1:index+1])
	root.right = consInPre(in[index+1:], pre[index+1:])

	return root
}

func indexOf(item int, arr []int) int {

	for i, val := range arr {
		if val == item {
			return i
		}
	}
	return -1
}

func postOrder(root *Node, post []int) []int {
	if root == nil {
		return post
	}

	post = postOrder(root.left, post)
	post = postOrder(root.right, post)
	post = append(post, root.item)
	// fmt.Println(post)
	return post
}

func consInPost(in []int, post []int) *Node {

	if in == nil {
		return nil
	}

	root := &Node{
		item: post[len(post)-1],
	}

	if len(in) == 1 {
		return root
	}

	index := indexOf(root.item, in)

	root.right = consInPost(in[index+1:], post[index:len(post)-1])
	root.left = consInPost(in[:index], post[:index])
	return root
}

func checkSymmetry(root *Node) bool {
	if root == nil {
		return true
	}
	return checkSymmetryHelper(root.left, root.right)
}

func checkSymmetryHelper(left *Node, right *Node) bool {

	if left == nil && right == nil {
		return true
	}
	if left != nil && right != nil {

		if left.item != right.item {
			return false
		} else {
			return checkSymmetryHelper(left.left, right.right) && checkSymmetryHelper(left.right, right.left)
		}

	} else {
		return false
	}

}

func flattenBtree(root *Node) {

	if root == nil {
		return
	}

	flattenBtree(root.left)
	flattenBtree(root.right)

	if root.left != nil {

		left := root.left
		root.left = nil

		if root.right != nil {
			right := root.right
			root.right = left

			var rightMost = root.right
			for rightMost.right != nil {
				rightMost = rightMost.right
			}
			rightMost.right = right
		} else {
			root.right = left
		}

	}
	return
}

func printTree(root *Node) {
	if root == nil {
		return
	}
	fmt.Print(root.item, " ")
	printTree(root.right)
}
