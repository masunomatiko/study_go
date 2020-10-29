// Run `./ex0414 tensorflow ranking`
// Open `http://localhost:8080/tensorflow/234`

package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
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

var issueListTemplate = template.Must(template.New("issueList").Parse(`
<h1>{{.Issues | len}} issues</h1>
<table>
<tr style='text-align: left'>
<th>#</th>
<th>State</th>
<th>User</th>
<th>Title</th>
</tr>
{{range .Issues}}
<tr>
	<td><a href='{{.CopyURL}}'>{{.Number}}</td>
	<td>{{.State}}</td>
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	<td><a href='{{.CopyURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

var issueTemplate = template.Must(template.New("issue").Parse(`
<h1>{{.Title}}</h1>
<dl>
	<dt>user</dt>
	<dd><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></dd>
	<dt>state</dt>
	<dd>{{.State}}</dd>
</dl>
<p>{{.Body}}</p>
`))

type IssueCopy struct {
	Issues         []Issues
	IssuesByNumber map[int]Issues
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

func GetIssues(owner, repo string) (ic IssueCopy, err error) {
	reqUrl := fmt.Sprintf("%s/repos/%s/%s/issues", APIURL, owner, repo)
	resp, err := http.Get(reqUrl)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return
	}
	var issues []Issues
	if err = json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		resp.Body.Close()
		return
	}
	resp.Body.Close()

	ic.Issues = issues
	ic.IssuesByNumber = make(map[int]Issues, len(issues))
	for _, issue := range issues {
		ic.IssuesByNumber[issue.Number] = issue
	}
	return
}

func (ic IssueCopy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.SplitN(r.URL.Path, "/", -1)
	numStr := pathParts[2]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(fmt.Sprintf("Issue number isn't a number: '%s'", numStr)))
		if err != nil {
			log.Printf("Error writing response for %s: %s", r, err)
		}
		return
	}
	issue, ok := ic.IssuesByNumber[num]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte(fmt.Sprintf("No issue '%d'", num)))
		if err != nil {
			log.Printf("Error writing response for %s: %s", r, err)
		}
		return
	}
	issueTemplate.Execute(w, issue)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: should set OWNER REPOSITORY")
		os.Exit(1)
	}
	owner := os.Args[1]
	repo := os.Args[2]

	issueCopy, err := GetIssues(owner, repo)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", issueCopy)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
