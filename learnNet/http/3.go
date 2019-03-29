package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func main()  {
	//要管理代理、TLS配置、keep-alive、压缩和其他设置，创建一个Transport：
	tr := &http.Transport{

		Dial: (&net.Dialer{
			Timeout: 10 * time.Second,
		}).Dial,
		MaxIdleConns: 2,

		//SSL连接专用 指定tls.Client所用的TLS配置信息，如果不指定，也会使用默认的配置
		//	TLSClientConfig: &tls.Config{RootCAs: pool}
		//是否取消长连接，默认false 即启用长连接
		DisableKeepAlives:false,
		//是否取消压缩默认值false，即启用压缩
		DisableCompression: true,
		//指定每个请求的目标主机之间最大非活跃连接（keep-alive）数量，如果不指定，默认使用
		MaxIdleConnsPerHost:10,

	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("http://test-api.zu.ke.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	result, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(result))
}