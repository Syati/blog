+++
Categories = ["Hugo"]
Description = "WordPress から github.io にブログを立ち上げる"
Tags = ["Blog", "github.io"]
date = "2015-06-05T21:55:31+09:00"
title = "Hugo Part 1 - Hugo で github にブログを立ち上げる"
comments = true
logo = "hugo"

+++

WordPress でブログをつくっていましたが、記事投稿が面倒くさなって続かない。
なぜかと考えた時にテキストで書いたものを簡単に POST できないからである。
過去に org2blog などを試し、簡単になった！！と喜んでいた時期もあったが、
WordPress の Version Up などで POST できなくり、それっきり・・・。

今回は、ついに上記のことを打開できる策を見つけたのだ。
それが **Hugo で github.io 上にブログを立ち上げることだ。**
何が良いかといえば、markdown でかけて、github に push するだけでブログ
が更新されることだ。
 <!--more-->
[Hugo Quickstart Guide](http://gohugo.io/overview/quickstart/) にならって
セットアップ手順を記していく。Part 1 では、まずはブログをローカルに構築する。
[Part 2]({{< ref "post/create_hugo_2.md" >}}) で、github.io にブログを立ち上げる。


## Step 1 インストールしましょう

### OSX の場合
 
以下のコマンドでおしまい。

~~~bash
$ brew install hugo
~~~

### Ubuntu の場合

[Hugo relase](https://github.com/spf13/hugo/releases) から適したものをDLして以下のコマンドを実行する。

~~~bash
$ sudo dpkg -i hugo_0.14_amd64.deb　# 自分に適したパッケージを選んでね
~~~


## Step 2 ブログをつくろう

以下のコマンドを実行して、カレントディレクトリにブログの雛形を作成できる。

~~~bash
$ hugo new site ./yourblog
~~~

雛形はこんな感じになる。

~~~bash
yourblog/
├── archetypes  # 新規記事を作成した際に利用される雛形を置く場所
├── config.toml # 設定を書くファイル
├── content     # 記事などが入る場所
├── data        # 今回は利用しない（サイトを生成するときに利用する DATA などを置く。詳しくは http://gohugo.io/extras/datafiles/ ）
├── layouts     # 今回は利用しない（サイトを生成するときの雛形）
└── static      # 今回は利用しない（サイトで利用する js, css, images などを置く）
~~~

## Step 3 新しい記事をつくろう

yourblog ディレクトリで以下のコマンド実行して、新規記事を作成する。

~~~bash
$ hugo new post/hello.md
~~~

こんな感じで新規記事が作成される。

~~~bash
content/
└── post
    └── hello.md
~~~
    

hello.md の中身はこんな感じ。

~~~markdown
+++
date = "2015-06-05T23:04:20+09:00"
draft = true # この行を消せば step 5 の -\-buildDrafts オプションはいりません
title = "hello"

+++

~~~

## Step 4 テーマをインストールしよう

自分好みのテーマをさがすために、テーマ一式いれてみる。yourblog ディレクトリで以下のコマンドを実行する。

~~~bash
$ git clone -\-recursive https://github.com/spf13/hugoThemes themes
~~~

## Step 5 ブログを見てみよう

とりあえず準備はととのったので、サーバーを立ち上げてブログを見てみる。

~~~bash
$ hugo server -\-theme=hyde -\-buildDrafts

1 of 1 draft rendered
0 future content 
1 pages created
0 paginator pages created
0 tags created
0 categories created
in 7 ms
Serving pages from /home/mizuki-y/Documents/yourblog/public
Web Server is available at http://127.0.0.1:1313/
~~~

http://127.0.0.1:1313/ にアクセスするとブログが見れるようになっている。

theme を以下のように変更して、自分の好きなものを選ぼう。

~~~bash
$ hugo server -\-theme=slim -\-buildDrafts
~~~


- コマンドオプションについて
    - **-\-theme**: themes ディレクトリに入っているディレクトリ名を指定してあげる
    - **-\-buildDrafts**: draft ステータスのものを表示するためのオプション

ブログ構築完成。

[Part 2]({{< ref "post/create_hugo_2.md" >}}) で、github.io にブログを立ち上げる。

