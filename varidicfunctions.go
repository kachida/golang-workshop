package main

import "fmt"

//Variadic functions can be called with any number of trailing arguments. For example, fmt.Println is a common variadic function.
//Here’s a function that will take an arbitrary number of ints as arguments.

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0

	//Within the function, the type of nums is equivalent to []int. We can call len(nums), iterate over it with range, etc.

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	
	//Variadic functions can be called in the usual way with individual arguments.
	sum(1,2)
	sum(1,2,3)

	//If you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this.
	nums := []int{1,3,4,5}
	sum(nums...)

	//Another key aspect of functions in Go is their ability to form closures, which we’ll look at next.

}