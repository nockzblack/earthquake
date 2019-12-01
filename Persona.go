package main

import (
	"fmt"
)


// Nodo is the class for the path/map
type Nodo struct {
	//person Persona

}
// Persona is defined
type Persona struct {
	speed int
	id int
	alivee bool
	currentNode *Nodo
}

func (p *Persona) walk(pos *Nodo) {
	p.currentNode = pos

}


func main() {
	p := Persona{5, 1, true, nil}
	fmt.Println("Hello Go")
	fmt.Println(p.speed)
}