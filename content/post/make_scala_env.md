+++
Categories = ["Emacs", "Scala"]
Description = ""
Tags = ["DevEnv"]
date = "2015-06-04T23:09:38+09:00"
title = "Scala の開発環境"

+++

# os

-   ubuntu 12.04
-   emacs 24.3.1

# scala 開発環境

## java のインストール

1.  jdk をインストール
    -   sudo apt-get install default-jdk
2.  以下のコマンドでエラーが無いことを確認
    -   java -version
    -   javac -version

## scala のインストール

1.  scala のダウンロード
    -   <http://www.scala-lang.org/download/>   
        -   sudo apt-get install scala でも可能だが version が古い
2.  解凍ファイルを usr/local に上書きする
3.  ターミナルから以下のコマンドで、インストールできているか確認する
    -   scala -version

# ENSIME 導入

## scala-mode2

-   以下から tool を clone する
    -   git clone <https://github.com/hvesalai/scala-mode2.git>
-   コピーした scala-mode に移動して make コマンドする
    -   カレントディレクトリで .elc ファイルを作成される
-   init.el に書く加える
    
        (add-to-list 'load-path "/path/to/scala-mode2/")
        (require 'scala-mode2)

## scala-mode 補助

-   yasnipet がない場合は、以下から clone する
    -   <https://github.com/capitaomorte/yasnippet>

# 参考サイト

-   [Scalaプログラミング入門](http://bach.istc.kobe-u.ac.jp/lect/ProLang/org/scala.html)
-   [Scalaの入門にあたってscala-modeを入れてみた](http://blog.iss.ms/2012/06/02/101357)
