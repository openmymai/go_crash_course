package main

import "fmt"

func main() {
	ids := []int{12, 42, 33, 43, 54}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids { // _ is iterator, but don't use it
		sum += id
	}
	fmt.Println(sum)

	// Range with maps
	emails := map[string]string{"Bob": "bob@gmail.com", "Mark": "mark@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s ", k, v)
	}
	fmt.Printf("\n")

	for k := range emails {
		fmt.Printf("Name: %s\n", k)
	}
}