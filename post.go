// Created by EldersJavas(EldersJavas&gmail.com)

package ddtvgo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func test() error {
	resp, err := http.Post("", "", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))
	return nil
}

func PostReq(urls string, params string) (string, error) {
	data1 := url.Values{"name": {""}, "id": {""}}
	resopne, err := http.PostForm(urls, data1)
	if err != nil {
		panic(err)
	}
	defer resopne.Body.Close()

	// 5. 一次性读取请求到的数据
	body, err := ioutil.ReadAll(resopne.Body)
	if err != nil {
		panic(err)
	}

	return string(body), nil
}
