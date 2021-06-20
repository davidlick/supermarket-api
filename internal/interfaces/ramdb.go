package interfaces

import "github.com/davidlick/supermarket-api/pkg/ramdb"

type RamDB interface {
	Get(column, key string) (r *ramdb.Record, err error)
	Select(column string) (rr []*ramdb.Record, err error)
	Insert(r *ramdb.Record) error
	Delete(r *ramdb.Record) error
}
