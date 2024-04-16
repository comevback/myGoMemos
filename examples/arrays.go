package main

import "fmt"

func arrays() {
	original := []int{10, 20, 30, 40, 50}

	// Method 1: Direct slicing
	slice1 := original[0:2] // Creates a slice from index 0 to 1, inclusive.

	// Method 2: Using the append function and a loop
	slice2 := make([]int, 0) // Creating an empty slice
	for i := 0; i < 2; i++ {
		slice2 = append(slice2, original[i])
	}

	// Print both slices
	fmt.Println("Slice 1:", slice1) // Output will be [10 20]
	fmt.Println("Slice 2:", slice2) // Output will be [10 20] as well
}

// two way to create a slice.
