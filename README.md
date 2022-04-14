# Beego で作ったごく初期機能の掲示板

## Link

- [100行未満かつGo標準ライブラリだけで作る掲示板](https://news.mynavi.jp/techplus/article/gogogo-9/)
- [[Go言語]Beegoフレームワークでアプリ開発 ～環境構築編～](https://selfnote.work/20191001/programming/golang-beego-environment/)

## 使い方

user idとgroup idを.envに書き込む  
【注意】このパラメータを指定してdocker-compose upしたら、コンテナ側で作られるユーザー権限が自分になれるかと思ったが上手く行っていない。

    cd beego-app
    bash ./make_env.sh

コンテナ作って起動する

    docker-compose build -d --build

下記URLでアクセスする。

    http://0.0.0.0:28080/bbs
