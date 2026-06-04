package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func Classify(n int) string {
	switch {
	case n < 0:
		return "neg"
	case n == 0:
		return "zero"
	default:
		return "pos"
	}
}

// ClassifyHandler は GET /classify?n=5 を受け取り、分類結果を本文に返す。
//
//	n=5  → 200 "pos"
//	n=-1 → 200 "neg"
//	n=0  → 200 "zero"
//	n が無い / 数値でない → 400 "n must be an integer"
func ClassifyHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("n")
	n, err := strconv.Atoi(q)
	if err != nil {
		http.Error(w, "n must be an integer", http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, Classify(n))
}

func main() {
	http.HandleFunc("/classify", ClassifyHandler)
	http.ListenAndServe(":8080", nil)
}
