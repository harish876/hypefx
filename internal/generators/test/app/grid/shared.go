package grid

import (
	"github.com/harish876/hypefx/components/dropdown"
	"github.com/harish876/hypefx/components/grid"
	"github.com/harish876/hypefx/components/input"
)

var (
	DEFAULT_PAGE_SIZE = 1
	DEFAULT_LIMIT     = 10
)

var Columns = []grid.GridColumn{
	{
		Typ:      grid.String,
		Label:    "Name",
		Key:      "Name",
		Renderer: "name",
	},
	{
		Typ:      grid.String,
		Label:    "Status",
		Key:      "Status",
		Renderer: "status",
		Editable: true,
		EditOptions: grid.GridEditOptions{
			EditType: grid.EditSelect,
			EditProps: grid.GridSelectEditProps{
				Id:    "Id",
				Name:  "status",
				Class: "mt-1 w-full rounded-md border-gray-200 bg-white text-sm text-gray-700 shadow-sm",
				Options: []dropdown.SelectOption{
					{Label: "Active", Value: "active"},
					{Label: "Inactive", Value: "inactive"},
				},
			},
		},
	},
	{
		Typ:      grid.String,
		Label:    "Role",
		Key:      "Position",
		Editable: true,
		EditOptions: grid.GridEditOptions{
			EditType: grid.EditInput,
			EditProps: grid.GridInputEditProps{
				Id:    "Id",
				Typ:   input.InputTypeText,
				Name:  "position",
				Class: "mt-1 w-full rounded-md border-gray-200 bg-white text-sm text-gray-700 shadow-sm",
			},
		},
	},
	{
		Typ:      grid.Array,
		Label:    "Badges",
		Key:      "Badges",
		Renderer: "badges",
		Editable: true,
		EditOptions: grid.GridEditOptions{
			EditType: grid.EditMultiSelect,
			EditProps: grid.GridMultiSelectEditProps{
				Id:   "Id",
				Name: "badges",
				Options: []dropdown.SelectOption{
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
		Typ:   grid.String,
		Label: "Email",
		Key:   "Email",
	},
}

var GridModel = grid.GridContext[grid.GridColumn]{
	Title:       "Customers",
	Subheading:  "",
	Description: "Example Gird with filters, pagination, export etc...",
	Columns:     Columns,
	IdField:     "Id",
	Url:         "/grid",
}
