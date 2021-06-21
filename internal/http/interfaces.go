package http

import "github.com/davidlick/supermarket-api/internal/produce"

type ProduceService interface {
	Add(items []produce.Item) error
	Remove(item produce.Item) error
	Get(produceCode string) (item produce.Item, err error)
	All() (items []produce.Item, err error)
}
