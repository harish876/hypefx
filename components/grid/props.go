package grid

import "github.com/harish876/hypefx/components/dropdown"

const (
	Integer = iota
	String
	Array
)

const (
	EditInput = iota
	EditSelect
	EditMultiSelect
)

// Shape of the data passed to the grid

type GridPagination struct {
	Current    int
	TotalPages int
	Limit      int
}

type GridContext[C any] struct {
	IdField     string         // Primary Identifier Field for the data being pased to the grid
	Version     int            // Grid Version number
	Title       string         // Title for the data being pased to the grid
	Subheading  string         // Subheading for the data being pased to the grid
	Description string         // Description for the data being pased to the grid
	Url         string         // Base URL for the data being pased to the grid
	Columns     []C            // List of columns for the data being pased to the grid
	PageOptions GridPagination // Pagination Options for the data being pased to the grid
}

type GridInputEditProps struct {
	Id    string // Id for the put request from the data. This should be the key of the unique identifier representing the data sent down
	Typ   string // Type of input component to be rendered
	Name  string // Name of the input for the put request i.e the key you recieve from the form value. Ideally this should be equal to the key in the database but can be changed if the parameter is passed
	Class string //  CSS class for styling the edit button
}

type GridSelectEditProps struct {
	Id      string // Id for the put request from the data
	Name    string // Name of the input for the put request i.e the key you recieve from the form value. Ideally this should be equal to the key in the database but can be changed if the parameter is passed
	Class   string
	Options []dropdown.SelectOption
}
type GridMultiSelectEditProps struct {
	Id                     string // Id for the put request from the data
	Name                   string // Name of the input for the put request i.e the key you recieve from the form value. Ideally this should be equal to the key in the database but can be changed if the parameter is passed
	Options                []dropdown.SelectOption
	DefaultSelectedOptions []dropdown.SelectOption
}

type GridEditProps interface{}

type GridEditOptions struct {
	EditType  int // EditType field specifies the type of editable (props.EditInput, props.EditMultiSelect,props.EditSelect)
	EditProps GridEditProps
}

type GridColumn struct {
	Typ         int             // Column Type (props.Integer,props.String, props.Array) - Mandatory
	Label       string          // Column Label Displayd on the UI (Usually a camelCased version of the data key) - Mandatory
	Key         string          // Actualy Data key to the render on the UI - Mandatory
	Filter      bool            // this can be many things todo - Optional
	Renderer    string          // Custom string values which can be associated to a render function in the row renderer function - Optional
	Editable    bool            // Boolean value to indicate whether the grid values can be edited - Optional
	EditOptions GridEditOptions // Edit options, whether the editable is an input field, a single select or a multiselect field - Optional
}
