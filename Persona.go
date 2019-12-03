package main
/*
import (
	"fmt"
	"sync"
	"time"
)



func main() {

	var wg sync.WaitGroup
	var myMap [100]*Node

	// making MyNodes
	for index := 0; index < len(myMap); index++ {
		myMap[index] = NewNode(false, false, 10)
		
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

*/
