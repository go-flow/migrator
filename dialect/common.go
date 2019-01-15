package dialect

import (
	"database/sql"
	"fmt"
)

type common struct {
	db *sql.DB
}

func init() {
	RegisterDialect("common", &common{})
}

func (common) Name() string {
	return "common"
}

func (c *common) SetDB(db *sql.DB) {
	c.db = db
}

func (c *common) DB() *sql.DB {
	return c.db
}

func (common) BindVar(i int) string {
	return "$$$"
}

func (common) Quote(key string) string {
	return fmt.Sprintf(`"%s"`, key)
}

func (c common) HasTable(tableName string) bool {
	var count int
	currentDatabase := c.CurrentDatabase()
	c.db.QueryRow("SELECT count(*) FROM INFORMATION_SCHEMA.TABLES WHERE table_schema = ? AND table_name = ?", currentDatabase, tableName).Scan(&count)
	return count > 0
}

func (common) SelectFromDummyTable() string {
	return ""
}

func (common) LastInsertIDReturningSuffix(tableNAme, columnName string) string {
	return ""
}

func (common) DefaultValueStr() string {
	return "DEFAULT VALUES"
}

func (c common) CurrentDatabase() (name string) {
	c.db.QueryRow("SELECT DATABASE()").Scan(&name)
	return
}