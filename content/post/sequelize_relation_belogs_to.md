+++
Categories = ["JavaScript"]
Description = "Sequelize でテーブル間のリレーション(belongs to)を作っていきたいと思う"
Tags = ["Sequelize", "ORM"]
comments = true
date = "2015-12-31T14:56:15+09:00"
title = "Sequelize Part 3 - リレーション belongs to"
draft = true

+++


<!--more-->

## 1. 設計

これをつくる！！

{{< figure class="image-half__center" src="UsersProjectsTable.jpg" title="ER図" >}}

## 2. モデル作成

### 2.1. 初期モデルの作成

~~~bash
$ sequelize model:create --name User --attributes firstName:string,lastName:string
$ sequelize model:create --name Project --attributes name:string
~~~

### 2.2. リレーション定義

モデルのリレーションは、**models/対象ファイル**の **classMethods 内の associate** に定義してあげる。リレーションの種類は以下の通り。

~~~javascript
// 変更前
// モデルファイルのリレーション定義箇所
    classMethods: {
      associate: function(models) {
        // associations can be defined here
      }
    }
    
// 変更後
// models/user.js
    classMethods: {
      associate: function(models) {
        User.belongsTo(models.Project);
      }
    }
    
// models/project.js
    classMethods: {
      associate: function(models) {
        Project.hasMany(models.User);
      }
    }
~~~


## 3. インスタンスを作成する

REPL を起動する。

~~~bash
$ node
~~~

以下を打ち込んで試してみる。

~~~javascript
// repl 内 共通部分
const models = require('./repl');
const User = models.User;
const Project = models.Project;
~~~

以降の 3.1, 3.2 で出力結果が異なることに注意

### 3.1. リレーションも一緒に作成 1

ユーザー作成、プロジェクト作成、リレーション付与を１つで行う。

~~~javascript
// repl 内
User.create({
    firstName: 'Suzuki', 
    lastName: 'Tarou',
    Project: {
        name: 'second project'
    }
}, {
    include: [ Project ]
}).then((user) => {
    console.log(JSON.stringify(user));
});
// 出力
// > Executing (default): INSERT INTO `Projects` (`id`,`name`,`updatedAt`,`createdAt`) VALUES (DEFAULT,'second project','2015-12-31 07:24:50','2015-12-31 07:24:50');
// Executing (default): INSERT INTO `Users` (`id`,`firstName`,`lastName`,`ProjectId`,`updatedAt`,`createdAt`) VALUES (DEFAULT,'Suzuki','Tarou',2,'2015-12-31 07:24:50','2015-12-31 07:24:50');
// {"id":2,"firstName":"Suzuki","lastName":"Tarou","Project":{"id":2,"name":"second project","updatedAt":"2015-12-31T07:24:50.000Z","createdAt":"2015-12-31T07:24:50.000Z"},"ProjectId":2,"updatedAt":"2015-12-31T07:24:50.000Z","createdAt":"2015-12-31T07:24:50.000Z"}
~~~

- 参考
    - [Creating with associations](http://docs.sequelizejs.com/en/latest/docs/associations/#creating-elements-of-a-belongsto-or-hasone-association)

### 3.2. リレーションも一緒に作成 2

ユーザー作成の後、プロジェクト作成、リレーション付与を行う。

~~~javascript
// ユーザ作成とリレーション付与
User.create({
    firstName: 'Yamada', 
    lastName: 'Tarou'
}).then((user) => {
    return user.createProject({
        name: 'second project'
    })
}).then((user) => {
    console.log(JSON.stringify(user));
});

// > Executing (default): INSERT INTO `Users` (`id`,`firstName`,`lastName`,`updatedAt`,`createdAt`) VALUES (DEFAULT,'Yamada','Tarou','2015-12-31 07:54:07','2015-12-31 07:54:07');
// Executing (default): INSERT INTO `Projects` (`id`,`name`,`updatedAt`,`createdAt`) VALUES (DEFAULT,'second project','2015-12-31 07:54:07','2015-12-31 07:54:07');
// Executing (default): UPDATE `Users` SET `ProjectId`=3,`updatedAt`='2015-12-31 07:54:07' WHERE `id` = 3
// {"id":3,"firstName":"Yamada","lastName":"Tarou","updatedAt":"2015-12-31T07:54:07.000Z","createdAt":"2015-12-31T07:54:07.000Z","ProjectId":3}
~~~

- 参考
    - [createAssociation](http://docs.sequelizejs.com/en/latest/api/associations/belongs-to/#createassociationvalues-options-promise)


### 3.3. 別々に作成後、リーレション付与

インスタンスを作成した後、リレーションを付与してあげる。

~~~javascript
// repl 内
// プロジェクト作成
Project.create({
    name: 'first project'
}).then((project) => console.log(JSON.stringify(project)));
// 出力
// > Executing (default): INSERT INTO `Projects` (`id`,`name`,`updatedAt`,`createdAt`) VALUES (DEFAULT,'first project','2015-12-31 06:18:49','2015-12-31 06:18:49');
// {"id":1,"name":"first project","updatedAt":"2015-12-31T06:18:49.000Z","createdAt":"2015-12-31T06:18:49.000Z"}

// ユーザ作成とリレーション付与
User.create({
    firstName: 'Yamada', 
    lastName: 'Tarou'
}).then((user) => {
    return user.setProject(1)
}).then((user) => {
    console.log(JSON.stringify(user));
});
// 出力
// > Executing (default): INSERT INTO `Users` (`id`,`firstName`,`lastName`,`updatedAt`,`createdAt`) VALUES (DEFAULT,'Yamada','Tarou','2015-12-31 07:07:30','2015-12-31 07:07:30');
// Executing (default): UPDATE `Users` SET `ProjectId`=1,`updatedAt`='2015-12-31 07:07:30' WHERE `id` = 1
// {"id":1,"firstName":"Yamada","lastName":"Tarou","updatedAt":"2015-12-31T07:07:30.000Z","createdAt":"2015-12-31T07:07:30.000Z","ProjectId":1}
~~~

- 参考
    - [setAssociation](http://docs.sequelizejs.com/en/latest/api/associations/belongs-to/#setassociationnewassociation-options-promise)






