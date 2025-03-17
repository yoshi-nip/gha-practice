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
