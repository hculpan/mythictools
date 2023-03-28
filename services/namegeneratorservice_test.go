package services

import "testing"

func TestInitGeneratorService(t *testing.T) {
	SaxonMaleNamesFilename = "../data/saxon-male-names.txt"
	SaxonFemaleNamesFilename = "../data/saxon-female-names.txt"

	err := InitNameGeneratorService()
	if err != nil {
		t.Error(err)
	}

	if len(SaxonMaleNames) != 301 {
		t.Errorf("expected 301 names, got %d names", len(SaxonMaleNames))
	}
}

func TestGenerateSaxonMaleNames(t *testing.T) {
	SaxonMaleNamesFilename = "../data/saxon-male-names.txt"
	SaxonFemaleNamesFilename = "../data/saxon-female-names.txt"

	err := InitNameGeneratorService()
	if err != nil {
		t.Error(err)
	}

	names := GenerateListSaxonMaleNames(10)
	if len(names) != 10 {
		t.Errorf("expected 10 names, got %d", len(names))
	}
}

func TestGenerateSaxonMaleNames100Times(t *testing.T) {
	SaxonMaleNamesFilename = "../data/saxon-male-names.txt"
	SaxonFemaleNamesFilename = "../data/saxon-female-names.txt"

	err := InitNameGeneratorService()
	if err != nil {
		t.Error(err)
	}

	for i := 0; i < 1000; i++ {
		names := GenerateListSaxonMaleNames(10)
		if len(names) != 10 {
			t.Errorf("expected 10 names, got %d", len(names))
		}
		if checkDuplicates(names) {
			msg := "found duplicates in names: ["
			for i, n := range names {
				if i > 0 {
					msg += ", "
				}
				msg += n
			}
			msg += "]"
		}
	}
}

func checkDuplicates(arr []string) bool {
	// Create a map to store the frequency of each element
	freq := make(map[string]int)

	// Loop through the array and update the frequency map
	for _, elem := range arr {
		freq[elem]++
	}

	// Check if any element appears more than once
	for _, count := range freq {
		if count > 1 {
			return true
		}
	}

	// If no element appears more than once, there are no duplicates
	return false
}
