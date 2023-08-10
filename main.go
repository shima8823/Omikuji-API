package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
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
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "Usage: omikuji <port>")
		os.Exit(1)
	}
	http.HandleFunc("/", handler)

	// HTTPサーバを起動する
	err := http.ListenAndServe(":"+args[0], nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
