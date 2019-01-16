package migrator

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestFileMigratorMysql(t *testing.T) {
	mp := "./testdata/migrations"
	dialect := "mysql"
	cs := "root:root@/test_migrator?multiStatements=true&readTimeout=1s&charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open(dialect, cs)
	if err != nil {
		t.Error(err)
	}

	migrator := NewFileMigrator(mp, dialect, db)

	if migrator.Path != mp {
		t.Errorf("expected path %s, got %s", mp, migrator.Path)
	}

	err = migrator.Status()
	if err != nil {
		t.Error(err)
	}

	err = migrator.Up()
	if err != nil {
		t.Error(err)
	}

	err = migrator.Down(1)
	if err != nil {
		t.Error(err)
	}

	err = migrator.Status()
	if err != nil {
		t.Error(err)
	}

	err = migrator.Reset()
	if err != nil {
		t.Error(err)
	}

	err = migrator.Down(3)
	if err != nil {
		t.Error(err)
	}

}
