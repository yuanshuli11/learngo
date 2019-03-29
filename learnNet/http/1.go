package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	//构造get请求
	resp, err := http.Get("http://test-api.zu.ke.com/v1/house/list?city_id=110000")

	if err != nil {
		return
	}
	//	fmt.Println(resp)
	//	fmt.Println(resp.Body)
	//	fmt.Println(ioutil.ReadAll(resp.Body))
	result, err := ioutil.ReadAll(resp.Body)

	//程序在使用完回复后必须关闭回复的主体。
	defer resp.Body.Close()
	fmt.Println(string(result))

	//post请求
	//resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)

	//postform
	//resp, err := http.PostForm("http://example.com/form",
	//	url.Values{"key": {"Value"}, "id": {"123"}})

}
