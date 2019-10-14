package main

// This will import the format library
import (
				"fmt"
				"math"
				"math/rand"
				"time")

// main ()is the only function that will run in the program
func main() {
	fmt.Println("=============================") 
	fmt.Println("[ Welcome Using GO program! ]") 
	fmt.Println("=============================")
	fmt.Println()
	fmt.Println("===============================")
  fmt.Println("LETS LEARN SOMETHING NEW TODAY!")
  fmt.Println("===============================")
	fmt.Println()
	cal()
	numgen()
	
	var num1 float64 = 8.5
	var num2 float64 = 6.5
	fmt.Println("4. The sum of", num1, "and", num2, "is")
	fmt.Println(add(num1,num2))
}

// This will execute via the main function
func cal() {
	fmt.Println ("1. Pi is", math.Pi)
	fmt.Println ("2. The square root of 36 is",math.Sqrt(36))
}

// This will generate a non-static number between 0-99
func numgen() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println ("3. A number between 0 and 99!",rand.Intn(100))
	fmt.Println()
}

func add(x float64, y float64) float64 {
	return x+y
}


