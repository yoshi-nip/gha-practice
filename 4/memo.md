## 継続的インテグレ-ションで大切にするべきこと

- クリーンに保つ
- 高速に実行する
- ノイズを減らす

## テストの運用について

- ユニットテストを中心にする
- たまに落ちるテストを誤魔化さない
- リファクタリングで落ちるテストを根絶する
- スローテストの実行タイミングをずらす(これは何かとのトレードオフの場合もあるので臨機応変に)


## テスト

ソースコードへのアクセスにチェックアウトが必要
言語のセットアップ
テストステップ

[go_test.yml](./bk_workflows/go_test.yml)

## 静的解析
タムアウト
デフォルトシェル
自動キャンセル

基本はタイムアウトつけた方がいい

[static_analysis.yml](./bk_workflows/static_analysis.yml)
[timeout.yml](./bk_workflows/timeout.yml)

## 環境変数
[environment.yml](./bk_workflows/environment.yml)
[environment_override.yml](./bk_workflows/environment_override.yml)
