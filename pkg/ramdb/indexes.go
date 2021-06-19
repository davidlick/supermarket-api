package ramdb

import "github.com/google/btree"

type index struct {
	tree   *btree.BTree
	column string
	table  *table
}
