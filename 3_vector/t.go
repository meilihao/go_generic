package main

import (
	"fmt"
)

func printSlice[T any](s []T) {
	fmt.Printf("%T\n",s)

	for _, v := range s {
		fmt.Printf("%v ", v)
	}
	fmt.Print("\n")
}

type vector[T any] []T

type Node struct{
	Id int
}

func main() {
	// Compiling Error
	// Cannot use generic type vector[T interface{}] without instantiation
	//vs0 := vector{1,2,3,4,5}

	vs := vector[int]{5,4,2,1}
	fmt.Printf("%T\n",vs)
	printSlice(vs)

	vs2 := vector[string]{"haha", "hehe"}
	fmt.Printf("%T\n",vs2)
	printSlice(vs2)

	vs3 := vector[*Node]{&Node{}, &Node{}}
	fmt.Printf("%T\n",vs3)
	printSlice(vs3)
}