package main

// STEP 02: httptest.NewServer で「本物のサーバ」越しにテストしよう
//
// $ go test -v -count=1
//
// [解説]
// httptest.NewServer(handler) は、空きポートで実際に listen する
// テスト用 HTTP サーバを起動し、*httptest.Server を返す。
//   - ts.URL … 起動したサーバのベース URL（例: http://127.0.0.1:54321）
//   - ts.Close() … サーバを停止（必ず defer で呼ぶ）
// http.Get(ts.URL + "/classify?n=5") のように実クライアントから叩ける。
//
// 問: STEP 01 の NewRecorder 版と何が違いますか？
//     - NewRecorder … ハンドラ関数を「直接呼ぶ」。ネットワーク無し・最速。
//     - NewServer   … 実際に listen し「HTTP 越しに」叩く。ミドルウェアや
//                      クライアント側の挙動も含めて検証できる。
//
// 問: どちらを使うべき？ → 単体のハンドラ検証は NewRecorder、
//     クライアント実装やサーバ全体の結合確認は NewServer が向く。
//
// 注: このテストは ClassifyHandler を直接渡しているので、ルーティング自体は
//     検証していない（どのパスでも同じハンドラに届く）。ルーティングまで
//     試すなら http.NewServeMux に登録したものを NewServer に渡す。

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClassifyHandler_OverHTTP(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(ClassifyHandler))
	defer ts.Close()

	cases := []struct {
		name     string
		path     string
		wantBody string
	}{
		{"positive", "/classify?n=5", "pos"},
		{"negative", "/classify?n=-3", "neg"},
		{"zero", "/classify?n=0", "zero"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := http.Get(ts.URL + tc.path)
			require.NoError(t, err) // 失敗時はここで停止（res が nil の defer panic を防ぐ）
			defer res.Body.Close()

			assert.Equal(t, http.StatusOK, res.StatusCode)

			body, err := io.ReadAll(res.Body)
			assert.NoError(t, err)
			assert.Equal(t, tc.wantBody, string(body))
		})
	}
}
