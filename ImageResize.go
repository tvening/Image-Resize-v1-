package main

import (
	"image/jpeg"
	"log"
	"os"

	"github.com/nfnt/resize"
)

// Voer hier het gewenste formaat in

const breedte = 50
const lengte = 50

// Voer hier het bestandsnaam in

const bestandsnaam = "".jpg""
const bestandout = "imageoutput.jpg"

func main() {
	// open "image.jpg"
	file, err := os.Open(bestandsnaam)
	if err != nil {
		log.Fatal(err)
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	m := resize.Resize(breedte, lengte, img, resize.Lanczos3)

	out, err := os.Create(bestandout)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
}
