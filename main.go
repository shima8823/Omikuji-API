package main

import (
	"math/rand"
	"net/http"
	"text/template"
)

type data struct {
	Name   string
	Result string
}

// おみくじの結果を保持するスライス
var results = []string{
	"大吉",
	"中吉",
	"小吉",
	"吉",
	"凶",
}

var tmpl = template.Must(template.New("data").
	Parse("<html><body>{{.Name}}さんの運勢は「<b>{{.Result}}</b>」です</body></html>"))

func handler(w http.ResponseWriter, r *http.Request) {
	// おみくじの結果をランダムに選ぶ
	n := rand.Intn(len(results))
	result := results[n]
	data := data{
		Name:   r.FormValue("p"),
		Result: result,
	}
	tmpl.Execute(w, data)
}

func main() {

	// HTTPハンドラを設定する
	http.HandleFunc("/", handler)

	// HTTPサーバを起動する
	http.ListenAndServe(":8080", nil)
}
