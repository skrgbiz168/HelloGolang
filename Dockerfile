FROM golang:1.23.3-alpine3.20

ENV CGO_ENABLED=0

# このコメント以後のコメントは[/go/src/app/]フォルダーにて実行させるように指定します
WORKDIR /go/src/app/

# ホストPCの[./back]フォルダーの配置されたソースコードを[/go/src/app/]フォルダーにコピーします
COPY ./src .

# OSのインストール済みのパッケージをバージョンアップし、必要なパッケージをインストールします
RUN apk upgrade --update && \
    apk --no-cache add git gcc musl-dev

# dlvツールのインストール
RUN go install github.com/go-delve/delve/cmd/dlv@latest
