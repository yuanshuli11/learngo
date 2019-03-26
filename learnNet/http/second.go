package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Transport

func main() {
	//要管理HTTP客户端的头域、重定向策略和其他设置，创建一个Client：
	//
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}
	////	resp, err := client.Get("http://test-api.zu.ke.com")
	//	resp, err := client.Get("http://ke.com")
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}

	req, err := http.NewRequest("GET", "http://10.26.27.128", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	//可以这样绑host
	req.Host = "yuantrade.ke.com"
	//测试环境增加noSign 免验签名
	req.Header.Add("noSign", `1`)
	resp, err := client.Do(req)

	result, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(result))
}

//当http状态码是 30x 会在执行跳转规则前执行此方法~
func redirectPolicyFunc(req *http.Request, via []*http.Request) error {

	fmt.Println("redirectPolicyFunc:", len(via))
	reCount := 3
	if len(via) >= reCount {
		return errors.New("stopped after too many redirects")
	}
	return nil
}
