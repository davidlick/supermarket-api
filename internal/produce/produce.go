package produce

import (
	"strings"

	"github.com/davidlick/supermarket-api/internal/interfaces"
	"github.com/davidlick/supermarket-api/pkg/ramdb"
)

type service struct {
	db interfaces.RamDB
}

// NewService creates a new produce service for storing produce items.
func NewService(db interfaces.RamDB) *service {
	return &service{
		db: db,
	}
}

// Add adds the Items to the database.
func (s *service) Add(items []Item) error {
	for _, item := range items {
		rec, err := ramdb.NewRecord(strings.ToLower(item.Code), KeyProduceCode, item)
		if err != nil {
			return err
		}

		err = s.db.Insert(rec)
		if err != nil {
			return err
		}
	}

	return nil
}

// Remove removes the item from the database.
func (s *service) Remove(item Item) error {
	rec, err := ramdb.NewRecord(strings.ToLower(item.Code), KeyProduceCode, item)
	if err != nil {
		return err
	}

	return s.db.Delete(rec)
}

// Get fetches the produceCode from the database.
func (s *service) Get(produceCode string) (item Item, err error) {
	rec, err := s.db.Get(KeyProduceCode, strings.ToLower(produceCode))
	if err != nil {
		return
	}

	err = rec.Deserialize(&item)
	return
}

// All returns all produce items stored in the database.
func (s *service) All() (items []Item, err error) {
	recs, err := s.db.Select(KeyProduceCode)
	if err != nil {
		return
	}

	for _, rec := range recs {
		var item Item
		err = rec.Deserialize(&item)
		if err != nil {
			return
		}

		items = append(items, item)
	}

	return
}
