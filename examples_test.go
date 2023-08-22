package main

import (
	"testing"
)

func TestPrintHelloWorld(t *testing.T) {
	string := PrintHelloWorld()
	if string != "Hello World!" {
		t.Errorf("Can't print Hello World!")
	}
}

func TestPalindromeString(t *testing.T) {
	const word string = "radar"
	var palindrome bool
	palindrome = true
	for i := 0; i < len(word)/2; i++ {
		i_reverse := len(word) - i - 1
		if word[i] != word[i_reverse] {
			palindrome = false
			break
		}
	}
	if palindrome != true {
		t.Errorf("%v isn't a palindrome because it doesn't read the same backwards as forwards", word)
	}
}

func TestPalindromeSlice(t *testing.T) {
	var word []string = []string{"m", "a", "d", "a", "m"}
	var polindrome bool

	polindrome = true
	for i := 0; i < len(word)/2; i++ {
		i_reverse := len(word) - i - 1
		if word[i] != word[i_reverse] {
			polindrome = false
			break
		}
	}
	if polindrome != true {
		t.Errorf("%v isn't a palindrome because it doesn't read the same backwards as forwards", word)
	}
}

func TestSumofSlice(t *testing.T) {
	var numbers []int = []int{1, 2, 8, 9}
	var result int
	for i := 0; i < len(numbers); i++ {
		result += numbers[i]
		t.Log(result)
	}

	const test_sum int = 20
	if result != test_sum {
		t.Errorf("result %v doesn't match the expected %v", result, test_sum)
	}
}
func TestSumofPositiveNumbersinSlice(t *testing.T) {
	var numbers []int = []int{1, -2, 8, -9, 5, -6}
	var result int
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > 0 {
			result += numbers[i]
		}
	}
	const test_sum int = 14
	if result != test_sum {
		t.Errorf("result %v doesn't match the expected %v", result, test_sum)
	}
}

func TestAppend(t *testing.T) {
	var structure structure
	structure.Append(6)
	if structure.GetValue(0) != 6 {
		t.Errorf("Can't append %v into structure", structure.GetValue(0))
	}
}

func TestGetSize(t *testing.T) {
	var structure structure
	var slice []int
	structure.Append(6)
	slice = append(slice, 6)
	structure.GetSize()
	if structure.GetSize() != len(slice) {
		t.Errorf("Structure size %v doesn't match slice size %v", structure.GetSize(), len(slice))
	}
}

func TestSwap(t *testing.T) {
	var structure structure
	structure.Append(3)
	structure.Append(2)
	structure.Append(5)
	structure.Swap(0, 2)
	if structure.GetValue(0) != 5 {
		t.Errorf("First member %v isn't 5", structure.GetValue(0))
	}
	if structure.GetValue(2) != 3 {
		t.Errorf("Second member %v isn't 3", structure.GetValue(1))
	}
}

func TestBubbleSort(t *testing.T) {
	var structure structure
	structure.Append(4)
	structure.Append(2)
	structure.Append(6)
	structure.Append(10)
	structure.BubbleSort()
	if structure.GetValue(0) != 2 {
		t.Errorf("Structure %v isn't sorted", structure)
	}
	if structure.GetValue(1) != 4 {
		t.Errorf("Structure %v isn't sorted", structure)
	}
	if structure.GetValue(2) != 6 {
		t.Errorf("structure %v isn't sorted", structure)
	}
	if structure.GetValue(3) != 10 {
		t.Errorf("structure %v isn't sorted", structure)
	}
}
func TestSelectionSort(t *testing.T) {
	var structure structure
	structure.Append(6)
	structure.Append(2)
	structure.Append(5)
	structure.Append(55)
	structure.Append(1)
	structure.Append(2)
	structure.Append(100)
	structure.Append(200)
	structure.SelectionSort()
	for i := 0; i < len(structure.slice)-1; i++ {
		if structure.slice[i] > structure.slice[i+1] {
			t.Errorf("Structure %v isn't sorted", structure)
		}

	}
}

func TestInsertionsort(t *testing.T) {
	var structure structure
	structure.Append(200)
	structure.Append(2)
	structure.Append(5)
	structure.Append(55)
	structure.Append(1)
	structure.Append(2)
	structure.Append(100)
	structure.Append(200)
	structure.Insertionsort()
	for i := 0; i < len(structure.slice)-1; i++ {
		if structure.slice[i] > structure.slice[i+1] {
			t.Errorf("Structure %v isn't sorted", structure)
		}

	}
}

func TestLinearSearch(t *testing.T) {
	var structure structure
	structure.Append(3)
	structure.Append(7)
	structure.Append(2)
	structure.Append(19)
	structure.SortwithRecursion()
	UserNumber := 2
	UserNumberIndex, found := structure.LinearSearch(UserNumber)
	if found == false {
		t.Errorf("Number %v isn't part of slice %v", UserNumber, structure)
	}
	if UserNumber != structure.GetValue(UserNumberIndex) {
		t.Errorf("Number %v not found on searched index in slice %v", UserNumber, structure)
	}
}

func TestBinarySearch(t *testing.T) {
	var structure structure
	structure.Append(77)
	structure.Append(43)
	structure.Append(90)
	structure.Append(8)
	structure.Append(19)
	structure.Append(99)
	structure.SortwithRecursion()
	UserNumber := 90
	UserNumberIndex, found := structure.BinarySearch(UserNumber)

	if found == false {
		t.Errorf("Number %v isn't part of slice %v", UserNumber, structure)
	}
	if UserNumber != structure.GetValue(UserNumberIndex) {
		t.Errorf("Number %v not found on searched index in slice %v", UserNumber, structure)
	}
}
