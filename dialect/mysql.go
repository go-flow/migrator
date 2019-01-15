package dialect

import "fmt"

type mysql struct {
	common
}

func init() {
	RegisterDialect("mysql", &mysql{})
}

func (mysql) Name() string {
	return "mysql"
}

func (mysql) Quote(key string) string {
	return fmt.Sprintf("`%s`", key)
}

func (mysql) SelectFromDummyTable() string {
	return "FROM DUAL"
}

func (mysql) DefaultValueStr() string {
	return "VALUES()"
}

func (c *mysql) MigrationExists(version string, tableName string) (bool, error) {
	var count int
	qStr := fmt.Sprintf("SELECT COUNT(*) FROM %s where version = ?", tableName)
	err := c.db.QueryRow(qStr, version).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (c *mysql) CreateMigrationTable(tableName string) error {
	q := fmt.Sprintf(`CREATE TABLE %s ( 
		version NVARCHAR(14) NOT NULL, 
		name NVARCHAR(255) NULL, 
		UNIQUE INDEX  schema_version_idx (version ASC));`, tableName)
	_, err := c.db.Exec(q)
	return err
}

func (c *mysql) CountRecords(tableName string) (int, error) {
	q := fmt.Sprintf("SELECT COUNT(*) FROM %s", tableName)
	var count int
	err := c.db.QueryRow(q).Scan(&count)
	return count, err
}

func (c *mysql) RemoveMigration(tableName string, version string) error {
	q := fmt.Sprintf("delete from %s where version = ?", tableName)
	_, err := c.db.Exec(q, version)
	return err
}
