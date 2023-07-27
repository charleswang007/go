package main

import (
	"fmt"
)

func myMessage() {
	fmt.Println("I just got executed!")
}

func familyName(fname string, age int) {
	fmt.Println("Hello", age, "year old", fname, "Refsnes")
}

func myFunction(x int, y int) (result int) {
	result = x + y
	return result
}

func factorial_recursion(x float64) (y float64) {
	if x > 0 {
		y = x * factorial_recursion(x-1)
	} else {
		y = 1
	}
	return
}

func main() {
	var student1 string = "John" // type is string
	var student2 = "Jane"        // type is inferred
	x := 2                       // type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	var a, b, c, d int = 1, 3, 5, 7 // declare multiple variables in the same line

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	const ( // Multiple Constants Declaration
		A int = 1
		B     = 3.14
		C     = "Hi!"
	)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)

	var i1, j1 string = "Hello", "World" // The Println() Function

	fmt.Println(i1, j1)

	var i2 string = "Hello" // The Println() Function
	var j2 int = 15

	fmt.Printf("i2 has value: %v and type: %T\n", i2, i2)
	fmt.Printf("j2 has value: %v and type: %T", j2, j2)

	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"} // Array
	var arr1 = [...]int{1, 2, 3}

	fmt.Println(cars)
	fmt.Println(arr1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"} // Slice
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)

	x1 := 20 // if
	x2 := 18
	if x1 > x2 {
		fmt.Println("x1 is greater than x2")
	}

	time := 20 // if-else
	if time < 18 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}

	for i := 0; i <= 100; i += 10 { // loop
		fmt.Println(i)
	}

	fruits := [3]string{"apple", "orange", "banana"}
	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}

	myMessage() // functions

	familyName("Liam", 3)
	familyName("Jenny", 14)
	familyName("Anja", 30)

	fmt.Println(myFunction(1, 2))

	fmt.Println(factorial_recursion(4)) // recursion
}
