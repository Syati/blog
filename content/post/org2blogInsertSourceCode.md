+++
Categories = ["Emacs"]
Description = "org2blog の記事にコードブロックを入れるをいれて綺麗にコードを見せる"
Tags = ["Blog", "org2blog"]
date = "2015-06-04T23:09:38+09:00"
title = "org2blog の記事にコードブロックを入れる"

+++

org2blog の記事投稿でコードも綺麗に投稿したい。ってことで先日 [org2blogからの投稿](http://syati.info/?p=1746) の続き。

## 必要なもの

1.  htmlize.el を load-path の通ったところに入れる。
    -   htmlize.el は [org-mode](http://orgmode.org/ja/index.html) をDLして、解凍した contrib/lisp に入っている。

2.  [SyntaxHighlighter Evolved](http://wordpress.org/extend/plugins/syntaxhighlighter/) をワードプレスに入れて有効化しておく。

3.  emacs の設定ファイル ( .emacs または .emacs.d/init.el ) の org2blog 設定に **:wp-code t** を書き加える。
    -   以下が .emacs.d/init.el に書き込んだ私の org2blog の設定
    
        ;;org2blog
        (require 'org2blog)
        (setq org2blog/wp-blog-alist
              '(("wordpress"
                 :url "your wordpress url"
                 :username "login name"
                 :wp-code t))) ;; insert this line

## 使ってみる

1.  org2blog/wp-new-entry で投稿記事を作成する。
2.  以下のようにコードブロックを挿入する（都合によりコードブロック内の＃は全角にしてありますが、# 半角にすること）。
    -   表記 **#+BEGIN\_SRC 言語　:syntaxhl [利用したいパラメータ（任意）] ソースコード #+END\_SRC**
        -   :syntaxhl にパラメータを与えない場合は、ワードプレスの管理画面->設定 SyntaxHighlighter の一般設定が用いられる。
        -   [パラメータはこちら](http://en.support.wordpress.com/code/posting-source-code/#configuration-parameters)
    
    <pre>
    ＃+BEGIN_SRC bash :syntaxhl
    echo "hoge" 
    ＃+END_SRC
    </pre>

これで org2blog の記事投稿でコードも綺麗に投稿できる。

## 補足

-   必要なもの 3. の設定を加えるのは、以下の通り <pre> タグが挿入されるため SyntaxHighlighter の表記方法 [language] code [/language] と合致しなくなるため。（以下は org2blog.el 内のコード）
    -   参考 [org2blog](https://github.com/punchagan/org2blog#posting-source-code-blocks)
    
        (defcustom org2blog/wp-use-sourcecode-shortcode nil
           "Non-nil means convert <pre> tags to WP sourcecode blocks.
         NOTE: htmlize.el available in org-mode's contrib directory should
         be on your emacs load-path for this to work."
           :group 'org2blog/wp
           :type 'boolean)

## 覚えておくと良いショートカット

-   コードブロックを挿入する際は **<キー<TAB>** をタイプするべし。
    -   以下は、挿入したいテンプレートのキー一覧
        -   参考 [Org-mode Tutorial / Cheat Sheet](http://emacsclub.github.com/html/org_tutorial.html)

<table border="2" cellspacing="0" cellpadding="6" rules="groups" frame="hsides">


<colgroup>
<col  class="left" />

<col  class="left" />
</colgroup>
<thead>
<tr>
<th scope="col" class="left">キー　</th>
<th scope="col" class="left">挿入したいテンプレート</th>
</tr>
</thead>

<tbody>
<tr>
<td class="left">s</td>
<td class="left">#+begin\_src &#x2026; #+end\_src</td>
</tr>


<tr>
<td class="left">e</td>
<td class="left">#+begin\_example &#x2026; #+end\_example</td>
</tr>


<tr>
<td class="left">q</td>
<td class="left">#+begin\_quote &#x2026; #+end\_quote</td>
</tr>


<tr>
<td class="left">v</td>
<td class="left">#+begin\_verse &#x2026; #+end\_verse</td>
</tr>


<tr>
<td class="left">c</td>
<td class="left">#+begin\_center &#x2026; #+end\_center</td>
</tr>


<tr>
<td class="left">l</td>
<td class="left">#+begin\_latex &#x2026; #+end\_latex</td>
</tr>


<tr>
<td class="left">L</td>
<td class="left">#+latex:</td>
</tr>


<tr>
<td class="left">h</td>
<td class="left">#+begin\_html &#x2026; #+end\_html</td>
</tr>


<tr>
<td class="left">H</td>
<td class="left">#+html:</td>
</tr>


<tr>
<td class="left">a</td>
<td class="left">#+begin\_ascii &#x2026; #+end\_ascii</td>
</tr>


<tr>
<td class="left">A</td>
<td class="left">#+ascii:</td>
</tr>


<tr>
<td class="left">i</td>
<td class="left">#+index: line</td>
</tr>


<tr>
<td class="left">I</td>
<td class="left">#+include: line</td>
</tr>
</tbody>
</table>

## その他、参考サイト

-   [Babel: Introduction](http://orgmode.org/worg/org-contrib/babel/intro.html)
-   [Emacs org-modeを使ってみる: (29) エクスポートオプション一覧](http://d.hatena.ne.jp/tamura70/20100304/org)
