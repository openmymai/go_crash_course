package main

import "fmt"

func main() {
	sum := adder()
	for i:=0;i<10;i++{
		fmt.Println(sum(i))
	}
}

func adder() func(x int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}