+++
Categories = ["Hugo"]
Description = "Disqus を用いて Hugo にコメント欄をいれる"
Tags = ["Blog", "github.io"]
comment = true
date = "2015-06-09T00:02:04+09:00"
title = "Hugo にコメント欄をいれる"

+++

[Part 1]({{< ref "post/create_hugo.md" >}}) 、 [Part 2]({{< ref "post/create_hugo.md" >}}) で
Hugo を用いて github.io にブログを立ち上げましたが、コメント欄が無いので導入したい！！

[Comments in Hugo](http://gohugo.io/extras/comments/) を参考にして早速、導入していこうとおもったが、
もっと楽にやりたいので、Part 1 で利用していた theme の hyde を変えて超簡単に導入する。

全てが終われば [yourblog](http://syati.github.io/yourblog/) のようにコメント欄が作成できる。

<!--more-->
**※part 1 で作成した yourblog があるものとして記す**

# Step 1 theme を変える

今まで通り、themes に中に theme を git clone してもいいのだが、
とりあえず今回は themes ディレクトリを利用しないことにする。

なので **config.toml** の**theme** 設定を削除する

<pre><code class="language-clike">
title = "My New Hugo Site"
baseurl = "http://syati.github.io/yourblog"
languageCode = "ja-jp"
theme = "hyde"  //この行はもう不要なので削除する

[params]
  description = "This is yourblog"
  author = "your name"

[taxonomies]
  category = "categories"
  tag = "tags"
</pre></code>

theme を [Syati/greyshade](https://github.com/Syati/greyshade) に変更する。
git clone して yourblog のメインテーマにする

<pre><code class="language-bash">
$ git clone git@github.com:Syati/greyshade.git
$ tree -L 1 #こんなディレクトリ構成
.
├── greyshade
└── yourblog

$ cd greyshade
$ cp -R archetypes images layouts static ../yourblog
$ cd ../yourblog
$ rm -rf themes # 必要無いのでとりあえず削除
$ hugo server # テーマが変わっていることが確認できる
</pre></code>

# Step 2 Disqus を導入

導入が楽なので [Disqus](https://disqus.com) を利用する。

Disqus に Sign up して http://[username].github.io でサイトを登録して
ゴニョゴニョすると、**Choose your platform** とかでる画面になるので
**Universal Code** をクリックすると、以下のコードをコメント欄を表示
させたいところに入れてね！と記載があるが、以下の **[disqus_shortname]** だけ覚えておく。

<pre><code class="language-markup">
\<div id="disqus_thread"\>\</div\>
\<script type="text/javascript"\>
    /* * * CONFIGURATION VARIABLES * * */
    var disqus_shortname = [disqus_shortname];
    
    /* * * DON'T EDIT BELOW THIS LINE * * */
    (function() {
        var dsq = document.createElement('script'); dsq.type = 'text/javascript'; dsq.async = true;
        dsq.src = '//' + disqus_shortname + '.disqus.com/embed.js';
        (document.getElementsByTagName('head')[0] || document.getElementsByTagName('body')[0]).appendChild(dsq);
    })();
\</script\>
\<noscript\>Please enable JavaScript to view the \<a href="https://disqus.com/?ref_noscript" rel="nofollow"\>comments powered by Disqus.\</a\>\</noscript\>
</pre></code>

# Step 3 config.toml を変更する

[テーマの設定](https://github.com/Syati/greyshade#setup) を config.toml にコピペして、編集。
Step 2 で覚えておいた **[disqus_shortname]** を設定に入れる。以下サンプル

<pre><code class="language-bash">
title = "My New Hugo Site"
baseurl = "http://syati.github.io/yourblog"
languageCode = "ja-jp"
canonifyurls = true

[author]
name = "your name"
\# email will use for gravatar
email = ""

[taxonomies]
category = "categories"

[params]
\# site description, will show under navigation
description = "This is yourblog"

\# RSS / Email (optional) subscription links (change if using something like Feedburner)
subscribe_rss = "/index.xml"
subscribe_email = ""

\# social links
facebook_user = ""
googleplus_user = ""
twitter_user = ""
github_user = ""
coderwall_user = ""
stackoverflow_user = ""
stackoverflow_user_id = ""
linkedin_user = ""
pinterest_user = ""
delicious_user = ""
pinboard_user = ""
quora_user = ""
instagram_user = ""
behance_user = ""
douban_user = ""

\# share links
facebook_like = true
twitter_tweet_button = true
google_plus_one = "true"
addthis_profile_id = ""

\# Disqus Comments
disqus_short_name = "[disqus_shortname]" #ここに part2 で覚えた[disqus_shortname]を入れる
disqus_show_comment_count = false

\# google analytics
google_analytics_tracking_id = ""
</pre></code>

# Step 4 記事にパラメータを加える

最後に記事のパラメータに **comments = true** といれれば、disqus を読み込んでコメント欄が入る。
ただし、localhost では読み込まないようにしているので、記事ページ下部 に **Comments**
とだけでてくる。

<pre><code class="language-markup">
+++
date = "2015-06-06T17:20:38+09:00"
title = "first"
comments = true

+++

first page.
</pre></code>

# Step 5 新規記事のテンプレートを変更しておく

毎回コマンド（ **hugo new post/new_post.md** ） で新規記事作成後に、
comments = true とパラメータ設定するのは面倒くさいので
yourblog/archetypes/default.md を以下のように変更しておくと楽ができる

<pre><code class="language-markup">
+++
Description = ""
Tags = []
Categories = []
comments = true
+++
</pre></code>


