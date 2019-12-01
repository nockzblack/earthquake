package main

import (
	"fmt"
	"sync"
	"time"
)

// Node is the class for the path/map
type Node struct {
	left         *Node
	right        *Node
	back         *Node
	front        *Node
	nextHop      *Node
	isExit       bool
	isBoder      bool
	stepsToExtit int
	name		 int
	mux          *sync.Mutex
}

func newNode(exit bool, border bool, steps int, name int) *Node {
	node := &Node{}
	node.left = nil
	node.right = nil
	node.back = nil
	node.front = nil
	node.nextHop = nil
	node.isExit = exit
	node.isBoder = border
	node.stepsToExtit = steps
	node.name = name
	node.mux = new(sync.Mutex)
	return node
}

// Persona is defined
type Persona struct {
	speed  float32
	id     int
	alivee bool
	pos    *Node
}

func (p *Persona) setPos(node *Node) {
	p.pos = node
	p.pos.mux.Lock()
	fmt.Printf("Person %d has locked the node %d\n", p.id, p.pos.name)
}

func (p *Persona) walk(wg *sync.WaitGroup) {
	
	for (p.pos.isExit != true) {
		fmt.Printf("Person %d is in node %d\n", p.id, p.pos.name)
		fmt.Printf("Person %d has started to move to %d\n", p.id, p.pos.nextHop.name)

		p.pos.nextHop.mux.Lock() // move right foot
		time.Sleep(time.Duration(p.speed) * time.Second)
		p.pos.mux.Unlock() // set free the origin node
		p.pos = p.pos.nextHop
		fmt.Printf("Person %d just moved to node %d\n", p.id, p.pos.name)
	}

	fmt.Printf("Person %d is out of danger\n", p.id)
	p.pos.mux.Unlock()
	//p.pos = nil

	defer wg.Done()
}

func main() {
	//p := Persona{5, 1, true, nil}
	fmt.Println("Hello Go")
	//fmt.Println(p.speed)

	var wg sync.WaitGroup
	var myMap [10]*Node

	// making nodes
	for index := 0; index < 10; index++ {
		myMap[index] = newNode(false, false, 10, index)
		
	}

	// making map -> next hope
	var next = myMap[9]
	for index := 8; index >= 0; index-- {
		myMap[index].nextHop = next
		next = myMap[index]
	}
	myMap[9].isExit = true;
	
	// making people
	var people [3]Persona
	for index := 1; index < 4; index++ {
		people[index-1] = Persona{float32(index/2), index, true, nil}
	}

	people[0].setPos(myMap[0])
	people[1].setPos(myMap[1])

	people[1].speed = 0.2
	people[0].speed = 2

	go people[0].walk(&wg)
	wg.Add(1)
	go people[1].walk(&wg)
	wg.Add(1)


	wg.Wait()

}
