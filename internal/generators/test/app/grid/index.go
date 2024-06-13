package grid

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/harish876/hypefx/components/grid"
	"github.com/harish876/hypefx/examples/db"
	"github.com/harish876/hypefx/internal/generators/template"
	"github.com/labstack/echo/v4"
)

func GridHandler(c echo.Context) error {

	page, pageErr := strconv.Atoi(c.QueryParam("page"))
	limit, limitErr := strconv.Atoi(c.QueryParam("limit"))

	if pageErr != nil {
		page = DEFAULT_PAGE_SIZE
	}

	if limitErr != nil {
		limit = DEFAULT_LIMIT
	}

	totalPages := len(db.GridData)/limit + 1
	offset := (page - 1) * limit

	var paginatedData []db.GridDataRow
	if offset < len(db.GridData) {
		end := offset + limit
		if end > len(db.GridData) {
			end = len(db.GridData)
		}
		paginatedData = db.GridData[offset:end]
	} else {
		paginatedData = []db.GridDataRow{}
	}

	var pageOptions = grid.GridPagination{
		Current:    page,
		TotalPages: totalPages,
		Limit:      limit,
	}
	return template.Render(c, http.StatusOK, grid.Grid(&GridModel, paginatedData, pageOptions))
}

// @post
func GridFilterHandler(c echo.Context) error {
	page, pageErr := strconv.Atoi(c.QueryParam("page"))
	limit, limitErr := strconv.Atoi(c.QueryParam("limit"))

	if pageErr != nil {
		page = DEFAULT_PAGE_SIZE
	}

	if limitErr != nil {
		limit = DEFAULT_LIMIT
	}

	filterKey := "name"
	filterValue := c.FormValue(filterKey)
	fmt.Println(filterValue)

	var data []db.GridDataRow

	for _, row := range db.GridData {
		fieldValue := getFieldStringValue(row, filterKey)
		if strings.Contains(strings.ToLower(fieldValue), strings.ToLower(filterValue)) {
			data = append(data, row)
		}
	}
	totalPages := int(math.Ceil(float64(len(data)) / float64(limit)))
	offset := (page - 1) * limit

	start := offset
	end := offset + limit
	if end > len(data) {
		end = len(data)
	}
	paginatedData := data[start:end]

	var pageOptions = grid.GridPagination{
		Current:    page,
		TotalPages: totalPages,
		Limit:      limit,
	}

	return template.Render(c, http.StatusOK, grid.Grid(&GridModel, paginatedData, pageOptions))
}

func getFieldStringValue(row db.GridDataRow, fieldName string) string {
	switch fieldName {
	case "name":
		return row.Name
	default:
		return ""
	}
}
