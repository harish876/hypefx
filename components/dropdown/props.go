package dropdown

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
