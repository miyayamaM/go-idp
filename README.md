# go-idp

### DB migration の更新

1. `/cmd/migration/migrations` に `YYYYMMDDHHSS_<migrationの説明>.up.sql` および `YYYYMMDDHHSS_<migrationの説明>.down.sql` 形式で マイグレーションファイルを作成する

2.

```shell

 go run ./cmd/migration db migrate

```

### ローカルサーバー

```shell

go run ./cmd/api

```
