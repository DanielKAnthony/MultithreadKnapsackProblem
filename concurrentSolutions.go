package main

// Function to generate solution on one thread
func oneThreadSolution(weights []int, values []int){
	go func() {
		defer wg.Done()
		knapsack1 = Knapsack{capacity: inputCapacity}
		bruteforceSolve(knapsack1.capacity, weights, values, &knapsack1)
	}()

	wg.Wait()
}

// Function to generate solution on two threads
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

// Function to generate solution on four threads
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

// Function to generate solution on eight threads
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

// Function to generate solution on sixteen threads
func sixteenThreadSolution(weights []int, values []int){
	last := len(items) - 1

	go func(){
		defer wg.Done()
		// skip item n - 3, skip item n - 2, skip item n - 1, skip item n
		knapsack1 = Knapsack{capacity: inputCapacity}
		bruteforceSolve(knapsack1.capacity, weights[:last - 3], values[:last - 3], &knapsack1)
	}()

	go func(){
		defer wg.Done()
		// skip item n - 3, skip item n - 2, skip item n - 1, add item n
		knapsack2 = Knapsack{capacity: inputCapacity}
		bruteforceSolve(knapsack2.capacity - weights[last], weights[:last - 3], values[:last - 3], &knapsack2)

		knapsack2.addItem(items[last])
	}()

	go func(){
		defer wg.Done()
		// skip item n - 3, skip item n - 2, add item n - 1, skip item n
		knapsack3 = Knapsack{capacity: inputCapacity}
		bruteforceSolve(knapsack3.capacity - weights[last - 1], weights[:last - 3], values[:last - 3], &knapsack3)

		knapsack3.addItem(items[last - 1])
	}()

	go func(){
		defer wg.Done()
		// skip item n - 3, skip item n - 2, add item n - 1, add item n
		knapsack4 = Knapsack{capacity: inputCapacity}
		bruteforceSolve(knapsack4.capacity - weights[last - 1] - weights[last], weights[:last - 3], values[:last - 3], &knapsack4)

		knapsack4.addItem(items[last - 1])
		knapsack4.addItem(items[last])
	}()

	go func(){
		defer wg.Done()
		// skip item n - 3, add item n - 2, skip item n - 1, skip item n
		knapsack5 = Knapsack{capacity: inputCapacity}
		bruteforceSolve(knapsack5.capacity - weights[last - 2], weights[:last - 3], values[:last - 3], &knapsack5)

		knapsack5.addItem(items[last - 2])
	}()

	go func(){
		defer wg.Done()
		// skip item n - 3, add item n - 2, skip item n - 1, add item n
		knapsack6 = Knapsack{capacity: inputCapacity}
		bruteforceSolve(knapsack6.capacity - weights[last - 2] - weights[last], weights[:last - 3], values[:last - 3], &knapsack6)

		knapsack6.addItem(items[last - 2])
		knapsack6.addItem(items[last])
	}()

	go func(){
		defer wg.Done()
		// skip item n - 3, add item n - 2, add item n - 1, skip item n
		knapsack7 = Knapsack{capacity: inputCapacity}
		bruteforceSolve(knapsack7.capacity - weights[last - 2] - weights[last - 1], weights[:last - 3], values[:last - 3], &knapsack7)

		knapsack7.addItem(items[last - 2])
		knapsack7.addItem(items[last - 1])
	}()

	go func(){
		defer wg.Done()
		// skip item n - 3, add item n - 2, add item n - 1, add item n
		knapsack8 = Knapsack{capacity: inputCapacity}
		bruteforceSolve(knapsack8.capacity - weights[last - 2] - weights[last - 1] - weights[last], weights[:last - 3], values[:last - 3], &knapsack8)

		knapsack8.addItem(items[last - 2])
		knapsack8.addItem(items[last - 1])
		knapsack8.addItem(items[last])
	}()

	go func(){
		defer wg.Done()
		// add item n - 3, skip item n - 2, skip item n - 1, skip item n
		knapsack9 = Knapsack{capacity: inputCapacity}
		bruteforceSolve(knapsack9.capacity - weights[last - 3], weights[:last - 3], values[:last - 3], &knapsack9)

		knapsack9.addItem(items[last - 3])
	}()

	go func(){
		defer wg.Done()
		// add item n - 3, skip item n - 2, skip item n - 1, add item n
		knapsack10 = Knapsack{capacity: inputCapacity}
		bruteforceSolve(knapsack10.capacity - weights[last - 3] - weights[last], weights[:last - 3], values[:last - 3], &knapsack10)

		knapsack10.addItem(items[last - 3])
		knapsack10.addItem(items[last])
	}()

	go func(){
		defer wg.Done()
		// add item n - 3, skip item n - 2, add item n - 1, skip item n
		knapsack11 = Knapsack{capacity: inputCapacity}
		bruteforceSolve(knapsack11.capacity - weights[last - 3] - weights[last - 1], weights[:last - 3], values[:last - 3], &knapsack11)

		knapsack11.addItem(items[last - 3])
		knapsack11.addItem(items[last - 1])
	}()

	go func(){
		defer wg.Done()
		// add item n - 3, skip item n - 2, add item n - 1, add item n
		knapsack12 = Knapsack{capacity: inputCapacity}
		bruteforceSolve(knapsack12.capacity - weights[last - 3] - weights[last - 1] - weights[last], weights[:last - 3], values[:last - 3], &knapsack12)

		knapsack12.addItem(items[last - 3])
		knapsack12.addItem(items[last - 1])
		knapsack12.addItem(items[last])
	}()

	go func(){
		defer wg.Done()
		// add item n - 3, add item n - 2, skip item n - 1, skip item n
		knapsack13 = Knapsack{capacity: inputCapacity}
		bruteforceSolve(knapsack13.capacity - weights[last - 3] - weights[last - 2], weights[:last - 3], values[:last - 3], &knapsack13)

		knapsack13.addItem(items[last - 3])
		knapsack13.addItem(items[last - 2])
	}()

	go func(){
		defer wg.Done()
		// add item n - 3, add item n - 2, skip item n - 1, add item n
		knapsack14 = Knapsack{capacity: inputCapacity}
		bruteforceSolve(knapsack14.capacity - weights[last - 3] - weights[last - 2] - weights[last], weights[:last - 3], values[:last - 3], &knapsack14)

		knapsack14.addItem(items[last - 3])
		knapsack14.addItem(items[last - 3])
		knapsack14.addItem(items[last])
	}()

	go func(){
		defer wg.Done()
		// add item n - 3, add item n - 2, add item n - 1, skip item n
		knapsack15 = Knapsack{capacity: inputCapacity}
		bruteforceSolve(knapsack15.capacity - weights[last - 3] - weights[last - 2] - weights[last - 1], weights[:last - 3], values[:last - 3], &knapsack15)

		knapsack15.addItem(items[last - 3])
		knapsack15.addItem(items[last - 2])
		knapsack15.addItem(items[last - 1])
	}()

	go func(){
		defer wg.Done()
		// add item n - 3, add item n - 2, add item n - 1, add item n
		knapsack16 = Knapsack{capacity: inputCapacity}
		bruteforceSolve(knapsack16.capacity - weights[last - 3] - weights[last - 2] - weights[last - 1] - weights[last], weights[:last - 3], values[:last - 3], &knapsack16)

		knapsack16.addItem(items[last - 3])
		knapsack16.addItem(items[last - 2])
		knapsack16.addItem(items[last - 1])
		knapsack16.addItem(items[last])
	}()

	wg.Wait()
}