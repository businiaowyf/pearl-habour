package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func printInOrder(root *Node) {
	if root != nil {
		printInOrder(root.Left)
		fmt.Printf("%v ", root.Val)
		printInOrder(root.Right)
	}
}

func height(root *Node) int {
	if root == nil {
		return 0
	}
	lh := 1 + height(root.Left)
	rh := 1 + height(root.Right)
	if lh >= rh {
		return lh
	} else {
		return rh
	}
}

func printLevel(root *Node, i int, direction string) {
	if root == nil {
		return
	}

	if i == 0 {
		fmt.Printf("%v ", root.Val)
	} else {
		if direction == "left" {
			printLevel(root.Left, i-1, direction)
			printLevel(root.Right, i-1, direction)
		} else {
			printLevel(root.Right, i-1, direction)
			printLevel(root.Left, i-1, direction)
		}
	}
}

func printLevelOrder(root *Node) {
	h := height(root)

	for i := 0; i < h; i++ {
		var direction string
		if i%2 == 0 {
			direction = "left"
		} else {
			direction = "right"
		}
		printLevel(root, i, direction)
		fmt.Println()
	}
}

func main() {
	root := &Node{Val: 20}

	root.Left = &Node{Val: 10}
	root.Right = &Node{Val: 30}

	root.Left.Left = &Node{Val: 5}
	root.Left.Right = &Node{Val: 15}
	root.Right.Left = &Node{Val: 25}
	root.Right.Right = &Node{Val: 35}

	root.Left.Left.Left = &Node{Val: 2}
	root.Left.Left.Right = &Node{Val: 7}
	root.Left.Right.Left = &Node{Val: 12}
	root.Left.Right.Right = &Node{Val: 17}

	printInOrder(root)
	fmt.Println()
	fmt.Println(height(root))
	printLevelOrder(root)
}
