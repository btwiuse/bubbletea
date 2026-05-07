package main

// A simple program that counts down from 5 and then exits.

import (
	"log"

	termimg "github.com/blacktop/go-termimg"
	boba "github.com/btwiuse/boba"
)

func main() {
	log.Println("Supported protocols:")
	protocols := termimg.DetermineProtocols()
	for _, p := range protocols {
		log.Println("-", p.String())
	}

	widget := termimg.NewImageWidgetFromImage(Image())

	p := boba.NewProgram(model{widget: widget, protocol: termimg.Kitty})
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
