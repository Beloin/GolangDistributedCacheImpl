// Package trees
package trees

import (
	"beloin.com/distributed-cache/pkg/datastructure"
)

type Color uint8

const (
	RED   = 0
	BLACK = 1
)

type Node[T datastructure.Comparable] struct {
	value T
	color Color

	parent *Node[T]

	left  *Node[T]
	right *Node[T]
}

func (n *Node[T]) Grandparent() *Node[T] {
	if n.parent == nil {
		return nil
	}

	return n.parent.parent
}

func (n *Node[T]) Uncle() *Node[T] {
	grand := n.Grandparent()
	if grand == nil {
		return nil
	}

	if grand.left == n.parent {
		return grand.right
	}

	return grand.left
}

// RedBlackTree Tree implementation
// 1. A node is Black or Red
// 2. Root is always Black
// 3. All nil leafs are considered Black leaf Nodes
// 4. Both branches of a Red Node are Black
// 5. Every path from a node to any descendant leaf has the same amount of black nodes
type RedBlackTree[T datastructure.Comparable] struct {
	root  *Node[T]
	Depth int
}

func NewRedBlackTree[T datastructure.Comparable]() *RedBlackTree[T] {
	return &RedBlackTree[T]{}
}

func (r *RedBlackTree[T]) getNextLeaf(n *Node[T]) *Node[T] {
	var last, curr *Node[T]
	for last, curr = r.root, r.root; curr != nil; {
		last = curr
		if curr.value.Compare(n.value) >= 0 {
			curr = curr.right
			continue
		}

		curr = curr.left
	}

	return last
}

func (r *RedBlackTree[T]) Insert(v T) {
	newNode := &Node[T]{
		value: v,
		color: RED,
	}

	if r.root == nil {
		newNode.color = BLACK
		r.root = newNode
		return
	}

	// TODO: Will leaf be nil?
	// Leaf as parent
	leaf := r.getNextLeaf(newNode)
	newNode.parent = leaf

	isRight := leaf.value.Compare(newNode.value) >= 0
	if isRight {
		leaf.right = newNode
	} else {
		leaf.left = newNode
	}

	if leaf.color == BLACK {
		return
	}

	// Fixing the tree

	// From here we can assure that newNode has Grandparent and an uncle (even if it's nil), since it's father is Red
	if newNode.Uncle().color == RED {
		newNode.repaint()
		return
	}

	if isRight {
		newNode.leftRotate()
		newNode = newNode.left
	} else {
		newNode.rightRotate()
		newNode = newNode.right
	}

	newNode.parent.color = BLACK
	grandparent := newNode.Grandparent()
	grandparent.color = RED

	if newNode == newNode.parent.left && newNode.parent == grandparent.left {
		newNode.parent.rightRotate()
	} else {
		newNode.parent.leftRotate()
	}
}

func (n *Node[T]) repaint() {
	// Paint both parent and uncle as BLACK, paint the Grandparent as RED, and then repaint-it
	// based on the Grandparent status
	for n.Uncle() != nil && n.Uncle().color == RED && n.parent.color == RED {
		n.Uncle().color = BLACK
		n.parent.color = BLACK

		n = n.Grandparent()
		n.color = RED

		if n.parent == nil {
			n.color = BLACK
			break
		}

		if n.parent.color == BLACK {
			break
		}
	}
}

func (n *Node[T]) leftRotate() {
	// On N3:
	//         N0                 N0
	//        /                  /
	//       N1                 N3
	//      /  \     ->        /  \
	//    N2    N3            N1   N4
	//         /  \          /  \
	//        N5   N4       N2   N5

	if n.parent == nil {
		return
	}

	parent := n.parent
	grandparent := parent.parent

	if parent.right != n {
		return
	}

	n.parent = grandparent
	if grandparent != nil {
		if grandparent.left == parent {
			grandparent.left = n
		} else {
			grandparent.right = n
		}
	}

	parent.right = n.left
	n.left.parent = parent
	n.left = parent
	parent.parent = n
}

func (n *Node[T]) rightRotate() {
	// On N2:
	//         N0                 N0
	//        /                  /
	//       N1                 N2
	//      /  \     ->        /  \
	//    N2    N3            N5   N1
	//   /  \                     /  \
	//  N5   N4                  N4   N3

	if n.parent == nil {
		return
	}

	parent := n.parent
	grandparent := parent.parent

	if parent.left != n {
		return
	}

	n.parent = grandparent
	if grandparent != nil {
		if grandparent.left == parent {
			grandparent.left = n
		} else {
			grandparent.right = n
		}
	}

	parent.left = n.right
	n.right.parent = parent
	n.right = parent
	parent.parent = n
}
