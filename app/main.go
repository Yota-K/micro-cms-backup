package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/micro-cms-backup/config"
)

// MEMO: サムネイル・カテゴリー・タグは一旦後回し
type Content struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	Description string `json:"description"`
}

type MicroCmsBlogs struct {
	Contents []Content `json:"contents"`
}

func main() {
	API_KEY, ENDPOINT := config.Init()
	if API_KEY == "" || ENDPOINT == "" {
		log.Fatal("API_KEYかエンドポイントが設定されていません")
	}

	result, err := reqApi(API_KEY, ENDPOINT)
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}

	createCsv(result)
}

func reqApi(API_KEY, ENDPOINT string) (*[]byte, error) {
	// リクエストを作成
	params := "blogs?fields=id,title,body,description&limit=9999"
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", ENDPOINT, params), nil)
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

func createCsv(jsonBytes *[]byte) {
	blogs := new(MicroCmsBlogs)

	// 構造体blogsにJSONのバイト列を格納する
	if err := json.Unmarshal(*jsonBytes, blogs); err != nil {
		log.Fatal(err)
		return
	}

	// MEMO: ループ内でCSVの書き込みを行うイメージ
	for _, blog := range blogs.Contents {
		fmt.Printf("id: %s\n", blog.Id)
		fmt.Printf("title: %s\n", blog.Title)
		fmt.Print("\n")
	}
}
