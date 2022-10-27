package reader

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
)

// Example struct to hold our CSV data
type ExampleInfo struct {
	Site    string
	Size    string
	Profile string
}

// Example struct to hold our TXT data
type ExampleProxy struct {
	IP string
}

func ReadCSV(path string) []ExampleInfo {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Convert raw data to a slice of structs
	exampleInfos := processCSVData(data)

	return exampleInfos
}

func ReadTXT(path string) []ExampleProxy {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var exampleProxies []ExampleProxy
	for scanner.Scan() {
		exampleProxy := ExampleProxy{
			IP: scanner.Text(),
		}
		exampleProxies = append(exampleProxies, exampleProxy)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return exampleProxies
}

func processCSVData(data [][]string) []ExampleInfo {
	var exampleInfo []ExampleInfo
	for i, line := range data {
		// Omit header line
		if i > 0 {
			var rec ExampleInfo
			for j, field := range line {
				if j == 0 {
					rec.Site = field
				} else if j == 1 {
					rec.Size = field
				} else if j == 2 {
					rec.Profile = field
				}
			}
			exampleInfo = append(exampleInfo, rec)
		}
	}
	return exampleInfo
}
