+++
Categories = ["Emacs", "CoffeeScript"]
Description = "Emacs で CoffeeScript を AutoComplete するメモ"
Tags = ["Env"]
date = "2015-06-04T23:09:38+09:00"
title = "Emacs で CoffeeScript を AutoComplete する"

+++

[Coffee-modeでauto-completeを使えるようにする](http://tatsuyano.github.io/blog/2013/03/19/coffee-mode-used-ac-dict/) を参考に、
coffee script で auto-complete 使えるようにして、ついでに jquery も auto-complete に追加してあげる。
<!--more-->
# 必要なパッケージ

-   M-x package-list-package から 以下をインストールする
    -   jquery-doc

# 設定

[Coffee-modeでauto-completeを使えるようにする](http://tatsuyano.github.io/blog/2013/03/19/coffee-mode-used-ac-dict/) を参考にして以下の部分を変更してあげる

~~~clike
(add-hook 'coffee-mode-hook
    '(lambda ()
        (jquery-doc-setup) ;; ここに jquery-doc を追加する
        (add-to-list 'ac-dictionary-files "~/.emacs.d/ac-dict/js2-mode") ;; ここの ~/emacs.d/~~ は自分の辞書ファイルがあるところにしてある
     ))
~~~
