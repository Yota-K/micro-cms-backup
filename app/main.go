package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

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

	// ファイル名の生成
	const timeFmt = "20060102"
	nowTime := time.Now().Format(timeFmt)
	filename := fmt.Sprintf("micro-cms-backup%s.csv", nowTime)

	csvfile, _ := os.Create(filename)
	defer csvfile.Close()
	writer := csv.NewWriter(csvfile)

	// ヘッダ行の書き込み処理
	writer.Write([]string{"id", "title", "body", "description"})

	// ループ内でCSVの書き込みを行う
	for _, blog := range blogs.Contents {
		record := []string{blog.Id, blog.Title, blog.Body, blog.Description}

		if err := writer.Write(record); err != nil {
			log.Fatal(err)
		}
		writer.Write(record)
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		log.Fatal(err)
	}
}
