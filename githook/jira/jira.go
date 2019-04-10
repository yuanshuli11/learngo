package jira

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type jiraList struct {
	issueTable string `json:"issueTable"`
}

type jiraIssueKeys map[string]string

func GetJiraTask() interface{} {

	url := "http://jira2.lianjia.com/rest/issueNav/1/issueTable"
	params := `startIndex=0&filterId=-1&jql=assignee%3DcurrentUser%28%29+AND+resolution+%3D+Unresolved+order+by+updated+DESC`
	data := post(url, params)
	str := []byte(data)
	var r interface{}
	json.Unmarshal(str, &r)
	details, ok := r.(map[string]interface{})
	if ok {
		for key, detail := range details {
			if key == "issueTable" {
				for k, v := range detail.(map[string]interface{}) {
					if k == "issueKeys" {
						return v
					}
				}
				break
			}

		}
	}

	return nil

}
  
func getAuth() (authString string) {
	file, err := ioutil.ReadFile(".pass")
	if err != nil {
		log.Fatal(err)
	}
	return string(file)
}
func post(url string, paramsStr string) string {
	retry := 3
	auth := "Basic " + getAuth()
	//要管理HTTP客户端的头域、重定向策略和其他设置，创建一个Client：
	client := &http.Client{
	}
	for i := 0; i < retry; i++ {
		req, err := http.NewRequest("POST", url, strings.NewReader(paramsStr))
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Add("x-atlassian-token", "nocheck")
		req.Header.Add("authorization", auth)
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode == 401 {
			fmt.Println(`your name or password may not correct!, use "sh copy_githooks.sh" to regenerate.`)
			os.Exit(1)
		} else if resp.StatusCode == 403 {
			fmt.Println(`you need goto http://jira2.lianjia.com/login.jsp to login`)
			os.Exit(1)
		} else if resp.StatusCode == 200 {
			result, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			return string(result)
		} else {
			fmt.Println(`jira2 error, retry...`)
			continue
		}
	}
	return "error"
}
