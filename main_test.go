package main

import "testing"

import "github.com/stretchr/testify/assert"

func TestParseData(t *testing.T) {
	fileNames := []string{"testfiles/data1.yml", "testfiles/data2.yml"}
	ymlPath := []string{"data", ".env"}
	returnedData := ParseFileData(fileNames, ymlPath)
	expectedValues := [][]string{}

	// These are the first two rows.
	row1 := []string{"a", "b"}
	row2 := []string{"a", "b"}

	// Append each row to the two-dimensional slice.
	expectedValues = append(expectedValues, row1)
	expectedValues = append(expectedValues, row2)
	assert.Equal(t, expectedValues, returnedData)
}

func TestComparingArraysThatAreEqual(t *testing.T) {
	expectedValues := [][]string{}
	fileNames := []string{"testfiles/data1.yml", "testfiles/data2.yml"}

	// These are the first two rows.
	row1 := []string{"a", "b"}
	row2 := []string{"a", "b"}

	expectedValues = append(expectedValues, row1)
	expectedValues = append(expectedValues, row2)

	testArraysAreEqual := CompareEnvArrays(expectedValues, fileNames)
	assert.Equal(t, true, testArraysAreEqual)
}

func TestComparingArraysThatArentEqualByKeys(t *testing.T) {
	expectedValues := [][]string{}

	// These are the first two rows.
	row1 := []string{"a", "c"}
	row2 := []string{"a", "b"}
	fileNames := []string{"testfiles/data1.yml", "testfiles/data2.yml"}

	expectedValues = append(expectedValues, row1)
	expectedValues = append(expectedValues, row2)

	testArraysAreEqual := CompareEnvArrays(expectedValues, fileNames)
	assert.Equal(t, false, testArraysAreEqual)
}

func TestComparingArraysThatArentEqualByLength(t *testing.T) {
	expectedValues := [][]string{}

	// These are the first two rows.
	row1 := []string{"a", "b", "c"}
	row2 := []string{"a", "b"}
	fileNames := []string{"testfiles/data1.yml", "testfiles/data2.yml"}

	expectedValues = append(expectedValues, row1)
	expectedValues = append(expectedValues, row2)

	testArraysAreEqual := CompareEnvArrays(expectedValues, fileNames)
	assert.Equal(t, false, testArraysAreEqual)
}
