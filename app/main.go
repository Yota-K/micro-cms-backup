package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/micro-cms-backup/config"
)

func main() {
	result, err := reqApi()
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}
	fmt.Println(string(*result))
}

func reqApi() (*[]byte, error) {
	API_KEY, ENDPOINT := config.Init()

	// リクエストを作成
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/blogs", ENDPOINT), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	req.Header.Set("X-API-KEY", API_KEY)

	// リクエストを実行
	res, err := new(http.Client).Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer res.Body.Close()

	// レスポンスの読み込み
	byteAry, err := ioutil.ReadAll(res.Body)

	return &byteAry, err
}
