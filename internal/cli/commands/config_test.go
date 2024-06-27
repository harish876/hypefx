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

func TestGetConfigV1(t *testing.T) {
	config, err := GetAllConfigV1()
	if err != nil {
		t.Fail()
	}
	fmt.Println("AppDir", config.AppDir)
}

func TestUpsertConfigV1(t *testing.T) {
	mp := make(map[string]string)
	mp["key1"] = "value2"
	if err := UpsertConfigV1("nested json", mp); err != nil {
		t.Fail()
	}
}

func TestDeleteConfigV1(t *testing.T) {
	key := "routing"
	if err := DeleteConfig(key); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Config deleted successfully.")
	}
}
