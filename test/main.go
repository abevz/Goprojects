package main

import (
	"fmt"
	"os"
)

type Server struct {
	Name string
	ID   int
}

// main function  
func main() {
	// js := Server{
	//	Name: "",
	//	ID:   0,
	//}
	a := 2
	b := 2
	m := []string{"aa"}
	c := 0 / len(m)

	fmt.Println("BAR:", os.Getenv("Bar"))
	fmt.Println(c)
	fmt.Println(len(m))
	fmt.Println("a is smaller than b" + fmt.Sprint(c))

	if a < b {
		fmt.Println("a is smaller than b")
	} else if a == 3 {
		fmt.Printf("kkk")
	}

	if a > b {
		fmt.Println("a is larger than b")
	}

	if a == b {
		fmt.Println("a is equal to b")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
