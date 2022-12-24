package database

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/photowey/popctl/configs"
)

const (
	TimeZone = "Asia%2FShanghai"
)

func ReverseEngineering() (*Database, error) {
	databasePtr, err := parseDatabase()

	return databasePtr, err
}

func parseDatabase() (*Database, error) {
	conf := configs.DatabaseFunc()
	dsn := fmt.Sprintf(DsnTemplate, conf.Username, conf.Password, conf.Host, conf.Port, conf.Database, TimeZone)
	driver, err := sql.Open(conf.Driver, dsn)
	if err != nil {
		return nil, err
	}
	err = driver.Ping()
	if err != nil {
		return nil, err
	}

	rows, err := driver.Query(TableInfoSQL)
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	if err != nil {
		return nil, err
	}

	database := &Database{
		Name:   conf.Database,
		Tables: make([]*Table, 0),
	}

	for rows.Next() {
		table := &Table{
			Fields: make([]*Field, 0),
		}
		err = rows.Scan(&table.Name)
		if err != nil {
			return nil, fmt.Errorf("scan table failed, err:%v", err)
		}
		err = reverseTable(driver, table)
		if err != nil {
			return nil, err
		}

		database.Tables = append(database.Tables, table)
	}

	return database, nil
}

func reverseTable(driver *sql.DB, table *Table) error {
	columnSQL := fmt.Sprintf(ColumnInfoSQL, table.Name)
	rows, err := driver.Query(columnSQL)
	if err != nil {
		return fmt.Errorf("query table:[%s] column failed, err:%v", table.Name, err)
	}
	for rows.Next() {
		field := &Field{}
		err = rows.Scan(&field.TableComment, &field.Name, &field.Type, &field.Comment, &field.TypLen, &field.TypLength, &field.NotNull, &field.PrimaryKey)
		if err != nil {
			// fmt.Printf("error sql: \n%s\n", columnSQL)
			return fmt.Errorf("scan table:[%s] column failed, err:%v", table.Name, err)
		}

		table.Fields = append(table.Fields, field)
	}

	return nil
}
