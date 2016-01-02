+++
Categories = ["JavaScript"]
Description = "Sequelize ではじめる Node ORM。とりあえず導入して、Node の REPL で試してみる"
Tags = ["Sequelize", "ORM"]
comments = true
date = "2015-12-29T12:38:07+09:00"
title = "Sequelize Part 1 - 導入してREPLで試してみる"

+++

ちょっと前まで MEAN とか流行っていた気がするんですが、今はどうなんですかね？実際にちょっとしたWEBサービスで利用してみると、MongoDB が思った以上に厄介だった。トランザクションが無いのが、どうしても辛い。後々考えると面倒くさいことになるのは目に見えているのでヤメヤメ。ということで従来通り RDB を利用、そして、ORM は **Sequelize** を利用したので、その辺をメモっていきたいと思う。

以降で説明するコードは、[01_repl](https://github.com/Syati/sequelize-sample/tree/master/01_repl) にあります。

<!--more-->

## 環境

- OSX 10.10.5
- Node 5.3.0
- MariaDB 10.1.9

## ORM 導入に際して

Node で RDB を利用するために、ORM のライブラリを検討。意外とあるもんですね。
以下に github の star と fork 数をまとめる(2015/12/29 調べ)。


| Lib                                                   | Star | Fork |
|-------------------------------------------------------|:----:|:----:|
| [waterline](https://github.com/balderdashy/waterline) | 2541 |   359|
| [bookshelf](https://github.com/tgriesser/bookshelf)   | 2824 |   228|
| [node-orm2](https://github.com/dresende/node-orm2)    | 1655 |   282|
| [sequelize](https://github.com/sequelize/sequelize)   | 5344 |  1160|

それぞれの Lib の差異はわかりませんが Star と Fork 数で判断して、**Sequelize で決定**。ドキュメントも見やすかったのでいいかな。

## 1. DB・ORM・ORM-CLI のインストール

### 1.1. DB インストール

mysql, posgre がある場合は不要。今回は mariadb を利用する。

~~~bash
$ brew update
$ brew install mariadb 

$ unset TMPDIR
$ mysql_install_db

## start mariadb
$ mysql.server start

## connect mariadb
$ mysql -uroot
~~~

- 参考
    - [Building MariaDB on Mac OS X Using Homebrew](https://mariadb.com/kb/en/mariadb/building-mariadb-on-mac-os-x-using-homebrew/)


### 1.2. ORM (Sequelize) インストール

~~~bash
$ npm install --save sequelize
$ npm install --save mysql // For both mysql and mariadb dialects
~~~

### 1.3. ORM-CLI (Sequelize CLI) インストール

初期設定、モデル作成、マイグレーション など便利なので CLI をいれる。
コマンドラインをグローバルで利用したいので **g** 付与する。

~~~bash
$ npm install -g sequelize-cli
~~~

## 2. 各種設定

現状の dir 構成は以下のとおり。
~~~bash
sample_prj
├── node_modules
└── package.json
~~~

### 2.1. DB 作成

今回利用する DB を **seq_dev** として、 mariadb に作成する。

~~~bash
$ mysql -uroot
## mariadb に接続後、今回利用するデータベースを作成する
$ create database seq_dev character set utf8;
~~~

### 2.2. 初期設定

**1.3** でインストールした Sequelize CLI を用いて初期設定する。

~~~bash
$ sequelize init #プロジェクトルートで実行。初期設定を作成する

# dir 構成は以下になる
sample_prj
├── config     # config.json に DB の初期設定が書き出される
├── migrations # cli で model:create した際、migration が自動作成される
├── models     # cli で model:create した際、model が自動作成される
├── node_modules
├── package.json
└── seeders    # cli で seed:create した際、seed が自動作成される

# 基本的に上記は .sequelizerc にてカスタマイズ可能
~~~

**config/config.json** に初期設定があるが、先ほど作成したデータベース名と異なっているので変更する。

~~~javascript
// サンプル config.json
{
  "development: {
    "username": "root",
    "password": null,
    "database": "seq_dev",
    "host": "127.0.0.1",
    "dialect": "mysql"
  }
}
~~~

## 3. はじめてのモデル作成

### 3.1. model の作成

models ディレクトリに model.js のように手動でモデルを作成しても良いのだが、CLI を利用してモデルを作成してあげると、マイグレーションも一緒に作成してくれるので、CLI を利用すると楽かもしれない。

**※ただし、[作成したモデルを変更しても該当マイグレーションを更新する術がない](http://stackoverflow.com/questions/27835801/how-to-auto-generate-migrations-with-sequelize-cli-from-sequelize-models)**ので、手動で変更してあげる必要がある。django みたいに model から migration が作成できたら良いんだけどな・・。


いざ、モデル作成。プロジェクトルートで以下のコマンドを叩く。

~~~bash
$ sequelize model:create --name User --attributes firstName:string,lastName:string

## models が以下のようになる
models
├── index.js
└── user.js
~~~

上記は、以下テーブルを作成することになる。

- Users
   - id (自動作成)
   - firstName
   - lastName
   - createdAt (自動作成)
   - updatedAt (自動作成)

なんで複数形？っていうのは、以下の通り。基本的に定義したモデルの複数形がテーブルになる。

> By default, sequelize will automatically transform all passed model names (first parameter of define) into plural.

id の自動作成は、明確に記載されてないんだけど（どこかにあるのかな）、以下が効いてると思われる。

> Sequelize will assume your table has a id primary key property by default.

createdAt, updatedAt の自動作成は通り。

> Sequelize will then automatically add the attributes createdAt and updatedAt.

自動作成を無効にしたい！！カスタマイズしたいということは、もちろん可能。以下の参考へ。

- 参考
    - [PrimaryKey](http://docs.sequelizejs.com/en/latest/docs/legacy/#primary-keys)
    - [Definition](http://docs.sequelizejs.com/en/latest/docs/models-definition)
    

### 3.2. models の index.js について

models dir には index.js が sequelize init によって作成されるが、これは models ディレクトリにある全ファイルを読み込んで、モデルを構築してくれる。その際に、モデル間のリレーションも一緒に構築する。利用する際には、index.js を読み込んで、該当モデルを利用することになる。(ES6 利用)

~~~javascript
// 利用例
const models = require('./models');

const User = models.User;
User.findAll().then((users) => console.log(users));
~~~

## 4. 試してみる

### 4.1. エントリポイントの作成

とりあえず色々試してみたいので、repl できるようにプロジェクトルートに **repl.js** とでもしてエントリポイントを作成する。

~~~javascript
// repl.js
const models = require('./models');

models.sequelize.sync({force:true}).then(() => {
  console.log("DEV DATA CREATED SUCCESSFULLY");
});

module.exports = models;
~~~

上記ですが、sync メソッドで、定義したモデルをDBに反映する。force:true オプションを与えることで、毎回 テーブルを Drop して Create してくれる。

**本番では migration を利用する**ことになると思いますが、開発環境や、とりあえず試めすくらいには 、スキームを変更するのが楽なので、force:true でも良い気がする。

- 参考
    - [sync](http://docs.sequelizejs.com/en/latest/api/sequelize/#sync)

### 4.2. Node REPLで試す

REPL を起動する。

~~~bash
$ node
~~~

以下を打ち込んで試してみる。

~~~javascript
// repl 内
const models = require('./repl');
const User = models.User;

// ユーザ作成
User.create({
    firstName: 'Yamada', 
    lastName: 'Tarou'
}).then((user) => console.log(JSON.stringify(user)));
// 出力
// > Executing (default): INSERT INTO `Users` (`id`,`firstName`,`lastName`,`updatedAt`,`createdAt`) VALUES (DEFAULT,'Yamada','Tarou','2015-12-30 03:10:38','2015-12-30 03:10:38');
// {"id":1,"firstName":"Yamada","lastName":"Tarou","updatedAt":"2015-12-30T03:10:38.000Z","createdAt":"2015-12-30T03:10:38.000Z"}

// ユーザ取得
User.findAll().then((users) => console.log(JSON.stringify(users)));
// 出力
// > Executing (default): SELECT `id`, `firstName`, `lastName`, `createdAt`, `updatedAt` FROM `Users` AS `User`;
// [{"id":1,"firstName":"Yamada","lastName":"Tarou","createdAt":"2015-12-30T03:10:38.000Z","updatedAt":"2015-12-30T03:10:38.000Z"}]

// ※出力を見やすくするため JSON.stringify しています
~~~

- その他メソッド
    - http://docs.sequelizejs.com/en/latest/api/model/


Project.create({name: "aho"});
let user = User.create({
    firstName: 'Yamada', 
    lastName: 'Tarou'
}).then((user) => user)
