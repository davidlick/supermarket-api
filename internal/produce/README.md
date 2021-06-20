# Produce Service

This produce service allows persisting produce in a database. It supports adding multiple produce items, removing a produce item, getting a produce item by produce code, and selecting all produce items from the database.

## Example

```go
// Create database and table.
db := ramdb.NewDatabase()
_ = db.CreateTable("produce", KeyProduceCode)

produceSvc := produce.NewService(db)

produceItem := produce.Item{
	Code: "A12T-4GH7-QPL9-3N4M",
	Name: "Lettuce",
	Price: money.New(346, "USD"),
}

_ = produceSvc.Add([]Item{produceItem})
_, _ = produceSvc.Get(produceItem.Code)
_, _ = produceSvc.All()
_ = produceSvc.Remove(produceItem)
```
