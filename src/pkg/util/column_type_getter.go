package util

import "DBGEN/src/pkg/entities"

func GetColumnType(columnType *entities.ColumnType) string {
	switch *columnType {
	case entities.ColumnInteger:
		return "INT "
	case entities.ColumnText:
		return "VARCHAR "
	case entities.ColumnBoolean:
		return "BOOLEAN "
	case entities.ColumnUUID:
		return "UUID "
	default:
		panic("Wrong type of")
	}
}
