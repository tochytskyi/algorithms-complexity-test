package bst

import (
	"log"
)

func Run() {
	var t Tree

	t.insert(5)
	t.insert(4)
	t.insert(1)
	t.insert(10)

	printPreOrder(t.root)

	log.Println("Try to find node with value 1")
	foundNode, err := find(t.root, 1)

	if err == nil {
		log.Printf("Node found with value %d ", foundNode.key)
	}

	log.Println("Try to find node with value 9")
	foundNode2, err2 := find(t.root, 9)

	if err2 == nil {
		log.Printf("Node found with value %d ", foundNode2.key)
	}
}

func printPreOrder(n *Node) {
	if n == nil {
		return
	} else {
		log.Printf("%d ", n.key)
		printPreOrder(n.left)
		printPreOrder(n.right)
	}
}
