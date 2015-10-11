+++
Categories = ["TypeScript"]
Description = "TypeScript + webpack でフロントエンドを実装する"
Tags = ["webpack"]
comments = true
date = "2015-10-10T18:37:23+09:00"
title = "TypeScript でフロントエンドを実装する"

+++

型が恋しいのです。仮引数見た時に、何が入るか簡単に知りたいのです。実引数の型が違った時には教えてほしいのです。そんなこんなで、TypeScript がお気に入りです。

今回は、どうやってフロントエンドで TypeScript を利用するか記していこうと思います。本記事のコードは以下のURLにあります。

https://github.com/Syati/typescript-sample/tree/master/webpack

<!--more-->

# Step 0: 事前準備

以下のコマンドを使用するのでグローバルにインストールしておく。

- Package managers
  - [npm](https://www.npmjs.com/): プロジェクトで利用するライブラリをインストールするために利用する。
  - [tsd](http://definitelytyped.org/tsd/): プロジェクトで利用するライブラリの定義ファイルをインストールするために利用する。

        ~~~bash
        npm install tsd -g
        ~~~

- typescript compiler
  - [tsc](http://www.typescriptlang.org/): ts ファイルをコンパイルするために利用する。

        ~~~bash
        npm install typescript -g
        ~~~

- Build tools
  - [webpack](https://webpack.github.io/docs/): フロントエンドでモジュールをロードするために利用する。

        ~~~bash
        npm install webpack -g
        ~~~


# Step 1: いろいろ初期化(プロジェクトの作成)

~~~bash
$ mkdir app
$ tsd init  # いろいろ尋ねてくるけど enter 連打
$ npm init

$ tree ./ -L 1
./
├── app             # ここに html, ts などを入れる
├── package.json    # npm init で作成される（ライブラリ管理）
├── tsd.json        # tsd init で作成される（ライブラリの定義を管理）
└── typings         # tsd init で作成される（ライブラリの定義のソースコードが入る）
~~~

# Step 2: ライブラリのインストール

npm と tsd を利用して、ライブラリ、ライブラリの定義をインストール。

~~~bash
# step 1 ./ で以下のコマンドを実行
## プロジェクトで利用するライブラリのインストール
$ npm install typescript --save
$ npm install jquery --save
## TypeScript で利用するにはもちろん定義が必要なので、tsd で jquery をインストール
$ tsd install jquery --save
~~~

# Step 3: app を作成

app の下に app.ts として以下を作成。

~~~~typescript
import * as $ from 'jquery';

$(() => {
    $(document.body).html("hello");
});
~~~~

app の下に index.html として以下を作成。 後に app.ts から bundle.js を作成するため、この段階で script を埋め込んでおく。

~~~markup
<!DOCTYPE html>
<html>
    <head>
        <meta charset="UTF-8"/>
        <script src="build/bundle.js"></script>
    </head>
    <body>
    </body>
</html>
~~~

# Step 4: コンパイルする

step 3 で作成した app.ts を js に変換するために tsc の設定を作成する。
**step 1 ./** で tsconfig.json として以下のように作成。

~~~javascript
{
  "compilerOptions": {
    "sourceMap": false,     # ソースマップ作成の有無
    "target": "ES5",        # target の設定
    "outDir": "app/build",  # 出力先
    "module": "commonjs"    # module のスタイル
  },
  "files": [                # ファイル
    "typings/tsd.d.ts",
    "app/app.ts"
  ]
}
~~~

とりあえずコンパイルしてみる。

~~~bash
# step 1 ./ で以下のコマンドを実行
tsc
~~~

app/build に app.js が作成されているので覗いてみる。**require** の記述があるが、もちろんフロントではつかえない。そこで、フロントでも require を利用するために、**webpack** を用いる。

~~~javascript
var $ = require('jquery');
$(function () {
    $(document.body).html("hello");
});
~~~

ここまでのプロジェクトディレクトリは以下のとおり。

~~~bash
./
├── package.json
├── tsconfig.json
├── tsd.json
├── app
│   ├── index.html
│   ├── app.ts
│   └── build
│       └── app.js
├── node_modules
│   └── jquery
│   └── typescript
└── typings
    ├── jquery
    │   └── jquery.d.ts
    └── tsd.d.ts
~~~

# Step 5: webpack 導入

フロントでも require を利用するために、**webpack** をインストールして、ついでに、TypeScript のファイルを処理できるように loader をいれる。

~~~bash
# webpack のインストールは、step 0 でインストールしたものとする
# typescript ファイルを処理できるように以下の webpack 用の loader をプロジェクトのライブラリに追加する
npm install ts-loader --save
~~~

webpack コマンドをたたけば require などの依存関係を解決してくれるのだが、いろいろとオプションを指定する必要がある。毎回コマンドの度に、それらを打ち込むのは面倒くさいので、設定ファイルに書いておいてあげる。

**step 1 ./** で webpack.config.js として以下のように webpack の設定を作成する。

~~~javascript
module.exports = {
  entry: './app/app.ts',
  output: {
    filename: './app/build/bundle.js'
  },
  resolve: {
    // Add `.ts` and `.tsx` as a resolvable extension.
    extensions: ['', '.webpack.js', '.web.js', '.ts', '.tsx', '.js']
  },
  module: {
    loaders: [
      // all files with a `.ts` or `.tsx` extension will be handled by `ts-loader`
      { test: /\.tsx?$/, loader: 'ts-loader' }
    ]
  }
}
~~~

これで簡単に webpack とコマンドをたたくだけ。

~~~bash
# step 1 ./ で以下のコマンドを実行
$ webpack
~~~

app/build に bundle.js が作成されているので覗いてみる。require が解決され、単一の js ファイルとして出力されていることがわかる。

これで index.html をブラウザーで覗いてみると、**hello** の文字が出力されている。

[ソースはこちらから](https://github.com/Syati/typescript-sample/tree/master/webpack)

# 参考

- [TypeScript and webpack](http://www.jbrantly.com/typescript-and-webpack/)
