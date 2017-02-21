package main

import (
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

var validPieces = []byte{'r', 'n', 'b', 'q', 'k', 'p', 'R', 'N', 'B', 'Q', 'K', 'P'}

func (options *Options) ParseCommandLineOptions() {
	var defaultFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1";

	flag.StringVar(&options.OutputFilename, "output", "out.png", "The output filename")
	flag.StringVar(&options.Fen, "fen", defaultFen, "The fen expression")
	flag.IntVar(&options.CellSize, "cellsize", DEFAULT_ICON_SIZE, "The board cell size")

	flag.Parse()
}

func main() {
	g_Options.ParseCommandLineOptions()
	var board, err = getBoardFromFen(g_Options.Fen)
	if err != nil {
		fmt.Println(err)
		os.Exit(65) // DATAERR according to /usr/include/sysexits.h
	}
	
	r := NewRasterBoardRenderer()
	r.DrawCompleteBoard(board, g_Options.OutputFilename, g_Options.CellSize)
}
