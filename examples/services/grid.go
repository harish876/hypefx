package services

import (
	"github.com/harish876/hypefx/examples/db"
)

func FilterById(data []db.GridDataRow, id string) db.GridDataRow {
	for _, row := range data {
		if row.Id == id {
			return row
		}
	}
	return db.GridDataRow{}
}

func UpdateById(data []db.GridDataRow, id string, newData db.GridDataRow) db.GridDataRow {
	for i, row := range data {
		if row.Id == id {
			data[i].Status = newData.Status
			data[i].Position = newData.Position
			return data[i]
		}
	}
	return db.GridDataRow{}
}

func DeleteById(id string) {
	var updatedData []db.GridDataRow
	for _, row := range db.GridData {
		if row.Id != id {
			updatedData = append(updatedData, row)
		}
	}

	db.GridData = updatedData

}
