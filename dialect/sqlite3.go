package dialect

type sqlite3 struct {
	common
}

func init() {
	RegisterDialect("sqlite3", &sqlite3{})
}

func (sqlite3) Name() string {
	return "sqlite3"
}

// HasTable check has table or not
func (s sqlite3) HasTable(tableName string) bool {
	var count int
	s.db.QueryRow("SELECT count(*) FROM sqlite_master WHERE type='table' AND name=?", tableName).Scan(&count)
	return count > 0
}

func (s sqlite3) CurrentDatabase() (name string) {
	var (
		ifaces   = make([]interface{}, 3)
		pointers = make([]*string, 3)
		i        int
	)
	for i = 0; i < 3; i++ {
		ifaces[i] = &pointers[i]
	}
	if err := s.db.QueryRow("PRAGMA database_list").Scan(ifaces...); err != nil {
		return
	}
	if pointers[1] != nil {
		name = *pointers[1]
	}
	return
}
