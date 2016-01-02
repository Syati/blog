+++
Categories = ["TypeScript"]
Description = "TypeScript でユニットテストする方法をメモる。加えてカバレッジ測定も行う"
Tags = ["test", "coverage"]
comments = true
date = "2015-10-31T14:00:39+09:00"
title = "TypeScript でユニットテストする。カバレッジもねっ!!"

+++

AltJs ってトランスパイルした後、テストを実行させる必要があるから・・、面倒くさいんだよね。なんか、あまり情報もないし、これだから AltJs は・・・と思うわけですが、やってみると案外大したことないもんです。

今回は、TypeScript のコードのユニットテスト、カバレッジを記していこうと思います。本記事のコードは以下のURLにあります。

https://github.com/Syati/typescript-sample/tree/master/011_test_sample

<!--more-->

## Step 0: 事前準備

### 流れ

1. ts でコードを書く
2. トランスパイルする（js と map を出力）
3. js をテストする
4. map を利用して**ts カバレッジ**を測る

### ライブラリ

**npm install -g mocha istanbul remap-istanbul** でそれぞれコマンドが利用できるようになります。

- Test framework
  - [mocha](https://www.npmjs.com/package/mocha)
      - テストフレームワークですね。他にも karma, jasmine など、たくさんありますがシンプルなので、とっつきやすかった。
- Coverage
  - [istanbul](https://www.npmjs.com/package/istanbul)
      - カバレッジを測定してくれる。カバレッジがあると、ユニットテストのモチベーションが上がりますよね！！
  - [remap-istanbul](https://www.npmjs.com/package/remap-istanbul)
      - istanbul だけでは、トランスパイル後（js）のカバレッジしか測れないので、map を利用して、ts のカバレッジを測る


## Step 1: とりあえずテストを書く

[プロジェクト](https://github.com/Syati/typescript-sample/tree/master/011_test_sample) を確認しておく。全体が理解しやすくなると思う。

テストは、単にコンストラクタで与えられた値を保持しているかどうか。

~~~typescript
// app/src/main.ts
export default class TestTarget {
    constructor(public name){
    }
}
~~~

~~~typescript
// app/test/main.test.ts
import assert from 'power-assert';
import TestTarget from '../src/main';

describe("TestTarget", () => {
    it("should have a name", () => {
        let testTarget = new TestTarget("test");
        assert.equal(testTarget.name, "test");
    });
});
~~~

すごいシンプルですね。


## Step 2: トランスパイルする

tsc コマンドを叩く前に設定ファイルを書いときましょう。以下の tsconfig.json になります。
とくに変わった点はないですね。重要なポイントとしては、**"sourceMap": true**です。
後々、ts のカバレッジを測るために必要になってきます。

~~~javascript
{
  "version": "1.6.2",
  "compilerOptions": {
    "module": "commonjs",
    "target": "es5",
    "sourceMap": true
  },
  "exclude": [
    "dist",
    "node_modules"
  ],
  "files": [
    "typings/tsd.d.ts",
    "app/src/main.ts",
    "app/test/main.test.ts"
  ]
}
~~~

プロジェクトルートで tsc コマンドを叩きます

~~~bash
$ tsc
~~~

## Step 3: テストする

Step 2 でトランスパイル済みなので、ts ファイル以外に js、map ファイルが存在すると思います。ディレクトリが汚れてあたふたするかもしれませんが、package.json に clean スクリプトを書くことで解決できるので、今は我慢。

プロジェクトルートでテストコマンドを叩いてみます。

~~~bash
$ mocha app/test/*.test.js

  TestTarget

    ✓ should have a name


  1 passing (8ms)
~~~

成功、終わりです。

味気ないですね。

ユニットテストのモチベーションをあげるためにも、やっぱりカバレッジですよ！！

## Step 4: カバレッジを測る

とくに必要なことはありません。Step 0 で説明したとおり **istanbul** を利用するだけ。
すこし異なっているのは、 mocha が _mocha を利用するということです。理由は、以下のとおり。

https://github.com/gotwarlost/istanbul/issues/44#issuecomment-16093330

~~~bash
$ istanbul cover _mocha -- ./app/test/*.test.js

  TestTarget

    ✓ should have a name


  1 passing (5ms)


=============================== Coverage summary ===============================
Statements   : 100% ( 6/6 )
Branches     : 100% ( 0/0 )
Functions    : 100% ( 2/2 )
Lines        : 100% ( 6/6 )
================================================================================
~~~

出力に、**coverage summary** が増えましたね。また、プロジェクトルートに **coverage** ディレクトリが出力されているので、中の index.html を見ましょう。通過したコード、100% の数字にやる気が出てきますね。ただし、これはトランスパイル後の js のカバレッジなので、あと一手間かけてあげましょう。

{{< figure src="coverage_js.png" title="トランスパイル後のカバレッジ" >}}

## Step 5: ソースマップから ts のカバレッジを測る

**coverage** ディレクトリにて以下のコマンドたたいてあげるだけです。ts のカバレッジが html-repot ディレクトリに出力されます。

~~~bash
$ remap-istanbul -i coverage.json -o html-report -t html
~~~

{{< figure src="coverage_ts.png" title="ts のカバレッジ" >}}


ts のカバレッジが測れました。ただ、一連のコマンドを毎回叩くのは疲れるので、**package.json** を用いて楽にしましょう。


＊図が Step 4 とデザインが違うのは、remap-istanbul で利用している istanbul のバージョンが古いからだと思われる

## Step 6: package.json で楽する

一連のコマンドを **package.json** に書いておきましょう。以下サンプル。

~~~javascript
{
  ...

  "scripts": {
    "postinstall": "tsd install",
    "clean": "find -E . -regex './app/.+\\.(js|map)$' -type f | xargs rm && rm -rf ./build",
    "pretest": "tsc",
    "test": "istanbul cover _mocha -- ./app/test/*.test.js",
    "posttest": "cd coverage && remap-istanbul -i coverage.json -o html-report -t html"
    },
    
    ...
}
~~~

大事なのは、 **1. pretest、2. test、3. posttest** です。

**npm test** を実行することで **1. 2. 3.**の順番で実行してくれます。今回の step2 - step5 をコマンド１つで済せられるということですね。

~~~bash
# 次回からはこれだけでよい
$ npm test 

# js とか map 削除したいとおもったら
$ npm clean
~~~

## その他

### .istanbul.yml でカバレッジレポートのディレクトリ指定する

**.istanbul.yml** に以下のように書いてあげることで、レポートの出力先を変えることが出来る。

~~~yaml
#.istanbul.yml
reporting:
    dir: ./docs/coverage
~~~

もっとカスタマイズしたい場合は、以下のコマンドからデフォルト設定など説明をみるといいかも。

~~~bash
$ istanbul help config
~~~
