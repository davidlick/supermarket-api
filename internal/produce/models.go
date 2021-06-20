package produce

import "github.com/Rhymond/go-money"

// Item models a produce item.
type Item struct {
	Code  string
	Name  string
	Price *money.Money
}
