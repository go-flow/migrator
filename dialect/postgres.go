package dialect

import "fmt"

type postgres struct {
	common
}

func init() {
	RegisterDialect("postgres", &postgres{})
	RegisterDialect("cloudsqlpostgres", &postgres{})
}

func (postgres) Name() string {
	return "postgres"
}

func (postgres) BindVar(i int) string {
	return fmt.Sprintf("$%v", i)
}

func (p postgres) HasTable(tableName string) bool {
	var count int
	p.db.QueryRow("SELECT count(*) FROM INFORMATION_SCHEMA.tables WHERE table_name = $1 AND table_type = 'BASE TABLE' AND table_schema = CURRENT_SCHEMA()", tableName).Scan(&count)
	return count > 0
}

// LastInsertIdReturningSuffix most dbs support LastInsertId, but postgres needs to use `RETURNING`
func (postgres) LastInsertIDReturningSuffix(tableName, columnName string) string {
	return fmt.Sprintf("RETURNING %v.%v", tableName, columnName)
}

func (p postgres) CurrentDatabase() (name string) {
	p.db.QueryRow("SELECT CURRENT_DATABASE()").Scan(&name)
	return
}
