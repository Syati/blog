+++
Categories = ["Emacs", "JavaScript"]
Description = "Emacs で JavaScript の開発環境を整えるメモ"
Tags = ["DevEnv"]
date = "2015-06-04T23:09:38+09:00"
title = "Emacs で JavaScript の開発環境を整える"

+++

開発環境はとっても大事なんですが、構築までに時間がかかりますよね。ということで私の環境を忘れないうちにメモ！！

# 前提

-   ubuntu 12.04 or OSX
-   emacs >= 24

# 構築

## パッケージ管理にmelpa追加

-   ~/.emacs.d/init.el に以下を追加
    
        (require 'package)
        (add-to-list 'package-archives
          '("melpa" . "http://melpa.milkbox.net/packages/") t)

## パッケージ管理からインストール

-   M-x package-list-package で以下の2点をインストールする（該当箇所で **i** をタイプしてマークをつけて **x** でインストールする）
    1.  flycheck
    2.  js2-mode
-   ~/.emacs.d/init.el に以下を追加
    
        (add-hook 'after-init-hook #'global-flycheck-mode)
        (add-to-list 'auto-mode-alist '("\\.js\\'" . js2-mode))

## jshint をインストール

上記をインストールすることで、javascript の文法エラーなどを捕まえてくれる。
-   npm をインストール
    -   ubuntu の場合
        
            sudo apt-get install npm
    -   OSX の場合
        
            brew install npm
-   jshint をインストール
    
        sudo npm install jshint -g
    
    -   オプション -g でグローバルで利用するということ。付けない場合は、カレントディレクトリにインストールされる。

# 試す

a.js とでもファイルを作ってみると emacsの下の帯に　(javascript-IDE FlyC) となっていることが確認できる。
後は適当にプログラムを書けば jshint がはしり、文法エラーなどを教えてくれる。

# 参考

-   [melpa](http://melpa.milkbox.net/#/getting-started)
-   [js2-mode](http://code.google.com/p/js2-mode/)
-   [npm](https://npmjs.org/)
