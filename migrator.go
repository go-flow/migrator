package migrator

import (
	"database/sql"
	"errors"
	"regexp"
)

// migrationRegEx defines regular expression for migration files
//
// example file `2018051570142_initial.up.sql`
var migrationRegEx = regexp.MustCompile(`(\d+)_([^\.]+)(\.[a-z]+)?\.(up|down)\.(sql)`)

// Migrator forms the basis of all migration systems.Migrator
// it does the actual heavy lifting of running migrations
type Migrator struct {
	dialect    string
	conn       *sql.DB
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
