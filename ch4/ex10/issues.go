package main

import (
	"fmt"
	"gopl/ch4/github"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	category := categorizeByAge(result.Items)

	fmt.Printf("%d issues:\n", result.TotalCount)
	for age, items := range category {
		fmt.Printf("%s:\n", age)
		for _, item := range items {
			fmt.Printf("\t#%-d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
}

func categorizeByAge(issues []*github.Issue) map[string][]*github.Issue {
	now := time.Now()
	category := make(map[string][]*github.Issue)
	for _, issue := range issues {
		k := key(now, issue.CreatedAt)
		category[k] = append(category[k], issue)
	}
	return category
}

func key(now, createdAt time.Time) string {
	monthAgo := now.AddDate(0, -1, 0)
	yearAgo := now.AddDate(-1, 0, 0)
	switch {
	case createdAt.After(monthAgo):
		return "less than a month old"
	case createdAt.After(yearAgo):
		return "less than a year old"
	}
	return "more than a year old"
}
