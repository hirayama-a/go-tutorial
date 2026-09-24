package chapter2

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

func Get(w http.ResponseWriter, _ *http.Request) {
	// mock API向けのGETリクエストを作成
	// クエリパラメータ age=25 もURLに付ける
	req, err := http.NewRequest(
		http.MethodGet,
		"http://mock-api/users?age=25",
		nil, // GETなのでリクエストボディはなし
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// README指定のヘッダを設定
	req.Header.Set("key", "dip")

	// mock APIへリクエスト送信
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// この関数終了時にレスポンスボディを閉じる
	defer resp.Body.Close()

	// mock APIから返ってきたレスポンスボディを全部読み取る
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// mock APIのステータスコードをそのまま返す
	w.WriteHeader(resp.StatusCode)

	// mock APIのレスポンスボディをそのまま返す
	if _, err := w.Write(body); err != nil {
		return
	}
}

func Post(w http.ResponseWriter, _ *http.Request) {
	// mock APIへ送るフォームデータを作成
	form := url.Values{}
	form.Set("name", "dip 次郎")
	form.Set("age", "24")

	// POSTリクエストを作成
	req, err := http.NewRequest(
		http.MethodPost,
		"http://mock-api/users",
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// README指定の認証用ヘッダ
	req.Header.Set("key", "dip")

	// mock実装に合わせてフォーム形式を指定
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// mock APIへリクエスト送信
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// mock APIのレスポンスボディを読み取る
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// mock APIのステータスコードをそのまま返す
	w.WriteHeader(resp.StatusCode)

	// mock APIのレスポンスボディをそのまま返す
	if _, err := w.Write(responseBody); err != nil {
		return
	}
}