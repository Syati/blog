+++
Categories = ["JavaScript"]
Description = "Angular controller の unit test をする"
Tags = ["Angular", "UnitTest"]
date = "2015-06-04T23:09:38+09:00"
title = "Angular test sample 1"

+++


js ってテストしにくいよねってことで今流行りの angular で unit test する。 
最初なので、事細かに説明するつもりで書いてみる。

サンプルは以下から git clone してください。ソースと突き合わせながら見るといいかも。

<https://github.com/Syati/angular-test-sample>
<!--more-->

# 準備

## ubuntu

- node & npm

~~~bash
$ sudo apt-get install nodejs
$ sudo apt-get install npm
~~~

## mac

- node & npm

~~~bash
$ brew install node
~~~

## 共通

- bower

~~~bash
$ sudo npm install -g bower
~~~

- karma

~~~bash
$ sudo npm install -g karma-cli
~~~

# 実施：サンプルを CLONE して TEST

github にサンプルを書いたので clone する。

~~~bash
$ git clone git@github.com:Syati/angular-test-sample.git
~~~

## prj 依存環境を入れる

~~~bash
$ cd your_clone_path/angular-test-sample/controller/recipe1
$ npm install
~~~

## test 実行する

~~~bash
$ npm test
~~~
    
以下のような出力で、テストが成功する。

~~~bash
WARN [karma]: Port 9876 in use
INFO [karma]: Karma v0.12.24 server started at http://localhost:9877/
INFO [launcher]: Starting browser PhantomJS
INFO [PhantomJS 1.9.8 (Linux)]: Connected on socket -9tnO3qPnHscs3zVjS4l with id 62869514
PhantomJS 1.9.8 (Linux): Executed 1 of 1 SUCCESS (0.038 secs / 0.007 secs)
~~~
    
# 説明

## directory
~~~bash
    recipe1
    ├── app
    │   ├── bower_compornents #1 bower js lib install dir
    │   ├── css
    │   ├── index.html
    │   ├── js
    │   │   └── app.js       #4 target js
    │   └── test              #2 test files
    │        └── app.spec.js  #4 test js
    ├── .bowerrc      #1 bower install dir 設定
    ├── bower.json    #1 bower install する js lib などの設定
    ├── karma.conf.js #3 angular のテスト設定
    ├── package.json  #2 npm lib などの設定 
    └── node_modules  #2 npm lib install dir
~~~
    
## #1 bower

js のパッケージ(jqueryなど)を DL してくれる便利パッケージ。

~~~bash
$ bower init  # コマンドで初期設定 bower.json を作成
$ bower install angular --save　# または --save-dev
~~~
 

デフォルトは **bower.json** がある場所がインストール先になるので **.bowerrc** にインストール先を書いておく。
上記を実施することで以下のファイルを作成する。

- bower.json
- .bowerrc

- その他
    - .gitignore に bower\_components を入れておく（bower install でいつでも同じ環境を構築できるので）。
    - bower のオプション **&#x2013;save** または **&#x2013;save-dev** を分けることで、本番用と開発用を分けることが出来る。
        - **bower install &#x2013;production** で &#x2013;save-dev で入れたパッケージはインストールせずに済むのである。
    - 参考
      -   <http://bower.io/docs/api>

## #2 npm

~~~bash
$ npm init 
$ npm install jasmine --save-dev　# または --save
~~~

上記を実施することで以下のファイルを作成する。

- package.json
- その他
    - .gitignore に node\_modules を入れておく（npm install でいつでも同じ環境を構築できるので）。
    - npm のオプション **&#x2013;save** または **&#x2013;save-dev** を分けることで、本番用と開発用を分けることが出来る(bower と同様ですね)。
        - **npm install &#x2013;production** で &#x2013;save-dev で入れたパッケージはインストールせずに済むのである。
        - 今回は test 用しかないので &#x2013;save-dev ですべてインストールする。
    - **package.json 内の scripts って便利**
    - 参考
        - <https://www.npmjs.org/doc/cli/npm-install.html>

## #3 karma  angular の unit test の設定

~~~bash
karma init  # コマンドで初期設定 karma.conf.js を作成できる
~~~

今回特に意識する場所としては以下の通り。

- basePath
    - 基準とするパスなので設定ファイルのある場所 **./** にする。
- files
    - ここで必要な js lib と テストのターゲット js とテスト js を読み込ませてあげる。
- browsers
    - init のデフォルトだと Chrome だが、テストするたびに Chrome が開くのは邪魔なので、PhantomJS にする。
        - [PhantomJS](http://phantomjs.org/)
- plugins
    - [jasmine](http://jasmine.github.io) を利用するために以下を追加する。
        - karma-jasmine
    - karma から PhantomJS を呼べるようにするために以下を追加する。
        - karma-phantomjs-launcher

## #4 test を書く

テストするうえで jasmine の理解は必須なので [tutorial](http://jasmine.github.io/2.1/introduction.html) で勉強しておく。
index.html をローカルで開けばわかると思うが button click で 'hello' とでてくるだけのもの。
テストする js は app.js である。

テストコードにコメント付与した。

~~~javascript
'use strict';
 
describe('Unit: app moudle ', function(){
  var scope; # scope がどこからでも呼べるように定義しておく

  beforeEach(function(){
    angular.mock.module('app'); # モジュールのモックを作成する
  });

  describe('MainCtrl', function(){
    # ここで DI する。今回は MainCtrl cotroller で GET, POST など何もしていないので
    # $rootScope と $controller のみ
    beforeEach(inject(function($rootScope, $controller){ 
      scope = $rootScope.$new(); # scope を作成して、上記で定義した scope に代入する

      # scope を渡してコントローラを作成する
      # これで scope に controller 内に定義した sayHello 関数などが入る 
      $controller('MainCtrl', {'$scope': scope}); 
    }));

    it('sayHello', function(){
      scope.sayHello();
      # scope.message が 'hello' となっているかテスト
      expect(scope.message).toBe('hello');
    });
  });
});
~~~

-   その他
    -   angular.mock.module を利用するうえで、bower.json に記述のある angular-mocks は必要。

# まとめ

controller のテストは意外と簡単。ただし、DI が増えてくるとハマるポイントも多くなってくるので、次はそんなサンプルを書いていくつもり。
