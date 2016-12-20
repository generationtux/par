package main

import "testing"
import "fmt"
import "github.com/stretchr/testify/assert"

func TestExample(t *testing.T) {
	fileNames := []string{"testfiles/data1.yml", "testfiles/data2.yml"}
	ymlPath := []string{"data", ".env"}
	returnedData := ParseFileData(fileNames, ymlPath)
	fmt.Println(returnedData)
	expectedValues := [][]string{}

	// These are the first two rows.
	row1 := []string{"a", "b"}
	row2 := []string{"a", "b"}

	// Append each row to the two-dimensional slice.
	expectedValues = append(expectedValues, row1)
	expectedValues = append(expectedValues, row2)
	assert.Equal(t, expectedValues, returnedData)
}
