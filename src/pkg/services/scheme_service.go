package services

import (
	"DBGEN/src/pkg/entities"
	"strconv"
)

func GenerateScheme(scheme *entities.Scheme) string {
	var query string = "CREATE SCHEMA [IF NOT EXISTS] " + scheme.Name + ";"
	for _, tab := range scheme.Tables {
		query += "CREATE TABLE [IF NOT EXISTS] " + tab.Name + "("
		for index, col := range tab.Columns {
			query += col.Name + "int" + strconv.Itoa(col.Size) + "NOT NULL"
			if index != len(tab.Columns)-1 {
				query += ","
			}
		}
		query += ")"
	}

	return query
}
