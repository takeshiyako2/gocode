package main

import (
    "fmt"
    "net/http"
    "log"
    "html/template"
    "time"
"strconv"
)

type TemplateData struct {
    Title string
    Datetime string
    Unixtime string
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    // 親テンプレートindex.htmlと、小テンプレートbody.htmlを用意
    tmpl := template.Must(template.ParseFiles("views/index.html", "views/body.html"))

    // タイトル
    title := "現在の時刻"

    // bodyに表示するコンテンツ 
    // 時刻オブジェクトを文字列に変換
    datetime := fmt.Sprint(time.Now())
//    unixtime := fmt.Sprint(time.Now().Unix())
    unixtime := strconv.FormatInt(time.Now().Unix(), 10)

    // テンプレートを実行して出力
    templatedata := TemplateData{title, datetime, unixtime}
    if err := tmpl.ExecuteTemplate(w, "base", templatedata); err != nil {
        fmt.Println(err)
    }
}

func main() {
    http.HandleFunc("/", HelloServer) 
    log.Printf("Start Go HTTP Server")
    err := http.ListenAndServe(":4000", nil)
    if err != nil {
       log.Fatal("ListenAndServe: ", err)
    }
}

