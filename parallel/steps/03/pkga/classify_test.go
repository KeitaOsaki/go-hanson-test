package pkga

// PART 2: -p フラグでパッケージ間の並列度を制御しよう
//
// pkgb/ も同じ構造になっている。
// steps/03/ ルートから以下のコマンドで実行しよう:
//
//   (1) デフォルト（パッケージ並列）
//       $ go test -v -count=1 ./pkga/ ./pkgb/
//
//   (2) パッケージを逐次実行（-p 1）
//       $ go test -v -p 1 -count=1 ./pkga/ ./pkgb/
//
// 問: (1) と (2) の実行時間の違いは何でしょう？
// 問: t.Parallel() を追加するとどう変わりますか？
//
// 実行時間の目安（各パッケージ 3 ケース × 1 秒）:
//
//   t.Parallel なし, -p 1 → pkga(3s) + pkgb(3s) = 約 6 秒
//   t.Parallel なし, -p 2 → pkga と pkgb を並列   = 約 3 秒
//   t.Parallel あり, -p 1 → pkga(1s) + pkgb(1s) = 約 2 秒
//   t.Parallel あり, -p 2 → pkga と pkgb を並列   = 約 1 秒

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
		{"negative", -1, "neg"},
		{"zero", 0, "zero"},
		{"positive", 5, "pos"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// TODO: t.Parallel() を追加しよう

			got := Classify(tc.in)
			assert.Equal(t, tc.want, got)
		})
	}
}
