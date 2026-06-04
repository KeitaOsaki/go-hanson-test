package main

// STEP 02: タイマー／タイムアウトを synctest でテストしよう
//
// $ go test -v -count=1
//
// [解説]
// WaitForSignal は time.After(d) でタイムアウトを実現している。
// もし synctest を使わずに「5 秒でタイムアウトする」ことを確かめると、
// テストは本当に 5 秒待たされる。
// synctest.Test のバブル内なら fake clock が時間を一気に進めるので、
// 5 秒のタイムアウトも、2 秒後に届くシグナルも、実時間ゼロで検証できる。
//
// [観察ポイント]
//   - TestWaitForSignal_Timeout : done を閉じない → タイムアウトで false
//   - TestWaitForSignal_Fired    : 別ゴルーチンが 2 秒後に done を閉じる → true
//
// 問: それぞれ d=5s なのに、テスト全体は何秒で終わりましたか？
//
// 問: バブル内で起動したゴルーチン（go func ...）はどう扱われる？
//     → synctest はバブル内の全ゴルーチンの「待ち」を監視し、
//       全員が待ちに入った時だけ fake clock を進める。
//       time.Sleep(2s) で待っているゴルーチンも fake clock で即座に進む。

import (
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWaitForSignal_Timeout(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		done := make(chan struct{}) // 誰も閉じない

		got := WaitForSignal(done, 5*time.Second)

		assert.False(t, got) // タイムアウトするはず
	})
}

func TestWaitForSignal_Fired(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		done := make(chan struct{})
		go func() {
			time.Sleep(2 * time.Second) // タイムアウト(5s)より前に
			close(done)
		}()

		got := WaitForSignal(done, 5*time.Second)

		assert.True(t, got) // シグナルが先に届くはず
	})
}
