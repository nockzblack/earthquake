package main

import (
	"fmt"
	"sync"
	"time"
)

// MyNode is the class for the path/map
type MyNode struct {
	left         *MyNode
	right        *MyNode
	back         *MyNode
	front        *MyNode
	nextHop      *MyNode
	isExit       bool
	isBoder      bool
	stepsToExtit int
	name		 int
	mux          *sync.Mutex
}

func newMyNode(exit bool, border bool, steps int, name int) *MyNode {
	auxNode := &MyNode{}
	auxNode.left = nil
	auxNode.right = nil
	auxNode.back = nil
	auxNode.front = nil
	auxNode.nextHop = nil
	auxNode.isExit = exit
	auxNode.isBoder = border
	auxNode.stepsToExtit = steps
	auxNode.name = name
	auxNode.mux = new(sync.Mutex)
	return auxNode
}

// Persona is defined
type Persona struct {
	speed  float32
	id     int
	isDeath bool
	pos    *MyNode
}

func (p *Persona) setPos(MyNode *MyNode) {
	p.pos = MyNode
	p.pos.mux.Lock()
	fmt.Printf("Person %d has locked the MyNode %d\n", p.id, p.pos.name)
}

func (p *Persona) walk(wg *sync.WaitGroup, exit chan int) {
	for  {
		select {
		case <-exit:
			fmt.Printf("Person %d has died\n", p.id)
			defer wg.Done()
			return 
			
		default:
			if (p.pos.isExit != true) {
				//fmt.Printf("Person %d is in MyNode %d\n", p.id, p.pos.name)
				//fmt.Printf("Person %d has started to move to %d\n", p.id, p.pos.nextHop.name)
		
				p.pos.nextHop.mux.Lock() // move right foot
				time.Sleep(time.Duration(p.speed) * time.Second)
				p.pos.mux.Unlock() // set free the origin MyNode
				p.pos = p.pos.nextHop
				//fmt.Printf("Person %d just moved to MyNode %d\n", p.id, p.pos.name)
			} else {
				fmt.Printf("Person %d is out of danger\n", p.id)
				p.pos.mux.Unlock()
				p.isDeath = false;
				defer wg.Done()
				return
			}
		}
	}
	
	
	
}

func main() {
	//p := Persona{5, 1, true, nil}
	fmt.Println("Hello Go")
	//fmt.Println(p.speed)

	var wg sync.WaitGroup
	var myMap [100]*MyNode

	// making MyNodes
	for index := 0; index < len(myMap); index++ {
		myMap[index] = newMyNode(false, false, 10, index)
		
	}

	// making map -> next hope
	var next = myMap[len(myMap)-1]
	for index := len(myMap)-2; index >= 0; index-- {
		myMap[index].nextHop = next
		next = myMap[index]
	}
	myMap[len(myMap)-1].isExit = true;
	
	// making people
	var people [10]Persona
	for index := 0; index < len(people); index++ {
		people[index] = Persona{float32(1), index, true, nil}
	}

	people[0].setPos(myMap[0])
	people[1].setPos(myMap[8])
	people[2].setPos(myMap[50])
	people[3].setPos(myMap[1])
	people[4].setPos(myMap[96])
	people[5].setPos(myMap[74])
	people[6].setPos(myMap[16])
	people[7].setPos(myMap[99])
	people[8].setPos(myMap[41])
	people[9].setPos(myMap[14])

	exit := make(chan int)

	go func() {
		timer := time.NewTimer(5500 * time.Millisecond)
		//fmt.Println("Timer hasnt' done")
		<-timer.C
		fmt.Println("Timer done")
		close(exit)
		defer wg.Done()
	}()
	wg.Add(1)
	
	for index := 0; index < len(people); index++ {
		go people[index].walk(&wg, exit)
		wg.Add(1)
	}

	
	


	
	wg.Wait()

}
