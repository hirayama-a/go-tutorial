package chapter1

import (
	"fmt"
	"net/http"
)

func GetEcho(w http.ResponseWriter, r *http.Request) {
	// FIXME: Getメソッドのアクセスか確認
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// FIXME: パラメータをFormに変換する
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// FIXME: パラメータを取得する
	form := r.Form

	// FIXME: レスポンスコード設定
	w.WriteHeader(http.StatusOK)

	// FIXME: パラメータをレスポンスに書き出す
	if _, err := fmt.Fprint(w, form); err != nil {
		return
	}
}
