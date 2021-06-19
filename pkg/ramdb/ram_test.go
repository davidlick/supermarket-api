package ramdb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabase_From(t *testing.T) {
	tests := []struct {
		test           string
		expectFunc     func(t *testing.T, db *database)
		expectedExists bool
	}{
		{
			test:           "it should set exist false if table not found",
			expectFunc:     func(t *testing.T, db *database) {},
			expectedExists: false,
		},
		{
			test: "it should set exist true if table is found",
			expectFunc: func(t *testing.T, db *database) {
				db.CreateTable("test_table")
			},
			expectedExists: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.test, func(t *testing.T) {
			db := NewDatabase()
			tc.expectFunc(t, db)

			table := db.From("test_table")

			assert.Equal(t, table.exists, tc.expectedExists)
		})
	}
}

func TestDatabase_CreateTable(t *testing.T) {
	tests := []struct {
		test          string
		expectFunc    func(t *testing.T, db *database)
		expectedError error
	}{
		{
			test: "it should error if the table already exists",
			expectFunc: func(t *testing.T, db *database) {
				db.CreateTable("test_table")
			},
			expectedError: ErrTableExists,
		},
		{
			test:       "it should not return an error on successful table creation",
			expectFunc: func(t *testing.T, db *database) {},
		},
	}

	for _, tc := range tests {
		t.Run(tc.test, func(t *testing.T) {
			db := NewDatabase()
			tc.expectFunc(t, db)

			err := db.CreateTable("test_table")

			assert.Equal(t, tc.expectedError, err)
		})
	}
}
