package main

import (
	"fmt"
	"math"
	"sort"
)

type Node struct {
	item  int
	left  *Node
	right *Node
	next  *Node
}

//
func main() {

	var root *Node
	// var root1 *Node
	root = insert(root, 41)
	root = insert(root, 61)
	root = insert(root, 2)
	root = insert(root, 10)
	root = insert(root, 30)
	root = insert(root, 50)

	// root1 = insert(root1, 41)
	// root1 = insert(root1, 61)
	// root1 = insert(root1, 2)
	// root1 = insert(root1, 10)
	// root1 = insert(root1, 30)
	// root1 = insert(root1, 50)

	root = insert(root, 400)
	root = insert(root, 1)
	root = insert(root, 13)
	root = insert(root, 32)
	root = insert(root, 100)
	root = insert(root, 11)
	root = insert(root, 250)
	root = insert(root, 90)

	// fmt.Println(root.item)
	// fmt.Println(Search(root, 2))
	// fmt.Println(findMin(root))
	// fmt.Println(findMax(root))
	// fmt.Println(findHeight(root))
	// fmt.Println("preorder")
	// preorder(root)
	// fmt.Println("\ninorder")
	// inorder(root)
	// fmt.Println("")
	// inorder(root1)
	// fmt.Println("\npostorder")
	// postorder(root)
	// fmt.Println("\nlevelorder")
	// levelorder(root)
	// printLevel(root)
	// leftView(root)
	// rightView(root)
	// spiralLevelorder(root)

	// fmt.Println("\n", checkBST(root))

	// root = delete(root, 30)
	// fmt.Println("\npreorder")
	// preorder(root)
	// fmt.Print("\n", inorderSucc(root, 13).item)

	// createTopView(root)
	// createBotView(root)

	// fmt.Println("diameter: ", diameter(root))

	// fmt.Println(checkHeightBalanced(root))

	// fmt.Println(alc(root, 250, 100))

	// fmt.Println(checkIdentical(root, root1))

	nextRightPointersOfTree(root)
	fmt.Println(root.left.right.next.item)

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

	return int(math.Max(
		float64(findHeight(root.left)),
		float64(findHeight(root.right)))) + 1

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

	postorder(root.left)
	postorder(root.right)
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

	for len(q) != 0 {

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

func leftView(root *Node) {

	if root == nil {
		return
	}
	var q []*Node

	printStat := true
	if len(q) == 0 {
		q = enqueue(q, root)
		q = enqueue(q, nil)
	}

	for len(q) != 0 {
		temp := q[0]
		q = dequeue(q)

		if temp != nil {
			if printStat {
				fmt.Print(temp.item, ", ")
				printStat = false
			}

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
			q = enqueue(q, nil)
			printStat = true
		}
	}
}

func rightView(root *Node) {

	if root == nil {
		return
	}
	var q []*Node

	printStat := true
	if len(q) == 0 {
		q = enqueue(q, root)
		q = enqueue(q, nil)
	}

	for len(q) != 0 {
		temp := q[0]
		q = dequeue(q)

		if temp != nil {
			if printStat {
				fmt.Print(temp.item, ", ")
				printStat = false
			}

			if temp.right != nil {
				q = enqueue(q, temp.right)
			}

			if temp.left != nil {
				q = enqueue(q, temp.left)
			}

		} else {
			if len(q) == 0 {
				break
			}
			fmt.Println(" ") // change level
			q = enqueue(q, nil)
			printStat = true
		}
	}
}

func createTopView(root *Node) {

	createTopBottomView(root, 0, 0)

}

func createBotView(root *Node) {

	createTopBottomView(root, 0, 1)
}

func createTopBottomView(root *Node, hd int, flag int) {

	if root == nil {
		return
	}
	type o struct {
		node *Node
		hd   int
	}
	m := make(map[int][]*Node)

	var q []o

	q = append(q, o{root, hd})

	if val, ok := m[hd]; ok {
		m[hd] = append(m[hd], root)
	} else {
		m[hd] = append(val, root)
	}

	for len(q) > 0 {
		hd = q[0].hd

		curr := q[0].node
		if curr.left != nil {

			if val, ok := m[hd-1]; ok {
				m[hd-1] = append(m[hd-1], curr.left)
			} else {
				m[hd-1] = append(val, curr.left)
			}
			q = append(q, o{curr.left, hd - 1})

		}
		if curr.right != nil {

			if val, ok := m[hd+1]; ok {
				m[hd+1] = append(m[hd+1], curr.right)
			} else {
				m[hd+1] = append(val, curr.right)
			}
			q = append(q, o{curr.right, hd + 1})

		}
		q = q[1:] // dequeue
	}

	var keys []int
	for i, _ := range m {
		keys = append(keys, i)
		// fmt.Println(i, ": ", val[0].item, " ")
	}

	sort.Ints(keys)

	if flag == 0 {
		for _, k := range keys {
			fmt.Print(m[k][0].item, " ")
		}
	}

	if flag == 1 {
		for _, k := range keys {
			fmt.Print(m[k][len(m[k])-1].item, " ")
		}
	}

}

func diameter(root *Node) int {

	if root == nil {
		return 0
	}
	// fmt.Println(findHeight(root.left))
	// return 0
	// fmt.Println(float64(findHeight(root.left) + findHeight(root.right) + 3))
	return int(math.Max(
		float64(findHeight(root.left)+findHeight(root.right)+3),
		math.Max(
			float64(diameter(root.left)),
			float64(diameter(root.right)))))
}

func checkHeightBalanced(root *Node) bool {

	if root.left == nil && root.right == nil {
		return true
	}

	if root.left != nil && root.right != nil {

		left := checkHeightBalanced(root.left)
		right := checkHeightBalanced(root.right)

		if !left || !right {
			return false
		}

		diff := findHeight(root.left) - findHeight(root.right)

		if diff < 0 {
			diff = diff * -1
		}

		if diff > 1 {
			return false
		} else {
			return true
		}

	} else {
		return false
	}

}

func alc(root *Node, a int, b int) *Node {

	if root == nil {
		return nil
	}
	if root.item == a || root.item == b {
		return root
	}
	var ls *Node
	var rs *Node
	if root.left != nil {
		ls = alc(root.left, a, b)
	}
	if root.right != nil {
		rs = alc(root.right, a, b)

	}

	if ls == nil {
		return rs
	}
	if rs == nil {
		return ls
	}
	return root

}

func checkIdentical(root1 *Node, root2 *Node) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil && root2 != nil {
		return false
	}
	if root2 == nil && root1 != nil {
		return false
	}
	if (root1.item == root2.item) && checkIdentical(root1.left, root2.left) && checkIdentical(root1.right, root2.right) {
		return true
	}
	return false

}

func nextRightPointersOfTree(root *Node) {
	if root == nil {
		return
	}

	var q []*Node

	if len(q) == 0 {
		q = enqueue(q, root)
		q = enqueue(q, nil)

	}

	var tempPrev *Node
	for len(q) != 0 {

		temp := q[0]
		q = dequeue(q)

		if tempPrev != nil && temp != nil {
			tempPrev.next = temp
		}
		tempPrev = temp

		if temp != nil {

			if temp.left != nil {
				q = enqueue(q, temp.left)
			}

			if temp.right != nil {
				q = enqueue(q, temp.right)
			}

		} else {
			if len(q) == 1 {
				break
			} else {
				q = enqueue(q, nil)
			}
		}
	}
}
