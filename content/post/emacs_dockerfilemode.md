+++
Categories = ["Emacs", "Docker"]
Description = "Emacs で dockerfile を編集するなら dockerfile-mode"
Tags = ["Env"]
comments = true
date = "2015-09-02T17:58:56+09:00"
title = "Emacs で Dockerfile を編集する"

+++

Docker に入門して、Dockerfile を書くことになったので、とりあえず dockerfile-mode 。

<!--more-->

# Install

-   M-x package-list-package で以下をインストールする
    - dockerfile-mode

-   ~/.emacs.d/init.el に以下を追加

    ~~~clike
    (autoload 'dockerfile-mode "dockerfile-mode" nil t)
    (add-to-list 'auto-mode-alist '("Dockerfile\\'" . dockerfile-mode))
    ~~~
