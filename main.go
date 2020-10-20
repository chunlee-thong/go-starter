package main

import (
	"fmt"
	chunlee "go-starter/src"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"time"

	"rsc.io/quote"
)

func main() {
	//print
	fmt.Println(quote.Go())
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Hello this is my age %g \n", math.Sqrt(7))
	//simple function
	sum := add(10, 23)
	fmt.Println("sum of 10 and 23 is:", sum)

	//multiple return function
	a, b := multipleReturn("apple", "book")
	fmt.Println(a, b)

	//named return value
	fmt.Println(split(17))

	// call function from other package
	mul := chunlee.Multiply(2, 3)
	fmt.Println("Multiplication:", mul)

	//variable
	var x int
	var good, bad bool
	var name, school, phone string
	var age, grade int = 24, 12
	const ME = "CHUNLEE"
	k := 3
	fmt.Println(x, good, bad, name, school, phone)
	fmt.Println(age, grade, k)

	//data type
	var (
		ToBe    bool       = false
		newName string     = "chunlee"
		height  float64    = 1.68
		newAge  int        = 24
		z       complex128 = cmplx.Sqrt(-5 + 12i)
	)
	fmt.Println(newName, height, newAge, ToBe, z)

	//loop
	total := 0

	for i := 0; i < 10; i++ {
		total += i
	}

	//init and post are optional
	//i:=1
	//for ;i<10;{
	//	i+=i
	//}
	fmt.Println(total)

	//control statement
	if total < 50 {
		fmt.Println("total is less than 50")
	} else {
		fmt.Println("total is more than 50")
	}

	if x := 10; x == 10 {
		fmt.Println("X is 10")
	}

	//simple switch
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("run on macos")
	case "linux":
		fmt.Println("run on linux")
	default:
		fmt.Println("run on other")
	}

	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 1:
		fmt.Println("tomorrow")
	case today + 2:
		fmt.Println("in 2 days")
	default:
		fmt.Println("too far away")
	}

	//defer function
	defer fmt.Println("this look like async await")
	fmt.Println("Hello meow")

}

func addition(x int, y int) int {
	return x + y
}

func add(x, y int) int {
	return x + y
}

func multipleReturn(a, b string) (string, string) {
	return "first param: " + a, ", second param: " + b
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
