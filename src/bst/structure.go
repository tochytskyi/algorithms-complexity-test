package bst

import (
	"errors"
)

type Tree struct {
	root *Node
}

func (tree *Tree) Insert(value int) {
	if tree.root == nil {
		tree.root = &Node{key: value}
	} else {
		tree.root.insert(value)
	}
}

func (tree *Tree) Remove(key int) {
	remove(tree.root, key)
}

func (tree *Tree) Min() *Node {
	if tree.root == nil {
		return nil
	}

	node := tree.root

	for node.left != nil {
		node = node.left
	}

	return node
}

func (tree *Tree) Max() *Node {
	if tree.root == nil {
		return nil
	}

	node := tree.root

	for node.right != nil {
		node = node.right
	}

	return node
}

type Node struct {
	key   int
	left  *Node
	right *Node
}

func (n *Node) insert(value int) {
	if value <= n.key {
		if n.left == nil {
			n.left = &Node{key: value}
		} else {
			n.left.insert(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{key: value}
		} else {
			n.right.insert(value)
		}
	}
}

func find(node *Node, value int) (*Node, error) {
	var next *Node

	if value < node.key {
		next = node.left
	} else if value > node.key {
		next = node.right
	} else {
		return node, nil
	}

	if next == nil {
		return &Node{}, errors.New("Node not found")
	}

	return find(next, value)
}

func remove(node *Node, key int) *Node {
	if node == nil {
		return nil
	}

	if key < node.key {
		node.left = remove(node.left, key)
		return node
	}

	if key > node.key {
		node.right = remove(node.right, key)
		return node
	}

	// key == node.key
	if node.left == nil && node.right == nil {
		node = nil
		return nil
	}
	if node.left == nil {
		node = node.right
		return node
	}
	if node.right == nil {
		node = node.left
		return node
	}

	leftmostrightside := node.right
	for {
		//find smallest value on the right side
		if leftmostrightside != nil && leftmostrightside.left != nil {
			leftmostrightside = leftmostrightside.left
		} else {
			break
		}
	}

	node.key = leftmostrightside.key
	node.right = remove(node.right, node.key)

	return node
}
