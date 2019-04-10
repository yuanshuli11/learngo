package main

import (
	"flag"
	"fmt"
	"learngo/githook/jira"
	"os"
	"strings"
)

var (
	m string //msg信息
)

func main() {
	flag.StringVar(&m, `m`, ``, `this is commit-msg`)
	flag.Parse()
	task := jira.GetJiraTask()
	arr := strings.Split(m, " ")
	issue := ""
	for _, item := range task.([]interface{}) {
		if arr[0] == item.(string)  {
			issue = item.(string)
			break
		}
	}

	if issue == "" {
		fmt.Println("your jira number should be one of these:\n", task)
	}
	os.Exit(0)
}
