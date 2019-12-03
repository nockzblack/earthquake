package main


import (
	"os"
	"fmt"
	"strconv"
)

func main() {

	args := os.Args[1:] // len(args) should be 2 (people and second to timeOut)

	if (len(args) != 2) {
		fmt.Println("parameters are wrong")
		fmt.Println("use: ./binaryName nPeople secondsToFinish")
		fmt.Println("example: ./earthqueake 5 10")
	} else {

		nPeople, errPeople := strconv.Atoi(args[0])
		timeout, errTimeout := strconv.Atoi(args[1])

		if errTimeout != nil || errPeople != nil {
			fmt.Println("There is some erros on parameters")
			fmt.Println("try to run like: ./earthqueake 5 10")
		} else {
			path := "/Users/fer/Desktop/AP Final Project/earthquake/mapitaWrande.csv"
			auxMap := newMapa(10,path, 16,12)
			auxMap.initializeMap()
			auxSimulation := newPersonManager(nPeople,timeout, auxMap);
			auxSimulation.runSimulation();
		}
	}

	
}