## コンテキスト
githubコンテキスト
runnerコンテキスト

がある

githubは多いけど、runnnerはそこまで多くない

[context.yml](/.github/workflows/context.yml)

## 環境変数
[environment.yml](/.github/workflows/environment.yml)
[environment_override.yml](/.github/workflows/environment_override.yml)


コンテキストを直接埋め込むのはアンチパターン
環境変数経由で伝える。
こういうのを中間環境変数という。

ルールとして
1. コンテキストはシェルコマンドに直接ハードコーディングせず、環境変数経由で渡す
2. 環境変数はダブルクオートで囲む`""`

## Variabels と Secrets
Variabels - 複数のworkflowでも使える環境変数だと思っていただければok

Secrets  - 秘匿性の高いものを扱う場合に使用する。複数のworkflowでも使える
- 登録した値は暗号化される
- ログ出力時にマスクされるが、空白などが入ると普通に出力されたりするので注意
- 登録後に値が全く確認できなくなる

[variables_secrets.yml](/.github/workflows/variables_secrets.yml)
