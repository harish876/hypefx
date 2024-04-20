package handlers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/harish876/hypefx/components/props"
	"github.com/harish876/hypefx/examples/db"
	"github.com/harish876/hypefx/examples/services"
	"github.com/harish876/hypefx/examples/views/grid"
	"github.com/labstack/echo/v4"
)

var (
	DEFAULT_PAGE_SIZE = 1
	DEFAULT_LIMIT     = 10
)

var columns = []props.GridColumn{
	{
		Typ:      props.String,
		Label:    "Name",
		Key:      "Name",
		Renderer: "name",
	},
	{
		Typ:      props.String,
		Label:    "Status",
		Key:      "Status",
		Renderer: "status",
		Editable: true,
		EditOptions: props.GridEditOptions{
			EditType: props.EditSelect,
			EditProps: props.GridSelectEditProps{
				Id:    "Id",
				Name:  "status",
				Class: "mt-1 w-full rounded-md border-gray-200 bg-white text-sm text-gray-700 shadow-sm",
				Options: []props.SelectOption{
					{Label: "Active", Value: "active"},
					{Label: "Inactive", Value: "inactive"},
				},
			},
		},
	},
	{
		Typ:      props.String,
		Label:    "Role",
		Key:      "Position",
		Editable: true,
		EditOptions: props.GridEditOptions{
			EditType: props.EditInput,
			EditProps: props.GridInputEditProps{
				Id:    "Id",
				Typ:   props.InputTypeText,
				Name:  "position",
				Class: "mt-1 w-full rounded-md border-gray-200 bg-white text-sm text-gray-700 shadow-sm",
			},
		},
	},
	{
		Typ:      props.Array,
		Label:    "Badges",
		Key:      "Badges",
		Renderer: "badges",
		Editable: true,
		EditOptions: props.GridEditOptions{
			EditType: props.EditMultiSelect,
			EditProps: props.GridMultiSelectEditProps{
				Id:   "Id",
				Name: "badges",
				Options: []props.SelectOption{
					{Label: "Design", Value: "Design"},
					{Label: "Product", Value: "Product"},
					{Label: "Marketing", Value: "Marketing"},
					{Label: "Engineering", Value: "Engineering"},
					{Label: "Analytics", Value: "Analytics"},
					{Label: "UI Design", Value: "UI Design"},
				},
			},
		},
	},
	{
		Typ:   props.String,
		Label: "Email",
		Key:   "Email",
	},
}

var gridCtx = props.GridContext[props.GridColumn]{
	Title:       "Customers",
	Subheading:  "",
	Description: "Example Gird with filters, pagination, export etc...",
	Columns:     columns,
	IdField:     "Id",
}

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

	var paginatedData []props.GridDataRow
	if offset < len(db.GridData) {
		end := offset + limit
		if end > len(db.GridData) {
			end = len(db.GridData)
		}
		paginatedData = db.GridData[offset:end]
	} else {
		paginatedData = []props.GridDataRow{}
	}

	var pageOptions = props.GridPagination{
		Current:    page,
		TotalPages: totalPages,
		Limit:      limit,
	}
	return Render(c, http.StatusOK, grid.Grid(gridCtx, paginatedData, pageOptions))
}

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

	var data []props.GridDataRow

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

	var pageOptions = props.GridPagination{
		Current:    page,
		TotalPages: totalPages,
		Limit:      limit,
	}

	return Render(c, http.StatusOK, grid.Grid(gridCtx, paginatedData, pageOptions))
}

func GridRowHandler(c echo.Context) error {
	id := c.Param("id")
	row := services.FilterById(db.GridData, id)
	return Render(c, http.StatusOK, grid.RenderRow(columns, row))
}

func UpdateGridRowHandler(c echo.Context) error {
	id := c.Param("id")
	var newData = props.GridDataRow{
		Id:       id,
		Status:   c.FormValue("status"),
		Position: c.FormValue("position"),
	}
	updatedRow := services.UpdateById(db.GridData, id, newData)
	fmt.Println(updatedRow)
	return Render(c, http.StatusOK, grid.RenderRow(columns, updatedRow))
}

func DeleteGridRowHandler(c echo.Context) error {
	id := c.Param("id")
	services.DeleteById(id)
	return c.NoContent(http.StatusOK)
}

func RenderEditGridHandler(c echo.Context) error {
	id := c.Param("id")
	row := services.FilterById(db.GridData, id)
	return Render(c, http.StatusOK, grid.EditRow(columns, row))
}

func getFieldStringValue(row props.GridDataRow, fieldName string) string {
	switch fieldName {
	case "name":
		return row.Name
	default:
		return ""
	}
}
