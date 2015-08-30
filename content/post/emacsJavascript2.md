+++
Categories = ["Emacs", "JavaScript"]
Description = "Emacs で Javascript なら js2-mode, tern-mode"
Tags = ["Env"]
date = "2015-08-31T13:38:35+09:00"
title = "Emacs で Javascript なら js2-mode, tern-mode"

+++

[Emacs で JavaScript の開発環境を整える]({{< ref "post/emacsJavascript.md" >}}) で、
Emacs と Javascript の開発環境を記しましたが、さらに改善できたので、メモメモ。

ドットで jquery も underscore も browser もしてくれる。 angular も!!

<!--more-->

# Step 1 tern のインストール・設定

1. npm で tern をインストールする

    ~~~bash
    $ npm install -g tern
    ~~~

2. ~/.tern-config を以下のような感じで作成

    ~~~javascript
    {
      "libs": [
        "browser",
        "jquery",
        "ecma5",
        "underscore"
      ],
      "plugins": {
        "angular": {},
        "Node": {}
      }
    }
    ~~~

# Step 2 emacs の package をインストール

1. M-x package-list-package で以下の2点をインストールする（該当箇所で **i** をタイプしてマークをつけて **x** でインストールする）
    1.  tern
    2.  js2-mode
    3.  tern-auto-complete
2. ~/.emacs.d/init.el に以下を追加

    ~~~clike
    (autoload 'js2-mode "js2-mode" nil t)
    (add-to-list 'auto-mode-alist '("\\.js\\'" . js2-mode))

    (add-hook 'js2-mode-hook
        (lambda ()
            (tern-mode t)))

    (eval-after-load 'tern
        '(progn
            (require 'tern-auto-complete)
            (tern-ac-setup)))
    ~~~

# Step 3 試す

~~~javascript
angular
  .module('MyApp', [])
  .controller('SomeCtrl', function($scope, $http) {
    # ここで $scope. とでも入れてみる
  })
~~~

# 参考

-   [tern](http://ternjs.net/doc/manual.html)
-   [js2-mode](http://code.google.com/p/js2-mode/)
-   [npm](https://npmjs.org/)
