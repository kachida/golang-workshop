// range_ iterates over elements in a variety of data
// structures. Let's see how to use `range` with some
// of the data structures we've already learned.

package main

import "fmt"

func main() {

	//Here we use range to sum the numbers in a slice. Arrays work like this too.

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	//range on arrays and slices provides both the index and value for each entry.
	// Above we didn’t need the index, so we ignored it with the blank identifier _. Sometimes we actually want the indexes though.

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index of 3 : ", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for key, value := range kvs {
		fmt.Printf("%s -> %s \n", key, value)
	}

	for k := range kvs {
		fmt.Println("key: ", k)
	}

	for i, c := range "go" {
		fmt.Printf("%d -> %c \n", i, c)
	}

}