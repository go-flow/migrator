package migrator

import (
	"database/sql"
	"errors"
	"regexp"

	"github.com/go-flow/migrator/dialect"
)

// migrationRegEx defines regular expression for migration files
//
// example file `2018051570142_initial.up.sql`
var migrationRegEx = regexp.MustCompile(`(\d+)_([^\.]+)(\.[a-z]+)?\.(up|down)\.(sql)`)

// New returns a new "blank" migrator.
//
// a blank Migrator should be only used as
// basis for a new type of migration system
func New(dialectName string, conn *sql.DB) Migrator {
	return Migrator{
		dialect: dialect.New(dialectName, conn),
		migrations: map[string]Migrations{
			"up":   Migrations{},
			"down": Migrations{},
		},
	}
}

// Migrator forms the basis of all migration systems.Migrator
// it does the actual heavy lifting of running migrations
type Migrator struct {
	dialect    dialect.Dialect
	schemaPath string
	migrations map[string]Migrations
}

// Up runs pending `up` migrations and applies them to the database
func (m Migrator) Up() error {
	return errors.New("Method not Implemented")
}

// Down runs pending `down` migrations and
//rolls back the database by the specified number of steps
func (m Migrator) Down(step int) error {
	return errors.New("Method not Implemented")
}

// Status prints out the status of applied/Pending migrations
func (m Migrator) Status() error {
	return errors.New("Method not Implemented")
}

// Reset executes all `down` migrations followed by the `up` migrations
func (m Migrator) Reset() error {
	return errors.New("Method not Implemented")
}
