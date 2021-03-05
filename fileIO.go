package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"path"
)

// These vars are stored for later reference to make the output file
var inputFile string // name of input file
var fullPath string // path to current directory

// This function prompts the user to enter the name of an input file, and populates the variables
// found in knapsackProblem.go.
func getInput() (error) {
	fmt.Println("Input file name: ")
	fmt.Scanln(&inputFile)

	var err error
	var file *os.File

	// get absolute path, append file name, and open
	fullPath, err = os.Getwd()
	if err != nil {
		fmt.Println("Error - could not get absolute path")
		return err
	}

	file, err = os.Open(fullPath + "\\" + inputFile)
	if err != nil {
		fmt.Println("Could not open the file. Make sure it is in the same directory.")
		return err
	}

	// read file and populate variables
	defer file.Close()

	scan := bufio.NewScanner(file)
	scan.Scan()

	var itemCount int
	itemCount, err = strconv.Atoi(scan.Text()) // Get number of items
	if err != nil {
		fmt.Println("Error parsing file.")
		return err
	}

	for i := 0; i<itemCount; i++ {
		scan.Scan()
		line := strings.Fields(scan.Text())

		var val, wt int
		val, err = strconv.Atoi(line[1])
		if err != nil {
			fmt.Println("Error parsing file.")
			return err
		}

		wt, err = strconv.Atoi(line[2])
		if err != nil {
			fmt.Println("Error parsing items.")
			return err
		}

		tempItem := Item{
			name: line[0],
			value: val,
			weight: wt,
		}

		items = append(items, tempItem)
	}

	scan.Scan()
	var capacity int
	capacity, err = strconv.Atoi(scan.Text())
	if err != nil {
		fmt.Println("Error parsing capacity.")
		fmt.Println(err)
		return err
	}
	
	inputCapacity = capacity
	return nil
}

// This function access the variables needed to popluate the .sol file.
func writeOutput(k Knapsack){
	var totalValue int

	outputFile := strings.TrimSuffix(inputFile, path.Ext(inputFile)) + ".sol" // strip inputFile extension and append .sol

	filePath, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error - could not output result")
		return
	}

	defer filePath.Close()

	// Write names of included items at the end
	for i := len(k.items) - 1; i>=0; i-- {
		totalValue += k.items[i].value

		if i == 0 {
			defer fmt.Fprintf(filePath, "%s", k.items[i].name)
		}else{
			defer fmt.Fprintf(filePath, " %s", k.items[i].name)
		}
	}

	// Write total value contained in the final knapsack
	fmt.Fprintf(filePath, "%d\n", totalValue)
}