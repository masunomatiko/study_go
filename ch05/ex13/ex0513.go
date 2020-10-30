package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/adonovan/gopl.io/ch5/links"
)

var domain string

func crawl(reqUrl string) error {
	fmt.Println(reqUrl)

	if domain == "" {
		p, err := url.Parse(reqUrl)
		if err != nil {
			log.Fatalf("crawl %s get: %v", err)
		}
		domain = p.Hostname()
		if strings.HasPrefix(domain, "www.") {
			domain = domain[4:]
		}
		fmt.Printf("Domain: %s \n\n", domain)
	}

	list, err := links.Extract(reqUrl)
	if err != nil {
		log.Print(err)
	}

	for _, l := range list {
		p, err := url.Parse(l)
		if err != nil {
			continue
		}
		if strings.Contains(p.Hostname(), domain) {
			resp, err := http.Get(l)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			file, err := os.Create(l)
			if err != nil {
				return err
			}
			_, err = io.Copy(file, resp.Body)
			if err != nil {
				return err
			}
			err = file.Close()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func main() {
	err := crawl(os.Args[1])
	if err != nil {
		log.Fatal("something went wrong: %v", err)
	}
}
