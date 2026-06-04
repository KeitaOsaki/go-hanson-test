package main

// STEP 03: 実行タイミングと -p フラグの挙動を観察しよう
//
// ================================================================
// PART 1: サブテストの「バリア」を観察しよう
// ----------------------------------------------------------------
// $ go test -v -count=1 -run TestClassify
//
// TestClassifySeq（逐次）と TestClassifyPar（並列）のログ出力順を比べよう。
//
// [手順]
// 1. まず TODO を追加せずに実行して TestClassifySeq のログ順を確認しよう
// 2. TestClassifyPar の TODO に t.Parallel() を追加して再度実行しよう
// 3. 2 つのテストのログ出力順を比較しよう
//
// [観察ポイント]
// - ">>> loop end" は各サブテストの ">>> [name] start" より前に出ますか？後ですか？
// - -v 出力の "=== PAUSE" と "=== CONT" に注目しよう
//
// [解説]
// t.Parallel() をサブテスト内で呼ぶと、そのサブテストは一時停止（PAUSE）し
// 親テストのループが全て終わった後に再開（CONT）される。
// これをサブテストの「バリア」と呼ぶ。
// ================================================================
//
// ================================================================
// PART 2: -p フラグでパッケージ間の並列度を制御しよう
// ----------------------------------------------------------------
// pkga/ と pkgb/ の classify_test.go を開いて確認しよう。
//
//   (1) デフォルト（パッケージ並列）
//       $ go test -v -count=1 ./pkga/ ./pkgb/
//
//   (2) パッケージを逐次実行（-p 1）
//       $ go test -v -p 1 -count=1 ./pkga/ ./pkgb/
//
// 問: (1) と (2) の実行時間の違いは何でしょう？
// 問: t.Parallel() を追加するとどう変わりますか？
// ================================================================

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 逐次処理（t.Parallel なし）
func TestClassifySeq(t *testing.T) {
	cases := []struct {
		name string
		in   int
		want string
	}{
		{"negative", -1, "neg"},
		{"zero", 0, "zero"},
		{"positive", 5, "pos"},
	}

	t.Log(">>> loop start")
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Logf(">>> [%s] start", tc.name)
			got := Classify(tc.in)
			assert.Equal(t, tc.want, got)
			t.Logf(">>> [%s] done", tc.name)
		})
	}
	t.Log(">>> loop end")
}

// サブテストの並列化
func TestClassifyPar(t *testing.T) {
	cases := []struct {
		name string
		in   int
		want string
	}{
		{"negative", -1, "neg"},
		{"zero", 0, "zero"},
		{"positive", 5, "pos"},
	}

	t.Log(">>> loop start")
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// TODO: t.Parallel() を追加しよう
			//       → ログの出力順がどう変わるか観察しよう

			t.Logf(">>> [%s] start", tc.name)
			got := Classify(tc.in)
			assert.Equal(t, tc.want, got)
			t.Logf(">>> [%s] done", tc.name)
		})
	}
	t.Log(">>> loop end")
}
