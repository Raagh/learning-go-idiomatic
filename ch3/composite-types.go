package main

import "fmt"

func main() {
	// var x [3]int
	// var x = [3]int{10, 20, 30}
	// var x = [12]int{1, 5: 4, 6, 10: 100, 15} // postionn 5: value 4
	// var x = [...]int{10, 20, 30}
	// var slice = []int{20, 20, 30}
	// var y = [3]int{10, 20, 30}
	//
	//  slice = append(slice, 32)
	//
	// fmt.Println(x[0])
	// fmt.Println(x == y)
	// fmt.Println(len(x))
	// fmt.Println(len(slice))

	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	y := x[:2]
	z := x[2:]
	fmt.Println(cap(x), cap(y), cap(z))
	x = append(x, 60)
	y = append(y, 30, 40, 50)
	z = append(z, 70)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

	jane := person{}
  jhon := person{ 
    name: "jhon",
    age: 30,
  }

  juan := struct {
    name string
    age int
  }{
    name: "juan",
    age: 17,
  }

	fmt.Println(jane)
	fmt.Println(jhon)
	fmt.Println(juan)
	fmt.Println(juan.name)
}

type person struct {
	name string
	age  int
}

