+++
Categories = []
Description = ""
Tags = []
comments = true
date = "2015-11-24T21:01:50+09:00"
title = "sequelize_belongs_to"
draft = true

+++

## model

~~~javascript
// models/user.js
export default function createModel(sequelize, DataTypes) {
  const User = sequelize.define("User", {
    username        : { type: DataTypes.STRING, allowNull: false },
    password        : { type: DataTypes.STRING, allowNull: false },
  }, {
    classMethods: {
      associate: (models) => {
        User.belongsTo(models.Group);
      }
    }
  });
  
  return User;
}
~~~

~~~javascript
// models/group.js
export default function createModel(sequelize, DataTypes) {
  const Group = sequelize.define("Group", {
    name        : { type: DataTypes.STRING, allowNull: false },
  });
  
  return Group;
}
~~~

## query

~~~javascript
// controller/user.js
export function create(req, res){
    //transaction つけるのがベター
    User.create(req.body).then((user) => {
        return user.setGroup(1);
    }).then((user) => {
        res.json(user);
    }).catch((err) => {
        return res.status(400).send(err);
    });
}
~~~

~~~bash
Executing (default): INSERT INTO "Users" ("id","username","password","updatedAt","createdAt") VALUES (DEFAULT,'yamada','$2a$10$23JSAwZv1DD29Nl.UlhMaOq3AT120fH41CKNyBuuurNrnt3b3uYXW','dasfdfas','fdsafdsa','2015-11-24 12:12:44.162 +00:00','2015-11-24 12:12:44.162 +00:00') RETURNING *;
Executing (default): UPDATE "Users" SET "UserStatusId"=1,"updatedAt"='2015-11-24 12:12:44.167 +00:00' WHERE "id" = 3
~~~



