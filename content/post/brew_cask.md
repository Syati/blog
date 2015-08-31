+++
Categories = ["OSX"]
Description = "OSX の App のインストールは cask が便利"
Tags = ["Env"]
comments = true
date = "2015-08-31T18:44:21+09:00"
title = "brew cask が便利"

+++

Web ページや App Store から App をダブルクリックしてインストールは面倒くさいので **cask** をいれて解決する。

<!--more-->

# インストールとコマンド

## インストール

**cask** コマンドを使えるようにするために以下のコマンドを実行する。

~~~bash
$ brew install caskroom/cask/brew-cask
~~~

## コマンド

検索

~~~bash
$ brew cask search chrome

==> Partial matches
chrome-devtools                    google-chrome                      google-chrome-dev
chrome-remote-desktop-host         google-chrome-beta
chromecast                         google-chrome-canary
~~~

インストール

~~~bash
$ brew cask install google-chrome
~~~

アンインストール

~~~bash
$ brew cask uninstall google-chrome
~~~

# その他

cask でインストールしたアプリは以下にインストールされ、ホームフォルダ(~/)の Applications からアプリへのシンボリックリンクが作成される。ルート(/)にインストールが必要なものはパスワードが求められる。

~~~bash
/opt/homebrew-cask/Caskroom/
~~~

以下、~/Applications の例

~~~bash
~/Applications
├── Alfred 2.app -> /opt/homebrew-cask/Caskroom/alfred/2.7.2_400/Alfred 2.app
├── Alfred Preferences.app -> /opt/homebrew-cask/Caskroom/alfred/2.7.2_400/Alfred 2.app/Contents/Preferences/Alfred Preferences.app
├── Chrome Apps.localized
├── Emacs.app -> /opt/homebrew-cask/Caskroom/emacs/24.5-1/Emacs.app
├── Firefox.app -> /opt/homebrew-cask/Caskroom/firefox/40.0.3/Firefox.app
├── FirefoxDeveloperEdition.app -> /opt/homebrew-cask/Caskroom/firefoxdeveloperedition-ja/latest/FirefoxDeveloperEdition.app
├── Google Chrome.app -> /opt/homebrew-cask/Caskroom/google-chrome/latest/Google Chrome.app
├── ShiftIt.app -> /opt/homebrew-cask/Caskroom/shiftit/1.6.3/ShiftIt.app
└── Visual Studio Code.app -> /opt/homebrew-cask/Caskroom/visual-studio-code/0.7.0/Visual Studio Code.app
~~~

# 参考

- [Homebrew Cask](http://caskroom.io/)
