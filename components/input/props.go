package input

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
