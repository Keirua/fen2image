package main

import (
    "fmt"
    "flag"
    "image"
    "image/color"
    "image/png"
    "image/draw"
    "os"
    "regexp"
    "strings"
    "errors"
)

var DEFAULT_ICON_SIZE = 60;

type Options struct {
    Fen     string
    OutputFilename     string
    CellSize int
}

var g_Options Options;

var fen = "8/8/8/4k3/5R2/8/8/3QK3 w - - 0 1";

var defaultBoard = [8][8]byte{  
 {'r','n','b','q','k','b','n','r'},
 {'p','p','p','p','p','p','p','p'},
 {' ',' ',' ',' ',' ',' ',' ',' '},
 {' ',' ',' ',' ',' ',' ',' ',' '},
 {' ',' ',' ',' ',' ',' ',' ',' '},
 {' ',' ',' ',' ',' ',' ',' ',' '},
 {'P','P','P','P','P','P','P','P'},
 {'R','N','B','Q','K','B','N','R'}};

var validPieces = []byte{'r','n','b','q','k','p', 'R','N','B','Q','K','P'};
var icons map[byte] *image.Image = make(map[byte] *image.Image);

func (options *Options) ParseCommandLineOptions() {
    flag.StringVar(&options.OutputFilename, "output", "out.png", "The output filename")
    flag.StringVar(&options.Fen, "fen", "8/8/8/8/8/8/8/8 w - - 0 0", "The fen expression")
    flag.IntVar(&options.CellSize, "cellsize", DEFAULT_ICON_SIZE, "The board cell size")

    flag.Parse();
}

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
    return &src;
}

func contains(s []byte, e byte) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func Rect(x1 int, y1, x2, y2 int, col color.RGBA, img *image.RGBA) {
    for x := x1; x<= x2; x++ {
        for y := y1; y <= y2; y++ {
            img.Set(x, y, col)
        }
    }
}

func DrawBackground(img *image.RGBA){
    // Lichess colors : 8ca2ad and dee3e6
    var whiteColor = color.RGBA{222, 227, 230, 255};
    var blackColor = color.RGBA{140,162, 173, 255};

    var s = g_Options.CellSize;

    for x := 0; x< 8; x++ {
        for y := 0; y < 8; y++ {
            color := blackColor;
            if (x+y)%2 == 0 {
                color = whiteColor;
            }

            Rect(x*s, y*s, (x+1)*s, (y+1)*s, color, img)
        }
    }
}

func DrawBoard(board [8][8]byte, img *image.RGBA){
    for x := 0; x< 8; x++ {
        for y := 0; y < 8; y++ {
            var piece = board[y][x];
            
            var isValidPiece = piece != ' ' && contains(validPieces, piece);
            var _, isLoaded = icons[piece];
            if isValidPiece && isLoaded {
                var coords = image.Point{-x*g_Options.CellSize,-y*g_Options.CellSize}
                draw.Draw(img, img.Bounds(), *(icons[piece]), coords, draw.Over);
            }
        }
    }
}

func loadIcons(){
    for _, s := range validPieces {
        var iconFile = "icons/"+ string(s) +"60.png";
        icons[s] = loadIcon(iconFile)
    }
}

func isValidFen(fen string) bool{
    var linePattern = "[1-9rnbqkpRNBQKP]+";
    var fenRegex = fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s/%s w|b - - \\d+ \\d+", linePattern, linePattern, linePattern, linePattern, linePattern, linePattern, linePattern, linePattern)
    var isValid, _ = regexp.MatchString(fenRegex, fen)
    return isValid
}

func getBoardLine(fenElement string) ([8]byte, error) {
    var line = [8]byte{};
    var pos = 0;
    for _, cellValue := range fenElement {
        if pos >= 8 {
            return line, errors.New(fmt.Sprintf("Oops: %s is not a valid fen element", fenElement))
            break;
        }
        if contains(validPieces, byte(cellValue)){
            line[pos] = byte(cellValue)
            pos = pos + 1
            continue;
        }  else {
            pos = pos + int((int(cellValue) - int('0')))
        }
    }

    return line, nil;
}

func getBoardFromFen(fen string) ([8][8]byte, error) {
    if false == isValidFen(fen) {
        return defaultBoard, errors.New("Invalid FEN !");
    }
    var splits = strings.Split(fen, " ")
    var lines = strings.Split(splits[0], "/")

    var board = [8][8]byte{};
    for y := 0; y < 8; y++ {
        var err error
        board[y], err = getBoardLine(lines[y])
        if err != nil {
            return defaultBoard, err
        }
    }

    return board, nil;
}

func main() {
    var fen = "4k3/r6Bp/8p/8/8/8/8/K6Q w - - 0 0"

    g_Options.ParseCommandLineOptions();
    loadIcons();

    img := image.NewRGBA(image.Rect(0, 0, 8*g_Options.CellSize, 8*g_Options.CellSize))

    f, _ := os.OpenFile(g_Options.OutputFilename, os.O_WRONLY|os.O_CREATE, 0600)
    defer f.Close()

    var board, err = getBoardFromFen(fen);
    DrawBackground(img);
    DrawBoard(board, img)

    if err != nil {
        fmt.Println(err)
        fmt.Println("Output will be a default board")
        os.Exit(65) // DATAERR according to /usr/include/sysexits.h
    }

    png.Encode(f, img)
}
