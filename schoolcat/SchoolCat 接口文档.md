# ***SchoolCat***	接口文档

**注册（post）**

47.103.210.124:7000/user/register

```
客户端

{

  "email":"12w23qwer.com",

  "pwd":"1qaz2wsx"

}

服务器

{

  "msg": "注册成功",

  "user_id": 2

}
{
	"msg":"邮箱已被注册"
}
```

登录（post）

47.103.210.124:7000/user/login

```
客户端

{

  "email":"12w23qwer.com",

  "pwd":"1qaz2wsx"

}

服务端

{

  "iconsrc": "null",

  "msg": "欢迎使用",

  "userid": 2,

  "username": ""

}
{
	"msg": "欢迎使用,admin",
	"userid":user0.ID,
	"auth":"12345",//管理员地任何操作都要在header里面加上这个键值对
	"username":user0.Username,
	"iconsrc":user0.IconSrc,
}
{
	"msg": "用户名或者密码有误"
}
```

添加和更改信息(post)

47.103.210.124:7000/user/info

```
客户端

{

  "ID":2,

  "username":"qwert",

  "gender":"",

  "school":"ncu",

  "resume":"个人简介",

  "icon_src":""

}

服务器

{

  "iconsrc": "",

  "msg": "操作成功",

  "username": "qwert"

}
```

用户添加share(post)

47.103.210.124:7000/user/newShare

```
客户端

{

  "user_id":2,

  "username":"qwert",

  "icon_src":"",

  "address":"信工楼",

  "content":"heloo there is a cat",

  "share_star":0,

  "share_images":[

​    {"src":"adcsdsd.com"},

​    {"src":"dwfefd.cn"},

​    {"src":"ufhdjdj"}

  ]

}

服务器

{

  "ShareID": 1,

  "msg": "分享成功"

}
```

用户删除share(delete)

47.103.210.124:7000/user/deleteShare

```
客户端(在header里面加)

"user_id":6,

"share_id":5

服务器

{

  "msg": "无权删除"

}

{

  "msg": "删除成功"

}
```



用户添加share评论(post)

47.103.210.124:7000/user/newShareComment

```
客户端

{

  "user_id":3,

  "username":"qwert123",

  "icon_src":"bhsdchjdq.sads",

  "share_id":3,

  "comment_star":0,

  "comment":"cbdsyugysdgvhsdgdvjhwvhb"

  }

服务器

{

  "msg": "评论成功",

  "user_comment": 3

}
```

用户删除share评论(delete)

47.103.210.124:7000/user/deleteShareComment

```
客户端(在header里面加)

"user_id":6,

"comment_id":5

服务器

{

  "msg": "无权删除"

}

{

  "msg": "删除成功"

}
```

用户搜索share(get)

47.103.210.124:7000/user/search?keywords=

```
客户端

服务器

{

  "shares": [

​    {

​      "ID": 2,

​      "CreatedAt": "2021-12-11T11:24:32.651+08:00",

​      "UpdatedAt": "2021-12-11T11:24:32.651+08:00",

​      "DeletedAt": **null**,

​      "user_id": 2,

​      "username": "qwert",

​      "icon_src": "",

​      "address": "信工ewesfd楼",

​      "content": "heloosddscdssdfd there is a cat",

​      "share_star": 0,

​      "share_images": [

​        {

​          "ID": 4,

​          "CreatedAt": "2021-12-11T11:24:32.653+08:00",

​          "UpdatedAt": "2021-12-11T11:24:32.653+08:00",

​          "DeletedAt": **null**,

​          "share_id": 2,

​          "src": "adcsdsd.com"

​        },

​        {

​          "ID": 5,

​          "CreatedAt": "2021-12-11T11:24:32.653+08:00",

​          "UpdatedAt": "2021-12-11T11:24:32.653+08:00",

​          "DeletedAt": **null**,

​          "share_id": 2,

​          "src": "dwfefd.cn"

​        },

​        {

​          "ID": 6,

​          "CreatedAt": "2021-12-11T11:24:32.653+08:00",

​          "UpdatedAt": "2021-12-11T11:24:32.653+08:00",

​          "DeletedAt": **null**,

​          "share_id": 2,

​          "src": "ufhsddscddsfsdvdjdj"

​        }

​      ],

​      "user_comment": [

​        {

​          "ID": 2,

​          "CreatedAt": "2021-12-11T11:42:27.193+08:00",

​          "UpdatedAt": "2021-12-11T11:42:27.193+08:00",

​          "DeletedAt": **null**,

​          "username": "qwert123",

​          "icon_src": "bhsdchjdq.sads",

​          "user_id": 3,

​          "share_id": 2,

​          "comment_star": 0,

​          "comment": "cbdsyugyfdsvssdgvhsdgdvjhwvhb"

​        }

​      ]

​    }

  ]

}


```

用户请求share(get)

47.103.210.124:7000/user/viewShare

客户端

服务器（和搜索格式一样）



用户请求自己的share(get)

47.103.210.124:7000/selfShare

客户端（在header里加）

"user_id":5

服务器（和搜索格式一样）



用户点赞shareComment(put)

47.103.210.124:7000/user/commentLike?comment_id=13&like=yes(no)

客户端//yes点赞，no取消点赞

服务器

{

  "shares": "ok"

}



用户点赞share(put)

47.103.210.124:7000/user/commentLike?share_id=13&like=yes(no)

客户端//yes点赞，no取消点赞

服务器

{

  "shares": "ok"

}





用户删除share(delete)

47.103.210.124:7000/user/deleteShare

客户端

服务器















用户删除share(delete)

47.103.210.124:7000/user/deleteShare

客户端

服务器







用户删除share(delete)

47.103.210.124:7000/user/deleteShare

客户端

服务器





用户删除share(delete)

47.103.210.124:7000/user/deleteShare

客户端

服务器



用户删除share(delete)

47.103.210.124:7000/user/deleteShare

客户端

服务器
