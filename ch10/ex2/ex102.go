package main

import (
	"archive/tar"
	"archive/zip"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

//  Usage:
// $ ./ex102 f=zip
// Contents of README:
// This is the source code repository for the Go programming language.
// $ ./ex102 f=tar
// Contents of small.txt:
// Kilts
// Contents of small2.txt:
// Google.com

func main() {
	var format string
	flag.Parse()
	flag.StringVar(&format, "f", "zip", "format to output")
	if err := OpenFile(format); err != nil {
		fmt.Printf("%s: %v\n", format, err)
		os.Exit(1)
	}
}

func OpenFile(format string) error {
	switch format {
	case "zip":
		r, err := zip.OpenReader("testdata/readme.zip")
		if err != nil {
			fmt.Printf("fail to get ReadCloser: %v", err)
			return err
		}
		defer r.Close()
		for _, f := range r.File {
			fmt.Printf("Contents of %s:\n", f.Name)
			rc, err := f.Open()
			if err != nil {
				errors.New("fail to open file")
				return err
			}
			_, err = io.CopyN(os.Stdout, rc, 68)
			if err != nil {
				fmt.Errorf("fail to read contents: %v", err)
				return err
			}
			fmt.Println()
			rc.Close()
			return nil
		}
	case "tar":
		r, err := os.Open("testdata/gnu.tar")
		if err != nil {
			err := fmt.Errorf("fail to open file")
			return err
		}
		defer r.Close()
		tr := tar.NewReader(r)
		for {
			hpr, err := tr.Next()
			if err == io.EOF {
				break
			}
			if err != nil {
				err := errors.New("fail to get Reader")
				return err
			}
			fmt.Printf("Contents of %s: \n", hpr.Name)
			if _, err := io.Copy(os.Stdout, tr); err != nil {
				err := fmt.Errorf("fail to read contents: %v", err)
				return err
			}
			fmt.Println()
			return nil
		}
	default:
		err := fmt.Errorf("archive package doesn't have format %s", format)
		return err
	}
	return nil
}
