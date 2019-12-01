package main

import (
	"fmt"
)

// Persona is defined
type Persona struct {
	speed int
	id int

}


func main() {
	p := Persona{5, 01}
	fmt.Println("Hello Go")
	fmt.Println(p.speed)
}