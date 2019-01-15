package dialect

import "fmt"

type mssql struct {
	common
}

func init() {
	RegisterDialect("mssql", &mssql{})
}

func (mssql) Name() string {
	return "mssql"
}

func (mssql) Quote(key string) string {
	return fmt.Sprintf(`[%s]`, key)
}

func (s mssql) HasTable(tableName string) bool {
	var count int
	currentDatabase := s.CurrentDatabase()
	s.db.QueryRow("SELECT count(*) FROM INFORMATION_SCHEMA.tables WHERE table_name = ? AND table_catalog = ?", tableName, currentDatabase).Scan(&count)
	return count > 0
}

// CurrentDatabase return current database name
func (s mssql) CurrentDatabase() (name string) {
	s.db.QueryRow("SELECT DB_NAME() AS [Current Database]").Scan(&name)
	return
}
