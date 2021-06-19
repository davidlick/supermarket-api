package ramdb

import "errors"

var (
	ErrNoTable     = errors.New("table does not exist")
	ErrTableExists = errors.New("table already exists")

	ErrNoRecord     = errors.New("record does not exist")
	ErrRecordExists = errors.New("record already exists")

	ErrNoIndex      = errors.New("index does not exist")
	ErrInvalidIndex = errors.New("invalid index column")
	ErrIndexExists  = errors.New("index already exists")
)
