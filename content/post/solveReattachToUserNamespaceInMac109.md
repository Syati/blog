+++
Categories = ["Shell", "OSX"]
Description = "tmux 起動時の reattach-to-user-namespace: unsupported new OS, trying as if it were 10.6-10.8 を解消 in Mac 10.9"
Tags = ["Tips", "tmux"]
date = "2015-06-04T23:09:38+09:00"
title = "reattach-to-user-namespace: unsupported new OS を解決する"

+++


ワーニングが出たら気になる。何が何でも消したくなってしまう。解決手順は以下のとおり。10.9に対応した reattach-to-user-namespace にするだけです。

## 手順

1.  <https://github.com/ChrisJohnsen/tmux-MacOSX-pasteboard> ここにいってソースコードをDLしてくる（git cloneでもOK）
2.  ダウンロードしたファイルのディレクトリで make する
    
        cd ~/Downloads/tmux-MacOSX-pasteboard
        make reattach-to-user-namespace
3.  パスの通っている reattach-to-usernamespace がどこにあるかを以下のコマンドで確認しておく
    
        which reattach-to-user-namespace
4.  makeで出来た reattach-to-user-namespace をパスの通っているものに上書きする

これでワーニングもサヨウナラ

## 参考

-   <https://github.com/ChrisJohnsen/tmux-MacOSX-pasteboard>
