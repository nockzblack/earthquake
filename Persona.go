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
	mux          *sync.Mutex
}

func newNode(exit bool, border bool, steps int) *Node {
	node := &Node{}
	node.left = nil
	node.right = nil
	node.back = nil
	node.front = nil
	node.nextHop = nil
	node.isExit = exit
	node.isBoder = border
	node.stepsToExtit = steps
	node.mux = new(sync.Mutex)
	return node
}

// Persona is defined
type Persona struct {
	speed  int
	id     int
	alivee bool
	pos    *Node
}

func (p *Persona) setPos(node *Node, wg *sync.WaitGroup) {
	p.pos = node
	fmt.Printf("In setPos from Persona %d\n", p.id)
	p.pos.mux.Lock()
	fmt.Printf("Person %d has locked the node\n", p.id)
	time.Sleep(2 * time.Second)
	defer p.pos.mux.Unlock()
	fmt.Printf("Person %d has unlocked the node\n", p.id)
	defer wg.Done()
}

func (p *Persona) walk(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Person %d has started to move\n", p.id)
	//p.pos.nextHop
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
	n := newNode(false, false, 10)

	for index := 1; index < 10; index++ {
		p := Persona{index, index, true, nil}
		fmt.Printf("Person %d has been created\n", index)
		go p.setPos(n, &wg)
		//go p.walk(&wg)
		wg.Add(1)
	}

	wg.Wait()

}
