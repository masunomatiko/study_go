package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

// Usage:
// $ ./ex0305| ./ex101 f=png > output.png
// Input format =  png
// $ ./ex0305| ./ex101 f=jpeg > output.jpeg
// Input format =  png
// $ ./ex0305| ./ex101 f=gif > output.gif
// Input format =  png
// $ ./ex0305| ./ex101  > output.gif
// Input format =  png

func main() {
	var format string
	flag.Parse()
	flag.StringVar(&format, "f", "jpeg", "format to output")

	if err := convert(os.Stdin, os.Stdout, format); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", format, err)
		os.Exit(1)
	}
}

func convert(in io.Reader, out io.Writer, format string) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format = ", kind)
	switch format {
	case "jpeg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "gif":
		return gif.Encode(out, img, &gif.Options{NumColors: 256})
	case "png":
		return png.Encode(out, img)
	default:
		err := fmt.Errorf("image package doesn't have format %s", format)
		return err
	}
}
