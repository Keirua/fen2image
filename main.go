package main

import (
	"image"
	"flag"
	"fmt"
	"os"
)

var DEFAULT_ICON_SIZE = 60

type Options struct {
	Fen            string
	OutputFilename string
	CellSize       int
}

var g_Options Options

var fen = "8/8/8/4k3/5R2/8/8/3QK3 w - - 0 1"

var defaultBoard = [8][8]byte{
	{'r', 'n', 'b', 'q', 'k', 'b', 'n', 'r'},
	{'p', 'p', 'p', 'p', 'p', 'p', 'p', 'p'},
	{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{'P', 'P', 'P', 'P', 'P', 'P', 'P', 'P'},
	{'R', 'N', 'B', 'Q', 'K', 'B', 'N', 'R'}}

var validPieces = []byte{'r', 'n', 'b', 'q', 'k', 'p', 'R', 'N', 'B', 'Q', 'K', 'P'}
var icons map[byte]*image.Image = make(map[byte]*image.Image)

func (options *Options) ParseCommandLineOptions() {
	flag.StringVar(&options.OutputFilename, "output", "out.png", "The output filename")
	flag.StringVar(&options.Fen, "fen", "8/8/8/8/8/8/8/8 w - - 0 0", "The fen expression")
	flag.IntVar(&options.CellSize, "cellsize", DEFAULT_ICON_SIZE, "The board cell size")

	flag.Parse()
}

func main() {
	var fen = "4k3/r6B/8/8/8/8/8/K6Q w - - 0 0"

	g_Options.ParseCommandLineOptions()
	var board, err = getBoardFromFen(fen)
	if err != nil {
		fmt.Println(err)
		os.Exit(65) // DATAERR according to /usr/include/sysexits.h
	}
	
	drawCompleteBoard(board, g_Options.OutputFilename, g_Options.CellSize)
}
