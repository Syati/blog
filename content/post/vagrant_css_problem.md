+++
Categories = ["Linux", "Platform"]
Description = ""
Tags = ["Automation Platform"]
date = "2015-06-04T23:09:38+09:00"
title = "vagrant、chef-solo で開発環境を構築する"

+++

最近、vagrant、chef-solo で開発環境を構築するのがブームなSyatiです。xampp、mamp とかで？？ができないとかよくありますが、面倒くさいので皆サーバーと同じ環境にして開発しましょうよ。そしたら、？？ができないなんてことがなくなりますので。

今回はまったのは、vagrant の共有フォルダに apache のバーチャルホストのドキュメントルートを設定して、ローカルでソースコードをいじり、ブラウザで確認している時である。

**CSSをいじって、ブラウザで確認しても反映されないのである。もちろんキャッシュクリア済み。やむなくvagrantを再起動すると反映されるのである**

<!--more-->

## 原因

httpd.conf の設定に問題があった。
以下の設定を加えるのである。EnableSendfile なんて httpd.conf 中のどこにも記述が無いんだが、デフォルトが On のため、明示的に加えてあげる必要がある。

    EnableSendfile Off

以下[ EnableSendfile](http://httpd.apache.org/docs/2.2/mod/core.html#enablesendfile)より引用

> このディレクティブはクライアンにファイルの内容を送るときに httpd がカーネルの sendfile サポートを使うかどうかを制御します。デフォルトでは、 例えば静的なファイルの配送のように、リクエストの処理にファイルの 途中のデータのアクセスを必要としないときには、Apache は OS が サポートしていればファイルを読み込むことなく sendfile を使って ファイルの内容を送ります。
> sendfile は read と send を別々に行なうことと、バッファの割り当てを 回避します。しかし、プラットフォームやファイルシステムの中には 運用上の問題を避けるためにこの機能を使用不可にした方が良い場合があります

おそらく、以下の部分が原因か！？それでファイルを変更しても、逐次更新されないんでしょう。
-   **Apache は OS が サポートしていればファイルを読み込むことなく sendfile を使って ファイルの内容を送ります。**

## 参考

-   <http://qiita.com/shoyan/items/12389d5beaa8695b1a53>

上記に nginx の場合も書いて有りますね。
