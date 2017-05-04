+++
Categories = ["Tool"]
Description = "Emacs ユーザーで Sequel Pro を利用しているなら ctrl+h ヘルプを backspace に変更することで、幸せになれるはず。"
Tags = ["Keybindings", "Emacs"]
comments = true
date = "2017-04-29T19:02:10+09:00"
title = "Sequel Pro の ctrl+h ヘルプを backspace に変更する"
logo = "hack"

+++

Sequel Pro 便利ですよね。簡単に select とか DB 作成、きれいな結果表示などなど。
ただ、emacs ユーザーの私にとってウザくてウザくて堪らないのが、**ctrl+h** を打つと、このソフトでは**ヘルプ**を開くということ。頼むから Backspace にしてくれ！！と。

調べてみると、上書き出来るではないか！！[github の issue](https://github.com/sequelpro/sequelpro/issues/1974)に書いてあった。ということで以下に、その方法記す。

<!--more-->

## 準備

とりあえず、デフォルトでは ctrl+h でこのヘルプがでてくる。

{{< figure class="image-half__center" src="popup_help.jpg" title="fig 1. help が出力される" >}}

なんで、ヘルプがでるかというと以下で ctrl+h にショートカットが割り当てられているため。

{{< figure class="image-half__center" src="default_shortcut.jpg" title="fig 2. default のショートカット" >}}

## 修正

### 1. システム環境設定 -> キーボード -> ショートカットタブ -> アプリケーション

{{< figure class="image-half__center" src="default_keyboard_shortcut.jpg" title="fig 3. アプリケーションのショートカット" >}}

### 2. + ボタンから新規追加する

大事なのは Sequel Pro で既存で割当てられいるメニューコマンドと同じにするということ。すなわち **準備 fig 2.** に表示されている **^H** のメニューコマンド **単語の MySQL ヘルプ** である。とりあえず **Shift + Command + H** にでも割り当てる。

{{< figure class="image-half__center" src="add_keyboard_shortcut.jpg" title="fig 4. Sequel Pro の既存ショートカットを上書きする" >}}


{{< figure class="image-half__center" src="added_keyboard_shortcut.jpg" title="fig 5. 登録完了" >}}

### 3. Sequel Pro でショートカットが上書きされたかを確認する

Sequel Pro のヘルプを確認する。変更されていることが確認できる。

{{< figure class="image-half__center" src="added_shortcut.jpg" title="fig 6. ショートカット上書き完了" >}}

後はクエリで ctrl+h で backspece できることを確認する。

## 終わりに

ctrl + h で *ヘルプに悩まされる日々、さようなら*
