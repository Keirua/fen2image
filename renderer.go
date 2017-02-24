package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"image/jpeg"
	"os"
	"strings"

	"github.com/golang/freetype"
	"io/ioutil"
)

type BoardRenderer interface {
	DrawCompleteBoard(board [8][8]byte, filename string, cellsize int)
}

type RasterBoardRenderer struct {
	icons map[byte]*image.Image
	/*context *freetype.Context
	fontSize float64*/
}

func NewRasterBoardRenderer() *RasterBoardRenderer {
	p := new(RasterBoardRenderer)
    p.icons = make(map[byte]*image.Image)
    p.loadIcons()
    // p.loadFont()
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

func (r RasterBoardRenderer) drawPieces(img* image.RGBA, board [8][8]byte, cellsize int, reverseBoard bool) *image.RGBA {
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

func (r RasterBoardRenderer) drawLetter(context freetype.Context, x int, s string) {
	fontSize := 30.0
	cellsize := 60
	y := 8
	xPos := cellsize/2-int(context.PointToFixed(fontSize)>>6)/4 + x*cellsize
	yPos := int(context.PointToFixed(fontSize)>>6)+ int((float64(y))*float64(cellsize))

	pt := freetype.Pt(xPos, yPos)
	context.DrawString(s, pt)
}

func (r RasterBoardRenderer) drawNumber(context freetype.Context, y int, s string) {
	fontSize := 30.0
	cellsize := 60
	x := 8

	xPos := int(context.PointToFixed(fontSize)>>6)/4 + x*cellsize
	yPos := int(context.PointToFixed(fontSize)>>6)/2+ int((float64(y)+0.5)*float64(cellsize))

	pt := freetype.Pt(xPos, yPos)
	context.DrawString(s, pt)
}

func (r RasterBoardRenderer) drawText(img* image.RGBA, cellsize int) *image.RGBA {
	fontBytes, _ := ioutil.ReadFile("./font/luxisr.ttf")
	/*if err != nil {
		//log.Println(err)
		return img
	}*/
	f, _ := freetype.ParseFont(fontBytes)
	/*if err != nil {
		//log.Println(err)
		return img
	}*/

	fontSize := 30.0
	context := freetype.NewContext()
	context.SetDPI(72)
	context.SetFont(f)
	context.SetFontSize(fontSize)
	context.SetSrc(image.Black)

	context.SetClip(img.Bounds())
	context.SetDst(img)
	
	// r.drawLetter(*context, 1,1,"a");
	for i := 0; i < 8; i++ {
		r.drawLetter(*context, i, string(rune('a'+i)));
		r.drawNumber(*context, i, string(rune('0'+i)));
	}

	return img
}

func (r RasterBoardRenderer) DrawCompleteBoard(board [8][8]byte, filename string, cellsize int, reverse bool, drawCellNames bool) {
	imageSize := 8*cellsize;
	if drawCellNames {
		imageSize = int(8.5*float64(cellsize));
	}

	img := image.NewRGBA(image.Rect(0, 0, imageSize, imageSize))
	img = r.drawPieces(img, board, cellsize, reverse)
	img = r.drawText(img, cellsize)

	if strings.Contains(strings.ToLower(filename), ".png") {
		f, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0600)
		defer f.Close()
		png.Encode(f, img)
	} else if strings.Contains(strings.ToLower(filename), ".jpg") {
		f, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0600)
		defer f.Close()
		jpegOptions := jpeg.Options{100}
		jpeg.Encode(f, img, &jpegOptions)
	} else {
		panic("invalid output file format")
	}
}