package cacher

import "beloin.com/distributed-cache/pkg/datastructure/trees"

type internalComparable struct {
	key   string
	value string
}

func (ic *internalComparable) Compare(other any) int {
	otheric, ok := other.(*internalComparable)
	if !ok {
		return 0
	}

	return len(ic.key) - len(otheric.key)
}

type RBCacher struct {
	rbtree trees.RedBlackTree[*internalComparable]
}

func (cacher *RBCacher) GetString(key string) (string, bool) {
	result := cacher.rbtree.Search(&internalComparable{key: key})
	if result != nil {
		return result.value, true
	}

	return "", false
}

func (cacher *RBCacher) SetString(key string, value string) bool {
	cacher.rbtree.Insert(&internalComparable{key: key, value: value})
	return true
}
