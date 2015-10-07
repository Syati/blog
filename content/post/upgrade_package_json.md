+++
Categories = ["JavaScript"]
Description = "package.json の dependencies をどうやって一括更新するのかなと思ってたらあった！！すごく楽だった。"
Tags = ["npm", "node"]
comments = true
date = "2015-10-07T22:00:32+09:00"
title = "package.json を一括更新"

+++

package.json の dependencies どうやって一括で更新するのかなと思ってたらあった！！

[npm-check-updates](https://www.npmjs.com/package/npm-check-updates)

これで簡単に更新できる！！

<!--more-->

# 更新する

以下の通り。すごく簡単だね。

~~~bash
$ npm install -g npm-check-updates
# package.json のあるフォルダで
$ ncu -u
 body-parser    ~1.8.1  →  ~1.14.1 
 cookie-parser  ~1.3.3  →   ~1.4.0 
 debug          ~2.0.0  →   ~2.2.0 
 jade           ~1.6.0  →  ~1.11.0 
 mongoose       ~4.0.8  →  ~4.1.10 
 morgan         ~1.3.0  →   ~1.6.1 
 serve-favicon  ~2.1.3  →   ~2.3.0
 
Upgraded /your/path/package.json
 
$ npm install
~~~

一括更新完了！！
