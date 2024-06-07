package commands

import (
	"fmt"
	"testing"
)

func TestUpsertConfig(t *testing.T) {
	key := "foo"
	value := "bar"

	if err := UpsertConfig(key, value); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Config updated successfully.")
	}
}

func TestDeleteConfig(t *testing.T) {
	key := "foo"

	if err := DeleteConfig(key); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Config deleted successfully.")
	}
}
