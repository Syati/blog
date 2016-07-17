+++
Categories = ["PHP", "OSX"]
Description = "でないようにする 'You are running composer with xdebug enabled. This has a major impact on runtime performance'"
Tags = ["composer"]
comments = true
date = "2016-07-13T01:01:50+09:00"
title = "うるさいよ 'You are running composer with xdebug enabled. This has a major impact on runtime performance'"
logo = "php"

+++

最近、php を触る機会が増えたので、メモメモ。
composer 叩いたら xdebug を enable にしてると遅いよっていう以下のメッセージが流れるので composer 叩いた時に出ないようにする。

~~~bash
You are running composer with xdebug enabled. This has a major impact on runtime performance. See https://getcomposer.org/xdebug
~~~

## 方法

方法 1. 2. どちらかでインストール、設定する

### 1. brew でインストールする

~~~bash
brew install composer

# .zshrc に alias を設定する
alias composer="php -n /usr/local/Cellar/composer/1.1.2/libexec/composer.phar"
~~~

**1.1.2** というバージョンのところは適宜修正する。

#### その他

以下のよう alias を設定すれば良いのではと思うかもしれないが、うまくコマンドが通らない

~~~bash
# .zshrc に alias を設定する
alias composer="php -n /usr/local/bin/composer"

#上記の設定だとコマンドを叩いた時に以下のような出力がでてしまう
$ composer
/usr/bin/env php -d allow_url_fopen=On -d detect_unicode=Off /usr/local/Cellar/composer/1.1.2/libexec/composer.phar "$@"
~~~

というのも **usr/local/bin/composer** のリンク先が **/usr/local/Cellar/composer/1.1.2/libexec/composer** でその中身が以下のようになっているため。

~~~bash
#!/usr/bin/env bash
/usr/bin/env php -d allow_url_fopen=On -d detect_unicode=Off /usr/local/Cellar/composer/1.1.2/libexec/composer.phar "$@"
~~~

### 2. composer の web からインストールする

~~~bash
wget https://getcomposer.org/download/1.1.3/composer.phar
mv ./composer.phar /usr/local/bin/composer
chmod +x /usr/local/bin/composer
~~~

~~~bash
# .zshrc に alias を設定する
alias composer="php -n /usr/local/bin/composer"
~~~

## 参考

http://stackoverflow.com/questions/31083195/disabling-xdebug-when-running-composer
