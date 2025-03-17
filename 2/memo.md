## Github Actionsの実行
[hello.yml](/.github/workflows/hello.yml)

## エラー関連

- インデント、タイポに気をつける
- YAML構文、アンカーやエイリアスは使えない


## 起動方法
手動実行
[manual.yml](/.github/workflows/manual.yml)

CLIからの例
```
gh workflow run manual.yml -f greeting=GoodBye
✓ Created workflow_dispatch event for manual.yml at main

To see runs for this workflow, try: gh run list --workflow=manual.yml
```

もろんGUIからもできる

choice型、指定以外のものだと弾かれる
[choice.yml](/.github/workflows/choice.yml)
```
gh workflow run choice.yml -f log-level=error
✓ Created workflow_dispatch event for choice.yml at main

To see runs for this workflow, try: gh run list --workflow=choice.yml
yoshi@YoshiMacBook-Pro gha-practice % gh workflow run choice.yml -f log-level=hogehoge
could not create workflow dispatch event: HTTP 422: Provided value 'hogehoge' for input 'log-level' not in the list of allowed values (https://api.github.com/repos/yoshi-nip/gha-practice/actions/workflows/150200086/dispatches)
```
