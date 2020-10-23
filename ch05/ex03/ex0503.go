package main

// ./ex0503 https://golang.org`

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}

	findTextNode(doc)
}

func findTextNode(n *html.Node) {
	if n.Type == html.ElementNode && (n.Data == "style" || n.Data == "style") {
		return
	}
	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}
	if n.FirstChild != nil {
		findTextNode(n.FirstChild)
	}
	if n.NextSibling != nil {
		findTextNode(n.NextSibling)
	}
}
