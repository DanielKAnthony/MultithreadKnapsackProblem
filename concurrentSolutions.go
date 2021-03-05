package main

func twoThreadSolution(weights []int, values []int){
	last := len(items) - 1

	go func() {
		defer wg.Done()
		// add D
		knapsack1 = Knapsack{capacity: inputCapacity}
		bruteforceSolve(knapsack1.capacity - weights[last], weights[:last], values[:last], &knapsack1)
		
		knapsack1.addItem(items[last])
	}()
	
	go func() {
		defer wg.Done()
		// skip D
		knapsack2 = Knapsack{capacity: inputCapacity}
		bruteforceSolve(knapsack2.capacity, weights[:last], values[:last], &knapsack2)
	}()
	
	wg.Wait()
}

func fourThreadSolution(weights []int, values []int){
	last := len(items) - 1

	go func(){
		defer wg.Done()
		// skip c and d
		knapsack1 = Knapsack{capacity: inputCapacity}
		bruteforceSolve(knapsack1.capacity, weights[:last - 1], values[:last - 1], &knapsack1)
	}()

	go func(){
		defer wg.Done()
		// add c, skip d
		knapsack2 = Knapsack{capacity: inputCapacity}
		bruteforceSolve(knapsack2.capacity - weights[last - 1], weights[:last - 1], values[:last - 1], &knapsack2)

		knapsack2.addItem(items[last - 1])
	}()

	go func(){
		defer wg.Done()
		// skip c, add d
		knapsack3 = Knapsack{capacity: inputCapacity}
		bruteforceSolve(knapsack3.capacity - weights[last], weights[:last - 1], values[:last - 1], &knapsack3)
		
		knapsack3.addItem(items[last])
	}()

	go func(){
		defer wg.Done()
		// add c and d
		knapsack4 = Knapsack{capacity: inputCapacity}
		bruteforceSolve(knapsack4.capacity - weights[last - 1] - weights[last], weights[:last - 1], values[:last - 1], &knapsack4)
		
		knapsack4.addItem(items[last - 1])
		knapsack4.addItem(items[last])
	}()

	wg.Wait()
}

func eightThreadSolution(weights []int, values []int){
	last := len(items) - 1

	go func(){
		defer wg.Done()
		// skip b, skip c, skip d
		knapsack1 = Knapsack{capacity: inputCapacity}
		bruteforceSolve(knapsack1.capacity, weights[:last - 2], values[:last - 2], &knapsack1)
	}()

	go func(){
		defer wg.Done()
		// skip b, skip c, add d
		knapsack2 = Knapsack{capacity: inputCapacity}
		bruteforceSolve(knapsack2.capacity - weights[last], weights[:last - 2], values[:last - 2], &knapsack2)

		knapsack2.addItem(items[last])
	}()

	go func(){
		defer wg.Done()
		// skip b, add c, skip d
		knapsack3 = Knapsack{capacity: inputCapacity}
		bruteforceSolve(knapsack3.capacity - weights[last - 1], weights[:last - 2], values[:last - 2], &knapsack3)

		knapsack3.addItem(items[last - 1])
	}()

	go func(){
		defer wg.Done()
		// skip b, add c, add d
		knapsack4 = Knapsack{capacity: inputCapacity}
		bruteforceSolve(knapsack4.capacity - weights[last - 1] - weights[last], weights[:last - 2], values[:last - 2], &knapsack4)

		knapsack4.addItem(items[last - 1])
		knapsack4.addItem(items[last])
	}()

	go func(){
		defer wg.Done()
		// add b, skip c, skip d
		knapsack5 = Knapsack{capacity: inputCapacity}
		bruteforceSolve(knapsack5.capacity - weights[last - 2], weights[:last - 2], values[:last - 2], &knapsack5)

		knapsack5.addItem(items[last - 2])
	}()

	go func(){
		defer wg.Done()
		// add b, skip c, add d
		knapsack6 = Knapsack{capacity: inputCapacity}
		bruteforceSolve(knapsack6.capacity - weights[last - 2] - weights[last], weights[:last - 2], values[:last - 2], &knapsack6)

		knapsack6.addItem(items[last - 2])
		knapsack6.addItem(items[last])
	}()

	go func(){
		defer wg.Done()
		// add b, add c, skip d
		knapsack7 = Knapsack{capacity: inputCapacity}
		bruteforceSolve(knapsack7.capacity - weights[last - 1] - weights[last - 2], weights[:last - 2], values[:last - 2], &knapsack7)

		knapsack7.addItem(items[last - 2])
		knapsack7.addItem(items[last - 1])
	}()

	go func(){
		defer wg.Done()
		// add b, add c, add d
		knapsack8 = Knapsack{capacity: inputCapacity}
		bruteforceSolve(knapsack8.capacity - weights[last - 2] - weights[last - 1] - weights[last], weights[:last - 2], values[:last - 2], &knapsack8)

		knapsack8.addItem(items[last - 2])
		knapsack8.addItem(items[last - 1])
		knapsack8.addItem(items[last])
	}()

	wg.Wait()	
}