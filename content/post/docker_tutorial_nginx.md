+++
Categories = ["Docker", "OSX"]
Description = "Docker で nginx を構築する"
Tags = ["Env", "Infra"]
comments = true
date = "2015-09-20T15:49:48+09:00"
title = "Docker Part 3 - Docker で nginx を構築する"

+++

[docker-machine コマンドの使用方法・流れ]({{< ref "post/docker_machine.md" >}}) で、docker vm の作成までは出来るようになったので、今回 docker vm に nginx container を立ち上げたいと思う。ここからが環境構築の本番。はやく MEAN とか LAMP 環境を構築したい思うが、急がば回れです。

以前、何度か [Docker入門 (全11回)]( http://dotinstall.com/lessons/basic_docker) で docker に入門しましたが、その後挫折。 入門して nginx 構築を出来たはいいが、 MEAN とか LAMP の開発環境構築するのがすごく困難な感じがして。[docker コマンド](https://docs.docker.com/reference/commandline/cli/) たくさんあるし、複数 container のリンクとか、データの永続化とか、ググればググるほど手強いぞ docker ってなってた・・・。

が、時が経過した今、楽になってますよ。情報も増えてるし、何より docker-compose （複数コンテナの構築）が便利だった。復習も兼ねて、まずは nginx を構築!!

ただし、以下を読むより動画 [Docker入門 (全11回)]( http://dotinstall.com/lessons/basic_docker) で学ぶほうが分かりやすい。

<!--more-->
    
# Step1 docker vm の作成(復習)

とりあえず docker vm を作成・利用する。忘れた方は、こちらで[docker-machine コマンドの使用方法・流れ]({{< ref "post/docker_machine.md" >}})

~~~bash
$ docker-machine create --driver virtualbox mynginx
$ eval "$(docker-machine env mynginx)"
$ docker-machine ls
NAME            ACTIVE   DRIVER       STATE     URL                         SWARM
mynginx         *        virtualbox   Running   tcp://192.168.99.100:2376
~~~

これで上記の vm にたいして docker コマンドを発行できので、叩いてみる。コンテナがないので何も表示されない。

~~~bash
$ docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES

# 参考）eval で vm を指定しなかった場合
$ docker ps
Get http:///var/run/docker.sock/v1.20/containers/json: dial unix /var/run/docker.sock: no such file or directory.
* Are you trying to connect to a TLS-enabled daemon without TLS?
* Is your docker daemon up and running?
~~~

# Step1.5 簡単な全体イメージ

以下で使用する用語の全体イメージは以下のような感じである。

**Dockerfile** から **Image** を作成、Image から **Container** を立ち上げる

# Step2 Dockerfile の作成

Dockerfile （ Image 作成の手順書）を作成する。dir 構成は以下の通り。

~~~bash
mynginx/
└── Dockerfile
~~~

Dockerfile は以下の通り。

~~~bash
# Docker file
FROM nginx:latest                        # イメージの指定
MAINTAINER mizuki-y<mizuki-y@syati.info> # 作成者の情報
~~~

上記で手順書と書いたが、２行しかない手順書である。
ただ単に、Image に[Official Repository nginx](https://hub.docker.com/_/nginx/)を利用しますと記述しただけ。
**RUN**、**CMD** とか他にも命令はあるが、今回はシンプルに。いますぐ知りたい！！という場合は、以下を参考に。

- [Dockerfileとdocker buildコマンドでDockerイメージの作成](http://www.atmarkit.co.jp/ait/articles/1407/08/news031.html)
- [Dockerfile reference](https://docs.docker.com/reference/builder/)

# Step3 Image の作成

Dockerfile から Image を 作成するため **build** コマンドを実施する。

~~~bash
# build コマンドのヘルプ見てみる
$ docker build --help

Usage:  docker build [OPTIONS] PATH | URL | -

Build a new image from the source code at PATH
## 多いので省略して、今回利用するものだけ
  -t, --tag=                      Repository name (and optionally a tag) for the image

# いろいろありますが、とりあえず OPTIONS -t （名前付け）だけ覚えればいいでしょう。PATH には Dockerfile パスを。
$ docker build -t syati:nginx .
~~~

# Step4 Image の確認

作成した Image を確認するため **images** コマンドを実施する。 Step2 で FROM しか指定していないので何も変わっていない（SIZE 同じ）が Dockerfile から Image を作成できている。

~~~bash
$ docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             VIRTUAL SIZE
syati               nginx               a232c3216f20        9 seconds ago       132.9 MB
nginx               latest              0b354d33906d        10 days ago         132.9 MB　   # ベースとなったイメージ

# 参考） -t オプションつけないと <none> になる
$ docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             VIRTUAL SIZE
<none>              <none>              dd06038a778d        10 minutes ago      132.9 MB
nginx               latest              0b354d33906d        10 days ago         132.9 MB
~~~

# Step5 Container の動作を確認する

作成した Image から Container を立ち上げるために **run** コマンドを実施する

~~~bash
# run コマンドのヘルプ見てみる
$ docker run --help

Usage:	docker run [OPTIONS] IMAGE [COMMAND] [ARG...]

Run a command in a new container
## 多いので省略して、今回利用するものだけ
  -d, --detach=false              Run container in background and print container ID
  -i, --interactive=false         Keep STDIN open even if not attached
  -P, --publish-all=false         Publish all exposed ports to random ports
  -p, --publish=[]                Publish a container's port(s) to the host

# コンテナを立ち上げる
$ docker run -P -d syati:nginx
5c70e0b84c91dc04d56fab9b9967765716008bfcb4d1fe102ba34d305994541b
## -P は コンテナで開いているポートと docker vm のポートをランダムにリンクする(以下の PORTS を見るとわかる）
## -d は コンテナをバックグラウンドで実行する(バックグラウンドで実行しないと次のコマンド打てなくなるから大事)

# ps コマンドで 動作を確認する
$ docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                                           NAMES
5c70e0b84c91        syati:nginx       "nginx -g 'daemon off"   46 seconds ago      Up 45 seconds       0.0.0.0:32771->80/tcp, 0.0.0.0:32770->443/tcp   trusting_lumiere


# 参考) -P でランダムなポートでリンクされても困っちゃうという場合は、以下のとおり -p を使って指定してあげる
$ docker run -p 8000:80 -d syati:nginx
2ac58799de9cfdceb02abf2b6050761801641b93c820097048cb46a9c0be3e6a

$ docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                           NAMES
2ac58799de9c        syati:nginx       "nginx -g 'daemon off"   3 seconds ago       Up 3 seconds        443/tcp, 0.0.0.0:8000->80/tcp   backstabbing_einstein
~~~

Dockerfile のことを少しでも知っていたら、以下の点を疑問に思うかもしれない。

1. Dockerfile に CMD を記述していないのに、 COMMAND に nginx -g って・・。
2. Dockerfile に EXPOSE を記述していないのに、 PORTSに 80 443 って・・。

これは利用イメージが [Official Repository nginx](https://hub.docker.com/_/nginx/) この [Dockerfile](https://github.com/nginxinc/docker-nginx/blob/7f3ef0927ec619d20181e677c97f991df0d7d446/Dockerfile) を利用して作成されているためである。

# Step6 nginx のデフォルトページを見る

Step1 で docker vm の ip を確認してポートを付与してあげる。 -p オプションでポートをした場合で考えると以下にアクセスすれば **Welcome to nginx!** が見れるはず

~~~bash
http://192.168.99.100:8000
~~~


# その他

## container の中身を確認する

~~~bash
$ docker run -t -i syati:nginx /bin/bash

# -t -i オプションは以下の通り
$ docker run --help
Usage:	docker run [OPTIONS] IMAGE [COMMAND] [ARG...]

Run a command in a new container
  -i, --interactive=false         Keep STDIN open even if not attached
  -t, --tty=false                 Allocate a pseudo-TTY
~~~
