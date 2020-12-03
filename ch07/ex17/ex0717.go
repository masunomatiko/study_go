package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var found bool
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			for _, v := range tok.Attr {
				if v.Name.Local == "class" && v.Value == "enumar" {
					found = true
				}
			}
		case xml.EndElement:
			found = false
		case xml.CharData:
			if found {
				fmt.Printf("%s\n", tok)
			}
		}
	}
}
