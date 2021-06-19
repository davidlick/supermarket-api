package ramdb

import "github.com/google/btree"

// Get validates that the table and index for the given column exists and searches for the given key in the tree.
func (t *table) Get(column, key string) (r *record, err error) {
	if !t.exists {
		return nil, ErrNoTable
	}

	if !t.HasIndex(column) {
		return nil, ErrNoIndex
	}

	return t.keyLookup(key, t.indexes[column])
}

// keyLookup tries to find the given key in the tree. It returns ErrNoRecord if not found.
func (t *table) keyLookup(key string, index *index) (r *record, err error) {
	item := &record{id: keyHash(key)}
	result := index.tree.Get(item)
	if result == nil {
		return nil, ErrNoRecord
	}

	return result.(*record), nil
}

// Select returns all of the records in the database sorted in ascending order by id.
func (t *table) Select(column string) (rr []*record, err error) {
	if !t.exists {
		return nil, ErrNoTable
	}

	if !t.HasIndex(column) {
		return nil, ErrNoIndex
	}

	t.indexes[column].tree.Ascend(func(item btree.Item) bool {
		r := item.(*record)
		rr = append(rr, r)
		return true
	})

	return
}

// Insert adds the record to the database. It returns ErrRecordExists if the record already exists. Insert is thread safe.
func (t *table) Insert(r *record) error {
	if !t.exists {
		return ErrNoTable
	}

	if !t.HasIndex(r.keyColumn) {
		return ErrNoIndex
	}

	index := t.indexes[r.keyColumn]

	if has := index.tree.Has(r); has {
		return ErrRecordExists
	}

	t.mutex.Lock()
	defer t.mutex.Unlock()

	index.tree.ReplaceOrInsert(r)
	return nil

}

// Delete removes the item from the database. It returns ErrNoRecord if the record does not exist and ErrNotDeleted if the removal fails. Delete is thread safe.
func (t *table) Delete(r *record) error {
	if !t.exists {
		return ErrNoTable
	}

	if !t.HasIndex(r.keyColumn) {
		return ErrNoIndex
	}

	index := t.indexes[r.keyColumn]

	if has := index.tree.Has(r); !has {
		return ErrNoRecord
	}

	t.mutex.Lock()
	defer t.mutex.Unlock()

	index.tree.Delete(r)
	return nil
}
