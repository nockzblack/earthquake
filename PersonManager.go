package main

import (
	"fmt"
	"sync"
	"time"
)

// PersonManager this struch handle the instatiation and walking
type PersonManager struct {
	wg *sync.WaitGroup
	mapa *Map
	people []Persona
	seconds int
	exit chan int
}


func (pm *PersonManager) initPeople() {
	for index := 0; index < len(pm.people); index++ {
		pm.people[index] = Persona{float32(index/2), index, true, nil}
	}
}



func (pm *PersonManager) setPeoplePos() {
	//var rows =  make([]int, pm.mapa.height)
	//var cols = make([]int, pm.mapa.height)

	//for index := 0; index < len(pm.people); index++ {}

	pm.people[0].setPos(pm.mapa.nodes[4][4])
	pm.people[1].setPos(pm.mapa.nodes[10][4])
	pm.people[2].setPos(pm.mapa.nodes[11][5])
	pm.people[3].setPos(pm.mapa.nodes[7][7])
	pm.people[4].setPos(pm.mapa.nodes[2][7])
}


func (pm *PersonManager) startTimer(timeToFinsh int) {
	pm.wg.Add(1)
	timer := time.NewTimer(time.Duration(timeToFinsh) * time.Second)
	//fmt.Println("Timer hasnt' done")
	<-timer.C
	fmt.Println("Timer done")
	close(pm.exit)
	defer pm.wg.Done()
}


func (pm *PersonManager) startWalking() {
	for index := 0; index < len(pm.people); index++ {
		go pm.people[index].walk(pm.wg, pm.exit)
		pm.wg.Add(1)
	}
}


func (pm *PersonManager) runSimulation() {
	pm.initPeople()
	pm.setPeoplePos()
	go pm.startTimer(pm.seconds)
	pm.startWalking()
	//pm.wg.Wait()
}


func newPersonManager(nPeople int, seconds int, mapa *Map)  *PersonManager{

	return &PersonManager {
		wg: new(sync.WaitGroup),
		mapa: mapa,
		seconds: seconds,
		people: make([]Persona, nPeople),
		exit: make(chan int),
	}
	
}






