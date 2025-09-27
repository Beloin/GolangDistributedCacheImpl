// Package trees
package trees

import (
	"beloin.com/distributed-cache/pkg/datastructure"
	"beloin.com/distributed-cache/pkg/datastructure/collections"
)

type Color uint8

const (
	BLACK = iota
	RED   = iota
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

func (r *RedBlackTree[T]) Minimum(node *Node[T]) *Node[T] {
	if node == nil {
		return nil
	}

	for node.left != nil {
		node = node.left
	}

	return node
}

func NewRedBlackTree[T datastructure.Comparable]() *RedBlackTree[T] {
	return &RedBlackTree[T]{}
}

func (r *RedBlackTree[T]) SearchNode(value T) *Node[T] {
	var node *Node[T]
	for curr := r.root; curr != nil; {
		compare := value.Compare(curr.value)
		if compare > 0 {
			curr = curr.right
			continue
		}

		if compare == 0 {
			node = curr
			break
		}

		curr = curr.left
	}

	return node
}

func (r *RedBlackTree[T]) Search(value T) T {
	node := r.SearchNode(value)
	if node == nil {
		var zero T
		return zero
	}

	return node.value
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

func (r *RedBlackTree[T]) Read(outSlice []T) []T {
	stack := collections.NewStack(1)

	node := r.root
	for node != nil {
		if node.right != nil {
			stack.Push(node.right)
		}

		if node.left != nil {
			stack.Push(node.left)
		}

		outSlice = append(outSlice, node.value)

		if stack.Empty() {
			node = nil
		} else {
			node = stack.Pop().(*Node[T])
		}
	}

	return outSlice
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
	r.insertFix(newNode)
}

func (tree *RedBlackTree[T]) insertFix(n *Node[T]) {
	// loop invariant: `n` is always RED
	for n.parent != nil && n.parent.color == RED { // Violation of rule 4.
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
					tree.leftRotate(n)                  // |
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

	var zero T

	node := tree.SearchNode(v)
	if node == nil {
		return zero
	}

	if node.left == nil && node.right == nil {
		if node.parent != nil {
			if node.parent.left == node {
				node.parent.left = nil
			} else {
				node.parent.right = nil
			}
		} else {
			tree.root = nil
		}

		// free(node)
		return node.value
	}

	// TODO: This is based that node is an INTERNAL Node (meaning is not a leaf)

	y := node
	yColor := y.color

	var x *Node[T]
	if node.left == nil {
		x = node.right
		tree.transplant(node, node.right)
	} else if node.right == nil {
		x = node.left
		tree.transplant(node, node.left)
	} else {
		y = tree.Minimum(node.right)
		yColor = y.color
		x = y.right
		
		if y != node.right {
			tree.transplant(y, y.right)
			y.right = node.right
			y.right.parent = y
		} else {
			// Creating a invalid node to prevent nil pointer
			// TODO: Possible problem when x is this, how to know if it's right or left of Y
			x = &Node[T]{color: BLACK}
			x.parent = y
			tree.transplant(node, y)
			y.left = node.left
			y.left.parent  = y
			y.color = node.color
		}
	}

	if yColor == BLACK {
		tree.deleteFix(x)
	}

	return node.value
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

func (tree *RedBlackTree[T]) deleteFix(node*Node[T])  {
	for node != tree.root && node.color == BLACK {
		if node == node.parent.left {
			w := node.parent.right
			
			if w.color == RED {
				w.color = BLACK                  //   | Case 01
				node.parent.color = RED          //   |
				tree.leftRotate(node.parent)
				w = node.parent.right
			}

			if w.left.color == BLACK && w.right.color == BLACK {
				w.color = RED                    //  | Case 02
				node = node.parent               //  |
			} else {
				if w.right.color == BLACK {
					w.left.color = BLACK           //  | Case 03
					w.color = RED                  //  |
					tree.rightRotate(w)
					w = node.parent.right
				}

				w.color = node.parent.color      //  | Case 04
				node.parent.color = BLACK        //  |

				w.right.color = BLACK
				tree.leftRotate(node.parent)
				node = tree.root
			}
		} else {
			w := node.parent.left

			if w.color == RED {
				w.color = BLACK
				node.parent.color = RED
				tree.rightRotate(node.parent)
				w = node.parent.left
			}

			if w.right.color == BLACK && w.left.color == BLACK {
				w.color = RED
				node = node.parent
			} else {
				if w.left.color == BLACK {
					w.right.color = BLACK
					w.color = RED
					tree.leftRotate(w)
					w = node.parent.left
				}

				w.color = node.parent.color
				node.parent.color = BLACK
				w.left.color = BLACK
				tree.rightRotate(node.parent)
				node = tree.root
			}
		}
	}

	node.color = BLACK
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
