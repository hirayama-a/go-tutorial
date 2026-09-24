package chapter2

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	// Chapter2のGetを呼び出すためのリクエストを作成
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "http://localhost/chapter2/get", nil)

	// Chapter2のGetを実行
	Get(w, r)

	// ステータスコードが200か確認
	assert.Equal(t, http.StatusOK, w.Code)

	// mock-apiから取得したユーザ情報が返っているか確認
	got := w.Body.String()
	assert.Contains(t, got, "dip 太郎")
	assert.Contains(t, got, "dip 花子")
	assert.Contains(t, got, `"age":25`)
}

func TestPost(t *testing.T) {
	// Chapter2のPostを呼び出すためのリクエストを作成
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "http://localhost/chapter2/post", nil)

	// Chapter2のPostを実行
	Post(w, r)

	// ステータスコードが200か確認
	assert.Equal(t, http.StatusOK, w.Code)

	// mock-apiへ登録したユーザ情報が返っているか確認
	got := w.Body.String()
	assert.Contains(t, got, "dip 次郎")
	assert.Contains(t, got, `"age":24`)
}