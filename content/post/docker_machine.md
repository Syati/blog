+++
Categories = ["Docker", "OSX"]
Description = "docker-machine コマンドの使用方法・流れ"
Tags = ["Env", "Infra"]
comments = true
date = "2015-09-09T23:00:19+09:00"
title = "docker-machine コマンドの使用方法・流れ"

+++

Docker すごい。 何がすごいって環境構築が楽すぎる。以前は、Vagrant と Chef でゴリゴリ書いてアプリケーションが動くように頑張ってましたが、そんな事しなくても大方用意されてる。あとはその組み合わせ環境を構築するだけ。懸念点であった Vagrant 、Chef のような provision が、遅いということは今のところ無い。最初にイメージを DL してコンテナ化するので、若干の時間はかかるものの、それでも早い。 具体的な環境構築については、後日記そうと思う。

今回は **docker-machine コマンド**（コンテナを動かす vm を管理するコマンド）の使用方法、流れを記す。


<!--more-->

# 前提

- OSX 10.10.5
- Docker Toolbox
    - Docker version 1.8.1, build d12ea79
    - docker-machine version 0.4.1 (e2c88d6)
- Virtual Box 5.0.3

# docker-machine コマンド、流れ

以下のレイヤーと対応するコマンドを覚えておくと、スムーズかもしれない。

| レイヤー             |  対応コマンド    |
| :------------------ |:---------------|
| Docker Container    | docker         |
| Docker VM           | docker-machine |
| OSX                 |                |

## Step 1 docker vm を確認する

作成した docker vm を確認する。普段の **ls** と同じ感じで使う頻度が高い。はじめは何もないので空である。

~~~bash
$ docker-machine ls

NAME      ACTIVE   DRIVER       STATE     URL                         SWARM
~~~

## Step 2 docker vm を作成する

docker vm を作成する。 以下のコマンドは、オプション driver で virtualbox を指定して、 vm 名を sample にしている。色々オプションがあるので、もっと詳しくという方は、[create](https://docs.docker.com/machine/reference/create/) を参考にしてください。

~~~bash
$ docker-machine create --driver virtualbox sample

Creating VirtualBox VM...
Creating SSH key...
Starting VirtualBox VM...
Starting VM...
To see how to connect Docker to this machine, run: docker-machine env sample
~~~

Step 1 の通り、作成した docker vm を確認してみる。vm が作成されているとともに、STATE を見ることで動作していること、URL を見ることで IP が割り振られていることが確認できる。

~~~bash
$ docker-machine ls

NAME      ACTIVE   DRIVER       STATE     URL                         SWARM
sample             virtualbox   Running   tcp://192.168.99.103:2376
~~~

## Step 3 docker vm を利用する

Step 2 まででは、対象の docker vm にコンテナを操作するコマンド(docker ps など)を実行できない。SSH で対象 vm に入れば可能だが、そんな面倒くさいことは毎回したくない。
そこで、以下のように env -> eval コマンドを実行して、利用する docker vm の環境変数を設定する。


~~~bash
$ docker-machine env sample

export DOCKER_TLS_VERIFY="1"
export DOCKER_HOST="tcp://192.168.99.103:2376"
export DOCKER_CERT_PATH="/Users/mizuki-y/.docker/machine/machines/sample"
export DOCKER_MACHINE_NAME="sample"
# Run this command to configure your shell:
# eval "$(docker-machine env sample)"
~~~

~~~bash
$ eval "$(docker-machine env sample)"
$ docker-machine ls

NAME      ACTIVE   DRIVER       STATE     URL                         SWARM
sample    *        virtualbox   Running   tcp://192.168.99.103:2376
~~~

これで今後利用する docker コマンドは docker vm である sample に発行される。

# その他

## docker vm の中身を確認する ssh で入る

~~~bash
$ docker-machine ssh sample

                        ##         .
                  ## ## ##        ==
               ## ## ## ## ##    ===
           /"""""""""""""""""\___/ ===
      ~~~ {~~ ~~~~ ~~~ ~~~~ ~~~ ~ /  ===- ~~~
           \______ o           __/
             \    \         __/
              \____\_______/
 _                 _   ____     _            _
| |__   ___   ___ | |_|___ \ __| | ___   ___| | _____ _ __
| '_ \ / _ \ / _ \| __| __) / _` |/ _ \ / __| |/ / _ \ '__|
| |_) | (_) | (_) | |_ / __/ (_| | (_) | (__|   <  __/ |
|_.__/ \___/ \___/ \__|_____\__,_|\___/ \___|_|\_\___|_|
Boot2Docker version 1.8.1, build master : 7f12e95 - Thu Aug 13 03:24:56 UTC 2015
Docker version 1.8.1, build d12ea79
~~~
