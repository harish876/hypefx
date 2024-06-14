package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//TODO: handle errors gracefully and logging it

func ReplaceFileContent(filePath, oldValue, newValue string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Read the file line by line
	var modifiedLines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Replace the string value
		newLine := strings.ReplaceAll(line, oldValue, newValue)
		modifiedLines = append(modifiedLines, newLine)
	}

	if err := scanner.Err(); err != nil {
		//fmt.Printf("Error scanning file: %v\n", err)
		return
	}

	// Close the file before writing
	file.Close()

	// Open the file again for writing
	file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		//fmt.Printf("Error opening file for writing: %v\n", err)
		return
	}
	defer file.Close()

	// Write the modified content back to the file
	writer := bufio.NewWriter(file)
	for _, line := range modifiedLines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			//fmt.Printf("Error writing to file: %v\n", err)
			return
		}
	}
	writer.Flush()
}
