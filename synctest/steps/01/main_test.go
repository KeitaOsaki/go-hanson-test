package main

// STEP 01: testing/synctest で「待たない」テストを書こう（Go 1.25 で正式化）
//
// [PART 1] まず素朴なテスト TestClassifyReal を実行してみよう
// $ go test -v -count=1 -run TestClassifyReal
// 問: Classify は 1 秒 Sleep する。3 ケースで合計何秒かかりましたか？
//
// [PART 2] synctest 版 TestClassifySynctest を実行してみよう
// $ go test -v -count=1 -run TestClassifySynctest
// 問: 同じ 1 秒 Sleep を 3 回呼んでいるのに、実時間は何秒でしたか？
//
// [解説]
// synctest.Test(t, f) は f を「バブル」と呼ぶ隔離環境で実行する。
// バブルの中では time が「偽の時計（fake clock）」になり、
// バブル内の全ゴルーチンが待ち（durably blocked）になると時計が一気に進む。
// そのため time.Sleep(1s) も実時間を消費せず即座に完了する。
// → 「時間に依存するコード」を、実時間を待たずに決定論的にテストできる。

import (
	"testing"
	"testing/synctest"

	"github.com/stretchr/testify/assert"
)

// 素朴な逐次テスト：実時間で 1 ケースあたり約 1 秒かかる。
func TestClassifyReal(t *testing.T) {
	cases := []struct {
		name string
		in   int
		want string
	}{
		{"negative", -1, "neg"},
		{"zero", 0, "zero"},
		{"positive", 5, "pos"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Classify(tc.in)
			assert.Equal(t, tc.want, got)
		})
	}
}

// synctest 版：fake clock により Sleep を待たずに即座に終わる。
func TestClassifySynctest(t *testing.T) {
	cases := []struct {
		name string
		in   int
		want string
	}{
		{"negative", -1, "neg"},
		{"zero", 0, "zero"},
		{"positive", 5, "pos"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				got := Classify(tc.in)
				assert.Equal(t, tc.want, got)
			})
		})
	}
}
