+++
Categories = ["Hugo"]
Description = "Hugo に google カスタム検索エンジンを導入して、記事検索を可能にする"
Tags = ["Blog", "github.io"]
comments = true
date = "2015-09-21T16:25:47+09:00"
title = "Hugo Part 5 - Hugo に Google カスタム検索エンジンを導入"
logo = "hugo"
+++

記事が多くなってくると、やっぱり検索したくなるもんです。Hugo で検索を使うには、[Tools](https://gohugo.io/tools) に書いてある通り以下の方法があります。

- [Google カスタム検索エンジン](https://cse.google.co.jp/cse/)
- 自前でインデックスを作成して提供する方法
    - [Hugoidx](https://github.com/blevesearch/hugoidx): Bleve を用いる
    - [Github Gist](https://gist.github.com/sebz/efddfc8fdcb6b480f567): grunt と lunr.js を用いる

当ページでは、[Github Gist](https://gist.github.com/sebz/efddfc8fdcb6b480f567)  を用いて検索を実現していますが、お手軽さでいうと、[Google カスタム検索エンジン](https://cse.google.co.jp/cse/) ですね。なので今回は、[Google カスタム検索エンジン](https://cse.google.co.jp/cse/) の導入方法を紹介。

<!--more-->

## Step 1 登録する

[Google カスタム検索エンジン](https://cse.google.co.jp/cse/) にアクセスして登録しましょう。

{{< figure src="step1.png" title="カスタム検索エンジンの登録" >}}

## Step 2 作成する

必要情報を入力しましょう。**検索するサイト**を入力すると、**検索エンジンの名前** も入力されるため枠で囲っていません。

{{< figure src="step2.png" title="カスタム検索エンジンの作成" >}}

## Step 3 確認する

作成された検索エンジンを確認しましょう。

{{< figure src="step3.png" title="カスタム検索エンジンの確認" >}}

## Step 4 設定する

**対象** のリストメニューの **デザイン** をえらんで、レイアウトやテーマなど色々設定しましょう。設定すると、**プレビュー**で確認できます。最後に**保存してコードを取得**します。

{{< figure src="step4.png" title="カスタム検索エンジンの設定" >}}

## Step 5 取得する

**コード** をコピーしときましょう。

{{< figure src="step5.png" title="カスタム検索エンジンのコードを取得" >}}


## Step 6 設置する

コピーしたコードを layouts の任意の html に貼り付けましょう。完成すると以下のように検索が可能になります。


{{< embedded_google_search >}}

