# ベースとなるDockerイメージ
FROM golang:1.17

# ソース格納先
ENV SRC_DIR=/go/src/github.com/

# ワーキングディレクトリの指定
WORKDIR $SRC_DIR

ENV APP_USER app
ARG USER_ID
ARG GROUP_ID
RUN groupadd --gid $GROUP_ID app && useradd -m -l --uid $USER_ID --gid $GROUP_ID $APP_USER

# 依存モジュールをインストール
RUN go get github.com/beego/bee
RUN go get -u github.com/astaxie/beego