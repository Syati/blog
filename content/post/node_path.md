+++
Categories = ["JavaScript", "Node"]
Description = "node で REPLを利用する際の環境パス"
Tags = ["Env", "REPL"]
date = "2015-06-04T23:09:38+09:00"
title = "node 利用するなら、まずは環境変数を設定しましょう"

+++


javascript を REPL で利用できるといえば **node** 。underscore.js を試したいと思いハマったメモ。

# はまった事

以下の通り、グローバルにunderscoreをいれてnodeを起動したのちunderscoreを読み込むとエラー。

~~~bash
$ sudo npm install undersocre -g
$ node 
    
require('underscore'); 
Error: Cannot find module 'underscore'
  at Function.Module._resolveFilename (module.js:338:15)
  at Function.Module._load (module.js:280:25)
  at Module.require (module.js:362:17)
  at require (module.js:378:17)
  at repl:1:1
  at REPLServer.self.eval (repl.js:109:21)
  at rli.on.self.bufferedCmd (repl.js:258:20)
  at REPLServer.self.eval (repl.js:116:5)
  at Interface.<anonymous> (repl.js:248:12)
  at Interface.EventEmitter.emit (events.js:96:17)
~~~

# 原因

npmでインストールしたnode\_modulesのパスを環境変数に設定していなかっただけ。

# 解決

1.  設定するパスを確認する

~~~bash
npm root -g
/usr/local/lib/node_modules
~~~

2.  利用ターミナル(bash なら .bashrc、zsh なら .zshrc ）に１の結果を以下のように追加して終了

~~~bash
export NODE_PATH=/usr/local/lib/node_modules
~~~
