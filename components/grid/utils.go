package grid

import (
	"fmt"
	"reflect"

	"github.com/harish876/hypefx/components/dropdown"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// getField retrieves the value of a field by name using reflection
func GetField[D any](data D, fieldName string) interface{} {
	// Get the reflect.Value of the struct
	value := reflect.ValueOf(data)

	// Get the field by name
	field := value.FieldByName(fieldName)

	// Check if the field is valid
	if !field.IsValid() {
		fmt.Println("Field not found:", fieldName)
		return ""
	}

	// Convert the field value to string and return
	if field.Kind() == reflect.Slice {
		if field.Len() > 0 {
			var arr = make([]string, 0)
			for i := 0; i < field.Len(); i++ {
				arr = append(arr, fmt.Sprintf("%s", field.Index(i)))
			}
			return arr

		} else {
			return make([]string, 0)
		}
	}

	return field
}

func GetMultiSelectOptions(fields []string) []dropdown.SelectOption {
	var options []dropdown.SelectOption
	for _, field := range fields {
		options = append(options, dropdown.SelectOption{Label: field, Value: field})
	}
	return options
}

func StartCase(s string) string {
	transformer := cases.Title(language.English)
	return transformer.String(s)
}
