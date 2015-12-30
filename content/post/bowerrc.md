+++
Categories = ["JavaScript"]
Description = "bowerrc を用いることでデフォルトのインストールディレクトリを変更できる"
Tags = ["bower"]
date = "2015-03-01T21:37:06+09:00"
title = "bowerrc で install ディレクトリを変更する"

+++

bower install をしたはいいが、デフォルトでインストールされるディレクトリ(bower\_components)を以下の例のように変えたい、そんな時は。

**例** 

-   project/
    -   js
        -   bower\_components
            -   jquery
    -   img
    -   css
    -   index.html
    -   bower.json

<!--more-->

## 解決策

以下のように **.bowerrc** を project ルートに作成してあげる。これを自分の好きなように書き換えることで、インストールディレクトリを変更できる。

~~~javascript
// .bowerrc
{
    "directory" : "js/bower_components"
}
~~~

以降、bower install すれば希望通り

### .bowerrc作成後

-   project/
    -   js
        -   bower\_components
            -   jquery
    -   img
    -   css
    -   index.html
    -   bower.json
    -   .bowerrc

## 参考

### デフォルトの動作

#### 初期ディレクトリ構成

-   project/
    -   js
    -   img
    -   css
    -   index.html

#### コマンド

project ルートで 以下のように bower コマンドを入れる

    bower init
    bower install jquery -save

#### コマンド後のディレクトリ構成

ディレクトリは以下のようになる。

-   project/
    -   bower\_components
        -   jquery
    -   js
    -   img
    -   css
    -   index.html
    -   bower.json
