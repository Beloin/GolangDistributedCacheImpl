// Package datastructure
package datastructure

type Comparable interface {

	// greater:  1
	// less:    -1
	// equals:   0
	Compare(any) int
}
