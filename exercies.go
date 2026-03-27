package main

import (
	"fmt"
	"strings"
)

func palindromeFromInput() {
	word := new(string)
	fmt.Print("Enter the word to check if its a Palindrome (press Enter when done): ")
	fmt.Scanf("%s", word)

	wordLen := len(*word)

	fmt.Printf("You entered: %s\n", *word)

	checkForEvenLen := wordLen % 2
	if checkForEvenLen == 0 {
		fmt.Printf("%s is a word with an even number of characters so it is not a Palindrome\n\n", *word)
		return
	} else {
		logger.Warnf("There are %d characters in %s\n", wordLen, *word)
	}
	arr := strings.Split(*word, "")

	var i int = 0
	var j int = len(arr) - 1

	for i < wordLen {
		if arr[j] != arr[i] {
			fmt.Print("Word is not a palindrome")
			return
		}
		j--
		i++
	}
	logger.Info("Word is a Palindrome!\n")
}

/* func ExercisesArraysAndSlices() {
	reverseArrayInPlace := func(arr []int) []int {
	n := len(arr)-1
	for i := range n {
		arr[n-i], arr[i] = arr[i], arr[n-i]
	}
	return arr
	}
	ExecReverseArrayInPlace := func(arr []int) {
		logger.Info("Reverse Array in-place")
		reverseArrayInPlace(arr[:])
		logger.Warn("After Change:",arr)
	}
	findSmallestAndLargest := func(arr []int) (int, int) {
		var (
			smallest int
			largest int
		)
		n := len(arr)-1

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if arr[i] > arr[j] {
					smallest = arr[j]
				}
				if arr[j] > arr[i] {
					largest = arr[j]
				}
			} 
		}
		return smallest, largest
	}
	ExecFindSmallestAndLargest := func(arr []int) {
		logger.Info("Find the smallest and largest from the given Array")
		s, l := findSmallestAndLargest(arr[:])
		logger.Warnf("Smallest: %d; Largest: %d",s,l)
	}
	findPeaks := func(arr []int) int {
		n := len(arr)
		peaks := 0
		// How would you set up a 'range' or 'for' loop to stay 
		// within the bounds so you don't hit a "nil pointer" or "out of bounds" 
		// when checking neighbors?
		for i := 2; i < n-1; i++ {
			// Your logic here: if arr[i] is bigger than left AND right...
			if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
				peaks++
			}
		}
		return peaks
	}
	ExecFindPeaks := func(arr []int) {
		peaks := findPeaks(arr) 
		logger.Warn("The peaks found were:", peaks)
	}
	selectionSort := func(arr []int) {
		n := len(arr)

		// 1. Outer loop: Iterate through the array (except the last element)
		for i := range n - 1 {
			
			// Assume the current position 'i' is the smallest for now
			minIndex := i

			// 2. Inner loop: Start from i+1 and go to the very end
			// This loop looks for a number smaller than arr[minIndex]
			for j := i + 1; j < n; j++ {
				// TODO: If arr[j] is smaller than arr[minIndex], 
				// update minIndex to be j
				if arr[j] < arr[minIndex] {
					minIndex = j
				}
			}

			// 3. Swap: Move the smallest number found to its correct spot (i)
			// TODO: Swap arr[i] and arr[minIndex]
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}

	}
	ExecSelectionSort := func(arr []int) {
		selectionSort(arr)
		logger.Warn("After: ", arr)
	}
	moveZerosToEnd := func(arr []int) {
		writeNext := 0 // This is our "spot tracker"

		// Loop through the array once
		for i := 0; i < len(arr); i++ {
			if arr[i] != 0 {
				// TODO: Move arr[i] into the arr[writeNext] spot
				// TODO: Increment writeNext
				arr[i], arr[writeNext] = arr[writeNext], arr[i]
				writeNext++
			}
		}

		// Now writeNext is sitting at the start of the "gap"
		// TODO: Use one more loop to fill the rest of the array with 0s
		for i := writeNext; i < len(arr); i++ {
			arr[i] = 0
		}
	}
	ExecMoveZerosToEnd := func(arr []int) {
		fmt.Println("Before:", arr)
		moveZerosToEnd(arr)
		fmt.Println("After: ", arr)
	}
	

	list := [...]int{11,53,0,42,5,72,4,69}
	logger.Debug("Original Array:",list)
	ExecReverseArrayInPlace(list[:]);
	ExecFindSmallestAndLargest(list[:])
	ExecFindPeaks(list[:])
	ExecSelectionSort(list[:])
	ExecMoveZerosToEnd(list[:])
} */

func xyz() {
	logger.Print("something")
	logger.Print(sepLine)
}