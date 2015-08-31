+++
Categories = ["Docker", "OSX"]
Description = "Docker について学んでいく"
Tags = ["Env", "Infra"]
date = "2015-08-31T18:39:00+09:00"
title = "Docker をインストールする"
draft = true
+++

巷で Docker と騒がれて随分経ったでしょうか。私の Mac が新しくなったのを気に、 Docker に入門したいと思います。何か新しい開発の度にゴミが入るのはちょっとね・・・・。

昔、Vagrant と Chef でローカル開発環境を構築していたこともありましたが provison が遅く、何だかなぁ〜と思っているうちに遠ざかっていました。Docker はコンテナ型で、さぞ軽いのだろうと思い、さっそくインストール。


# 前提

- OSX 10.10.5

# Step 1 Install VirtualBox

VirtualBox5.0.2 だと動かないので、以下からテストビルドを入れる

https://www.virtualbox.org/wiki/Testbuilds

5.0.2 だった場合 [エラー１]({{< ref "#aaa-1" >}})

- [Docker Toolbox付属のVirtualBox5.0.2では動かないので5.0.3を手動で入れること](http://qiita.com/tukiyo3/items/c912fe9e403706964995)

# Step 2 Install Docker Toolbox

~~~bash
$ brew cask install dockertoolbox
~~~


# Error

## virtualBox 5.0.2 時の start.sh を叩いた場合 {#aaa-1}

~~~bash

Started machines may have new IP addresses. You may need to re-run the `docker-machine env` command.

Setting environment variables for machine default...


                        ##         .
                  ## ## ##        ==
               ## ## ## ## ##    ===
           /"""""""""""""""""\___/ ===
      ~~~ {~~ ~~~~ ~~~ ~~~~ ~~~ ~ /  ===- ~~~
           \______ o           __/
             \    \         __/
              \____\_______/


host is not running
docker is configured to use the default machine with IP
For help getting started, check out the docs at https://docs.docker.com

default is not running. Please start this with docker-machine start default

$ docker-machine start default
exit status 1
Started machines may have new IP addresses. You may need to re-run the `docker-machine env` command.
~~~

## virtualBox 5.0.3 にした後、以前作成した vm のエラー

以前、作成したのは削除して、あらためて start.sh を叩きましょう。

~~~bash
$ docker-machine env default
open /Users/mizuki-y/.docker/machine/machines/default/ca.pem: no such file or directory
~~~
