package bst

import (
	"errors"
)

type Tree struct {
	root *Node
}

func (t *Tree) insert(value int) {
	if t.root == nil {
		t.root = &Node{key: value}
	} else {
		t.root.insert(value)
	}
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
	if value == node.key {
		return node, nil
	}

	var next *Node

	if value <= node.key {
		next = node.left
	} else {
		next = node.right
	}

	if next == nil {
		return &Node{}, errors.New("Node not found")
	}

	return find(next, value)
}
