package main

import "fmt"

// Append function givn two lists, add the second list to the end of the first list
func append(array1 []int, array2 []int) []int {
	s1 := make([]int, len(array1))
	// printSlice("s1", s1)
	slLen := len(s1)
	appendedLen := slLen + len(array2)
	// make as much space as needed for the items of the second array
	if appendedLen > cap(s1) { // if necessary, reallocate
		newS1 := make([]int, (appendedLen+1)*2)
		copy(newS1, s1)
		s1 = newS1
		// printSlice("s1", s1)
	}
	// Add items in array2 into array1 starting after the last item in the first array
	s1 = s1[0:appendedLen]
	copy(s1[0:slLen], array1)
	copy(s1[slLen:appendedLen], array2)
	return s1
}

// Concat function combines multiple lists
func concat(array1 []int, listOfArrays [][]int) []int {
	listOfArraysLen := len(listOfArrays)
	for i := 0; i < listOfArraysLen; i++ {
		fmt.Printf("What will be added %v\n", listOfArrays[i])
		array1 = append(array1, listOfArrays[i])
		fmt.Printf("Array after adding %v\n", array1)
	}
	return array1
}

type funcDef func(int) bool

// Filters return a list of items that meet a given predicate
func filter(predFunc funcDef, arr []int) []int {
	var meetsNeeds []int
	for _, item := range arr {
		if predFunc(item) == true {
			// newItem := int(item)
			itemSlice := []int{item}
			meetsNeeds = append(meetsNeeds, itemSlice)
		}
	}
	return meetsNeeds
}

// Length returns the total number of items in a list
func length(arr []int) int {
	var length int
	for _, num := range arr {
		fmt.Println(num)
		length++
	}
	return length
}

type funcApply func(int) int

// Map function given a list returns a result list after applying a given function
func implementationmap(fn funcApply, list []int) []int {
	var resultList []int
	newNumList := []int{0}
	newNum := 0
	for _, num := range list {
		newNum = fn(num)
		newNumList[0] = newNum
		resultList = append(resultList, newNumList)
	}
	return resultList
}

type funcreducer func(int, int) int

// Foldl function given a list, a function, an accumulator returns a list of items reduced by the accumulator value through a function adding the accumlator to the left
func foldl(fn funcreducer, list []int, accumulator int) int {
	for _, num := range list {
		accumulator = fn(accumulator, num)
		fmt.Println("accumulator", accumulator)
	}
	return accumulator
}

// Foldr given a function, a list, and an initial accumulator, fold (reduce) each item into the accumulator from the right using function(item, accumulator)
func foldr(fn funcreducer, list []int, accumulator int) int {
	for _, num := range list {
		accumulator = fn(num, accumulator)
		fmt.Println("accumulator", accumulator)
	}
	return accumulator
}

// Reverse func given a list, return a list with all the original items, but in reversed order
func reverse(list []int) []int {
	fmt.Println("Running reverse function")
	forward := 0
	backward := -1
	temp := 0
	// Chnage strategy because go doesn't have negative indexes like Python for simplicity and prevention of bugs
	for forward < length(list)/2 && backward >= -(length(list)/2) {
		fmt.Println("before temp", temp) // 0
		temp = list[forward]
		fmt.Println("after temp", temp)                         // 1
		fmt.Println("before back in the front ", list[forward]) // 1
		list[forward] = list[length(list+ backward)backward]
		fmt.Println("after back in the front ", list[forward])  // 7
		fmt.Println("before front is now back", list[backward]) // 7
		list[backward] = temp
		fmt.Println("before front is now back", list[backward]) // 1
		// go to next pair of values
		forward++
		backward

	}
	return list

}

// Problem Link: https://exercism.io/my/solutions/7141754fe059438790dd75474ffe65e8

func main() {
	// array1 := []int{1, 2}
	// array2 := []int{2, 3, 4, 5}
	// fmt.Println(append(array1, array2))
	// array1 := []int{1, 2}
	// listOfArrays := [][]int{[]int{3}, []int{}, []int{4, 5, 6}}
	// fmt.Println(concat(array1, listOfArrays))
	// fn := func(n int) bool { return n%2 == 1 }
	// arr := []int{1, 2, 3, 4, 5}
	// fmt.Println(filter(fn, arr))
	// arr := []int{}
	// fmt.Println(length(arr))
	// list := []int{1, 3, 5, 7}
	// fn := func(x int) int { return x + 1 }
	// fmt.Println(implementationmap(fn, list))
	// list := []int{1, 2, 3, 4}
	// fn := func(x, y int) int { return x + y }
	// accumulator := 5
	// fmt.Println(foldl(fn, list, accumulator))
	// list := []int{1, 2, 3, 4}
	// fn := func(x, y int) int { return x + y }
	// accumulator := 5
	// fmt.Println(foldr(fn, list, accumulator))
	list := []int{1, 3, 5, 7}
	fmt.Println(reverse(list))
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
