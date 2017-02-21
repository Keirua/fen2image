package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

type BoardRenderer interface {
	DrawCompleteBoard(board [8][8]byte, filename string, cellsize int)
}

type RasterBoardRenderer struct {
	icons map[byte]*image.Image
}

func NewRasterBoardRenderer() *RasterBoardRenderer {
	p := new(RasterBoardRenderer)
    p.icons = make(map[byte]*image.Image)
    p.loadIcons()
    return p
}

func (r RasterBoardRenderer) loadIcons() {
	for _, s := range validPieces {
		var iconFile = "icons/" + string(s) + "60.png"
		r.icons[s] = r.loadIcon(iconFile)
	}
}

func (r RasterBoardRenderer) loadIcon(inputFilename string) *image.Image {
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

func (r RasterBoardRenderer) rect(x1 int, y1, x2, y2 int, col color.RGBA, img *image.RGBA) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			img.Set(x, y, col)
		}
	}
}

func (r RasterBoardRenderer) drawPieces(board [8][8]byte, cellsize int, reverseBoard bool) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, 8*cellsize, 8*cellsize))

	// Lichess colors : 8ca2ad and dee3e6
	var whiteColor = color.RGBA{222, 227, 230, 255}
	var blackColor = color.RGBA{140, 162, 173, 255}

	var s = cellsize;

	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			// drawCell
			cellColor := blackColor
			if (x+y)%2 == 0 {
				cellColor = whiteColor
			}
			r.rect(x*s, y*s, (x+1)*s, (y+1)*s, cellColor, img)

			// drawPiece
			var piece = board[y][x]
			if reverseBoard {
				piece = board[7-y][7-x]
			}

			var _, isLoaded = r.icons[piece]
			if isChessPieceOrPawn(piece) && isLoaded {
				var coords = image.Point{-x * cellsize, -y * cellsize}
				draw.Draw(img, img.Bounds(), *(r.icons[piece]), coords, draw.Over)
			}
		}
	}

	return img
}

func (r RasterBoardRenderer) DrawCompleteBoard(board [8][8]byte, filename string, cellsize int, reverse bool) {
	img := r.drawPieces(board, cellsize, reverse)

	f, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()

	png.Encode(f, img)
}