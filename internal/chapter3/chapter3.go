package chapter3

import (
	"encoding/json"
	"io"
	"net/http"
)

// ユーザ情報APIのレスポンス
type User struct {
	UserID int64  `json:"user_id"`
	Name   string `json:"name"`
	Age    int64  `json:"age"`
}

// 案件情報
type Entry struct {
	Name   string `json:"name"`
	UserID int64  `json:"user_id"`
	Salary int64  `json:"salary"`
}

// 案件情報APIのレスポンス
type EntriesResponse struct {
	Entries []Entry `json:"entries"`
}

func Get(w http.ResponseWriter, r *http.Request) {
	// 1. ユーザ情報取得APIへGET
	userReq, err := http.NewRequest(
		http.MethodGet,
		"http://mock-api/users?name=dip%20%E5%A4%AA%E9%83%8E",
		nil,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userReq.Header.Set("key", "dip")

	userResp, err := http.DefaultClient.Do(userReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer userResp.Body.Close()

	// ユーザAPIのJSONを読み取る
	userBody, err := io.ReadAll(userResp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// JSON → Goの構造体へ変換
	var users []User
	if err := json.Unmarshal(userBody, &users); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// dip 太郎が取れなければ終了
	if len(users) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	userID := users[0].UserID

	// 2. 案件情報取得APIへGET
	entryReq, err := http.NewRequest(
		http.MethodGet,
		"http://mock-api/entries",
		nil,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	entryReq.Header.Set("key", "dip")

	entryResp, err := http.DefaultClient.Do(entryReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer entryResp.Body.Close()

	// 案件APIのJSONを読み取る
	entryBody, err := io.ReadAll(entryResp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// JSON → Goの構造体へ変換
	var entriesResponse EntriesResponse
	if err := json.Unmarshal(entryBody, &entriesResponse); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 3. ユーザIDが一致する案件だけ抽出
	var result []Entry
	for _, entry := range entriesResponse.Entries {
		if entry.UserID == userID {
			result = append(result, entry)
		}
	}

	// 4. 抽出結果をJSONで返す
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(result); err != nil {
		return
	}
}