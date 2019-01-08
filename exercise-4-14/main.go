package main

import (
	"fmt"
	"github"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

var result *github.IssuesSearchResult
var issueListTemplate *template.Template
var issueTemplate *template.Template
var userTemplate *template.Template

func init() {
	// call github api search issues
	var err error
	result, err = github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	// parse issues template
	issueListTemplate = template.Must(template.ParseFiles("C:/Users/admin/go/src/test/issueList.html"))
	issueTemplate = template.Must(template.ParseFiles("C:/Users/admin/go/src/test/issue.html"))
	userTemplate = template.Must(template.ParseFiles("C:/Users/admin/go/src/test/user.html"))

}

func main() {
	// Add handle funcs for 3 pages
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/issue", issueHandler)
	http.HandleFunc("/", issueListHandler)
	// Run web server
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// handles list issue page
func issueListHandler(w http.ResponseWriter, r *http.Request) {
	if err := issueListTemplate.Execute(w, result); err != nil {
		log.Fatal(err)
	}
}

// handles issue page
func issueHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	number, err := strconv.Atoi(r.Form["number"][0])
	if err != nil {
		log.Fatal(err)
	}
	var issue *github.Issue
	for _, item := range result.Items {
		if item.Number == number {
			issue = item
		}
	}
	if issue != nil {
		if err := issueTemplate.Execute(w, issue); err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Fprintln(w, "Issue not found")
	}
}

// handles user page
func userHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	name := r.Form["name"][0]
	var user *github.User
	for _, item := range result.Items {
		if current := item.User; current.Login == name {
			user = item.User
			break
		}
	}
	if user != nil {
		if err := userTemplate.Execute(w, user); err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Fprintln(w, "User not found")
	}
}
