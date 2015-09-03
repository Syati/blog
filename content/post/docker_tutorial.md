+++
Categories = []
Description = ""
Tags = []
comments = true
date = "2015-09-03T15:54:19+09:00"
title = "docker_tutorial"
draft = true

+++


# docker-machine 利用するコマンド

## docker vm を作成する

~~~bash
$ docker-machine create --driver virtualbox sample

Creating VirtualBox VM...
Creating SSH key...
Starting VirtualBox VM...
Starting VM...
To see how to connect Docker to this machine, run: docker-machine env sample
~~~


~~~bash
$ docker-machine ls

NAME      ACTIVE   DRIVER       STATE     URL                         SWARM
sample             virtualbox   Running   tcp://192.168.99.103:2376
~~~

## docker vm を利用する

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


## docker vm に ssh で入る

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
