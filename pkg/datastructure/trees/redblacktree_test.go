package trees

import (
	"reflect"
	"testing"
)

type TestInt int

func (t TestInt) Compare(other any) int {
	otherInt := other.(TestInt)
	if t < otherInt {
		return -1
	}
	if t > otherInt {
		return 1
	}
	return 0
}

func TestRedBlackTree_InsertSingle(t *testing.T) {
	tree := NewRedBlackTree[TestInt]()
	tree.Insert(TestInt(42))

	if tree.root == nil {
		t.Error("Tree should have root after insertion")
	}
	if tree.root.value != TestInt(42) {
		t.Errorf("Root value should be 42, got %v", tree.root.value)
	}
	if tree.root.color != BLACK {
		t.Error("Root should be black")
	}
	if tree.root.left != nil {
		t.Error("Root left should be nil")
	}
	if tree.root.right != nil {
		t.Error("Root right should be nil")
	}
}

func TestRedBlackTree_ReadTree(t *testing.T) {
	tree := NewRedBlackTree[TestInt]()
	tree.Insert(TestInt(42))

	arr := make([]TestInt, 0, 1)
	arr = tree.Read(arr)

	if len(arr) == 0 {
		t.Error("No item read")
	}
}

func TestRedBlackTree_SmallTreeRead(t *testing.T) {
	// Tree:
	//            2B
	//           /  \
	//          1B  4B
	//             /  \
	//            3R   5R
	//
	// Result order (stack)
	//  [ 2, 1, 4, 3, 5 ]
	tree := NewRedBlackTree[TestInt]()
	tree.Insert(TestInt(1))
	tree.Insert(TestInt(2))
	tree.Insert(TestInt(3))
	tree.Insert(TestInt(4))
	tree.Insert(TestInt(5))

	arr := make([]TestInt, 0, 1)
	arr = tree.Read(arr)

	if len(arr) == 0 {
		t.Error("No item read")
	}

	expected := []TestInt{2, 1, 4, 3, 5}

	if !reflect.DeepEqual(expected, arr) {
		t.Errorf("Expected %+v != %+v", expected, arr)
	}
}

func TestRedBlackTree_SmallTreeReadLeft(t *testing.T) {
	// Tree:
	//            5B
	//          /    \
	//         3B     \
	//        /  \     \
	//       2B   4B    \
	//      /            \
	//     1R            11B
	//                  /    \
	//                 /     21B
	//                /     /   \
	//               /    20R   22R
	//              8R
	//             /  \
	//           6B    10B
	//                /
	//               9R
	//
	// Result order (stack)
	//  [ 5, 3, 2, 1, 4, 11, 8, 6, 10, 9, 21, 20, 22 ]
	tree := NewRedBlackTree[TestInt]()
	tree.Insert(TestInt(10))
	tree.Insert(TestInt(11))
	tree.Insert(TestInt(21))
	tree.Insert(TestInt(5))
	tree.Insert(TestInt(4))
	tree.Insert(TestInt(6))
	tree.Insert(TestInt(3))
	tree.Insert(TestInt(2))
	tree.Insert(TestInt(8))
	tree.Insert(TestInt(1))
	tree.Insert(TestInt(9))
	tree.Insert(TestInt(20))
	tree.Insert(TestInt(22))

	arr := make([]TestInt, 0, 1)
	arr = tree.Read(arr)

	if len(arr) == 0 {
		t.Error("No item read")
	}

	expected := []TestInt{5, 3, 2, 1, 4, 11, 8, 6, 10, 9, 21, 20, 22}
	if !reflect.DeepEqual(expected, arr) {
		t.Errorf("Expected %+v != %+v", expected, arr)
	}
}

// TODO: Implement tests based on this
//
//
// func TestRedBlackTree_New(t *testing.T) {
// 	tree := NewRedBlackTree[TestInt]()
// 	if tree == nil {
// 		t.Error("NewRedBlackTree should not return nil")
// 	}
// 	if tree.root != nil {
// 		t.Error("New tree should have nil root")
// 	}
// 	if tree.Depth != 0 {
// 		t.Error("New tree should have depth 0")
// 	}
// }
//
// func TestRedBlackTree_InsertSingle(t *testing.T) {
// 	tree := NewRedBlackTree[TestInt]()
// 	tree.Insert(TestInt(42))
//
// 	if tree.root == nil {
// 		t.Error("Tree should have root after insertion")
// 	}
// 	if tree.root.value != TestInt(42) {
// 		t.Errorf("Root value should be 42, got %v", tree.root.value)
// 	}
// 	if tree.root.color != BLACK {
// 		t.Error("Root should be black")
// 	}
// 	if tree.root.left != nil {
// 		t.Error("Root left should be nil")
// 	}
// 	if tree.root.right != nil {
// 		t.Error("Root right should be nil")
// 	}
// }
//
// func TestRedBlackTree_InsertMultiple(t *testing.T) {
// 	tree := NewRedBlackTree[TestInt]()
// 	values := []TestInt{50, 30, 70, 20, 40, 60, 80}
//
// 	for _, v := range values {
// 		tree.Insert(v)
// 	}
//
// 	// Basic structure validation
// 	if tree.root == nil {
// 		t.Fatal("Tree should have root")
// 	}
//
// 	// Root should be black
// 	if tree.root.color != BLACK {
// 		t.Error("Root must be black")
// 	}
//
// 	// Verify root value (should be 50 or one of the middle values)
// 	if tree.root.value != TestInt(50) && tree.root.value != TestInt(40) && tree.root.value != TestInt(60) {
// 		t.Errorf("Unexpected root value: %v", tree.root.value)
// 	}
// }
//
// func TestRedBlackTree_GetNextLeaf(t *testing.T) {
// 	tree := NewRedBlackTree[TestInt]()
// 	tree.Insert(TestInt(50))
// 	tree.Insert(TestInt(30))
// 	tree.Insert(TestInt(70))
//
// 	// Test getting next leaf for various values
// 	testNode := &Node[TestInt]{value: TestInt(25)}
// 	leaf := tree.GetNextLeaf(testNode)
// 	if leaf == nil {
// 		t.Fatal("Should find a leaf")
// 	}
// 	if leaf.value != TestInt(30) {
// 		t.Errorf("Expected leaf value 30, got %v", leaf.value)
// 	}
//
// 	testNode2 := &Node[TestInt]{value: TestInt(60)}
// 	leaf2 := tree.GetNextLeaf(testNode2)
// 	if leaf2.value != TestInt(70) {
// 		t.Errorf("Expected leaf value 70, got %v", leaf2.value)
// 	}
//
// 	testNode3 := &Node[TestInt]{value: TestInt(80)}
// 	leaf3 := tree.GetNextLeaf(testNode3)
// 	if leaf3.value != TestInt(70) {
// 		t.Errorf("Expected leaf value 70, got %v", leaf3.value)
// 	}
// }
//
// func TestRedBlackTree_ColorProperties(t *testing.T) {
// 	tree := NewRedBlackTree[TestInt]()
// 	values := []TestInt{10, 5, 15, 3, 7, 12, 18}
//
// 	for _, v := range values {
// 		tree.Insert(v)
// 	}
//
// 	// Verify root is black
// 	if tree.root.color != BLACK {
// 		t.Error("Root must be black")
// 	}
//
// 	// Verify no red node has red children (property 4)
// 	verifyNoRedRedViolation(t, tree.root)
// }
//
// func verifyNoRedRedViolation[T datastructure.Comparable](t *testing.T, node *Node[T]) {
// 	if node == nil {
// 		return
// 	}
//
// 	if node.color == RED {
// 		if node.left != nil && node.left.color == RED {
// 			t.Errorf("Red node %v has red left child %v", node.value, node.left.value)
// 		}
// 		if node.right != nil && node.right.color == RED {
// 			t.Errorf("Red node %v has red right child %v", node.value, node.right.value)
// 		}
// 	}
//
// 	verifyNoRedRedViolation(t, node.left)
// 	verifyNoRedRedViolation(t, node.right)
// }
//
// func TestRedBlackTree_EmptyTree(t *testing.T) {
// 	tree := NewRedBlackTree[TestInt]()
//
// 	// GetNextLeaf on empty tree should return nil?
// 	// Current implementation may panic, so we'll test carefully
// 	if tree.root != nil {
// 		t.Error("Empty tree should have nil root")
// 	}
// }
