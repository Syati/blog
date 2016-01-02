+++
Categories = ["JavaScript"]
Description = "Sequelize のリレーションを学ぶ。belongsTo, hasOne, hasMany, belongsToMany の4つだけ。後、オプションを少し学べばやりたいことは大方出来るようになる。"
Tags = ["Sequelize", "ORM"]
comments = true
date = "2016-01-02T22:00:15+09:00"
title = "Sequelize Part 2 - リレーションについて"

+++
    
[Sequelize Part 1]({{< ref "post/sequelize_getting_started.md" >}}) で、なんとなくわかってきた！！と思うので、今回は、リレーションについて学ぶ。後々、WEBアプリケーションフレームワークから利用する際、ここら辺の理解は必須ですからね。

<!--more-->

## 1. リレーションの種類

リレーションを理解するためには、[Relations / Associations](http://docs.sequelizejs.com/en/latest/docs/associations/) を読み込まないと始まらない。結構ボリュームがあって、心折れそうになりますが！！実は大したことはないはず・・・。まずは、少しだけ用語の説明。

> When calling a method such as User.hasOne(Project), we say that the User model (the model that the function is being invoked on) is the source and the Project model (the model being passed as an argument) is the target.

単なる呼び方ですね。先に出てくるモデルが **source** 。後に出てくるモデルが **target** ですよってことですね。

~~~javascript
// 以下の場合
User.hasOne(Project)
// User モデルが source
// Project モデルが target
~~~

これは、外部キーがどちらに追加されるのかを明記するために利用する。リレーションは、以降の4種類（メソッド）ある。

### 1.1. belongsTo 

**プレイヤ belongs to チーム** という場合。

~~~javascript
var Player = this.sequelize.define('Player', {/* ...attributes */})
  , Team  = this.sequelize.define('Team', {/* ...attributes */});

Player.belongsTo(Team);
~~~

上記は、以下のテーブルになる。

- Players
    - id
    - ...attributes
    - createdAt
    - updatedAt
    - **TeamId (belongsTo によって追加される)**
- Teams
    - id
    - ...attributes
    - createdAt
    - updatedAt

[belongsTo](http://docs.sequelizejs.com/en/latest/docs/associations/#belongsto) では **source (Player)** に外部キー（**TeamId**）が追加される。また、Player インスタンスは、[belongsTo - api](http://docs.sequelizejs.com/en/latest/api/associations/belongs-to/) メソッドを持つ。

**※外部キーが camelCase** ！？という方は、[2.2. underscored: 外部キーの命名規則を camelCase から underscored にする](#2-2-underscored-外部キーの命名規則を-camelcase-から-underscored-にする:4ab89ffec193a533a37a4e4a7c6ae107) を参照のこと。


### 1.2. hasOne

**ユーザ has one プロフィール** という場合。

~~~javascript
var User = sequelize.define('User', {/* ... */})
var Profile = sequelize.define('Profile', {/* ... */})
 
User.hasOne(Profile)
~~~

上記は、以下のテーブルになる。

- Users
    - id
    - ...
    - createdAt
    - updatedAt
- Profiles
    - id
    - ...
    - createdAt
    - updatedAt
    - **UserId (hasOne によって追加される)**
    
[hasOne]((http://docs.sequelizejs.com/en/latest/docs/associations/#hasone) では **target (Profile)** に外部キー（**UserId**）が追加される。**belongsTo とその点が違うことに注意**。また、User インスタンスは [hasOne - api](http://docs.sequelizejs.com/en/latest/api/associations/has-one/) メソッドを持つ。

### 1.3. hasMany

**プロジェクト has many ユーザ** という場合。

~~~javascript
var Project = sequelize.define('Project', {/* ... */})
var User = sequelize.define('User', {/* ... */})
 
Project.hasMany(User)
~~~

上記では、以下のテーブルになる。

- Projects
    - id
    - ...
    - createdAt
    - updatedAt
- Users
    - id
    - ...
    - createdAt
    - updatedAt
    - **ProjectId (hasMany によって追加される)**

[hasMany](http://docs.sequelizejs.com/en/latest/docs/associations/#one-to-many-associations) では **target (User)** に外部キー（**ProjectId**）が追加される。これだけ見ると **hasOne と同じじゃないか**と思うかもしれないが、インスタンスが持つメソッドが異なる。Project インスタンスは [hasMany - api](http://docs.sequelizejs.com/en/latest/api/associations/has-many/) メソッドを持つ。

### 1.4. belongsToMany

**プロジェクトは複数のユーザを持ち、ユーザは複数のプロジェクトに属する** という **多対多** の場合。

~~~javascript
var Project = sequelize.define('Project', {/* ... */})
var User = sequelize.define('User', {/* ... */})

Project.belongsToMany(User, {through: 'UserProjects'});
User.belongsToMany(Project, {through: 'UserProjects'});
~~~

上記は、以下のテーブルになる。

- Projects
    - id
    - ...
    - createdAt
    - updatedAt
- Users
    - id
    - ...
    - createdAt
    - updatedAt
- **UserProjects   (through オプションで指定した名前のテーブル)**
    - **UserId    (belongsToMany によって追加される)**
    - **ProjectId (belongsToMany によって追加される)**
    - createdAt
    - updatedAt


[belongsToMany](http://docs.sequelizejs.com/en/latest/docs/associations/#belongs-to-many-associations) では、**through で指定した名前のテーブル**に双方のキー（ProjectId, UserId）が追加される。また、それぞれのインスタンスに以下のとおり(または [belongsToMany - api](http://docs.sequelizejs.com/en/latest/api/associations/belongs-to-many/)) メソッドが追加される。

> This will add methods **getUsers, setUsers, addUser,addUsers to Project**, and **getProjects, setProjects and addProject, addProjects to User**.

上記の UserProjects （中間テーブル）を見ると分かると思いますが **id** はありません。理由は以下のとおり。UserId, ProjectId の複合キーで十分なため、別にプライマリキー設ける必要ないでしょって話かな。

> By default the code above will add ProjectId and UserId to the UserProjects table, and remove any previously defined primary key attribute - the table will be uniquely identified by the combination of the keys of the two tables, and there is no reason to have other PK columns.

ただし UserProjects モデルを以下のように定義することで、**id** を追加することや、**status** カラムを中間テーブルに追加することができる。

~~~javascript
var Project = sequelize.define('Project', {/* ... */})
var User = sequelize.define('User', {/* ... */})

var UserProjects = sequelize.define('UserProjects', {
    id: {
      type: DataTypes.INTEGER,
      primaryKey: true,
      autoIncrement: true
    },
    status: DataTypes.STRING
})

User.belongsToMany(Project, { through: UserProjects })
Project.belongsToMany(User, { through: UserProjects })
~~~

## 2. その他、オプション

### 2.1. 片方向と双方向

以下の違いを覚えておく。

- 片方向（one way）
    - belongsTo
    - hasOne
    - hasMany
- 双方向 (two way)
    - belongsToMany

片方向とは、言葉の通りだが、一方からしか取得できないということである。例えば、**プレイヤ belongs to チーム** という場合で考えると、**プレイヤからチームを取得することは可能**だが、**チームからプレイヤを取得することは不可能**ということである。以下コードサンプル。

~~~javascript
var Player = this.sequelize.define('Player', {/* ...attributes */})
  , Team  = this.sequelize.define('Team', {/* ...attributes */});

Player.belongsTo(Team);

...

// インスタンス利用
// plyaer インスタンスから team を取得することは可能だが
player.getTeam()
// 以下のように team から プレイヤを取得することはできない
// team.getPlayers()
~~~

なので、**片方向のメソッドを利用して、双方向から取得できるようにする**には、以下のようにモデルを定義してあげる必要がある。

~~~javascript
var Player = this.sequelize.define('Player', {/* ...attributes */})
  , Team  = this.sequelize.define('Team', {/* ...attributes */});

Player.belongsTo(Team); // belongsTo を利用
Team.hasMany(Player);   // hasMany を利用
...

// インスタンス利用
// plyaer インスタンスから team を取得することは可能かつ
player.getTeam()
// team から players を取得することが可能。
team.getPlayers()
~~~

双方向については、上記のようなことをする必要はない。双方向を利用した際に、インスタンスにメソッドが追加されるため。[1.4. belongsToMany](http://localhost:1313/post/sequelize_relation/#1-4-belongstomany:4ab89ffec193a533a37a4e4a7c6ae107) 参照のこと。

### 2.2. underscored: 外部キーの命名規則を camelCase から underscored にする

**Foreign keys** は camelCase になります。

> The default casing is camelCase however if the source model is configured with underscored: true the foreignKey will be snake_case.

もし、**underscored** のほうがいいよねって場合は、以下のようにしてあげる。
    
~~~javascript
var Player = this.sequelize.define('Player', {
    /* ... */
}, {
    underscored: true
});

var Team  = this.sequelize.define('Team', {/* ... */});

Player.belongsTo(Team);
~~~

上記は、以下のテーブルになる。

- Players
    - id
    - ...
    - createdAt
    - updatedAt
    - **team_id (camelCase ではなく _ になる)**
- Teams
    - id
    - ...
    - createdAt
    - updatedAt

### 2.3. as: target のモデル名を任意に変更する

**as** オプションに任意の名前を与えることで、target のモデル名を変更できる

~~~javascript
var User = this.sequelize.define('User', {/* ... */})
  , Project = sequelize.define('Project', {/* ... */})

User.belongsTo(Project, {as: 'Prj'});
~~~

上記は、以下のテーブルになる。

- Users
    - id
    - ...
    - createdAt
    - updatedAt
    - **PrjId (ProjectId ではなく as で指定した名前Id になる)**
- Projects
    - id
    - ...
    - createdAt
    - updatedAt

**※ここで注意する必要があるのは** リレーション(belongsToなど)を追加したときにインスタンスに追加されるメソッドである。今回の場合は User インスタンスのもつメソッドが以下になる。

~~~javascript
// as なしの場合、以下のメソッドだが（デフォルト）
user.getProject();

// as ありの場合、以下のメソッド(指定した名前)になる
user.getPrj()
~~~

### 2.4. foreignKey: 外部キーの名前を任意に変更する

**foreignKey** オプションに任意の名前を与えることで、Sequelize の外部キーの命名規則を完全無視できる

~~~javascript
var User = this.sequelize.define('User', {/* ... */})
  , Project = sequelize.define('Project', {/* ... */})

User.belongsTo(Project, {foreignKey: 'fk_project'});
~~~

上記は、以下のテーブルになる。

- Users
    - id
    - ...
    - createdAt
    - updatedAt
    - **fk_project (ProjectId ではなく foreignKey で指定した名前になる)**
- Projects
    - id
    - ...
    - createdAt
    - updatedAt

**※ここで注意する必要があるのは** as で指定した場合とは異なり、インスタンスのもつメソッドは変わらないということ。

~~~javascript
// foreginKey あり・なしに関わらず、以下のメソッド（デフォルト）
user.getProject();
~~~

### 2.5. targetKey: 外部キーをプライマリキーから任意のカラムに変更する

以下にある通り、デフォルトの外部キーは target (または source)のプライマリキーになります。belongsTo としか書いてないけど、それ以外も基本同じ。

> By default the target key for a belongsTo relation will be the target primary key. To override this behavior, use the targetKey option.

例えば、Project の name カラムを外部キーにする場合。

~~~javascript
var User = this.sequelize.define('User', {/* ... */});
var Project = sequelize.define('Project', {
    name: { type: DataTypes.STRING, unique: true}
});

User.belongsTo(models.Project, {foreignKey: 'fk_projectname', targetKey: 'name'});
~~~

上記は、以下のテーブルになる。

- Users
    - id
    - ...
    - createdAt
    - updatedAt
    - fk_projectname
- Projects
    - id
    - name
    - createdAt
    - updatedAt
    
インスタンスのもつメソッドは [2.4. foreignKey: 外部キーの名前を任意に変更する](#2-4-foreignkey-外部キーの名前を任意に変更する:4ab89ffec193a533a37a4e4a7c6ae107) と変わりなし

