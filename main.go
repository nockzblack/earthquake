package main

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
	"strconv"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var args []string

func main() {
	// Everything in program will run here
	args = os.Args
	args = args[1:]
	pixelgl.Run(run)
}

func run() {
	fmt.Println(args)
	if len(args) != 2 {
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
			path := "mapitaWrande.csv"
			auxMap := newMapa(10, path, 16, 12)
			auxMap.initializeMap()
			auxSimulation := newPersonManager(nPeople, timeout, auxMap)

			if auxSimulation == nil {
				fmt.Println("imposible to run a simulation")
			} else {
				// Configuration of window: tittle to show, size and refresh
				cfg := pixelgl.WindowConfig{
					Title:  "<3  -Earthquake simulator-  <3",
					Bounds: pixel.R(0, 0, 800, 800),
					VSync:  false,
				}
				// Create the window
				win, err := pixelgl.NewWindow(cfg)
				if err != nil {
					panic(err)
				}

				pic, err := loadPicture("bat.png")
				if err != nil {
					panic(err)
				}
				sprite := pixel.NewSprite(pic, pic.Bounds())

				// |  					  |
				// V  Drawing everything  V
				imd := imdraw.New(nil)
				imd.Color = colornames.Black
				numVertical := len(auxMap.nodes)
				numHorizontal := len(auxMap.nodes[0])

				vertSum := 800 / (numVertical)
				horizontalSum := 800 / (numHorizontal)
				vposition := 0
				hposition := 0

				// Draw horizontal lines
				for i := 0; i < numHorizontal+1; i++ {
					imd.Push(pixel.V(float64(hposition), 0))
					imd.Push(pixel.V(float64(hposition), 800))
					imd.Line(10)
					hposition += horizontalSum
				}
				if hposition-horizontalSum < 800 {
					imd.Push(pixel.V(800, 0))
					imd.Push(pixel.V(800, 800))
					imd.Line(float64(800 - hposition + horizontalSum))
				}

				// Draw Vertical lines
				for i := 0; i < numVertical+1; i++ {
					imd.Push(pixel.V(0, float64(vposition)))
					imd.Push(pixel.V(800, float64(vposition)))
					imd.Line(10)
					vposition += vertSum
				}
				if vposition-vertSum < 800 {
					imd.Push(pixel.V(0, 0))
					imd.Push(pixel.V(800, 0))
					imd.Line(float64(800 - vposition + vertSum))
				}

				vposition = 800
				hposition = 0
				// Draw empty spaces
				for i := 0; i < len(auxMap.nodes); i++ {
					hposition = 0
					vposition -= vertSum
					for j := 0; j < len(auxMap.nodes[0]); j++ {
						if auxMap.nodes[i][j] == nil {
							imd.Color = colornames.Black
							imd.Push(pixel.V(float64(hposition), float64(vposition)))
							imd.Push(pixel.V(float64(hposition+horizontalSum), float64(vposition+vertSum)))
							imd.Rectangle(0)
						} else if auxMap.nodes[i][j].isExit {
							imd.Color = colornames.Firebrick
							imd.Push(pixel.V(float64(hposition), float64(vposition)))
							imd.Push(pixel.V(float64(hposition+horizontalSum), float64(vposition+vertSum)))
							imd.Rectangle(0)
						}
						hposition += horizontalSum
					}
				}

				auxSimulation.runSimulation()
				//ubicacionPersona := pixel.V(100,100)
				//batch := pixel.NewBatch(&pixel.TrianglesData{}, pic)

				//mat := pixel.IM.Moved(ubicacionPersona)
				//ubicacionPersona.Y = 0
				channel := auxSimulation.exit
				for !win.Closed() {
					select {
					case <-channel:
						//fmt.Println("Hey")
						auxSimulation.damageReport()
						return

					default:
						win.Clear(colornames.Ghostwhite)
						imd.Draw(win)
						//batch.Draw(win)
						//sprite.Draw(win, mat)
						//persona.Draw(win)
						for i := 0; i < len(auxSimulation.people); i++ {
							if auxSimulation.people[i].isDeath == true {
								x := auxSimulation.people[i].pos.position.X
								y := float64(len(auxMap.nodes)) - 1 - auxSimulation.people[i].pos.position.Y
								mat := pixel.IM.Moved(pixel.V(float64(horizontalSum)*x+float64(horizontalSum/2), float64(vertSum)*y+float64(vertSum/2)))
								sprite.Draw(win, mat.Scaled(pixel.V(float64(horizontalSum)*x+float64(horizontalSum/2), float64(vertSum)*y+float64(vertSum/2)), 0.2))
								//println(x, y)
							}
						}
						win.Update()
					}
				}

			}

		}
	}
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}
