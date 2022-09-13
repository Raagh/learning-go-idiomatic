package main

import "fmt"

func main() {
	// var x [3]int
	// var x = [3]int{10, 20, 30}
	// var x = [12]int{1, 5: 4, 6, 10: 100, 15} // postionn 5: value 4
	var x = [...]int{10, 20, 30}
	var y = [3]int{10, 20, 30}

	fmt.Println(x[0])
	fmt.Println(x == y)
	fmt.Println(len(x))
}
