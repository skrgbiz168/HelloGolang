FROM golang:1.23.3-alpine3.20

# このコメント以後のコメントは[/go/src/]フォルダーにて実行させるように指定します
WORKDIR /go/src/

# ホストPCの[./back]フォルダーの配置されたソースコードを[/go/src/]フォルダーにコピーします
COPY ./ .

# OSのインストール済みのパッケージをバージョンアップし、必要なパッケージをインストールします
RUN apk upgrade --update && \
    apk --no-cache add git gcc musl-dev

# dlvツールのインストール
RUN go install github.com/go-delve/delve/cmd/dlv@latest
