package main

import (
	"fmt"
	"sync"
	"time"
)

/***** Variables *****/
var items []Item // items given in input file
var totalValue int // total value in the optimal solution
var inputCapacity int // knapsack capcity on last line of input file
var wg sync.WaitGroup // waitgroup variable - threads added in func main() and used in ConcurrentSolutions..go

var knapsack1, knapsack2, knapsack3, knapsack4, knapsack5, knapsack6, knapsack7, knapsack8 Knapsack


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

func Max(a int, b int) int{
	if a > b {
		return a
	}
	return b
}

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
	inputProcess := getInput()

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


	wg.Add(8) // THREADS

	eightThreadSolution(weights, values)

	//FOR 2 THREADS
	// finalKnapsack := evaluateKnapsacks([]Knapsack{knapsack1, knapsack2})
	//FOR 4 THREADS
	// finalKnapsack := evaluateKnapsacks([]Knapsack{knapsack1, knapsack2, knapsack3, knapsack4})
	// FOR 8 THREADS
	finalKnapsack := evaluateKnapsacks([]Knapsack{knapsack1, knapsack2, knapsack3, knapsack4, knapsack5, knapsack6, knapsack7, knapsack8})


	time.Sleep(1 * time.Microsecond) // slow down program to avoid "0s" being logged - MONOTONIC CLOCK *******************************************************
	end := time.Now();
	fmt.Printf("Total runtime: %s\n", end.Sub(start))


	writeOutput(finalKnapsack)
}