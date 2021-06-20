# ramdb

RamDB is an implementation of an in-memory database with a simple API for selecting and querying the database. It uses b-trees as the underlying storage mechanism which allows fast searches and mutations.

All database commands are safe for concurrent operations.

## Example

```go
// Create database and table.
db := ramdb.NewDatabase()
_ = db.CreateTable("hotdogs", "frank_id")

// Mm, hotdogs.
type HotDog struct {
	FrankId string
	Condiments []string
	Brat bool
}

indog := HotDog{
	FrankId: "1",
	Condiments: []string{
		"kraut",
		"mustard",
	},
	Brat: true,
}

// Create a new Record and insert.
rec, _ := ramdb.NewRecord("1", "frank_id", indog)
_ = db.From("hotdogs").Insert(rec)


// Query data out.
ro, _ := db.From("hotdogs").Get("frank_id", "1")

var outdog HotDog
_ = ro.Deserialize(&outdog)

fmt.Printf("%+v\n", outdog)

// &HotDog{"1" ["kraut", "mustard"] true}
```
