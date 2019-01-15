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
