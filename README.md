#  二郎スタンプラリー BE

## 概要

- golang v1.20
- gorm(予定) → 使わずに頑張りたい

## 環境構築

Dokcer, DockerComposeがインストールされている前提です。

以下のコマンドで[Air](https://github.com/cosmtrek/air)が起動します。

```bash
docker compose up -d
```

`http://localhost:3000/health`をブラウザで開いて、`{"message":"OK"}`が表示されることを確認してください。

Airはファイルの変更検知してホットリロードします。
