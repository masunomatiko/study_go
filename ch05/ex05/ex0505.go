package main

// ./ex0505 https://golang.org`

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	url := os.Args[1]
	words, images, _ := CountWordsAndImages(url)
	fmt.Printf("Number of words: %d\nNumber of images: %d\n", words, images)
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
	}
	doc, err := html.Parse(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing html: %s\n", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return words, images, nil
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode {
		if n.Data == "style" || n.Data == "script" {
			return
		} else if n.Data == "img" {
			images++
		}
	} else if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		for _, line := range strings.Split(text, "\n") {
			if line != "" {
				words += len(strings.Split(line, " "))
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i := countWordsAndImages(c)
		words += w
		images += i
	}
	return
}
