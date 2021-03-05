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

func (k *Knapsack) setCapacity(capacity int){
	k.capacity = capacity
}

func (k *Knapsack) addItem(item Item){
	k.items = append(k.items, item)
}

func (k *Knapsack) removeItems(start int, end int){
	var newList []Item

	for i := 0; i<len(k.items); i++ {
		if i >= start && i <= end{
			continue
		}

		newList = append(newList, k.items[i])
	}

	k.items = newList
}