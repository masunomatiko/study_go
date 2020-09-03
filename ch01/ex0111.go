package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	fname := "ex0111.txt"
	out, err := os.Create(fname)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	urls := getURLs("top-1m.txt")

	start := time.Now()
	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, ch)
	}
	for range urls {
		fmt.Fprintln(out, <-ch)
	}
	fmt.Fprintf(out, "%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Sprintf("While reading %s: %v", url, err)
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%2fs %7d %s", secs, nbytes, url)

}

func getURLs(filepath string) []string {
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File %s could not read: %v\n", filepath, err)
		os.Exit(1)
	}

	// 関数return時に閉じる
	defer f.Close()

	// Scannerで読み込む
	urls := make([]string, 0, 1000000)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		url := scanner.Text()
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		urls = append(urls, url)
	}
	if err = scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "File %s scan error: %v\n", filepath, err)
	}
	return urls
}
