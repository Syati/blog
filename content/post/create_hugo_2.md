+++
Categories = ["Hugo"]
Description = "WordPress から github.io にブログを立ち上げる"
Tags = ["Blog", "github.io"]
date = "2015-06-07T21:55:31+09:00"
title = "Hugo で github にブログを立ち上げる Part 2"
comments = true

+++

[Part 1]({{< ref "post/create_hugo.md" >}}) では、ブログをローカルに構築したので、今回は github に repository をつくってブログを立ち上げる。
[Hosting on GitHub Pages](http://gohugo.io/tutorials/github-pages-blog/) を参考にして、セットアップ手順を記していく。

全てが終われば[yourblog](http://syati.github.io/yourblog/)のように作成できる。

 <!--more-->
- **※注意**
    - ソースコード内の **&quot;** が、二重引用符（始）、二重引用符（終）に 変わっているので  
      コピペしてもバグります。コピペする場合は修正してください。

# Step 1 github に blog repository をつくる

github にアクセスして、repository をつくる。repository name がそのまま URL になる（以下のようなURL）。

**http://[username].githug.io/[repo]**

- [] の意味
    - [username] : github の username
    - [repo] : github に作成する repository name

# Step 2 blog の設定をかいておく

**※以降、github に yourblog として repostitory を作成したとして記す。また、part 1 で作成した yourblog があるものとする**

part 1 で作成した yourblog ディレクトリ内の **config.toml** を開いて設定を書く（以下参考）。

<pre><code class="language-bash">
title = "My New Hugo Site"
baseurl = "http://[username].github.io/yourblog"
languageCode = "ja-jp"
theme = "hyde"　# part 1 で選んだ theme を設定する
canonifyurls = true # 相対パスではなく baseurl を基点とした絶対パスにする

[params]
  description = "This is yourblog"
  author = "your name"

[taxonomies]
  category = "categories"
  tag = "tags"
</pre></code>

- 参考
    - [Configuring Hugo](http://gohugo.io/overview/configuration/) 
    - [TOMLノススメ](http://mojavy.com/blog/2013/04/26/toml/)
    - [toml-lang/toml](https://github.com/toml-lang/toml)


# Step 3 repository に push する

yourblog ディレクトリで以下のコマンドを実行する。

<pre><code class="language-bash">
$ git init # part 1 で作成した yourblog をリポジトリにする
$ git remote add origin git@github.com:[username]/yourblog.git # remote を設定する
$ git pull origin master
$
$ rm -rf public # このディレクトリは git subtree を利用して管理するので削除する
$ git add -A
$ git commit -m "Add hugo template"
$ git push origin master
</pre></code>

# Step 4 gh-pages ブランチを作成する

gh-pages ブランチには hugo で作成されたコンテンツ（ public ）のみ置いて、
関係のないもの（ archetypes, themes, etc ）は紛らわしいので管理したくない。

そこで、master から独立したブランチを作成するために、orphan オプションを
つけて履歴のないブランチを作成する。

<pre><code class="language-bash">
$ git checkout -\-orphan gh-pages   # orphan ブランチ 作成
$ git rm -\-cached $(git ls-files)  # 要らないので、全て管理対象からすべて外す
$ git add README.md                # README.md だけいれておく
$ git commit -m "initial commit on gh-pages branch"
$ git push origin gh-pages
</pre></code>

master に戻って git subtree を利用して、gh-pages ブランチを master の public に取り込む。
subtree って何？って思う方は、下部の参考を見てください。

<pre><code class="language-bash">
$ git checkout master
$ git subtree add -\-prefix=public git@github.com:[username]/yourblog.git gh-pages -\-squash
$ git subtree pull -\-prefix=public git@github.com:[username]/yourblog.git gh-pages
</pre></code>

hugo コマンドで public を生成して以下のように push していくことで、
master に変更を加え、gh-pages にも変更を加える事が出来る。

<pre><code class="language-bash">
$ hugo
$ git add -A
$ git commit -m "Updating site"
$ git push origin master
$ git subtree push -\-prefix=public git@github.com:Syati/yourblog.git gh-pages
</pre></code>

基本的に新しい記事を書いて github.io に公開（ Deploy ）する際、上記の手順を踏むことになる。
毎回これするのはめんどくさいので、Step 5 に進む。

- 参考
    - [gitの空ブランチを作る](http://qiita.com/akiko-pusu/items/7c0a99b8cb37882d2cfe)
    - [Git Submodule の代替: Git Subtree](http://japan.blogs.atlassian.com/2014/03/alternatives-to-git-submodule-git-subtree/)
    - [git-subtree移行メモ](http://qiita.com/shogo82148/items/04b29b195dbbb373152f)
    - [submoduleとsubtree-mergingの使い分け](http://qiita.com/marutanm/items/d02e7d5ff8ed7c2c4b95)


# Step 5 めんどさい Deploy はスクリプトを利用する

[Hosting on GitHub Pages](http://gohugo.io/tutorials/github-pages-blog/) の **deploy.sh** に書かれている以下を利用する。

<pre><code class="language-bash">
\# !/bin/bash

echo -e "\033[0;32mDeploying updates to GitHub...\033[0m"

\# Build the project. 
hugo

\# Add changes to git.
git add -A

\# Commit changes.
msg="rebuilding site `date`"
if [ $# -eq 1 ]
  then msg="$1"
fi
git commit -m "$msg\"

\# Push source and build repos.
git push origin master
git subtree push -\-prefix=public git@github.com:[username]/yourblog.git gh-pages　
</pre></code>

[spencerlyon2/hugo_gh_blog](https://github.com/spencerlyon2/hugo_gh_blog) のレポジトリに deploy.sh
があるので　DL してきて修正するのがはやいかも。

deploy.sh を yourblog ディレクトリに保存したのち、実行できるように以下を実行する。

<pre><code class="language-bash">
$ chmod +x deploy.sh
</pre></code>

これで新しい記事を書いて github.io に公開（ Deploy ）する際は、以下のコマンドを実行するのみで良い。

<pre><code class="language-bash">
$ ./deploy.sh
</pre></code>






