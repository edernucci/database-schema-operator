package helpers

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	dbv1 "github.com/edernucci/database-schema-operator/api/v1"
)

func connect() (*sql.DB, error) {
	return sql.Open("sqlserver", "sqlserver://sa:ItsSecret2020@localhost?database=master")
}

func exec(sql string, db *sql.DB) (int64, error) {
	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec()
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

// CheckTable is CheckTable
func CheckTable(t string) (bool, error) {
	db, err := connect()
	defer db.Close()
	if err != nil {
		return false, err
	}

	queryText := fmt.Sprintf("select 1 from [sys].[sysobjects] where [xtype] = 'u' and [name] = '%s'", t)
	i, err := exec(queryText, db)
	if err != nil {
		return false, err
	}

	return (i > 0), nil
}

// CreateTable is CreateTable
func CreateTable(tableName string, columns []dbv1.Column) (int64, error) {
	db, err := connect()
	defer db.Close()
	if err != nil {
		return 0, err
	}
	var str strings.Builder
	str.WriteString(fmt.Sprintf("create table [%s] (", tableName))
	for _, column := range columns {
		str.WriteString(fmt.Sprintf("[%s] %s,", column.Name, column.Type))
	}
	str.WriteString(")")

	log.Println(str.String())

	return exec(str.String(), db)
}

// UpdateColumns is UpdateColumns
func UpdateColumns(tableName string, columns []dbv1.Column) (int64, error) {
	db, err := connect()
	defer db.Close()
	if err != nil {
		return 0, err
	}

	for _, column := range columns {
		str := fmt.Sprintf("alter table [%s] alter column [%s] %s", tableName, column.Name, column.Type)
		log.Println(str)
		exec(str, db)
	}

	return 0, nil
}
