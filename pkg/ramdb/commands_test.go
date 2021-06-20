package ramdb

import (
	"fmt"
	"sort"
	"sync"
	"testing"

	"github.com/google/btree"
	"github.com/stretchr/testify/assert"
)

func TestTable_Get(t *testing.T) {
	tests := []struct {
		test           string
		tableConfig    func() *table
		expectedRecord *Record
		expectedError  error
	}{
		{
			test: "it should return ErrNoTable if an invalid table is supplied",
			tableConfig: func() *table {
				return &table{}
			},
			expectedError: ErrNoTable,
		},
		{
			test: "it should return ErrNoIndex if no index exists for column",
			tableConfig: func() *table {
				return &table{
					exists: true,
				}
			},
			expectedError: ErrNoIndex,
		},
		{
			test: "it should return ErrNoRecord if no Record was found for key",
			tableConfig: func() *table {
				return &table{
					exists: true,
					indexes: map[string]*index{
						"test_column": &index{
							tree: btree.New(5),
						},
					},
				}
			},
			expectedError: ErrNoRecord,
		},
		{
			test: "it should return a Record when one is found",
			tableConfig: func() *table {
				tbl := &table{
					exists: true,
					indexes: map[string]*index{
						"test_column": &index{
							tree: btree.New(5),
						},
					},
				}
				rec, err := NewRecord("test_key", "test_column", struct{}{})
				if err != nil {
					t.Error(err)
				}
				tbl.indexes["test_column"].tree.ReplaceOrInsert(rec)
				return tbl
			},
			expectedRecord: &Record{
				serialized: []uint8{0x7b, 0x7d},
				key:        "test_key",
				keyColumn:  "test_column",
				id:         0x92488e1e3eeecdf9,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.test, func(t *testing.T) {
			tbl := tc.tableConfig()

			Record, err := tbl.Get("test_column", "test_key")

			assert.Equal(t, tc.expectedRecord, Record)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestTable_Select(t *testing.T) {
	tests := []struct {
		test            string
		tableConfig     func() (*table, []*Record)
		expectedRecords []*Record
		expectedError   error
	}{
		{
			test: "it should return ErrNoTable if an invalid table is supplied",
			tableConfig: func() (*table, []*Record) {
				return &table{
					mutex: &sync.Mutex{},
				}, nil
			},
			expectedError: ErrNoTable,
		},
		{
			test: "it should return ErrNoIndex if no index exists for column",
			tableConfig: func() (*table, []*Record) {
				return &table{
					exists:  true,
					mutex:   &sync.Mutex{},
					indexes: make(map[string]*index),
				}, nil
			},
			expectedError: ErrNoIndex,
		},
		{
			test: "it should return all Records in the database",
			tableConfig: func() (*table, []*Record) {
				tbl := &table{
					exists: true,
					mutex:  &sync.Mutex{},
					indexes: map[string]*index{
						"test_column": &index{
							tree: btree.New(5),
						},
					},
				}

				expectedRecords := make([]*Record, 0)
				for i := 0; i < 10; i++ {
					key := fmt.Sprintf("key-%d", i)
					rec, err := NewRecord(key, "test_column", struct{}{})
					if err != nil {
						t.Error(err)
					}
					tbl.indexes["test_column"].tree.ReplaceOrInsert(rec)
					expectedRecords = append(expectedRecords, rec)
				}

				sort.Slice(expectedRecords, func(a, b int) bool {
					return expectedRecords[a].id < expectedRecords[b].id
				})

				return tbl, expectedRecords
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.test, func(t *testing.T) {
			tbl, expectedRecords := tc.tableConfig()

			rr, err := tbl.Select("test_column")

			assert.Equal(t, expectedRecords, rr)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestTable_Insert(t *testing.T) {
	tests := []struct {
		test          string
		tableConfig   func() *table
		expectedError error
	}{
		{
			test: "it should return ErrNoTable if an invalid table is supplied",
			tableConfig: func() *table {
				return &table{
					mutex: &sync.Mutex{},
				}
			},
			expectedError: ErrNoTable,
		},
		{
			test: "it should return ErrNoIndex if no index exists for column",
			tableConfig: func() *table {
				return &table{
					exists:  true,
					mutex:   &sync.Mutex{},
					indexes: make(map[string]*index),
				}
			},
			expectedError: ErrNoIndex,
		},
		{
			test: "it should return ErrRecordExists if a Record with key exists",
			tableConfig: func() *table {
				tbl := &table{
					exists: true,
					mutex:  &sync.Mutex{},
					indexes: map[string]*index{
						"test_column": &index{
							tree: btree.New(5),
						},
					},
				}

				rec, err := NewRecord("test_key", "test_column", struct{}{})
				if err != nil {
					t.Error(err)
				}
				tbl.indexes["test_column"].tree.ReplaceOrInsert(rec)

				return tbl
			},
			expectedError: ErrRecordExists,
		},
		{
			test: "it should return no error if successful",
			tableConfig: func() *table {
				return &table{
					exists: true,
					mutex:  &sync.Mutex{},
					indexes: map[string]*index{
						"test_column": &index{
							tree: btree.New(5),
						},
					},
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.test, func(t *testing.T) {
			tbl := tc.tableConfig()
			rec, err := NewRecord("test_key", "test_column", struct{}{})
			if err != nil {
				t.Error(err)
			}

			err = tbl.Insert(rec)

			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestTable_Delete(t *testing.T) {
	tests := []struct {
		test          string
		tableConfig   func() *table
		expectedError error
	}{
		{
			test: "it should return ErrNoTable if an invalid table is supplied",
			tableConfig: func() *table {
				return &table{
					mutex: &sync.Mutex{},
				}
			},
			expectedError: ErrNoTable,
		},
		{
			test: "it should return ErrNoIndex if no index exists for column",
			tableConfig: func() *table {
				return &table{
					exists:  true,
					mutex:   &sync.Mutex{},
					indexes: make(map[string]*index),
				}
			},
			expectedError: ErrNoIndex,
		},
		{
			test: "it should return ErrNoRecord if the Record does not exist",
			tableConfig: func() *table {
				return &table{
					exists: true,
					mutex:  &sync.Mutex{},
					indexes: map[string]*index{
						"test_column": &index{
							tree: btree.New(5),
						},
					},
				}
			},
			expectedError: ErrNoRecord,
		},
		{
			test: "it should return no error if the item was deleted",
			tableConfig: func() *table {
				tbl := &table{
					exists: true,
					mutex:  &sync.Mutex{},
					indexes: map[string]*index{
						"test_column": &index{
							tree: btree.New(5),
						},
					},
				}
				rec, err := NewRecord("test_key", "test_column", struct{}{})
				if err != nil {
					t.Error(err)
				}

				tbl.indexes["test_column"].tree.ReplaceOrInsert(rec)
				return tbl
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.test, func(t *testing.T) {
			tbl := tc.tableConfig()
			rec, err := NewRecord("test_key", "test_column", struct{}{})
			if err != nil {
				t.Error(err)
			}

			err = tbl.Delete(rec)

			assert.Equal(t, tc.expectedError, err)
		})
	}
}
