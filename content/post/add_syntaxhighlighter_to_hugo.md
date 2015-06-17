+++
Categories = ["Hugo"]
Description = "Hugo に SyntaxHighlighter を導入""
Tags = ["Blog", "github.io"]
comments = true
date = "2015-06-17T22:42:39+09:00"
title = "Hugo に SyntaxHighlighter を導入"

+++

ソースコードをブログに載せる上で欠かせないのが SyntaxHighlighter ですよね。
というこで [Syntax Highlighting](http://gohugo.io/extras/highlighting/) を
参考にして早速導入する。


<!--more-->

- **※注意**
    - [Hugo にコメント欄をいれる]({{< ref "post/add_disqus_to_hugo.md" >}}) を実施したとして記す。
    - ソースコード内の **&quot;** が、二重引用符（始）、二重引用符（終）に 変わっているので  
      コピペしてもバグります。コピペする場合は修正してください。

# Step 1 Client-side の SyntaxHighlighter ライブラリ取得

導入方法として **Server-side** と **Client-side** の方法がありますが、
今回は後者を利用します。また、ライブラリーはたくさんありますが、**[Prism](http://prismjs.com/)**
を利用します。

Prism のHPにいって利用したい好きな Themes と Languages にチェックをいれて、
ページ下部の DOWNLOAD JS と DOWNLOAD CSS から DL しましょう。
ちなみに当ページの Prism のテーマは Default(Okaidia) を利用しています。

# Step 2 対象ファイルコピー

DL したファイルを以下のようにコピーしてあげる。

<pre><code class="language-bash">
$ cp prism.js yourblog/static/js/
$ cp prism.css yourblog/static/css/
</pre></code>

# Step 3 custom 設定を加える

以下のように yourblog/layouts/partials/custom に移動して
**対象ファイル１**と**対象ファイル２**を編集します。

<pre><code class="language-bash">
$ cd yourblog/layouts/partials/custom
$ tree -L 1
.
├── after_footer.html  # 対象ファイル１
├── asides
├── footer.html
├── head.html          # 対象ファイル２
├── header.html
└── navigation.html
</pre></code>

以下のように対象ファイル１に１行追加。

<pre><code class="language-markup">
\<script src="{{ .Site.BaseUrl }}/js/prism.js">\</script>
</pre></code>


以下のように対象ファイル２に１行追加。

<pre><code class="language-markup">
\<link href="{{ .Site.BaseUrl }}/css/prism.css" media="screen, projection" rel="stylesheet" type="text/css">
</pre></code>

これで設定は終了。

# Step 4 コードをいれてみる

新しい記事でも作成して、以下のような感じでコードをいれると

<pre><code class="language-markup">
\<pre>\<code class="language-c">
\#include \<stdio.h>

int main(void)
{
    printf("hello, world\n");
    return;
}
\</pre>\</code>
</pre></code>

あら素敵！！というふうになる。

<pre><code class=“language-c”>
\#include \<stdio.h>

int main(void)
{
    printf("hello, world\n");
    return;
}
</pre></code>


- 注意
    - markdown で利用されている **#** などは **\\#** でエスケープする必要があります




