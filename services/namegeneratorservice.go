package services

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

var SaxonMaleNamesFilename string = "./data/saxon-male-names.txt"
var SaxonFemaleNamesFilename string = "./data/saxon-female-names.txt"

var SaxonMaleNames []string = []string{}
var SaxonFemaleNames []string = []string{}

var rnd *rand.Rand

func InitNameGeneratorService() error {
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

	names, err := readFileIntoArray(SaxonMaleNamesFilename)
	if err != nil {
		return err
	}

	SaxonMaleNames = names

	names, err = readFileIntoArray(SaxonFemaleNamesFilename)
	if err != nil {
		return err
	}

	SaxonFemaleNames = names

	return nil
}

func isDuplicate(name string, names []string) bool {
	result := false

	for i := 0; i < len(names); i++ {
		result = name == names[i]
		if result {
			break
		}
	}

	return result
}

func generateListOfNames(count int, list []string) []string {
	result := make([]string, count)

	for n := 0; n < count; n++ {
		for {
			name := generateName(list)
			if !isDuplicate(name, result) {
				result[n] = name
				break
			}
		}
	}

	return result
}

func GenerateListSaxonMaleNames(count int) []string {
	return generateListOfNames(count, SaxonMaleNames)
}

func GenerateListSaxonFemaleNames(count int) []string {
	return generateListOfNames(count, SaxonFemaleNames)
}

func generateName(list []string) string {
	return list[rnd.Intn(len(list))]
}

func readFileIntoArray(filename string) ([]string, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file into a slice of strings
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
