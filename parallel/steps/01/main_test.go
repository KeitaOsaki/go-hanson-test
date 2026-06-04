package main

// STEP 01: テストの実行と並列化
//
// [PART 1] まずそのまま実行してみよう
// $ go test -v -count=1
// 問: 合計何秒かかりましたか？
//
// [PART 2] 各テスト関数の先頭に t.Parallel() を追加してみよう
// 問: 実行時間はどう変わりましたか？
// 問: なぜそのような結果になるのでしょう？

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClassifyNegative(t *testing.T) {
	got := Classify(-1)
	assert.Equal(t, "neg", got)
}

func TestClassifyZero(t *testing.T) {
	got := Classify(0)
	assert.Equal(t, "zero", got)
}

func TestClassifyPositive(t *testing.T) {
	got := Classify(5)
	assert.Equal(t, "pos", got)
}
