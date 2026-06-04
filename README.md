# 体で学ぶGoの並列テストの仕組み！

Goのテストツールがどのように動いているのかを、手を動かしながら学ぶハンズオンです。

### 進め方

トピックごとにディレクトリが分かれており、それぞれの中に `steps/` として
STEP が順番に並んでいます。各 STEP は独立した Go モジュール（`module gocon2026`,
Go 1.25.7）になっているので、その STEP のディレクトリに移動してテストを実行・
実験しながら進めてください。

各 STEP には `// TODO:` や問い・観察ポイントのコメントが書かれています。
それらに沿って実際にテストを動かし、出力（経過時間や `=== PAUSE` / `=== CONT`
など）を読みながら理解を深めていきます。

唯一の外部依存は `github.com/stretchr/testify`（assert）です。

### 内容について

3つのトピックを通して、Goのテストの実行と並列化の仕組みを学びます。

#### `parallel/` — テストの並列化

- `parallel/steps/01/` — テストの実行と、トップレベルのテスト関数を `t.Parallel()` で並列化する
- `parallel/steps/02/` — テーブル駆動テストと、サブテスト（`t.Run` + `t.Parallel`）の並列化
- `parallel/steps/03/` — サブテストの「バリア」（PAUSE/CONT）の観察と、`-p` フラグによる
  パッケージ間の並列度の制御。パッケージ並列を示すための同一パッケージ `pkga/` と `pkgb/` を含む

#### `httptest/` — HTTPハンドラのテスト（サーバ側）

- `httptest/steps/01/` — `httptest.NewRecorder` + `httptest.NewRequest` でハンドラを直接ユニットテストする
- `httptest/steps/02/` — `httptest.NewServer` で実際のHTTP越しに同じハンドラをテストする

#### `synctest/` — `testing/synctest`

- `synctest/steps/01/` — `synctest.Test` のバブル / 仮想クロックで `time.Sleep` を実時間ゼロに
  畳み込む。実時間テストと synctest テストを対比する
- `synctest/steps/02/` — 仮想クロックの下で、タイマー / タイムアウト（`time.After` と
  バックグラウンドの goroutine）を決定的にテストする
- `synctest/steps/03/` — `synctest.Test` + `synctest.Wait` で「遅い」かつ「flaky」を同時に解決する。
  TTLキャッシュ（`time.AfterFunc`）を題材に、`time.Sleep` で仮想クロックを進めてタイマーを発火させ、
  `synctest.Wait` で非同期の削除 goroutine が終わるまで待つ

