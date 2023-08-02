package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// Read input
	fmt.Print("Enter the names: (Enter empty line to confirm)... ")

	scanner := bufio.NewScanner(os.Stdin)

	var lines []string
	for {
		scanner.Scan()
		line := scanner.Text()

		// break the loop if line is empty
		if len(line) == 0 {
			break
		}
		lines = append(lines, line)
	}

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	// Read csv
	var nameMap map[string]string = CSVFileToMap("names.csv")

	for _, eachLine := range lines {
		// Map input to csv
		var loweredInput = strings.ToLower(eachLine)
		var trimmedInput = strings.Join(strings.Fields(loweredInput), "")
		var sanitisedInput = strings.Replace(trimmedInput, "-", "", -1)
		var output string

		if nameMap[sanitisedInput] == "" {
			output = "Not Found: " + sanitisedInput
		} else {
			output = nameMap[sanitisedInput]
		}

		// Print results
		fmt.Println(output)
	}
}

func CSVFileToMap(filePath string) map[string]string {
	var returnMap = make(map[string]string)

	csvfile, err := os.Open(filePath)
	if err != nil {
		log.Fatalln("Error occured opening csv file", err)
	}

	r := csv.NewReader(csvfile)

	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		returnMap[record[0]] = record[1]
	}

	return returnMap
}
