package main

import (
	"fmt"
	"time"
	"sync"
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

func (p *Persona) walk(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Person %d has started to move\n", p.id)

	time.Sleep(time.Duration(p.speed) * time.Second)
	fmt.Printf("Person %d just moved to this plase\n", p.id)
	//fmt.Println("Took this spot")
	//fmt.Printf("Person %d  just realeased this place\n", p.id)

}


func main() {
	//p := Persona{5, 1, true, nil}
	fmt.Println("Hello Go")
	//fmt.Println(p.speed)

	var wg sync.WaitGroup
	for index := 1; index < 11; index++ {
		p := Persona{index,index,true, nil}
		go p.walk(&wg)
		wg.Add(1)
	}

	wg.Wait()

}