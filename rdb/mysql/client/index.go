package client

import (
	"github.com/ideal-rucksack/datasource-golang-plugin/rdb"
)

type Client struct {
}

func (c Client) Databases() ([]string, error) {
	return []string{"db1", "db2"}, nil
}

func (c Client) Tables(database string) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (c Client) Columns(database, table string) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (c Client) ExecuteQuery(databaseName string, query string) ([]map[string]interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (c Client) Execute(database, sql string) error {
	//TODO implement me
	panic("implement me")
}

func (c Client) Ping() error {
	//TODO implement me
	panic("implement me")
}

func (c Client) TableSchema(database, table string) (rdb.TableSchema, error) {
	//TODO implement me
	panic("implement me")
}
