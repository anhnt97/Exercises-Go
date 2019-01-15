package main

import (
	"fmt"
	"github"
	"log"
	"os"
	"time"
)

func main() {
	results, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	resultsLessAMonth := github.IssuesSearchResult{}
	resultsLessAYear := github.IssuesSearchResult{}
	resultsMoreAYear := github.IssuesSearchResult{}
	timeNow := time.Now()
	for _, item := range results.Items {
		timeCreate := item.CreateAt
		diff := timeNow.Sub(timeCreate)
		days := int(diff.Hours() / 24)
		if days <= 30 {
			resultsLessAMonth.Items = append(resultsLessAMonth.Items, item)
		} else if days <= 365 {
			resultsLessAYear.Items = append(resultsLessAYear.Items, item)
		} else {
			resultsMoreAYear.Items = append(resultsMoreAYear.Items, item)
		}
	}
	fmt.Println("Issues less than a month old : ")
	for _, item := range resultsLessAMonth.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
	fmt.Println("Issues less than a year old : ")
	for _, item := range resultsLessAYear.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
	fmt.Println("Issues more than a year old  :")
	for _, item := range resultsMoreAYear.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
