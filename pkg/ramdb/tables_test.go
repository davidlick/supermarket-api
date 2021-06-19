package ramdb

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTable_CreateIndex(t *testing.T) {
	tests := []struct {
		test          string
		table         table
		column        string
		expectedError error
	}{
		{
			test:          "it should return ErrInvalidIndex if column is empty",
			expectedError: ErrInvalidIndex,
		},
		{
			test: "it should return ErrIndexExists if an index exists for column",
			table: table{
				mutex: &sync.Mutex{},
				indexes: map[string]*index{
					"test_column": &index{},
				}},
			column:        "test_column",
			expectedError: ErrIndexExists,
		},
		{
			test: "it should create an index successfully",
			table: table{
				mutex:   &sync.Mutex{},
				indexes: make(map[string]*index),
			},
			column: "test_column",
		},
	}

	for _, tc := range tests {
		t.Run(tc.test, func(t *testing.T) {
			err := tc.table.CreateIndex(tc.column)

			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestTable_HasIndex(t *testing.T) {
	tests := []struct {
		test        string
		tableConfig func() *table
		expectedHas bool
	}{
		{
			test: "it should return true if index exists",
			tableConfig: func() *table {
				return &table{
					indexes: map[string]*index{
						"test_column": &index{},
					},
				}
			},
			expectedHas: true,
		},
		{
			test: "it should return false if index does not exist",
			tableConfig: func() *table {
				return &table{
					indexes: make(map[string]*index),
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.test, func(t *testing.T) {
			tbl := tc.tableConfig()

			has := tbl.HasIndex("test_column")

			assert.Equal(t, tc.expectedHas, has)
		})
	}
}
