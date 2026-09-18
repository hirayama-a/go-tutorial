package main

import (
	"log"
	"net/http"

	"github.com/dip-dev/go-tutorial/internal/chapter1"
	"github.com/dip-dev/go-tutorial/internal/chapter2"
	"github.com/dip-dev/go-tutorial/internal/chapter3"
)

func main() {
	mux := http.NewServeMux()

	// Chapter1: エコーAPI
	mux.HandleFunc("/echo", chapter1.GetEcho)

	// Chapter2: mock APIへGETするAPI
	mux.HandleFunc("/chapter2/get", chapter2.Get)

	// Chapter2: mock APIへPOSTするAPI
	mux.HandleFunc("/chapter2/post", chapter2.Post)

	// Chapter3: ユーザ情報と案件情報を取得して対象案件を返すAPI
	mux.HandleFunc("/chapter3", chapter3.Get)

	// 8080ポートでHTTPサーバー起動
	if err := http.ListenAndServe("0.0.0.0:8080", mux); err != nil {
		log.Fatalf("failed to launch service: %+v", err)
	}
}