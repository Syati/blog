+++
Categories = ["Emacs"]
Description = "org2blog で毎回ログインする際に求められるパスワード入力を省略する"
Tags = ["Blog", "org2blog"]
date = "2015-06-04T23:09:38+09:00"
title = "org2blog でブログに接続する際のパスワードを .netrc に保存しておく"

+++

org2blog で毎回ログインする際に求められるパスワード入力がめんどくさいが、init.el に書いておくのもちょっと&#x2026;とおもっている場合は、.netrc を使うのがお勧め

# 手順

1.  ホームに以下のコマンドで .netrc を作成する
    
        echo "machine myblog login myusername password myrealpassword" > ~/.netrc

2.  myusername と myrealpassword を自分がログインするときのユーザー名とパスワードに書き換える

3.  emacs の設定ファイル（ .emacs.d/init.el ) に以下を書き加える
    
        ;; org2blog
        (require 'org2blog-autoloads)
        (require 'netrc) ;; or nothing if already in the load-path
        (setq org2blog/wp-use-sourcecode-shortcode t) ;;syntaxhl
        (setq blog (netrc-machine (netrc-parse "~/.netrc") "myblog" t))
        (setq org2blog/wp-blog-alist
               '(("my-blog"
                  :url "your wordpress url"
                  :username (netrc-get blog "login")
                  :password (netrc-get blog "password")
                  :tags-as-categories nil)))

# 参考

-   [org2blog](https://github.com/punchagan/org2blog#posting-source-code-blocks)
