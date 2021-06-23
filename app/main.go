package main

import (
	"fmt"
	_ "log"
	_ "net/http"
	_ "io/ioutil"
	_ "encoding/json"

	"github.com/micro-cms-backup/config"
)

func main() {
	API_KEY, ENDPOINT := config.Init()
	fmt.Println(API_KEY)
	fmt.Println(ENDPOINT)

	// client := &http.Client{}
	// header := http.Header{}
	// header.Set("X-API-KEY:", API_KEY)
	// 
	// // リクエストを作成
	// req, err := http.NewRequest("Get", ENDPOINT, nil)
	// if err != nil {
	// 	log.Fatal("リクエスト失敗")
	// }
	// 
	// // リクエストを実行
	// resp, err := client.Do(req)
	// if err != nil {
	// 	log.Fatal("リクエストの実行に失敗")
	// }
	// 
	// fmt.Println(resp.StatusCode)
	// 
	// // レスポンスの読み込み
	// // MEMO: "io/ioutil"・・・データの読み書きに使うパッケージをまとめたやつ
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal("レスポンスの読み込みに失敗しました")
	// }
	// 
	// // MEMO: bodyの型は[]byte
	// // JSONに変換する必要がある
	// // 参考: https://blog.kazu634.com/labs/golang/2020-03-07-how-to-use-net-http/
	// fmt.Println(body)
}
