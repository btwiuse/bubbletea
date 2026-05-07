package main

import (
	"bytes"
	"fmt"
	_ "embed"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"

	"github.com/dolmen-go/kittyimg"
)

//go:embed example.png
var examplePNG []byte

func Render() string {
	img, err := RenderImage(bytes.NewReader(examplePNG))
	if err != nil {
		panic(err)
	}

	header := "here is an example image encoded in kitties graphics protocol:"
	footer := "if you see nothing, it means your terminal doesn't support it"

	output := fmt.Sprintf("%s\n\n%s\n\n%s\n\n", header, img, footer)
	return output
}

// RenderString reads an image from r and returns Kitty protocol output.
func RenderImage(r io.Reader) (string, error) {
	var buf bytes.Buffer
	if err := kittyimg.Transcode(&buf, r); err != nil {
		return "", err
	}
	return buf.String(), nil
}
