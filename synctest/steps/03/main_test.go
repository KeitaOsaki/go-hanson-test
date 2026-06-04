package main

// STEP 03: synctest.Test + synctest.Wait で「遅い」と「flaky」を“同時に”解決しよう
//
// 題材: TTL 付きキャッシュ。Set した値は ttl 経過後、time.AfterFunc が起動する
// 「別ゴルーチン」によって自動削除される。
// → テストしたいこと:「ttl が過ぎたら値が消えていること」。
//
// 素朴な方法は time.Sleep で ttl 分だけ待ってから Get で確認すること。
// だがこれは 2 つの課題を同時に抱える：
//
//   (A) 実行時間が伸びる … ttl 分だけ毎回“本物の時間”を待つ（ttl=1s なら 1 秒）。
//   (B) flaky になる      … 削除は別ゴルーチン。ttl ちょうど待っても、削除ゴルーチンが
//                            走り終わる前に Get するとまだ値が残っていて、たまに失敗する。
//                            安全のため余分に待つ → さらに遅くなり (A) が悪化する。
//
// → 素朴な Sleep では「待ちを短く（速いが flaky）」か「待ちを長く（安定だが遅い）」の
//    トレードオフから逃げられない。(A) と (B) は同時には消せない。
//
// synctest.Test + synctest.Wait はこれを“同時に”解決する：
//   - fake clock が ttl の待ちを即座に進める          → 速い（課題A解決）
//   - synctest.Wait() が削除ゴルーチンの完了を確実に待つ → 決定論的・flaky でない（課題B解決）
//
// ★ ポイント: time.Sleep は「fake clock を進める」役、synctest.Wait は
//   「進めた結果として起動した非同期処理が“終わる”のを待つ」役。両者で 1 セット。
//   （synctest.Wait はゴルーチンが durably blocked になるまで待つ。Sleep 中の
//     ゴルーチンも durably blocked なので、Wait“だけ”では完了を待てない点に注意）
//
// ----------------------------------------------------------------
// [PART 1] 素朴版の遅さを体感しよう
//   $ go test -v -count=1 -run TestCache_Sleep
//   問: ttl=1s。テストは何秒かかりましたか？
//
// [PART 2] flaky（結果が不安定）を再現しよう
//   safeMargin を 0 にする（= ttl ちょうどしか待たない）と、削除ゴルーチンが
//   走り終わる前に Get してしまう。多くの回で FAIL し、まれに PASS する
//   ——つまり結果がゴルーチンのスケジューリング次第で安定しない。
//   $ go test -count=50 -run TestCache_Sleep
//   問: 結果は毎回同じになりますか？ 安定しないのはなぜでしょう？
//
// [PART 3] synctest 版を実行しよう
//   $ go test -v -count=1 -run TestCache_Synctest
//   問: 実行時間は？ -count=200 でも安定して PASS しますか？
// ================================================================

import (
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/assert"
)

const ttl = 1 * time.Second

// 削除ゴルーチンの完了を「祈って」余分に待つ安全マージン。
// 0 にすると ttl ちょうどしか待たず、削除完了前に Get して flaky になりうる。
const safeMargin = 50 * time.Millisecond

// 【課題】素朴な time.Sleep 待ち。
// safeMargin を足して安定させているが、その分だけ毎回遅い（課題A）。
// safeMargin=0 にすると速くなるが削除ゴルーチンと競合して flaky（課題B）。
func TestCache_Sleep(t *testing.T) {
	c := NewCache()
	c.Set("k", "v", ttl)

	// ttl 経過＋削除ゴルーチン完了を“祈って” safeMargin だけ余分に待つ
	time.Sleep(ttl + safeMargin) // TODO: safeMargin を 0 にして -count=50 で flaky を観察しよう

	_, ok := c.Get("k")
	assert.False(t, ok) // 期限切れで消えているはず
}

// 【解決】synctest.Test + synctest.Wait。
// fake clock で ttl を即座に進め、Wait で削除ゴルーチンの完了を確実に待つ。
// → 速い（実時間ゼロ）かつ決定論的（flaky でない）。2 つの課題を同時に解決。
func TestCache_Synctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		c := NewCache()
		c.Set("k", "v", ttl)

		// (1) fake clock を ttl だけ進める → AfterFunc が発火し削除ゴルーチンが起動
		time.Sleep(ttl)
		// (2) 起動した削除ゴルーチンが「完了」するまで確実に待つ（勘のマージンは不要）
		synctest.Wait()

		_, ok := c.Get("k")
		assert.False(t, ok)
	})
}
