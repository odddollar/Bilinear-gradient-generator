package main

import (
	"fmt"
	"github.com/schollz/progressbar"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
	"time"
)

// define pixel type
type pixel struct {
	red, green, blue int
}

func main() {
	// get start time
	start := getTime()

	// create random time seed
	rand.Seed(time.Now().UnixNano())

	// create array
	array := [512][512]pixel{}

	// create progress bar
	bar := progressbar.Default(768420)

	// assign random values to 4 corners
	array[0][0].red = rand.Intn(255)
	array[0][0].green = rand.Intn(255)
	array[0][0].blue = rand.Intn(255)

	array[0][511].red = rand.Intn(255)
	array[0][511].green = rand.Intn(255)
	array[0][511].blue = rand.Intn(255)

	array[511][0].red = rand.Intn(255)
	array[511][0].green = rand.Intn(255)
	array[511][0].blue = rand.Intn(255)

	array[511][511].red = rand.Intn(255)
	array[511][511].green = rand.Intn(255)
	array[511][511].blue = rand.Intn(255)

	// top row red
	for x := 1; x < 511; x++ {
		array[0][x].red = getValueRed(float64(x), 0, array)
		_ = bar.Add(1)
	}
	// middle chunk red
	for y := 1; y < 511; y++ {
		for x := 0; x < 512; x++ {
			array[y][x].red = getValueRed(float64(x), float64(y), array)
			_ = bar.Add(1)
		}
	}
	// bottom row red
	for x := 1; x < 511; x++ {
		array[511][x].red = getValueRed(float64(x), 511, array)
		_ = bar.Add(1)
	}

	// top row green
	for x := 1; x < 511; x++ {
		array[0][x].green = getValueGreen(float64(x), 0, array)
		_ = bar.Add(1)
	}
	// middle chunk green
	for y := 1; y < 511; y++ {
		for x := 0; x < 512; x++ {
			array[y][x].green = getValueGreen(float64(x), float64(y), array)
			_ = bar.Add(1)
		}
	}
	// bottom row green
	for x := 1; x < 511; x++ {
		array[511][x].green = getValueGreen(float64(x), 511, array)
		_ = bar.Add(1)
	}

	// top row blue
	for x := 1; x < 511; x++ {
		array[0][x].blue = getValueBlue(float64(x), 0, array)
		_ = bar.Add(1)
	}
	// middle chunk blue
	for y := 1; y < 511; y++ {
		for x := 0; x < 512; x++ {
			array[y][x].blue = getValueBlue(float64(x), float64(y), array)
			_ = bar.Add(1)
		}
	}
	// bottom row blue
	for x := 1; x < 511; x++ {
		array[511][x].blue = getValueBlue(float64(x), 511, array)
		_ = bar.Add(1)
	}

	// print array
	//for y := 0; y < 512; y++ {
		//fmt.Println(array[y])
	//}

	// create and save image
	createImage(array)

	// print time taken to complete
	fmt.Printf("Time to complete: %vs", float32(getTime()-start))
}

func createImage(matrix [512][512]pixel) {
	// create image parameters
	upLeft := image.Point{0, 0}
	lowRight:= image.Point{512, 512}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// iterate through pixels
	for y := 0; y < 512; y++ {
		for x := 0; x < 512; x++ {
			// set colour
			red := uint8(matrix[y][x].red)
			green := uint8(matrix[y][x].green)
			blue := uint8(matrix[y][x].blue)
			col := color.RGBA{red, green, blue, 0xff}
			img.Set(x, y, col)
		}
	}

	// save image
	f, _ := os.Create("image.png")
	_ = png.Encode(f, img)
}

func getValueRed(posX, posY float64, array [512][512]pixel) int {
	var calc float64
	calc = (((511-posX)*(511-posY))/(511*511)*float64(array[0][0].red)) + ((posX*(511-posY))/(511*511)*float64(array[0][511].red)) + (((511-posX)*posY)/(511*511)*float64(array[511][0].red)) + ((posX*posY)/(511*511)*float64(array[511][511].red))

	return int(math.Round(calc))
}

func getValueGreen(posX, posY float64, array [512][512]pixel) int {
	var calc float64
	calc = (((511-posX)*(511-posY))/(511*511)*float64(array[0][0].green)) + ((posX*(511-posY))/(511*511)*float64(array[0][511].green)) + (((511-posX)*posY)/(511*511)*float64(array[511][0].green)) + ((posX*posY)/(511*511)*float64(array[511][511].green))

	return int(math.Round(calc))
}

func getValueBlue(posX, posY float64, array [512][512]pixel) int {
	var calc float64
	calc = (((511-posX)*(511-posY))/(511*511)*float64(array[0][0].blue)) + ((posX*(511-posY))/(511*511)*float64(array[0][511].blue)) + (((511-posX)*posY)/(511*511)*float64(array[511][0].blue)) + ((posX*posY)/(511*511)*float64(array[511][511].blue))

	return int(math.Round(calc))
}

func getTime() float64 {return float64(time.Now().Unix())}