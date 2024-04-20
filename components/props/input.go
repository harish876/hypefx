package props

const (
	InputTypeText     = "text"
	InputTypePassword = "password"
	InputTypeNumber   = "number"
	InputTypeEmail    = "email"
	InputTypeCheckbox = "checkbox"
)

type InputProps struct {
	Id       string
	Typ      string
	Name     string
	Value    string
	Label    string
	Class    string
	Required bool
	Pattern  string
}

type SelectOption struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type SelectProps struct {
	Id           string
	Name         string
	Class        string
	DefaultValue SelectOption
	Options      []SelectOption
}
