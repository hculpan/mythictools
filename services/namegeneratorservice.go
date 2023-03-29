package services

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

var SaxonMaleNamesFilename string = "./data/saxon-male-names.txt"
var SaxonFemaleNamesFilename string = "./data/saxon-female-names.txt"

var CelticMaleNamesFilename string = "./data/celtic-male-names.txt"
var CelticFemaleNamesFilename string = "./data/celtic-female-names.txt"

var SaxonMaleNames []string = []string{}
var SaxonFemaleNames []string = []string{}

var CelticMaleNames []string = []string{}
var CelticFemaleNames []string = []string{}

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

	names, err = readFileIntoArray(CelticMaleNamesFilename)
	if err != nil {
		return err
	}

	CelticMaleNames = names

	names, err = readFileIntoArray(CelticFemaleNamesFilename)
	if err != nil {
		return err
	}

	CelticFemaleNames = names

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

func GenerateListCelticMaleNames(count int) []string {
	return generateListOfNames(count, SaxonMaleNames)
}

func GenerateListCelticFemaleNames(count int) []string {
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

var twoSyllables = []string{"Ash", "Birch", "Chelm", "Dor", "Ead", "Fer", "Grim", "Hert", "Irn", "Ken", "Luf", "Mild", "Nun", "Osb", "Padd", "Quen", "Ric", "Sige", "Tun", "Ulf", "Vort", "Wic", "Xen", "Yff", "Zeb"}
var threeSyllables = []string{"bur", "cun", "dun", "ecc", "farn", "gild", "helm", "inn", "keld", "lind", "morc", "nath", "old", "penn", "quar", "ridd", "stan", "thorp", "ulf", "vinn", "wick", "york", "zand"}

func GenerateTownName() string {
	numSyllables := rnd.Intn(2) + 2
	var name string
	for i := 0; i < numSyllables; i++ {
		if i == 0 {
			name += twoSyllables[rnd.Intn(len(twoSyllables))]
		} else {
			name += threeSyllables[rnd.Intn(len(threeSyllables))]
		}
	}
	return name
}

func GenerateTownNames(count int) []string {
	result := []string{}

	for i := 0; i < count; i++ {
		result = append(result, GenerateTownName())
	}

	return result
}
