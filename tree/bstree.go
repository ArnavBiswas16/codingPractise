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
	root = insert(root, 41)
	root = insert(root, 61)
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
	fmt.Println("\nlevelorder")
	levelorder(root)
	printLevel(root)
	fmt.Println("\n", checkBST(root))

	root = delete(root, 30)
	fmt.Println("\npreorder")
	preorder(root)
	fmt.Print("\n", inorderSucc(root, 13).item)
	spiralLevelorder(root)
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

func Search(root *Node, data int) *Node {
	if root == nil {
		return nil
	} else if root.item == data {
		return root
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

func FindMinNode(root *Node) *Node {
	if root == nil {
		return nil
	}

	if root.left != nil {
		return FindMinNode(root.left)
	} else {
		return root
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

func levelorder(root *Node) {
	if root == nil {
		return
	}
	var q []*Node

	if len(q) == 0 {
		q = enqueue(q, root)
	}

	for len(q) != 0 {

		temp := q[0]
		q = dequeue(q)

		fmt.Print(temp.item, ", ")

		if temp.left != nil {
			q = enqueue(q, temp.left)
		}

		if temp.right != nil {
			q = enqueue(q, temp.right)
		}

	}
}

func printLevel(root *Node) {

	if root == nil {
		return
	}
	var q []*Node

	if len(q) == 0 {
		q = enqueue(q, root)
		q = enqueue(q, nil)
	}

	i := 0
	for len(q) != 0 {
		i++
		temp := q[0]
		q = dequeue(q)

		if temp != nil {
			fmt.Print(temp.item, ", ")

			if temp.left != nil {
				q = enqueue(q, temp.left)
			}
			if temp.right != nil {
				q = enqueue(q, temp.right)
			}

		} else {
			if len(q) == 0 {
				break
			}
			fmt.Println(" ") // change level
			// reurn
			q = enqueue(q, nil)
		}
	}
}

func enqueue(queue []*Node, node *Node) []*Node {
	queue = append(queue, node) // Simply append to enqueue.
	return queue
}

func dequeue(queue []*Node) []*Node {

	if len(queue) != 0 {
		return queue[1:] // Slice off the element once it is dequeued.
	}
	return queue
}

func checkBST(root *Node) bool {
	if root == nil {
		return false
	}
	var q []*Node

	if len(q) == 0 {
		q = enqueue(q, root)
	}

	for len(q) != 0 {
		curr := q[0]
		if curr.left != nil && curr.left.item >= curr.item {
			return false
		}
		if curr.right != nil && curr.right.item < curr.item {
			return false

		}
		if curr.left != nil {
			q = enqueue(q, curr.left)

		}
		if curr.right != nil {
			q = enqueue(q, curr.right)

		}
		q = dequeue(q)
	}
	return true
}

func delete(root *Node, item int) *Node {
	// fmt.Println(item)
	if root == nil {
		return root
	} else if item < root.item {
		root.left = delete(root.left, item)
	} else if item > root.item {
		root.right = delete(root.right, item)
	} else if item == root.item {
		//case 1  if leaf node
		if root.left == nil && root.right == nil {
			root = nil
		} else if root.left == nil { // 1 child
			root = root.right
		} else if root.right == nil { // 1 child
			root = root.left
		} else { // both child
			temp := FindMinNode(root.right)
			root.item = temp.item
			root.right = delete(root.right, root.item)
		}
	}

	return root
}

func inorderSucc(root *Node, item int) *Node {

	if root == nil {
		return root
	}
	current := Search(root, item)
	if current == nil {
		return nil
	} else {
		if current.right != nil {
			return FindMinNode(current.right)
		} else {

			var succ *Node
			ancestor := root
			for ancestor != current {
				if current.item < ancestor.item {
					succ = ancestor
					ancestor = ancestor.left
				} else {
					ancestor = ancestor.right
				}
			}
			return succ
		}
	}
}

func spiralLevelorder(root *Node) {

	var s1 []*Node
	var s2 []*Node

	s1 = push(s1, root)

	for len(s1) > 0 {

		for len(s1) > 0 {
			var s1Item *Node

			s1, s1Item = pop(s1)

			fmt.Print(s1Item.item, ", ")

			if s1Item.left != nil {
				s2 = push(s2, s1Item.left)
			}
			if s1Item.right != nil {
				s2 = push(s2, s1Item.right)
			}

		}

		for len(s2) > 0 {

			var s2Item *Node

			s2, s2Item = pop(s2)

			fmt.Print(s2Item.item, ", ")

			if s2Item.right != nil {
				s1 = push(s1, s2Item.right)
			}
			if s2Item.left != nil {
				s1 = push(s1, s2Item.left)
			}
		}
	}

}

func push(s []*Node, item *Node) []*Node {

	s = append(s, item)
	return s
}

func pop(s []*Node) ([]*Node, *Node) {

	p := s[len(s)-1]
	s = s[:len(s)-1]
	return s, p
}
