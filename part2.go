package main

import "fmt"

//slice = vector(C++) 
type Stack struct { // Define class called stack.
	
	data []int //defined slice of integers datatype called data.

}

func (s *Stack) Push(x int) { //defined function__receiver(related to type stack)__method name with parameter
	s.data = append(s.data, x) // adds value to the end of the slice called data
}

func (s *Stack) Pop() int { //defined function__receiver(related to type Stack)__method no parameters
	n := len(s.data) // n is number of element in the Stack
	val := s.data[n-1] //val is the number of element minus 1.
	s.data = s.data[:n-1]  //redefines the number of element in the Stack by removing the last element.
	return val //returns the new number of elements from the Stack
}

func main() {

	var s Stack //defined list of variable with datatype of last input which are all integers
	            //created a Stack called s

	for i := 0; i < 101; i++ { //for loop up to 100 adding each number to the Stack and displays number Pushed

		s.Push(i)
		fmt.Println(i)
	}

	// for loops 100 times removing all numbers from Stack.
	for i := 0; i < 100; i++ {
		fmt.Println(s.Pop())

	}
	return 

}
