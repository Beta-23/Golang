package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.609

// value receivers
type car struct {
	gas_pedal uint16 // min 0 max 65500
	break_pedal uint16
	steering_wheel int16 // -32k -> +32k
	top_speed_kmh float64
}
// Kmh per hour method
func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax)
}

// Miles per hour method
func (c car) mph() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax/kmh_multiple)
}

// pointer receiver
func (c *car) new_top_speed(newspeed float64) {
	c.top_speed_kmh = newspeed
}

// function that does the same as the method, not optimized
func newer_top_speed(c car, speed float64) car {
	c.top_speed_kmh = speed
	return c
}

// main ()is the only function that will run in the program
func main() {
	a_car := car { gas_pedal: 65000, 
								break_pedal: 0,
								steering_wheel: 1256,
								top_speed_kmh: 220.5, 
							}
	fmt.Println("<===========================>")
	fmt.Println(a_car.kmh(),"[ KMH ] --<%%>--")
	fmt.Println("============================") 						
	fmt.Println(a_car.mph(),"[ MPH ] --<@@>--")
	fmt.Println("=============================")
	
	a_car = newer_top_speed(a_car, 500)
	fmt.Println("<===========================>")
	fmt.Println(a_car.kmh(),"[ KMH ] --<%%>--")
	fmt.Println("============================") 						
	fmt.Println(a_car.mph(),"[ MPH ] --<@@>--")
	fmt.Println("=============================")
	
}