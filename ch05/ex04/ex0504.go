package main

// ./ex0502 https://golang.org`

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var keyMapper = map[string]string{"a": "href", "img": "src", "script": "src", "link": "href"}

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

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	data := [4]string{"a", "img", "script", "link"}
	for _, d := range data {
		if n.Type == html.ElementNode && n.Data == d {
			for _, a := range n.Attr {
				if a.Key == keyMapper[d] {
					links = append(links, a.Val)
				}
			}
		}
	}

	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}
	return links
}
