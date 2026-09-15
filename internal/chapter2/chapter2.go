package chapter2

import (
	"bytes"
	"io"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest(
		http.MethodGet,
		"http://mock-api/users?age=25",
		nil,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	req.Header.Set("key", "dip")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)

	if _, err := w.Write(body); err != nil {
		return
	}
}

func Post(w http.ResponseWriter, r *http.Request) {
	body := []byte(`{
		"name": "dip 次郎",
		"age": 24
	}`)

	req, err := http.NewRequest(
		http.MethodPost,
		"http://mock-api/users",
		bytes.NewBuffer(body),
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	req.Header.Set("key", "dip")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)

	if _, err := w.Write(responseBody); err != nil {
		return
	}
}