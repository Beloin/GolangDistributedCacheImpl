package cacher

import "beloin.com/distributed-cache/pkg/datastructure/trees"

type internalComparable struct {
	v string
}

func (ic *internalComparable) Compare(other any) int {
	otheric, ok := other.(*internalComparable)
	if !ok {
		return 0
	}

	return len(ic.v) - len(otheric.v)
}

type RBCacher struct {
	rbtree trees.RedBlackTree[*internalComparable]
}
