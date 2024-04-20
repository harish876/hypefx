package services

import (
	"github.com/harish876/hypefx/components/props"
	"github.com/harish876/hypefx/examples/db"
)

func FilterById(data []props.GridDataRow, id string) props.GridDataRow {
	for _, row := range data {
		if row.Id == id {
			return row
		}
	}
	return props.GridDataRow{}
}

func UpdateById(data []props.GridDataRow, id string, newData props.GridDataRow) props.GridDataRow {
	for i, row := range data {
		if row.Id == id {
			data[i].Status = newData.Status
			data[i].Position = newData.Position
			return data[i]
		}
	}
	return props.GridDataRow{}
}

func DeleteById(id string) {
	var updatedData []props.GridDataRow
	for _, row := range db.GridData {
		if row.Id != id {
			updatedData = append(updatedData, row)
		}
	}

	db.GridData = updatedData

}
