## 使用している主な言語
- Go
- Posgre
- MaterialUI
- firebase

## 検証方法
使用しているディレクトリに移動
```
docker exec -it go-app /bin/sh
```

以下コマンドを実行(一例です)
```
go run main.go

GO_ENV=dev go run main.go
```

## DBの構築について
docker-compose up -d --buildをした後、以下コマンドで構築してください
```
RUN GO_ENV=dev go run migrate/migrate.go
```

## 参考URL
https://liginc.co.jp/blog/tech/638764\

大元のGitHub<br>
https://github.com/GomaGoma676/echo-rest-api