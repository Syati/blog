+++
Categories = ["OSX"]
Description = "私の OSX の環境をメモっておく"
Tags = ["Env"]
date = "2015-06-04T23:09:38+09:00"
title = "Development Environment in my OSX"

+++

仕事で Mac を利用することになった Syati です。main os が Ubuntu だからか、Windows より全然開発環境を構築しやすい。
ただ、やることはたくさんあった。ストレスが無い開発環境を構築することは難しいものです・・・。ある程度、構築できたので記録しておく。
<!--more-->

## 開発環境

-   xcode:開発環境
    -   プラグイン
        -   command line tools (os x mountain lion) for xcode
-   fink:package management のソフト
    -   上記を用いて入れたソフトウェア
        -   coreutil-default
            -   prefix に g がない方
            -   /sw/bin にコマンドが入るので、パスの優先度を上げておく
        -   wget
        -   tmux
    -   問題：mac の tmux では pbcopy/pbpaste で クリップボードの共有ができない
    -   解決：[X環境のクリップボードやOS Xのペーストボードとtmuxのバッファを連携する方法](http://d.hatena.ne.jp/tmatsuu/20111220/1324399472)
        -   autoconf
            -   古かったので結局自分で入れた
-   emacs24.2:エディター
    -   デフォルトは emacs22.1 のため導入する。ただし、パッチを入れないと最大化等できない。以下のサイト参照のこと
        -   <http://mgrace.info/?p=1032>
        -   <http://sakito.jp/emacs/emacs24.html>
-   デフォルトターミナルの変更
    -   コマンド **chsh** で zsh にする
-   iTerm2：ターミナル
    -   デフォルトのターミナルに比べて、なんか見やすかったから入れといた
        -   <http://www.iterm2.com/#/section/home>
    -   meta-key が使えるように以下を実施する
        -   <http://d.hatena.ne.jp/kitokitoki/20111129/p1>
    -   zsh + tmux を使っていると、当該ソフトの autocomplete とか tab とか、あまり必要性を感じないんだが、どうなんだろうか。なんかメリットあったら教えて下さい。

# ブラウジング

-   adobe pdf
-   adobe flash
-   firefox:ブラウザ
    -   meta key が効かない問題あり・・・。現在対処不明
    -   アドオンのインストール
        -   keysnail
    -   plugin(従来と異なる点だけ)
        -   Set Mac
            -   firefox で Ctl + space が利用できるようになる
                -   導入しない場合 Ctrl + space では、 右クリックになる
    -   問題：acrobat reader をインストールしたのに、firefox から pdf の DL がされない、見れない、見れたと思ったら文字化けして困った
    -   解決：firefox を選択した後、上部のメニューバーの「 firefox 」-> 「環境設定」-> 「アプリケーションタブ」を選択して以下の通りする
        -   ファイルの種類：Adobe Pdf document
        -   取り扱い方法　：Adobe Readerを使用（標準設定）
    -   ここが、Adobe Acrobat NPAPI Plug-in `~` とかになっていたのが原因だった

## キーバインド設定

-   HyperSwitch：ウィンドウ切り替え
    -   mac の command + tab だと、アプリケーション単位の切り替えしかできないことに加えて、main os の ubuntu とキーバインド alt + tab で揃えたいので導入する
        -   <http://bahoom.com/hyperswitch>
    -   問題：HyperSwitch の設定から command + tab -> alt + tab に変更したが、ターミナルではうまく働かない
    -   解決：ターミナルを選択した後、上部のメニューバーの「ターミナル」 -> 「キーボード入力のセキュリティを保護」のチェックを外す
        -   「キーボード入力のセキュリティを保護」が原因だったけど、これ何？
-   KeyRemap4MacBook:キーバインド変更
    -   キーボードの「英数」「かな」がうざいので、 以下の通り違うキーに変更する
        1.  KeyRemap4MacBookを立ち上げて、Change Key タブから For Japanese メニューを選択する
        2.  Change EISUU Key メニュー選択して、好きなキーバインドに変更する
        3.  Change KANA key も同様に実施する
-   IMEの変換（ひらがなと英数）を Ctrl + / にしたいので以下を実施する
    1.  「アップルボタン」 -> 「システム環境設定」 -> 「キーボード」-> 「キーボードショートカット」 -> 「キーボードと文字入力」を選択する
    2.  「前の入力ソース」をダブルクリック後、Ctrl + / をタイプする

## その他

-   inconsolate:フォント
-   takao:フォント
-   QuickSilver：ランチャー
    -   どこからでもキーボードでアプリケーションを起動したいため導入する
        -   <http://quicksilver.en.softonic.com/mac>
-   ShifIt:ウィンドウサイズ変更アプリ
    -   Windows の windowボタン + カーソル で ウィンドを半分にしたり拡大したり出来たと思いますが、それを可能にする。
        -   <http://code.google.com/p/shiftit/>
        -   [ウィンドウのサイズ変更や移動をホットキーから『ShiftIt』](http://veadardiary.blog29.fc2.com/blog-entry-2764.html)
-   マウスのホイール操作が逆なので戻す
    1.  「アップルボタン」 -> 「システム環境設定」 -> 「マウス」-> 「スクロールの方向：ナチュラル」のチェックを外す
