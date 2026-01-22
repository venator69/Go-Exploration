package main

import (
	"fmt"
)

// defining a struct
type Person struct {
	name   string
	age    int
	job    string
	salary int
}

// hello worlding
func main() {
	fmt.Println("Hello World!")
	fmt.Println("")

	// Variable declaration and initialization
	var (
		a int    = 6
		b string = "Hello"
	)
	c, d := 7, "World!"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%s \n\n", d)

	// printf
	var i = 15.5
	var txt = "Hello World!"

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%v%%\n", i)
	fmt.Printf("%T\n", i)

	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n\n", txt)

	// Using arrays (cannot be different types in same array)

	var arr1 = [4]int{1, 2, 3}
	arr2 := [...]int{4, 5, 6, 7, 8}
	arr3 := [5]int{1: 10, 2: 40}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(len(arr3))
	fmt.Println("")

	// Using slices (can be different types in same slice)
	arr4 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr4[2:4]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))
	fmt.Print("\n\n")

	// Append to slice
	myslice1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = append(myslice1, 20, 21)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n\n\n", cap(myslice1))

	// conditional statements
	time := 22
	if time < 10 {
		fmt.Println("Good morning.")
	} else if time < 20 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}

	// switch statement
	day := 4

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	}

	daytype := 6
	switch daytype {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}

	// for loop
	for i := 0; i < 9; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
		if i == 6 {
			break
		}
	}
	fmt.Println("")
	// enumerate over slice
	fruits := []string{"apple", "orange", "banana"}
	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}
	// functions with parameters
	familyName("Liam")
	familyName("Jenny")
	familyName("Anja")

	// functions with mult return values
	fmt.Println(myFunction(5, "Hello"))

	// to ommit return value
	result, _ := myFunction(10, "Hi")
	fmt.Println(result)
	fmt.Println("")

	// using struct
	var pers1 Person
	var pers2 Person

	// Pers1 specification
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000

	// Pers2 specification
	pers2.name = "Cecilie"
	pers2.age = 24
	pers2.job = "Marketing"
	pers2.salary = 4500

	// Access and print Pers1 info
	fmt.Println("Name: ", pers1.name)
	fmt.Println("Age: ", pers1.age)
	fmt.Println("Job: ", pers1.job)
	fmt.Println("Salary: ", pers1.salary)

	// Access and print Pers2 info
	fmt.Println("Name: ", pers2.name)
	fmt.Println("Age: ", pers2.age)
	fmt.Println("Job: ", pers2.job)
	fmt.Println("Salary: ", pers2.salary)

	// using maps
	var x = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	// using make to create map
	y := make(map[string]int)
	y["Oslo"] = 1
	y["Bergen"] = 2
	y["Trondheim"] = 3
	y["Stavanger"] = 4

	delete(x, "year")
	// add elements to map
	x["color"] = "Red" // Adding an element

	// access elements from map
	_, ok4 := x["model"] // Only checking for existing key and not its value
	fmt.Println(ok4)
	// delete element from map

	fmt.Printf("x\t%v\n", x)
	fmt.Printf("y\t%v\n", y)

	// took as reference
	z := x
	z["color"] = "Blue"
	fmt.Printf("x\t%v\n", x)
	fmt.Printf("z\t%v\n", z)

	// iterate over map
	for k, v := range x {
		fmt.Printf("%v : %v, ", k, v)
	}
}

// functions with parameters
func familyName(fname string) {
	fmt.Println("Hello", fname, "Refsnes")
}

// functions with multiple return values
func myFunction(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}
