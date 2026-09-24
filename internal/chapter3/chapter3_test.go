package chapter3

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	// Chapter3のGetを呼び出すためのリクエストを作成
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "http://localhost/chapter3", nil)

	// Chapter3のGetを実行
	Get(w, r)

	// ステータスコードが200か確認
	assert.Equal(t, http.StatusOK, w.Code)

	// dip太郎(user_id: 123456)に紐づく案件だけ返っているか確認
	got := w.Body.String()

	assert.Contains(t, got, "案件情報1")
	assert.Contains(t, got, `"user_id":123456`)
	assert.Contains(t, got, `"salary":123456`)

	// dip花子(user_id: 234567)の案件は含まれていないことを確認
	assert.NotContains(t, got, "案件情報2")
	assert.NotContains(t, got, `"user_id":234567`)
}