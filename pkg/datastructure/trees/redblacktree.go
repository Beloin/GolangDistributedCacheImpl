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
		if n.value.Compare(curr.value) >= 0 {
			curr = curr.right
			continue
		}

		curr = curr.left
	}

	return last
}

func (r*RedBlackTree[T]) Read(outSlice []*T)  {
	
}

// TODO: Follow the book!
func (r *RedBlackTree[T]) Insert(v T) {
	newNode := &Node[T]{
		value: v,
		color: RED,
	}

	if r.root == nil {
		newNode.color = BLACK // Fix violation of ruile 2.
		r.root = newNode
		return
	}

	parent := r.getNextLeaf(newNode)
	newNode.parent = parent

	isRight := newNode.value.Compare(parent.value) >= 0
	if isRight {
		parent.right = newNode
	} else {
		parent.left = newNode
	}

	if parent.color == BLACK {
		return
	}

	// From here we can assure that newNode has Grandparent and an uncle (even if it's nil), since it's father is Red

	// Fixing the tree
	r.fix(newNode)
}

func (tree *RedBlackTree[T]) fix(n *Node[T]) {
	// loop invariant: `n` is always RED
	for n.parent != nil || n.parent.color == RED { // Violation of rule 4.
		if n.parent == n.Grandparent().left {
			uc := n.Grandparent().right

			if uc != nil && uc.color == RED {
				n.parent.color = BLACK                   // |  Case 01
				uc.color = BLACK                         // |
				n.Grandparent().color = RED
				n = n.Grandparent()
			} else {
				if n == n.parent.right {
					n = n.parent                           // | Case 02
					tree.leftRotate(n)                     // |
				}

				n.parent.color = BLACK                   // | Case 03
				n.Grandparent().color = RED              // |
				tree.rightRotate(n.Grandparent())
			}

		} else {
			uc := n.Grandparent().left

			if uc != nil && uc.color == RED {
				n.parent.color = BLACK
				uc.color = BLACK
				n.Grandparent().color = RED
				n = n.Grandparent()
			} else {
				if n == n.parent.left {
					n = n.parent
					tree.rightRotate(n)
				}

				n.parent.color = BLACK
				n.Grandparent().color = RED
				tree.leftRotate(n.Grandparent())
			}
		}
	}

	tree.root.color = BLACK
}

func (tree *RedBlackTree[T]) Delete(v T) T {
	// newNode := &Node[T]{
	// 	value: v,
	// }
	
	// Zero value return
	var zero T
	return zero
}

func (tree *RedBlackTree[T]) transplant(u, v *Node[T]) {
	if u.parent == nil {
		tree.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}

	v.parent = u.parent
}

func (tree *RedBlackTree[T]) leftRotate(x *Node[T]) {
	// On N1:
	//         N0                 N0
	//        /                  /
	//       N1                 N3
	//      /  \     ->        /  \
	//    N2    N3            N1   N4
	//         /  \          /  \
	//        N5   N4       N2   N5

	y := x.right
	x.right = y.left

	if y.left != nil {
		y.left.parent = x
	}

	y.parent = x.parent

	if y.parent == nil {
		tree.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.left = x
	x.parent = y
}

func (tree *RedBlackTree[T]) rightRotate(x *Node[T]) {
	// On N1:
	//         N0                 N0
	//        /                  /
	//       N1                 N2
	//      /  \     ->        /  \
	//    N2    N3            N5   N1
	//   /  \                     /  \
	//  N5   N4                  N4   N3

	y := x.left
	x.left = y.right

	if y.right != nil {
		y.right.parent = x
	}

	y.parent = x.parent
	if x.parent == nil {
		tree.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.right = x
	x.parent = y
}
