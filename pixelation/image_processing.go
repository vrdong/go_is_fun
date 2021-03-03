package pixelation

import (
	"image"
	color2 "image/color"
	draw2 "image/draw"
)

func ProcessImage(startX, startY, sizeX, sizeY, squareSize, goroutineIncrement int, p1 image.Point, p2 image.Point, res draw2.Image) {
	var temp image.Image
	var color color2.Color

	for x := startX; x < sizeX; x = x + goroutineIncrement {
		for y := startY; y < sizeY; y = y + squareSize {
			// process image
			if isInside(x, y, p1, p2) {
				temp = image.NewRGBA(image.Rect(x, y, x+squareSize, y+squareSize))
				color = averageColor(x, y, x+squareSize, y+squareSize, res)
				draw2.Draw(res, temp.Bounds(), &image.Uniform{C: color}, image.Point{X: x, Y: y}, draw2.Src)
			}
			//temp = image.NewRGBA(image.Rect(x, y, x+squareSize, y+squareSize))
			//color = averageColor(x, y, x+squareSize, y+squareSize, res)
			//draw2.Draw(res, temp.Bounds(), &image.Uniform{C: color}, image.Point{X: x, Y: y}, draw2.Src)

		}
	}
}

func isInside(x int, y int, p1 image.Point, p2 image.Point) bool {
	if x > p1.X && x < p2.X && y > p1.Y && y < p2.Y {
		return true
	}
	return false
}

const convertRGB = 0x101
const alpha = 255

func averageColor(startX int, startY int, sizeX int, sizeY int, img draw2.Image) color2.Color {
	var redBucket, greenBucket, blueBucket uint32
	var red, green, blue uint32
	var area uint32

	area = uint32((sizeX - startX) * (sizeY - startY))

	for x := startX; x < sizeX; x++ {
		for y := startY; y < sizeY; y++ {
			red, green, blue, _ = img.At(x, y).RGBA()
			redBucket += red
			greenBucket += green
			blueBucket += blue
		}
	}

	redBucket = redBucket / area
	greenBucket = greenBucket / area
	blueBucket = blueBucket / area
	return color2.RGBA{
		R: uint8(redBucket / convertRGB),
		G: uint8(greenBucket / convertRGB),
		B: uint8(blueBucket / convertRGB),
		A: alpha,
	}
}
