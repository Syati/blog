+++
Categories = ["Docker", "OSX"]
Description = "Mac に Docker Toolbox をインストールして Docker を学ぶ"
Tags = ["Env", "Infra"]
comments = true
date = "2015-09-03T16:00:00+09:00"
title = "Docker Part 1 - Mac に Docker をインストールする"

+++

巷で Docker と騒がれて随分経ったでしょうか。私の Mac が新しくなったのを気に Docker に入門したいと思います。何か新しい開発の度にゴミが入るのはちょっと・・・、とも思いますしね。

昔、Vagrant と Chef でローカル開発環境を構築していたこともありましたが provison が遅く、何だかなぁ〜と思っているうちに遠ざかっていました。Docker はコンテナ型で、vagrant より軽いのかな。

<!--more-->

## 前提

- OSX 10.10.5

## Step 1 Install VirtualBox

VirtualBox5.0.2 だと Step 3 で動かなくなるので、[テストビルド](https://www.virtualbox.org/wiki/Testbuilds)から入れる。5.0.2 だった場合、 [エラー1]({{< ref "#error-1" >}})が出る。

- 参考
    - [Docker Toolbox付属のVirtualBox5.0.2では動かないので5.0.3を手動で入れること](http://qiita.com/tukiyo3/items/c912fe9e403706964995)

## Step 2 Install Docker Toolbox

~~~bash
$ brew cask install dockertoolbox
~~~

[cask](({{< ref "post/brew_cask.md" >}}) って思う方は、公式から [Docker Toolbox](https://www.docker.com/toolbox) をインストールしましょう。詳しい説明は、[Install Docker Mac OS X](https://docs.docker.com/mac/step_one/) 。インストールされる場所は以下のとおり、/usr/local/bin なので、path が通っていればコマンドが通るはずです。

> By default, the standard Docker Toolbox installation:
>    installs binaries for the Docker tools in /usr/local/bin

コマンド例

~~~bash
$ docker --version
Docker version 1.8.1, build d12ea79
~~~

## Step 3 とりあえず起動

[Installation](http://docs.docker.com/mac/step_one/#step-3-verify-your-installation) に書いてあるとおり、以下のコマンドで、とりあえず起動、コマンドを実行してみる。[エラー]({{< ref "#error" >}})が無いことをもって、インストール・起動ができたことの確認を終了する。

~~~bash
$ bash '/Applications/Docker Quickstart Terminal.app/Contents/Resources/Scripts/start.sh'
$ docker run hello-world

Hello from Docker.
This message shows that your installation appears to be working correctly.

To generate this message, Docker took the following steps:
 1. The Docker client contacted the Docker daemon.
 2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
 3. The Docker daemon created a new container from that image which runs the
    executable that produces the output you are currently reading.
 4. The Docker daemon streamed that output to the Docker client, which sent it
    to your terminal.

To try something more ambitious, you can run an Ubuntu container with:
 $ docker run -it ubuntu bash

Share images, automate workflows, and more with a free Docker Hub account:
 https://hub.docker.com

For more examples and ideas, visit:
  https://docs.docker.com/userguide/
~~~

## Error {#error}

### VirtualBox 5.0.2 時で start.sh を叩いた場合 {#error-1}

以下のとおり、vm が立ち上がってくれない。

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

### VirtualBox 5.0.3 にした後、以前作成した vm のエラー {#error-2}

設定が取れないので、以前作成した vm は削除して、あらためて start.sh を叩きましょう。

~~~bash
$ docker-machine env default
open /Users/mizuki-y/.docker/machine/machines/default/ca.pem: no such file or directory
~~~
