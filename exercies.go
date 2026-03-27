package main

func reverArrayInPlace(arr []int) []int {
	n := len(arr)-1
	for i := range n {
		arr[n-i], arr[i] = arr[i], arr[n-i]
	}
	return arr
}
func reverArrayInPlaceTest(arr []int) {
	logger.Info("Reverse Array in-place")
	reverArrayInPlace(arr[:])
	logger.Warn("After Change:",arr)
}
func findSmallestAndLargest(arr []int) (int, int) {
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
func findSmallestAndLargestTest(arr []int) {
	logger.Info("Find the smallest and largest from the given Array")
	s, l := findSmallestAndLargest(arr[:])
	logger.Warnf("Smallest: %d; Largest: %d",s,l)
}

func ExerciesArraysAndSlices() {
	list := [...]int{11,53,42,5,72,4,69}
	logger.Debug("Original Array:",list)
	reverArrayInPlaceTest(list[:]);
	findSmallestAndLargestTest(list[:])
}