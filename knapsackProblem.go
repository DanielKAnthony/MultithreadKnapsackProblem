package main

import (
	"fmt"
	"sync"
	"time"
)

var items []Item // items given in input file
var inputCapacity int // knapsack capcity on last line of input file
var wg sync.WaitGroup // waitgroup variable - threads added in func main() and used in concurrentSolutions.go

//Knapsacks to be populated by the threads generated in concurrentSolutions.go
var knapsack1, knapsack2, knapsack3, knapsack4, knapsack5, knapsack6, knapsack7, knapsack8, 
knapsack9, knapsack10, knapsack11, knapsack12, knapsack13, knapsack14, knapsack15, knapsack16 Knapsack

// A variant of the recursive approach to the knapsack problem.
func bruteforceSolve(W int, wt []int, val []int, knapsack *Knapsack) int{

	if (len(wt) == 0 || W == 0) {
		return 0 
	}
	last := len(wt)-1

	if wt[last] > W { 
		return bruteforceSolve(W, wt[:last], val[:last], knapsack)	 

	} else {
		currPacked := len(knapsack.items)
		packedItems := items[last].value + bruteforceSolve(W - items[last].weight, wt[:last], val[:last], knapsack)

		currOut := len(knapsack.items)
		out := bruteforceSolve(W, wt[:last], val[:last], knapsack)

		if packedItems > out{
			if len(knapsack.items) > currOut{
				knapsack.removeItems(currOut, len(knapsack.items))
			}

			knapsack.addItem(items[last])
			return packedItems

		}else{
			if currOut > currPacked{
				knapsack.removeItems(currPacked, currOut)
			}

			return out
		}
	}
}

// After all valid knapsacks have been populated by the threads, return the knapsack with the greatest value.
func evaluateKnapsacks(kArr []Knapsack) Knapsack{
	var maxIndex, currMax int

	for i := 0; i<len(kArr); i++{
		tempMax := 0
		tempWt := 0

		for j := 0; j<len(kArr[i].items); j++{
			tempMax += kArr[i].items[j].value
			tempWt += kArr[i].items[j].weight
		}

		if (tempWt <= inputCapacity) && tempMax > currMax{
			currMax = tempMax
			maxIndex = i
		}
	}

	return kArr[maxIndex]
}

func main(){
	inputProcess := getInput() // parse input file - see fileIO.go

	// stop program if error occurred
	if inputProcess != nil{
		return
	}

	var weights, values []int
	for _, it := range(items){
		weights = append(weights, it.weight)
		values = append(values, it.value)
	} // get weights and values from input items
	
	start := time.Now();

	wg.Add(16) // THREADS - can change to either 1, 2, 4, 8, or 16

	// based on the number of threads added above, call the corresponding funtion method
	sixteenThreadSolution(weights, values)

	/*** Uncomment one of the lines below corresponding to the number of threads used. Leave all other variants disabled. ***/

	// FOR 1 THREAD
	// finalKnapsack := knapsack1
	//FOR 2 THREADS
	// finalKnapsack := evaluateKnapsacks([]Knapsack{knapsack1, knapsack2})
	//FOR 4 THREADS
	// finalKnapsack := evaluateKnapsacks([]Knapsack{knapsack1, knapsack2, knapsack3, knapsack4})
	// FOR 8 THREADS
	// finalKnapsack := evaluateKnapsacks([]Knapsack{knapsack1, knapsack2, knapsack3, knapsack4, knapsack5, knapsack6, knapsack7, knapsack8})
	// FOR 16 THREADS
	finalKnapsack := evaluateKnapsacks([]Knapsack{knapsack1, knapsack2, knapsack3, knapsack4, knapsack5, knapsack6, knapsack7, knapsack8,
		knapsack9, knapsack10, knapsack11, knapsack12, knapsack13, knapsack14, knapsack15, knapsack16})

	time.Sleep(1 * time.Microsecond) // slow down program to avoid "0s" being logged
	end := time.Now();
	fmt.Printf("Total runtime: %s\n", end.Sub(start))

	writeOutput(finalKnapsack)
}