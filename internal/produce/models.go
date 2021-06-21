package produce

import "github.com/Rhymond/go-money"

// Item models a produce item.
type Item struct {
	Code  string       `json:"code"`
	Name  string       `json:"name"`
	Price *money.Money `json:"price"`
}
