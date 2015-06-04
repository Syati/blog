+++
Categories = ["Windows", "Emacs", "Linux"]
Description = "Cygwin、Zsh、NTEmacs で Windows に Linux環境を構築をする"
Tags = ["DevEnv"]
date = "2015-06-04T23:09:38+09:00"
title = "Cygwin + Zsh + NTEmacs で作る Linux環境 in Windows"

+++


最近のメインOSは、Ubuntu で Windows とはお別れしているのですが、否応なく Windows を利用しなければいけない時がある Syati です。 そんな時、欠かせないのが Cygwin、Zsh、NTEmacs です。もっと欲を言えば AutoHotkey もあるのですが、それは、また次の機会に。今回は、Cygwin、Zsh、NTEmacs で Windows に Linux環境を構築をする。

# 1. インストール

1.  [cygwin](http://www.cygwin.com/) のインストール
    1.  注意がでるけどインストールディレクトリをドライブのRoot( F:/ ) にでもする
2.  デフォルトから追加で、とりあえずインストールするパッケージは以下の通り（あとは好きなパッケケージを入れる）
    1.  Archive
        1.  p7zip
        2.  unzip
        3.  zip
    2.  Devel
        1.  gcc-core
        2.  gcc-g++
    3.  Net
        1.  openssh
    4.  Shells
        1.  zsh
    5.  Web
        1.  wget
3.  [gnupack emacs only](http://sourceforge.jp/projects/gnupack/) でインストール
    1.  ダウンロードしたファイルを解凍
    2.  解凍されたディレクトリの中身(bin,etc,info などなど）をすべてコピーして cygwin をインストールしたディレクトリの F:/usr/local/ に上書きする

# 2. cygwin のデフォルト Shell を zsh に切り替え、初期設定する

1.  /etc/passwd をエディタで開いて /bin/bash を /bin/zsh に置換する
2.  Cygwin が MS-DOS形式のパスに対して Warning を出力するので消しておく
    -   [cygwin がMS-DOS形式のパスをWarningするので消した](http://d.hatena.ne.jp/takuya_1st/20110423/1303586388)
3.  Cygwin Terminal ショートカットから Terminal を起動すると zsh startup が立ち上がるので (0) Exit, creating the &#x2026;.. を選択（0 をタイプする）
4.  HOME に .zshrc が作成されていると思うが、中身を消して、以下のサイトから設定ファイルをコピペして利用する
    1.  [漢のzsh 24 グッバイ野郎ども! コピペではじめるzshファイナル](http://news.mynavi.jp/column/zsh/024/index.html)
    2.  zsh についても詳しく書かれているので、勉強しておく

# 3. Emacs の作業場所(HOME) を設定する

1.  F:/usr/local/bin の中にある runemacs.exe のショートカットを作成 Desktop にでも貼り付ける
2.  My Computer で右クリック -> プロパティ -> 詳細設定タブ -> 環境変数 -> ○○のユーザー環境変数 の新規でHOMEを作成する
    -   変数名： HOME
    -   変数値： F:/home/ユーザー名/
3.  1.で作成したショートカットを右クリックしてプロパティから、作業フォルダを %HOME% と入力する（%HOME% は 2. 作成した変数名HOMEの変数値を入れるという意味）
4.  emacsを起動して Ctrl + x, Ctrl + f をタイプしたあとに ~/ を入力して、HOME に移動するか確認して終了

# 4. Emacs の環境を設定する (設定ファイルは ~/.emacs.d/init.el )

Emacs で設定する環境変数は、以下の通り３つある。本設定が上手くいっていない/理解していないと command not found で悩むかも。
1.  PATH
    -   emacs で shell (M-x shell)を利用する際に使う。端末(cmd, minttyなど)で利用する場合と同じ。
        -   ※ M は Alt キー のこと
2.  exec-path
    -   emacs のコマンド（grep, shell, diff, dired-mode 中の圧縮/解凍など）を利用する際に使う。
3.  load-path
    -   emacs-lisp（\*.el、\*.elc) を利用する際に使う。

## PATH, exec-path に同じ設定をする

cygwin から emacs を立ち上げた場合と、GUIから emacs を立ち上げた場合は、異なる環境変数が利用されるため、混乱を避けるためにも PATH, exec-path に同じ設定をする
1.  違いを確かめて見たい場合は、双方で立ち上げた emacs の **scratch** で以下をタイプして、それぞれの行末で（Ctrl-j) をタイプして式を評価して見ましょう。
    1.  (getenv "PATH")　
    2.  exec-path
2.  修正するには、emacs の設定ファイルに以下のように PATH と exec-path を記述する
    
        (let ((my-emacs-path (list
                 "/bin"
                 "/usr/bin"
                 "/usr/local/bin")))
         (setq exec-path my-emacs-path)
         (setenv "PATH" (mapconcat 'identity my-emacs-path ";")))

### 補足

cygwin から呼び出した場合は、/etc/profile の PATH に加えて、Windows の環境変数 PATH 、/home/ユーザー名/シェル設定ファイル（ .zshrc または .bashrc ) のPATH が読み込まれる （優先されるのは、先に記述されている方）。GUIからの場合は、Windows の環境変数 PATH が exec-path にも適用される。

### 参考

-   <http://ergoemacs.org/emacs/emacs_env_var_paths.html>

## GUI から立ち上げる emacs の M-x shell を コマンドプロンプトから zsh にする

cygwin から emacs を立ち上げて、 M-x shell で呼び出されるのは、[2. cygwin のデフォルト Shell を zsh に切り替え、初期設定する] を設定していた場合 zsh が呼び出されるが GUI の emacs から M-x shell で呼び出されるのはコマンドプロンプトになる。理由は、cygwin からの場合は、cygwinの環境変数のSHELL(zsh.exe)が利用される一方、GUIからは、NT-emacs のデフォルト環境変数SHELL（cmdproxy.exe）が利用されるため。
1.  設定ファイルに以下を書き込む
    
        (setq shell-file-name "zsh")
        (setenv "SHELL" shell-file-name) 
        (setq explicit-shell-file-name shell-file-name)

### 参考

-   [NTEmacsWithCygwin](http://emacswiki.org/emacs/NTEmacsWithCygwin)
-   <http://flex.ee.uec.ac.jp/texi/emacs-jp/emacs-jp_202.html>

## デフォルトの load-path に自分がインストールする emacs-lisp（\*.el、\*.elc) 置場を追記しておく

1.  ~/.emacs.d/lisp を自分がインストールする emacs-lisp（\*.el、\*.elc) 置場として、load-path に追加する。
    -   ~/.emacs.d/lisp 以下のサブディレクトリも自動で追記してくれる
    
        (let ((default-directory "~/.emacs.d/lisp/"))
          (normal-top-level-add-subdirs-to-load-path))

### 参考

-   <http://emacswiki.org/emacs/LoadPath>

# まとめ

以下にここまで記述した init.el の設定を残しておく。

    ;; Set PATH and exec-path
    (let ((my-emacs-path (list
                          "/bin"
                          "/usr/bin"
                          "/usr/local/bin")))
      (setq exec-path my-emacs-path)
      (setenv "PATH" (mapconcat 'identity my-emacs-path ";")))
    
    ;; Add my lisp dir to load-path
    (let ((default-directory "~/.emacs.d/lisp/"))
      (normal-top-level-add-subdirs-to-load-path))            
    
    ;; Set SHELL Value
    (setq shell-file-name "zsh")
    (setenv "SHELL" shell-file-name) 
    (setq explicit-shell-file-name shell-file-name)

日本語の設定や、キーカスタマイズなど、その他設定することは、まだまだたくさんありますが、それはまた次回ということで Windows に Linux環境の構築を〆る。
