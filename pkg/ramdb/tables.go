package ramdb

import (
	"sync"

	"github.com/google/btree"
)

type table struct {
	exists  bool
	mutex   *sync.Mutex
	indexes map[string]*index
}

// CreateIndex creates an index for onColumn.
func (t *table) CreateIndex(column string) error {
	if column == "" {
		return ErrInvalidIndex
	}

	if _, found := t.indexes[column]; found {
		return ErrIndexExists
	}

	idx := &index{
		tree:   btree.New(5),
		column: column,
		table:  t,
	}

	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.indexes[column] = idx
	return nil
}

// HasIndex returns true if an index exists for the column and false if it does not.
func (t *table) HasIndex(column string) bool {
	if _, found := t.indexes[column]; found {
		return true
	}

	return false
}
