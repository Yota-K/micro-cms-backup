package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	_ "encoding/json"

	"github.com/micro-cms-backup/config"
)

func main() {
	API_KEY, ENDPOINT := config.Init()

	// リクエストを作成
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/blogs", ENDPOINT), nil)
	if err != nil {
		log.Fatal("request error")
	}

	req.Header.Set("X-API-KEY", API_KEY)

	// リクエストを実行
	res, err := new(http.Client).Do(req)
	if err != nil {
		log.Fatal("response error")
	}
	defer res.Body.Close()

	// レスポンスの読み込み
	byteAry, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Read error")
	}

	fmt.Println(string(byteAry))
}
