package trees

// TODO: Implement tests based on this

// import (
// 	"testing"
//
// 	"beloin.com/distributed-cache/pkg/datastructure"
// )
//
// // TestInt implements Comparable for testing
// type TestInt int
//
// func (t TestInt) Compare(other datastructure.Comparable) int {
// 	otherInt := other.(TestInt)
// 	if t < otherInt {
// 		return -1
// 	}
// 	if t > otherInt {
// 		return 1
// 	}
// 	return 0
// }
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
