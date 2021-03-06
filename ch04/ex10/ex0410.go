package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"time"
)

const IssueURL = "https://api.github.com/search/issues"
const DateFormat = "2020/09/29"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issues
}
type Issues struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:html_url`
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssueURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", err)
	}
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	items := result.Items
	sort.Slice(items, func(i, j int) bool { return items[i].CreatedAt.After(items[j].CreatedAt) })

	idx := 0
	fmt.Println("-----Issued within a month-----")
	for _, item := range result.Items {
		if item.CreatedAt.After(time.Now().AddDate(0, -1, 0)) {
			idx = 0
			fmt.Printf("#%-5d %s %9.9s %.55s\n", item.Number, item.CreatedAt, item.User.Login, item.Title)
		} else if item.CreatedAt.After(time.Now().AddDate(-1, 0, 0)) {
			if idx == 0 {
				fmt.Println("\n-----Issued within a year-----")
			}
			idx = 1
			fmt.Printf("#%-5d %s %9.9s %.55s\n", item.Number, item.CreatedAt, item.User.Login, item.Title)

		} else {
			if idx == 1 {
				fmt.Println("\n-----Issued more than a year ago-----")
			}
			idx = 2
			fmt.Printf("#%-5d %s %9.9s %.55s\n", item.Number, item.CreatedAt, item.User.Login, item.Title)

		}
	}
}
