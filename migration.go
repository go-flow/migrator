package migrator

import (
	"database/sql"
	"fmt"
)

// Migration handles data for a given database migration
type Migration struct {
	// Path to the migration
	//
	// for Example ``./migrations/2018051570142_initial.up.sql`
	Path string

	//Version of the migration
	//
	// for example 2`018051570142`
	Version string

	// Name of the migration
	//
	// for exmample `initial`
	Name string

	// Direction of the migration
	//
	// for ecample `up`
	Direction string

	// Runner function which executes migration
	Runner func(Migration, *sql.DB) error
}

// Run executes the migration.
//
// Returns error if runner is not defined,
// and returns result from Runner execution (error)
func (m Migration) Run(conn *sql.DB) error {
	if m.Runner == nil {
		return fmt.Errorf("migration runner not defined for %s", m.Path)
	}
	return m.Runner(m, conn)
}

// Exists checks if migration exists in DB
func (m Migration) Exists(conn *sql.DB, migrationTableName string) (bool, error) {
	return false, nil
}

// Migrations collection
type Migrations []Migration

// Len returns number of migrations
func (m Migrations) Len() int {
	return len(m)
}

// Less checks if migration version at index i is lesss than migration at index j
func (m Migrations) Less(i, j int) bool {
	return m[i].Version < m[j].Version
}

// Swap migrations at index i and j
func (m Migrations) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
