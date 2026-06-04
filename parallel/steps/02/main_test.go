package main

// STEP 02: テーブル駆動テストとサブテストの並列化
//
// $ go test -v -count=1
//
// PART 1: STEP 01 と同じ検証をできるようにしよう！
// PART 2: 並列化しよう！
//
// 問: 追加前と後で実行時間はどう変わりましたか？
// 問: step01 との違いは何でしょう？
//
// PART 2でのログの期待値
// ~/steps/01  go test -v -count=1
// === RUN   TestClassifyNegative
// === PAUSE TestClassifyNegative
// === RUN   TestClassifyZero
// === PAUSE TestClassifyZero
// === RUN   TestClassifyPositive
// === PAUSE TestClassifyPositive
// === CONT  TestClassifyNegative
// === CONT  TestClassifyPositive
// === CONT  TestClassifyZero
// --- PASS: TestClassifyNegative (1.00s)
// --- PASS: TestClassifyZero (1.00s)
// --- PASS: TestClassifyPositive (1.00s)

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClassify(t *testing.T) {
	cases := []struct {
		name string
		in   int
		want string
	}{
		// cases? ココが怪しい
		{"zero", 0, "zero"},
	}
	// 並列化 STEP01で何を追加しましたか？
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Classify(tc.in)
			assert.Equal(t, tc.want, got)
		})
	}
}
