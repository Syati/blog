+++
Categories = ["Emacs"]
Description = "Emacs を 256色で起動するメモ"
Tags = ["DevEnv"]
date = "2015-06-04T23:09:38+09:00"
title = "Emacs を 256色で起動する"

+++

最近は、さくらVPSで色々と遊んでいますので、そのメモ。

.bashrc 又は .zshrc に以下を書き込むことで、emacs を256色で起動できる。

    alias emacs='TERM=xterm-256color emacs'

emacs が256色で起動出来ているかどうか確認するには、 emacs M-x から
以下のコマンドを実行しましょう。
