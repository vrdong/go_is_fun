package main

import (
	"fmt"
	pixelation2 "go_is_fun/pixelation"
	"image"
	"image/draw"
	"log"
	"math"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	pixelation()
}

func pixelation() {
	var wg sync.WaitGroup
	var sizeX, sizeY int
	var img image.Image
	var res *image.RGBA
	var goroutineCount = 10
	var goroutineIncrement int

	// reading image file
	imagePath, squareSize, processingMode, point1, point2 := readCommandLine()

	img = pixelation2.OpenImage(imagePath)

	// getting image size
	sizeX = img.Bounds().Size().X
	sizeY = img.Bounds().Size().Y

	if point2.X == point2.Y && point2.X == 0 {
		point2.X = sizeX
		point2.Y = sizeY
	}
	// checking for errors
	commandLineErrorCheck(sizeX, sizeY, squareSize, processingMode)

	// creating a mask for the result
	res = image.NewRGBA(image.Rect(0, 0, sizeX, sizeY))
	draw.Draw(res, res.Bounds(), img, image.Point{}, draw.Src)

	// setting the # of goroutines according to the size of image
	if processingMode == "M" {
		goroutineCount = int(math.Ceil(float64(sizeX) / float64(squareSize)))
	}

	fmt.Println("The number of goroutines: ", goroutineCount)

	// adding goroutines to the wait group
	wg.Add(goroutineCount)

	// defining how much we should step over the x-axis
	goroutineIncrement = goroutineCount * squareSize

	fmt.Println("Processing the image...")
	for i := 0; i < goroutineCount; i++ {
		go func(i int) {
			defer wg.Done()
			start := time.Now()
			//fmt.Print(i * squareSize, 0, sizeX, sizeY, squareSize, goroutineCount)
			pixelation2.ProcessImage(i*squareSize, 0, sizeX, sizeY, squareSize, goroutineIncrement, point1, point2, res)
			elapsed := time.Since(start)
			log.Println("Completed in: ", elapsed)
		}(i)
	}

	// saving the image
	defer pixelation2.SaveImage(imagePath, res)

	// making main goroutine wait until other goroutines finish
	wg.Wait()
}

// Checks and panics in case of errors
func check(err error) {
	if err != nil {
		log.Panicln(err)
	}
}

// Reads command line arguments
func readCommandLine() (string, int, string, image.Point, image.Point) {
	var imagePath, processingMode, size, x1, y1, x2, y2 string
	var newX1, newY1, newX2, newY2 int
	var err error

	if len(os.Args) < 4 {
		log.Fatalln("Command not found. Please enter according to the template" +
			" fileName.jpg squareSize processingMode [S or M]" +
			" For example: go run main.go monalisa.jpg 30 M")
	}
	imagePath = os.Args[1]
	size = os.Args[2]
	processingMode = os.Args[3]

	if len(os.Args) == 8 {
		x1 = os.Args[4]
		y1 = os.Args[5]
		x2 = os.Args[6]
		y2 = os.Args[7]
		newX1, err = strconv.Atoi(x1)
		check(err)
		newY1, err = strconv.Atoi(y1)
		check(err)
		newX2, err = strconv.Atoi(x2)
		check(err)
		newY2, err = strconv.Atoi(y2)
		check(err)
	}
	squareSize, err := strconv.Atoi(size)
	check(err)

	return imagePath, squareSize, processingMode,
		image.Point{
			X: newX1,
			Y: newY1,
		}, image.Point{
			X: newX2,
			Y: newY2,
		}
}

// Handles errors for the square size and processing mode
func commandLineErrorCheck(sizeX, sizeY, squareSize int, processingMode string) {
	if squareSize > sizeX || squareSize > sizeY || squareSize <= 1 {
		log.Fatalln("Out of bounds or small square size. Please change the size of the square")
	}
	if processingMode != "M" && processingMode != "S" {
		log.Fatalln("Wrong processing mode. Please enter S for single or M for multi-threaded mode.")
	}
}
