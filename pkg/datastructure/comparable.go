// Package datastructure
package datastructure

type Comparable interface {
	// greater:   >0
	// less:      <0
	// equals:    =0
	Compare(any) int
}
