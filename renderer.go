package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func loadIcon(inputFilename string) *image.Image {
	infile, err := os.Open(inputFilename)
	if err != nil {
		// replace this with real error handling
		panic(err)
	}
	defer infile.Close()

	// Decode will figure out what type of image is in the file on its own.
	// We just have to be sure all the image packages we want are imported.
	src, _, err := image.Decode(infile)
	if err != nil {
		// replace this with real error handling
		panic(err)
	}
	return &src
}

func Rect(x1 int, y1, x2, y2 int, col color.RGBA, img *image.RGBA) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			img.Set(x, y, col)
		}
	}
}

func DrawBackground(img *image.RGBA, cellsize int) {
	// Lichess colors : 8ca2ad and dee3e6
	var whiteColor = color.RGBA{222, 227, 230, 255}
	var blackColor = color.RGBA{140, 162, 173, 255}

	var s = cellsize;

	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			color := blackColor
			if (x+y)%2 == 0 {
				color = whiteColor
			}

			Rect(x*s, y*s, (x+1)*s, (y+1)*s, color, img)
		}
	}
}

func DrawBoard(board [8][8]byte, img *image.RGBA, cellsize int) {

	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			var piece = board[y][x]

			var _, isLoaded = icons[piece]
			if isChessPieceOrPawn(piece) && isLoaded {
				var coords = image.Point{-x * cellsize, -y * cellsize}
				draw.Draw(img, img.Bounds(), *(icons[piece]), coords, draw.Over)
			}
		}
	}
}

func loadIcons() {
	for _, s := range validPieces {
		var iconFile = "icons/" + string(s) + "60.png"
		icons[s] = loadIcon(iconFile)
	}
}

func drawCompleteBoard(board [8][8]byte, filename string, cellsize int) {
	loadIcons()

	img := image.NewRGBA(image.Rect(0, 0, 8*cellsize, 8*cellsize))

	f, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()

	DrawBackground(img, cellsize)
	DrawBoard(board, img, cellsize)

	png.Encode(f, img)
}