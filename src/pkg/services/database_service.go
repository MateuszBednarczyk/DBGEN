package services

import (
	"DBGEN/src/pkg/entities"
	"DBGEN/src/pkg/util"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func GenerateDatabase(database *entities.Database, c echo.Context) error {
	var query string
	query += "CREATE DATABASE " + database.Name + "; "
	for _, tab := range database.Tables {
		query += "CREATE TABLE " + tab.Name + " ( "
		query += "ID " + util.GetColumnType(&tab.PrimaryKey) + "UNIQUE NOT NULL, "
		for index, col := range tab.Columns {
			query += col.Name
			switch col.ColumnType {
			case entities.ColumnInteger:
				query += " INT "
			case entities.ColumnText:
				query += " VARCHAR "
			case entities.ColumnBoolean:
				query += " BOOLEAN "
			case entities.ColumnUUID:
				query += " UUID "
			}
			query += strconv.Itoa(col.Size) + " NOT NULL"
			if index != len(tab.Columns)-1 {
				query += ","
			}
		}
		query += "); "
	}

	for _, fk := range database.ForeignKeys {
		if fk.Source.ColumnType != fk.RelatedTo.ColumnType {
			return c.JSON(http.StatusBadRequest, "Please check if tables types equals")
		}
		query += "ALTER TABLE " + fk.Name + " ADD CONSTRAINT " + fk.Source.Name + " REFERENCES TO " + fk.RelatedTo.Name + "; "
	}

	return c.JSON(http.StatusOK, query)
}
