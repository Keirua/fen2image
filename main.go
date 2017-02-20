package main

import (
    "fmt"
    "flag"
    "image"
    "image/color"
    "image/png"
    "image/draw"
    "os"
)

type Options struct {
    Fen     string
    OutputFilename     string
}

var g_Options Options;

var fen = "8/8/8/4k3/5R2/8/8/3QK3 w - - 0 1";
var ICON_SIZE = 60;

func (options *Options) ParseCommandLineOptions() {
    flag.StringVar(&options.OutputFilename, "output", "out.png", "The output filename")
    flag.StringVar(&options.Fen, "fen", "8/8/8/8/8/8/8/8 w - - 0 0", "The fen expression")

    flag.Parse();
}

func loadIcon(inputFilename string) image.Image {
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
    return src;
}

func parseFen(fen string){

}

// 8ca2ad
// dee3e6

func Rect(x1 int, y1, x2, y2 int, col color.RGBA, img *image.RGBA) {
    for x := x1; x<= x2; x++ {
        for y := y1; y <= y2; y++ {
            img.Set(x, y, col)
        }
    }
}

func DrawBoard(img *image.RGBA){
    var whiteColor = color.RGBA{0, 128, 192, 255};
    var blackColor = color.RGBA{96, 96, 96, 255};

    for x := 0; x<= 8; x++ {
        for y := 0; y <= 8; y++ {
            color := blackColor;
            if (x+y)%2 == 0 {
                color = whiteColor;
            }

            Rect(x*ICON_SIZE, y*ICON_SIZE, (x+1)*ICON_SIZE, (y+1)*ICON_SIZE, color, img)
        }
    }
}

func main() {
    g_Options.ParseCommandLineOptions();

    img := image.NewRGBA(image.Rect(0, 0, 8*ICON_SIZE, 8*ICON_SIZE))
    img.Set(2, 3, color.RGBA{255, 0, 0, 255})

    f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
    bishopIcon := loadIcon("icons/b60.png")
    defer f.Close()

    DrawBoard(img);
    draw.Draw(img, img.Bounds(), bishopIcon, image.Point{0,0}, draw.Over);

    //Rect(32,100,100,130, color.RGBA{0, 128, 255, 255}, img);

    png.Encode(f, img)

    fmt.Println("Success !")
}
