package main

// ./ex0502 https://golang.org`

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

	var links = make(map[string]int)
	visitAndMap(links, doc)
	for k, v := range links {
		fmt.Printf("%s: %d\n", k, v)
	}
}

func visitAndMap(links map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		links[n.Data]++
	}
	if n.FirstChild != nil {
		visitAndMap(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		visitAndMap(links, n.NextSibling)
	}
}
