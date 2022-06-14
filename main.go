package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
	"time"
)

func main() {
	// create random time seed
	rand.Seed(time.Now().UnixNano())

	// create array
	redArray := [512][512]int{}
	greenArray := [512][512]int{}
	blueArray := [512][512]int{}

	// assign random values to 4 corners
	redArray[0][0] = rand.Intn(255)
	greenArray[0][0] = rand.Intn(255)
	blueArray[0][0] = rand.Intn(255)

	redArray[0][511] = rand.Intn(255)
	greenArray[0][511] = rand.Intn(255)
	blueArray[0][511] = rand.Intn(255)

	redArray[511][0] = rand.Intn(255)
	greenArray[511][0] = rand.Intn(255)
	blueArray[511][0] = rand.Intn(255)

	redArray[511][511] = rand.Intn(255)
	greenArray[511][511] = rand.Intn(255)
	blueArray[511][511] = rand.Intn(255)

	// top row red
	for x := 1; x < 511; x++ {
		redArray[0][x] = getValue(float64(x), 0, &redArray)
	}
	// middle chunk red
	for y := 1; y < 511; y++ {
		for x := 0; x < 512; x++ {
			redArray[y][x] = getValue(float64(x), float64(y), &redArray)
		}
	}
	// bottom row red
	for x := 1; x < 511; x++ {
		redArray[511][x] = getValue(float64(x), 511, &redArray)
	}

	// top row green
	for x := 1; x < 511; x++ {
		greenArray[0][x] = getValue(float64(x), 0, &greenArray)
	}
	// middle chunk green
	for y := 1; y < 511; y++ {
		for x := 0; x < 512; x++ {
			greenArray[y][x] = getValue(float64(x), float64(y), &greenArray)
		}
	}
	// bottom row green
	for x := 1; x < 511; x++ {
		greenArray[511][x] = getValue(float64(x), 511, &greenArray)
	}

	// top row blue
	for x := 1; x < 511; x++ {
		blueArray[0][x] = getValue(float64(x), 0, &blueArray)
	}
	// middle chunk blue
	for y := 1; y < 511; y++ {
		for x := 0; x < 512; x++ {
			blueArray[y][x] = getValue(float64(x), float64(y), &blueArray)
		}
	}
	// bottom row blue
	for x := 1; x < 511; x++ {
		blueArray[511][x] = getValue(float64(x), 511, &blueArray)
	}

	// create and save image
	createImage(redArray, greenArray, blueArray)
}

func createImage(redArray, greenArray, blueArray [512][512]int) {
	// create image parameters
	upLeft := image.Point{X: 0, Y: 0}
	lowRight := image.Point{X: 512, Y: 512}
	img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

	// iterate through pixels
	for y := 0; y < 512; y++ {
		for x := 0; x < 512; x++ {
			// set colour
			red := uint8(redArray[y][x])
			green := uint8(greenArray[y][x])
			blue := uint8(blueArray[y][x])
			col := color.RGBA{R: red, G: green, B: blue, A: 0xff}
			img.Set(x, y, col)
		}
	}

	// save image
	f, _ := os.Create("image.png")
	_ = png.Encode(f, img)
}

func getValue(posX, posY float64, array *[512][512]int) int {
	calc := (((511 - posX) * (511 - posY)) / (511 * 511) * float64(array[0][0])) + ((posX * (511 - posY)) / (511 * 511) * float64(array[0][511])) + (((511 - posX) * posY) / (511 * 511) * float64(array[511][0])) + ((posX * posY) / (511 * 511) * float64(array[511][511]))

	return int(math.Round(calc))
}
