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
var ICON_SIZE = 32;

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

func main() {
    g_Options.ParseCommandLineOptions();

    img := image.NewRGBA(image.Rect(0, 0, 8*ICON_SIZE, 8*ICON_SIZE))
    img.Set(2, 3, color.RGBA{255, 0, 0, 255})

    f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
    bishopIcon := loadIcon("icons/bbws.png")
    defer f.Close()

    draw.Draw(img, img.Bounds(), bishopIcon, image.Point{0,0}, draw.Src);

    png.Encode(f, img)

    fmt.Println("Success !")
}
