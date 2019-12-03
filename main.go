package main

func main() {
	path := "/Users/fer/Desktop/AP Final Project/earthquake/mapitaWrande.csv"
	auxMap := newMapa(10,path, 16,12)
	auxMap.initializeMap()
	auxSimulation := newPersonManager(5,5, auxMap);
	auxSimulation.runSimulation();
}