package main

// STEP 01: httptest.NewRecorder でハンドラを単体テストしよう
//
// [PART 1] まずそのまま実行してみよう
// $ go test -v -count=1
// 問: サーバ（http.ListenAndServe）を起動していないのに、なぜハンドラをテストできるのでしょう？
//
// [解説]
// httptest.NewRecorder() は http.ResponseWriter の「録画機」を返す。
// ハンドラに渡して直接呼び出すと、書き込まれたステータス・本文を
// rec.Result() から取り出して検証できる。ネットワークも実サーバも要らない。
//
// httptest.NewRequest(method, target, body) は *http.Request を組み立てる。
// クエリは target に含めればよい（例: "/classify?n=5"）。
//
// [PART 2] 異常系のテストを追加しよう
// 下の cases には「正常系」しか無い。"n が無い" ケースを追加してみよう。
// 問: そのとき期待するステータスコードは？

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClassifyHandler(t *testing.T) {
	cases := []struct {
		name       string
		target     string
		wantStatus int
		wantBody   string
	}{
		{"positive", "/classify?n=5", http.StatusOK, "pos"},
		{"negative", "/classify?n=-1", http.StatusOK, "neg"},
		{"zero", "/classify?n=0", http.StatusOK, "zero"},
		// TODO: "n が無い" 異常系のケースを追加しよう（例: target "/classify"）
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tc.target, nil)
			rec := httptest.NewRecorder()

			ClassifyHandler(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			assert.Equal(t, tc.wantStatus, res.StatusCode)

			body, err := io.ReadAll(res.Body)
			assert.NoError(t, err)
			assert.Equal(t, tc.wantBody, string(body))
		})
	}
}
