package main

type Item struct{
	name string
	value int
	weight int
}

type Knapsack struct{
	capacity int
	items []Item // items in knapsack
}

// define value for knapsack capacity
func (k *Knapsack) setCapacity(capacity int){
	k.capacity = capacity
}

// append item struct to knapsack item list
func (k *Knapsack) addItem(item Item){
	k.items = append(k.items, item)
}

// remove range of items in a knapsack. Used in the recursive approach
func (k *Knapsack) removeItems(start int, end int){
	var newList []Item

	for i := 0; i<len(k.items); i++ {
		if i >= start && i < end{
			continue
		}

		newList = append(newList, k.items[i])
	}

	k.items = newList
}