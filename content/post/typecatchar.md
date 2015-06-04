+++
Categories = ["Design"]
Description = ""
Tags = ["Fonts", "Google Fonts"]
date = "2015-06-04T23:09:38+09:00"
title = "Google Fonts をローカルで利用する"

+++

Google Fonts をローカル環境で使いたいと思い調べたら、意外と簡単だった。以下は参考のほぼ転記。

# 環境

-   OS:ubuntu 12.04

# インストール

    sudo add-apt-repository ppa:andrewsomething/typecatcher
    sudo apt-get update
    sudo apt-get install typecatcher

メニュー　-> アクセサリー -> TypeCatchar で起動した後、好きなフォントを選んでインストール

# その他

ubuntu 13.10 以上は、以下の通り

    sudo apt-get install typecatcher

# 参考

-   <http://www.webupd8.org/2013/08/easily-download-and-install-google-web.html>
