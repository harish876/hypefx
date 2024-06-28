package commands

import (
	"fmt"
	"testing"
)

func TestSetConfig(t *testing.T) {
	config, _ := GetConfig()
	if err := SetConfig(config); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Config updated successfully.")
	}
}
func TestGetConfig(t *testing.T) {
	m, err := GetConfig()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Config updated successfully.", m)
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
