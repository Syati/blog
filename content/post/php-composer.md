+++
Categories = ["PHP", "OSX"]
Description = "でないようにする 'You are running composer with xdebug enabled. This has a major impact on runtime performance'"
Tags = ["composer"]
comments = true
date = "2016-07-13T01:01:50+09:00"
title = "うるさいよ 'You are running composer with xdebug enabled. This has a major impact on runtime performance'"

+++

最近、php を触る機会が増えたので、メモメモ。
composer 叩いたら xdebug を enable にしてると遅いよっていう以下のメッセージが流れるので composer 叩いた時に出ないようにする。

<!--more-->

~~~bash
You are running composer with xdebug enabled. This has a major impact on runtime performance. See https://getcomposer.org/xdebug
~~~

## 方法

zsh を利用しているので .zshrc に以下を加える。

~~~bash
alias composer="php -n /usr/local/bin/composer"
~~~

## 参考

http://stackoverflow.com/questions/31083195/disabling-xdebug-when-running-composer
