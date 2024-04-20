package props

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

type GridDataRow struct {
	Id       string
	Name     string
	Status   string
	Position string
	Email    string
	Badges   []string
	Img      string
}

type GridData []GridDataRow

type GridPagination struct {
	Current    int
	TotalPages int
	Limit      int
}

type GridContext[C any] struct {
	IdField     string
	Version     int
	Title       string
	Subheading  string
	Description string
	Url         string
	Columns     []C
	PageOptions GridPagination
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
	Options []SelectOption
}
type GridMultiSelectEditProps struct {
	Id                     string // Id for the put request from the data
	Name                   string // Name of the input for the put request i.e the key you recieve from the form value. Ideally this should be equal to the key in the database but can be changed if the parameter is passed
	Options                []SelectOption
	DefaultSelectedOptions []SelectOption
}

type GridEditProps interface{}

type GridEditOptions struct {
	EditType  int
	EditProps GridEditProps
}

type GridColumn struct {
	Typ         int
	Label       string
	Key         string
	Filter      bool // this can be many things todo
	Renderer    string
	Editable    bool
	EditOptions GridEditOptions
}
