package helpers

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	"strings"

	dbv1 "github.com/edernucci/database-schema-operator/api/v1"
)

func connect(options *dbv1.DatabaseSpec) (*sql.DB, error) {
	connectionString := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s", options.User, options.Password, options.Server, options.Name)
	return sql.Open("sqlserver", connectionString)
}

func exec(sql string, db *sql.DB) (int64, error) {
	stmt, err := db.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec()
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

// CheckTable is CheckTable
func CheckTable(t string, options *dbv1.DatabaseSpec) (bool, error) {
	db, err := connect(options)
	if err != nil {
		return false, err
	}
	defer db.Close()

	queryText := fmt.Sprintf("select 1 from [sys].[sysobjects] where [xtype] = 'u' and [name] = '%s'", t)
	i, err := exec(queryText, db)
	if err != nil {
		return false, err
	}

	return i > 0, nil
}

// CreateTable is CreateTable
func CreateTable(tableName string, columns []dbv1.Column, options *dbv1.DatabaseSpec) (int64, error) {
	db, err := connect(options)
	if err != nil {
		return 0, err
	}
	defer db.Close()
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
func UpdateColumns(tableName string, columns []dbv1.Column, options *dbv1.DatabaseSpec) (int64, error) {
	db, err := connect(options)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	for _, column := range columns {
		str := fmt.Sprintf("alter table [%s] alter column [%s] %s", tableName, column.Name, column.Type)
		log.Println(str)
		_, err := exec(str, db)
		if err != nil {
			log.Println(err)
		}
	}

	return 0, nil
}
