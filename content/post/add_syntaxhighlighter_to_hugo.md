+++
Categories = ["Hugo"]
Description = "Hugo に SyntaxHighlighter を導入"
Tags = ["Blog", "github.io"]
comments = true
date = "2015-06-17T22:42:39+09:00"
title = "Hugo Part 4 - Hugo に SyntaxHighlighter を導入"

+++

ソースコードをブログに載せる上で欠かせないのが SyntaxHighlighter ですよね。
というこで [Syntax Highlighting](http://gohugo.io/extras/highlighting/) を
参考にして早速導入する。


<!--more-->

# Step 1 Client-side の SyntaxHighlighter ライブラリ取得

導入方法として **Server-side** と **Client-side** の方法がありますが、
今回は後者を利用します。また、ライブラリーはたくさんありますが、**[Prism](http://prismjs.com/)**
を利用します。

Prism のHPにいって利用したい好きな Themes と Languages にチェックをいれて、
ページ下部の DOWNLOAD JS と DOWNLOAD CSS から DL しましょう。
ちなみに当ページの Prism のテーマは Default(Okaidia) を利用しています。

# Step 2 対象ファイルコピー

DL したファイルを以下のようにコピーしてあげる。

~~~bash
$ cp prism.js yourblog/static/js/
$ cp prism.css yourblog/static/css/
~~~

# Step 3 custom 設定を加える

以下のように yourblog/layouts/partials/custom に移動して
**対象ファイル１**と**対象ファイル２**を編集します。

~~~bash
$ cd yourblog/layouts/partials/custom
$ tree -L 1
.
├── after_footer.html  # 対象ファイル１
├── asides
├── footer.html
├── head.html          # 対象ファイル２
├── header.html
└── navigation.html
~~~

以下のように対象ファイル１に１行追加。

~~~markup
<script src="{{ .Site.BaseURL }}/js/prism.js"></script>
~~~



以下のように対象ファイル２に１行追加。

~~~markup
<link href="{{ .Site.BaseURL }}/css/prism.css" media="screen, projection" rel="stylesheet" type="text/css">
~~~

これで設定は終了。

# Step 4 コードをいれてみる

新しい記事でも作成して、以下のような感じでコードをいれると。
~ 半角でいれると解釈されるため 〜 全角で入れています。

~~~markup
〜〜〜c
#include <stdio.h>

int main(void)
{
    printf("hello, world\n");
    return;
}
〜〜〜
~~~

あら素敵！！というふうになる。

~~~c
#include <stdio.h>

int main(void)
{
    printf("hello, world\n");
    return;
}
~~~





