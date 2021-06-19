package ramdb

type database struct {
	tables map[string]*table
}

// NewDatabase initializes a new database with no tables.
func NewDatabase() *database {
	return &database{
		tables: make(map[string]*table),
	}
}

// From selects a table for running commands.
func (db *database) From(tablename string) *table {
	t, ok := db.tables[tablename]
	if !ok {
		return &table{}
	}

	return t
}

// CreateTable creates a new table in the database with indexes for each column specified.
func (db *database) CreateTable(tablename string, indexOnColumns ...string) error {
	if _, found := db.tables[tablename]; found {
		return ErrTableExists
	}

	tbl := &table{
		exists: true,
	}

	for _, onColumn := range indexOnColumns {
		err := tbl.CreateIndex(onColumn)
		if err != nil {
			return err
		}
	}

	db.tables[tablename] = tbl
	return nil
}
