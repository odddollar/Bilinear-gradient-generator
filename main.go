package main

import (
	"fmt"
	"github.com/schollz/progressbar/v3"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
	"time"
)

func main() {
	// get start time
	start := getTime()

	// create random time seed
	rand.Seed(time.Now().UnixNano())

	// create array
	redArray := [512][512]int{}
	greenArray := [512][512]int{}
	blueArray := [512][512]int{}

	// create progress bar
	bar := progressbar.Default(1048564)

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
		redArray[0][x] = getValue(float64(x), 0, redArray)
		_ = bar.Add(1)
	}
	// middle chunk red
	for y := 1; y < 511; y++ {
		for x := 0; x < 512; x++ {
			redArray[y][x] = getValue(float64(x), float64(y), redArray)
			_ = bar.Add(1)
		}
	}
	// bottom row red
	for x := 1; x < 511; x++ {
		redArray[511][x] = getValue(float64(x), 511, redArray)
		_ = bar.Add(1)
	}

	// top row green
	for x := 1; x < 511; x++ {
		greenArray[0][x] = getValue(float64(x), 0, greenArray)
		_ = bar.Add(1)
	}
	// middle chunk green
	for y := 1; y < 511; y++ {
		for x := 0; x < 512; x++ {
			greenArray[y][x] = getValue(float64(x), float64(y), greenArray)
			_ = bar.Add(1)
		}
	}
	// bottom row green
	for x := 1; x < 511; x++ {
		greenArray[511][x] = getValue(float64(x), 511, greenArray)
		_ = bar.Add(1)
	}

	// top row blue
	for x := 1; x < 511; x++ {
		blueArray[0][x] = getValue(float64(x), 0, blueArray)
		_ = bar.Add(1)
	}
	// middle chunk blue
	for y := 1; y < 511; y++ {
		for x := 0; x < 512; x++ {
			blueArray[y][x] = getValue(float64(x), float64(y), blueArray)
			_ = bar.Add(1)
		}
	}
	// bottom row blue
	for x := 1; x < 511; x++ {
		blueArray[511][x] = getValue(float64(x), 511, blueArray)
		_ = bar.Add(1)
	}

	// create and save image
	createImage(redArray, greenArray, blueArray, bar)

	// print time taken to complete
	fmt.Printf("Time to complete: %vs", float32(getTime()-start))
}

func createImage(redArray, greenArray, blueArray [512][512]int, bar *progressbar.ProgressBar) {
	// create image parameters
	upLeft := image.Point{0, 0}
	lowRight:= image.Point{512, 512}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// iterate through pixels
	for y := 0; y < 512; y++ {
		for x := 0; x < 512; x++ {
			// set colour
			red := uint8(redArray[y][x])
			green := uint8(greenArray[y][x])
			blue := uint8(blueArray[y][x])
			col := color.RGBA{red, green, blue, 0xff}
			img.Set(x, y, col)

			// increase progress bar
			_ = bar.Add(1)
		}
	}

	// save image
	f, _ := os.Create("image.png")
	_ = png.Encode(f, img)
}

func getValue(posX, posY float64, array [512][512]int) int {
	var calc float64
	calc = (((511-posX)*(511-posY))/(511*511)*float64(array[0][0])) + ((posX*(511-posY))/(511*511)*float64(array[0][511])) + (((511-posX)*posY)/(511*511)*float64(array[511][0])) + ((posX*posY)/(511*511)*float64(array[511][511]))

	return int(math.Round(calc))
}

func getTime() float64 {return float64(time.Now().Unix())}