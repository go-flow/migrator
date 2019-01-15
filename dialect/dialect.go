package dialect

import (
	"database/sql"
	"fmt"
	"reflect"
)

// Dialect defines set of methods needed
// by migrator to interact with different database engines
//
// It contains behaviors that differ across SQL database
type Dialect interface {
	// Name get dialect's name
	Name() string

	// SetDB set db for dialect
	SetDB(db *sql.DB)

	// DB gets dialect db
	DB() *sql.DB

	// BindVar return the placeholder for actual values in SQL statements, in many dbs it is "?", Postgres using $1
	BindVar(i int) string
	// Quote quotes field name to avoid SQL parsing exceptions by using a reserved word as a field name
	Quote(key string) string

	// HasTable check has table or not
	HasTable(tableName string) bool

	// SelectFromDummyTable return select values, for most dbs, `SELECT values` just works, mysql needs `SELECT value FROM DUAL`
	SelectFromDummyTable() string

	// LastInsertIdReturningSuffix most dbs support LastInsertId, but postgres needs to use `RETURNING`
	LastInsertIDReturningSuffix(tableName, columnName string) string

	// DefaultValueStr
	DefaultValueStr() string

	// CurrentDatabase return current database name
	CurrentDatabase() string

	// MigrationExists  checks if migration version exists in database table
	MigrationExists(version string, tableName string) (bool, error)

	// CreateMigrationTable creates migrations table in database
	CreateMigrationTable(tableName string) error

	// CountRecords retunrs number of rows in provided table
	CountRecords(tableName string) (int, error)

	// RemoveMigration deletes migration version from database table
	RemoveMigration(tableName string, version string) error
}

var dialectsMap = map[string]Dialect{}

// New creates dialect instance for given dialect name and db connection
func New(name string, db *sql.DB) Dialect {
	if value, ok := dialectsMap[name]; ok {
		dialect := reflect.New(reflect.TypeOf(value).Elem()).Interface().(Dialect)
		dialect.SetDB(db)
		return dialect
	}

	fmt.Printf("`%v` is not officially supported, running under compatibility mode.\n", name)
	commonDialect := &common{}
	return commonDialect
}

// RegisterDialect register new dialect
func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}
