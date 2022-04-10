# go
FROM golang:1.18.0-alpine
# apkのアップデートとgitのインストール
RUN apk add --update && apk add git
# appディレクトリの作成
RUN mkdir /src

# AIRのインストール
RUN go install github.com/cosmtrek/air@latest

# ワーキングディレクトリの設定
WORKDIR /src
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD ../src /src

RUN go mod tidy

CMD ["air", "-c", ".air.toml"]
