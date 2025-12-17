package main

import (
	"embed"
	_ "embed"
	"image"
	"image/png"
	"io/fs"

	"github.com/sysdeep/gothic/gothic"
)

//go:embed images
var imagesFS embed.FS

func main() {
	ir := gothic.NewInterpreter(initGUI)
	<-ir.Done
}

func initGUI(ir *gothic.Interpreter) {
	ir.UploadImage("bg", loadPNG(imagesFS, "images/background.png"))
	ir.Eval(`ttk::label .l -image bg`)
	ir.Eval(`pack .l -expand true`)
}

func loadPNG(storage fs.FS, filename string) image.Image {

	f, err := storage.Open(filename)
	if err != nil {
		panic(err)
	}

	img, err := png.Decode(f)
	if err != nil {
		panic(err)
	}
	return img
}
