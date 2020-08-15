package main

import (
	"fmt"
	"math"
)

type Node struct {
	item  int
	left  *Node
	right *Node
}

func main() {

	var root *Node
	root = insert(root, 2)
	root = insert(root, 10)
	root = insert(root, 30)
	root = insert(root, 400)
	root = insert(root, 1)
	root = insert(root, 13)
	root = insert(root, 32)
	root = insert(root, 100)
	fmt.Println(root.item)
	fmt.Println(Search(root, 2))
	fmt.Println(findMin(root))
	fmt.Println(findMax(root))
	fmt.Println(findHeight(root))
	fmt.Println("preorder")
	preorder(root)
	fmt.Println("\ninorder")
	inorder(root)
	fmt.Println("\npostorder")
	postorder(root)

}

func insert(root *Node, item int) *Node {

	if root == nil {
		root := createNewNode(root, item)
		return root
	} else if item <= root.item {
		root.left = insert(root.left, item)
	} else {
		root.right = insert(root.right, item)
	}
	return root
}

func createNewNode(root *Node, item int) *Node {

	node := &Node{
		item: item,
	}
	return node
}

func Search(root *Node, data int) bool {
	if root == nil {
		return false
	} else if root.item == data {
		return true
	} else if data <= root.item {
		return Search(root.left, data)
	} else {
		return Search(root.right, data)
	}
}

func findMin(root *Node) int {
	if root == nil {
		return -1
	}

	if root.left != nil {
		return findMin(root.left)
	} else {
		return root.item
	}
}

func findMax(root *Node) int {

	if root == nil {
		return -1
	}

	if root.right != nil {
		return findMax(root.right)
	} else {
		return root.item
	}
}

func findHeight(root *Node) int {
	if root == nil {
		return -1
	}
	if root.left == nil && root.right == nil {
		return 0
	}

	return int(math.Max(float64(findHeight(root.left)), float64(findHeight(root.right))) + 1)

}

func preorder(root *Node) {

	if root == nil {
		return
	}
	fmt.Print(root.item, ", ")

	if root.left != nil {
		preorder(root.left)
	}
	if root.right != nil {
		preorder(root.right)
	}
}

func inorder(root *Node) {
	if root == nil {
		return
	}

	inorder(root.left)

	fmt.Print(root.item, ", ")

	inorder(root.right)

}

func postorder(root *Node) {
	if root == nil {
		return
	}

	inorder(root.left)
	inorder(root.right)
	fmt.Print(root.item, ", ")

}