package main

import "fmt"

//hello world
func PrintHelloWorld() string {
	fmt.Println("Hello World!")
	return "Hello World!"
}

//for loop example
func ListSlice[T any](slice []T) {
	for i := 0; i < len(slice)-1; i++ {
		fmt.Print(slice[i])
	}
}

//example of structure with []int member
type structure struct {
	slice []int
}

//function that appends value into slice
func (s *structure) Append(value int) {
	s.slice = append(s.slice, value)
}

//function that returns value of an index in slice
func (s structure) GetValue(index int) int {
	return s.slice[index]
}

//function that returns size (len) of a slice
func (s structure) GetSize() int {
	return len(s.slice)

}

//function that swaps two members of a slice
func (s *structure) Swap(index0 int, index1 int) {
	var temp int

	temp = s.slice[index1]
	s.slice[index1] = s.slice[index0]
	s.slice[index0] = temp

}

//bubble sort example
func (s *structure) BubbleSort() {
	for z := 0; z < len(s.slice)-1; z++ {
		for i := 0; i < len(s.slice)-1; i++ {
			if s.slice[i] > s.slice[i+1] {
				var temp int
				temp = s.slice[i]
				s.slice[i] = s.slice[i+1]
				s.slice[i+1] = temp
			}
		}
	}
}

//selection sort example
func (s *structure) SelectionSort() {
	for i := 0; i < len(s.slice); i++ {
		var min int
		min = i
		for j := i + 1; j < len(s.slice)-1; j++ {
			if s.slice[j] < s.slice[min] {
				min = j
			}
		}
		s.Swap(i, min)
	}
}

//insertion sort example
func (s *structure) Insertionsort() {
	for i := 1; i < len(s.slice); i++ {
		j := i
		for j > 0 {
			if s.slice[j-1] > s.slice[j] {
				s.slice[j-1], s.slice[j] = s.slice[j], s.slice[j-1]
			}
			j = j - 1
		}
	}
}

//example of recursion
func (s *structure) SortwithRecursion() {
	s.slice = RecursionSort(s.slice)
}

func RecursionSort(unsortedSlice []int) []int {
	if len(unsortedSlice) == 0 {
		return []int{}
	}

	var smallestValue int = unsortedSlice[0]
	for i := 0; i < len(unsortedSlice); i++ {
		if unsortedSlice[i] < smallestValue {
			smallestValue = unsortedSlice[i]
		}
	}

	for i := len(unsortedSlice) - 1; i >= 0; i-- {
		if unsortedSlice[i] == smallestValue {
			unsortedSlice = append(unsortedSlice[0:i], unsortedSlice[i+1:len(unsortedSlice)]...)
		}
	}
	var sortedArray []int
	sortedArray = append(sortedArray, smallestValue)
	sortedArray = append(sortedArray, RecursionSort(unsortedSlice)...)
	return sortedArray
}

//example of linear search of sorted slice
func (s structure) LinearSearch(UserNumber int) (index int, found bool) {
	for i := 0; i < len(s.slice); i++ {
		if UserNumber == s.slice[i] {
			return i, true
		}
	}
	return 0, false
}

//example of binary search of sorted slice
func (s structure) BinarySearch(UserNumber int) (mid int, found bool) {
	low := 0
	high := len(s.slice) - 1

	for {
		mid = (low + high) / 2
		if UserNumber == s.slice[mid] {
			return mid, true
		}
		if UserNumber > s.slice[mid] {
			low = mid + 1
		}
		if UserNumber < s.slice[mid] {
			high = mid - 1
		}
		if low > high {
			break
		}
	}
	return 0, false
}
