package main

import (
    "fmt"
    "flag"
    "image"
    "image/color"
    "image/png"
    "os"
)

type Options struct {
    Fen     string
    OutputFilename     string
}

var g_Options Options;

var fen = "8/8/8/4k3/5R2/8/8/3QK3 w - - 0 1";

func (options *Options) ParseCommandLineOptions() {
    flag.StringVar(&options.OutputFilename, "output", "out.png", "The output filename")
    flag.StringVar(&options.Fen, "fen", "8/8/8/8/8/8/8/8 w - - 0 0", "The fen expression")

    flag.Parse();
}

func parseFen(fen string){

}

func main() {
    g_Options.ParseCommandLineOptions();

    img := image.NewRGBA(image.Rect(0, 0, 100, 50))
    img.Set(2, 3, color.RGBA{255, 0, 0, 255})

    f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
    defer f.Close()
    png.Encode(f, img)

    fmt.Println("Success !")
}
