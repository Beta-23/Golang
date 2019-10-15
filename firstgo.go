package main

// This will import the format library
import (
				
				"fmt"
				"bytes"
				"math"
				"math/rand"
				"time"
				)

// main ()is the only function that will run in the program
func main() {
	fmt.Println("=============================") 
	fmt.Println("[ Welcome To GO programing! ]") 
	fmt.Println("=============================")
	fmt.Println()
	l := NewLife(40, 15)
	for i := 0; i < 25; i++ {
		l.Step()
		fmt.Print("\x0c", l) // Clear screen and print field.
		time.Sleep(time.Second / 30)
	}
	fmt.Println("===============================")
  fmt.Println("FIBONACCI NUMBERS!")
  fmt.Println("===============================")
  f := fib()
	// Function calls are evaluated left-to-right.
	fmt.Println(f(), f(), f(), f(), f())
	fmt.Println()
	fmt.Println("===============================")
  fmt.Println("LETS LEARN SOMETHING NEW TODAY!")
  fmt.Println("===============================")
	fmt.Println()
	cal()
	numgen()
	
	// The := short assignment statement can be used in place of a var declaration with implicit (no) type
	num1, num2 := 8.5, 6.1
	fmt.Println("4. The sum of", num1, "and", num2, "is")
	fmt.Println(add(num1,num2))
	w1, w2 := "Hi", "there, way to Go!"
	fmt.Println("----------------------")
	fmt.Println(multiple(w1,w2))
	fmt.Println("----------------------")
}

// fib returns a function that returns
// successive Fibonacci numbers.
func fib() func() int {
	a, b := 0, 25
	return func() int {
		a, b = b, a+b
		return a
	}
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

// Adding function with datatype of float and precision of 64bits
func add(x,y float64) float64 {
	return x+y
}

// Adding function with multiple stings parameters
func multiple(a,b string) (string, string) {
	return a,b
}

// Field represents a two-dimensional field of cells.
type Field struct {
	s    [][]bool
	w, h int
}

// NewField returns an empty field of the specified width and height.
func NewField(w, h int) *Field {
	s := make([][]bool, h)
	for i := range s {
		s[i] = make([]bool, w)
	}
	return &Field{s: s, w: w, h: h}
}

// Set sets the state of the specified cell to the given value.
func (f *Field) Set(x, y int, b bool) {
	f.s[y][x] = b
}

// Alive reports whether the specified cell is alive.
// If the x or y coordinates are outside the field boundaries they are wrapped
// toroidally. For instance, an x value of -1 is treated as width-1.
func (f *Field) Alive(x, y int) bool {
	x += f.w
	x %= f.w
	y += f.h
	y %= f.h
	return f.s[y][x]
}

// Next returns the state of the specified cell at the next time step.
func (f *Field) Next(x, y int) bool {
	// Count the adjacent cells that are alive.
	alive := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (j != 0 || i != 0) && f.Alive(x+i, y+j) {
				alive++
			}
		}
	}
	// Return next state according to the game rules:
	//   exactly 3 neighbors: on,
	//   exactly 2 neighbors: maintain current state,
	//   otherwise: off.
	return alive == 3 || alive == 2 && f.Alive(x, y)
}

// Life stores the state of a round of Conway's Game of Life.
type Life struct {
	a, b *Field
	w, h int
}

// NewLife returns a new Life game state with a random initial state.
func NewLife(w, h int) *Life {
	a := NewField(w, h)
	for i := 0; i < (w * h / 4); i++ {
		a.Set(rand.Intn(w), rand.Intn(h), true)
	}
	return &Life{
		a: a, b: NewField(w, h),
		w: w, h: h,
	}
}

// Step advances the game by one instant, recomputing and updating all cells.
func (l *Life) Step() {
	// Update the state of the next field (b) from the current field (a).
	for y := 0; y < l.h; y++ {
		for x := 0; x < l.w; x++ {
			l.b.Set(x, y, l.a.Next(x, y))
		}
	}
	// Swap fields a and b.
	l.a, l.b = l.b, l.a
}

// String returns the game board as a string.
func (l *Life) String() string {
	var buf bytes.Buffer
	for y := 0; y < l.h; y++ {
		for x := 0; x < l.w; x++ {
			b := byte(' ')
			if l.a.Alive(x, y) {
				b = '*'
			}
			buf.WriteByte(b)
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}



