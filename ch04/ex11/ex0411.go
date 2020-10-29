package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"sort"
	"strings"
	"time"
)

const APIURL = "https://api.github.com"
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

// for simplicity, only title is `required`
// note: json is case-sensitive
type NewIssue struct {
	Title string `json:"title""`
	Body  string `json:"body"`
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	reqUrl := fmt.Sprintf("%s/masunomatiko/search/issues?q=%s", APIURL, q)

	resp, err := http.Get(reqUrl)
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

func ReadIssues(q string) {
	result, err := SearchIssues([]string{q})
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

func CreateIssues() error {

	reqUrl := fmt.Sprintf("%s/masunomatiko/study_go/issues", APIURL)

	tmpfile, err := ioutil.TempFile(os.TempDir(), "c_")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Invoke VSCode to fill in the body of the content")
	cmd := exec.Command("code", tmpfile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	content, err := os.Open(tmpfile.Name())
	if err != nil {
		log.Fatal(err)
	}

	reqBody := bytes.NewBuffer(make([]byte, 128))

	var issue NewIssue

	reader := bufio.NewReader(content)
	title, _ := reader.ReadString('\n')
	issue.Title = strings.TrimSpace(title)

	body := []byte{}
	b, e := reader.ReadByte()
	for e == nil {
		body = append(body, b)
		b, e = reader.ReadByte()
	}
	issue.Body = string(body)

	data, err := json.Marshal(issue)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	reqBody = bytes.NewBuffer(data)

	req, err := http.NewRequest("POST", reqUrl, reqBody)
	if err != nil {
		log.Fatalf("Create new request failed: %v\n", err)
		os.Exit(1)
	}

	req.SetBasicAuth(os.Getenv("USERNAME"), os.Getenv("PASSWORD"))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusCreated {
		return err
	}
	return nil
}

func UpdateIssues(issueNum string) error {
	reqUrl := fmt.Sprintf("%s/masunomatiko/study_go/issues/%d", APIURL, issueNum)
	resp, err := http.Get(reqUrl)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("search query failed: %s", err)
	}
	var result Issues
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return err
	}

	tmpfile, err := ioutil.TempFile(os.TempDir(), "c_")
	if err != nil {
		log.Fatal(err)
	}
	tmpfile.WriteString(result.Title)
	tmpfile.Write([]byte("\n"))
	tmpfile.WriteString(result.Body)
	tmpfile.Close()

	fmt.Println("Invoke VSCode to fill in the body of the content")
	cmd := exec.Command("code", tmpfile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	content, err := os.Open(tmpfile.Name())
	if err != nil {
		log.Fatal(err)
	}

	reqBody := bytes.NewBuffer(make([]byte, 128))

	var issue NewIssue

	reader := bufio.NewReader(content)
	title, _ := reader.ReadString('\n')
	issue.Title = strings.TrimSpace(title)

	body := []byte{}
	b, e := reader.ReadByte()
	for e == nil {
		body = append(body, b)
		b, e = reader.ReadByte()
	}
	issue.Body = string(body)

	data, err := json.Marshal(issue)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	reqBody = bytes.NewBuffer(data)

	req, err := http.NewRequest("POST", reqUrl, reqBody)
	if err != nil {
		log.Fatalf("Create new request failed: %v\n", err)
		os.Exit(1)
	}

	req.SetBasicAuth(os.Getenv("USERNAME"), os.Getenv("PASSWORD"))
	req.Header.Set("Content-Type", "application/json")

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusCreated {
		return err
	}
	return nil
}

func CloseIssues(issueNum string) error {
	reqUrl := fmt.Sprintf("%s/masunomatiko/study_go/issues/%s", APIURL, issueNum)

	resp, err := http.PostForm(
		reqUrl,
		url.Values{"state": {"closed"}},
	)
	if err != nil {
		log.Fatal(err)
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	fmt.Println("All issues closed.")
	return nil
}

func main() {
	var mode string
	var issueNum string
	var title string

	flag.Parse()
	flag.StringVar(&mode, "m", "", "create/read/update/delete")
	flag.StringVar(&issueNum, "i", "", "issueNum")
	flag.StringVar(&title, "t", "", "title")

	switch mode {
	case "create":
		CreateIssues()
	case "read":
		ReadIssues(title)
	case "update":
		UpdateIssues(issueNum)
	case "delete":
		CloseIssues(issueNum)
	default:
		log.Fatal("select mode to use api")
	}
}
