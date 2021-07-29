package handler

import (
	"create_sp/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) getTables(c echo.Context) error {
	var tables []model.Table
	// return c.JSON(http.StatusOK, "test")
	rows, err := h.db.Raw("SHOW TABLES;").Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	defer rows.Close()
	for rows.Next() {
		var table model.Table

		rows2, err := h.db.Raw("SHOW COLUMNS FROM table1").Rows()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		defer rows2.Close()
		var cols []model.Column
		for rows2.Next() {
			var col model.Column
			rows2.Scan(&col.Field, &col.Type, &col.Null, &col.Key, &col.Default, &col.Extra)
			cols = append(cols, col)
		}

		rows.Scan(&table.Name)
		table.Columns = cols
		tables = append(tables, table)
	}

	return c.JSON(http.StatusOK, tables)
}
