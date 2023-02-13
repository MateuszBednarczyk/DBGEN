package services

import (
	"DBGEN/src/pkg/entities"
	"strconv"
)

func GenerateScheme(database *entities.Database) string {
	var query string
	query += "CREATE DATABASE " + database.Name + "; "
	scheme := database.Scheme
	query += "CREATE SCHEMA [IF NOT EXISTS] " + scheme.Name + ";"
	for _, tab := range scheme.Tables {
		query += "CREATE TABLE [IF NOT EXISTS] " + tab.Name + "("
		for index, col := range tab.Columns {
			switch col.ColumnType {
			case entities.ColumnInteger:
				query += "INT "
			case entities.ColumnText:
				query += "VARCHAR "
			case entities.ColumnBoolean:
				query += "BOOLEAN "
			case entities.ColumnUUID:
				query += "UUID "
			}
			query += strconv.Itoa(col.Size) + " NOT NULL"
			if index != len(tab.Columns)-1 {
				query += ","
			}
		}
		query += ")"
	}

	return query
}
