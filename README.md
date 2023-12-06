# HackWeek

# 接口说明
## 1.1：注册

- 请求路径：api/v1/user/register
- 请求方法：POST
- 请求体

|   参数    |           含义           |              备注              |
| :-------: | :----------------------: | :----------------------------: |
| user_name |         用户姓名         | 只能包含汉字、大小写字母、数字 |
| password  | 经前端加密过后的用户密码 |  经前端加密，后端存储时也加密  |

```
{
     "user_name":"lxy",
     "password":"123456"
}
```



- 响应体

|  参数  |         含义         | 备注 |
| :----: | :------------------: | ---- |
| status | 后端响应返回的状态码 |      |
|  data  |     数据或者信息     |      |
|  msg   |       响应信息       |      |
| error  |     详细错误信息     |      |

```
{
    "status": 200,
    "data": "操作成功",
    "msg": "操作成功",
    "error": ""
}
```



## 1.2：登录

- 请求路径：api/v1/user/login
- 请求方法：POST
- 请求体

|   参数    |   含义   |              备注              |
| :-------: | :------: | :----------------------------: |
| user_name | 用户姓名 | 只能包含汉字、大小写字母、数字 |
| password  | 用户密码 |  经前端加密，后端存储时也加密  |

```
{
     "user_name":"lxy",
     "password":"123456"
}
```



- 响应体

|    参数     |     含义     |     备注     |
| :---------: |:----------:| :----------: |
|   status    | 后端响应返回的状态码 |              |
|    data     |    用户数据    |              |
| data->user  |    用户信息    |              |
| data->token |  认证token   | 有效期为一天 |
|     msg     |    响应信息    |              |
|    error    |   详细错误信息   |              |

```
{
    "status": 200,
    "data": {
        "user": {
            "id": 2,
            "user_name": "lxy",
            "create_at": 1701557314
        },
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwidXNlcm5hbWUiOiJseHkiLCJleHAiOjE3MDE2NDM3MjUsImlzcyI6InRvLWRvLWxpc3QifQ.XOQFVHIcdttR0BpQFANUcbLmgstmZhQh5BM6hwm2Qek"
    },
    "msg": "操作成功",
    "error": ""
}
```

## 2.1：更改密码

- 请求路径：api/v1/user/updatePwd
- 请求方法：POST
- 请求体(需要token认证)

|    参数     | 含义  |         备注         |
|:---------:|:---:| :------------------: |
| updatePwd | 新密码 |  |

```
{
    "originPwd":"1234567",
    "updatePwd":"123456"
}
```



- 响应体

|  参数  |         含义         | 备注 |
| :----: | :------------------: | :--: |
| status | 后端响应返回的状态码 |      |
|  data  |     数据或者信息     |      |
|  msg   |       响应信息       |      |
| error  |     详细错误信息     |      |

```
{
    "status": 200,
    "data": "操作成功",
    "msg": "操作成功",
    "error": ""
}
```


## 2.2：更改用户个人信息

- 请求路径：api/v1/user/updateInfo
- 请求方法：POST
- 请求体(需要token认证)

|   参数   | 含义  |         备注         |
| :------: |:---:| :------------------: |
| update_name | 用户名字信息 | 后续可能有其他的信息 |

```
{
    "update_name":"test11"
}
```



- 响应体

|  参数  |         含义         | 备注 |
| :----: | :------------------: | :--: |
| status | 后端响应返回的状态码 |      |
|  data  |     数据或者信息     |      |
|  msg   |       响应信息       |      |
| error  |     详细错误信息     |      |

```
{
    "status": 200,
    "data": "操作成功",
    "msg": "操作成功",
    "error": ""
}
```

## 2.3：得到登录用户的详细信息

- 请求路径：api/v1/user/info
- 请求方法：Get
- 请求体(需要token认证)：无


- 响应体

|  参数  |     含义     |    备注     |
| :----: |:----------:|:---------:|
| status | 后端响应返回的状态码 |           |
|  data  |   登录用户数据   | 后续可能会添加信息 |
|  msg   |    响应信息    |           |
| error  |   详细错误信息   |           |

```
{
    "status": 200,
    "data": {
        "id": 15,
        "user_name": "test11",
        "create_at": 1701754717
    },
    "msg": "操作成功",
    "error": ""
}
```

## 3.1：获取星火大模型V3.0生成的故事
- 请求路径：api/v1/user/info
- 请求方法：Post
- 请求体

|  参数  |     含义     |    备注     |
| :----: |:----------:|:---------:|
| mood| 故事的心情 | |
| keywords| 故事的关键词 | 关键词之间用+连接 |
```
e.g.
{
    "mood" : "开心"
    "keywords":"室友+电脑"
}
```
- 响应体

|  参数  |     含义     |    备注     |
| :----: |:----------:|:---------:|
| story| 根据关键词和心情生成的故事(200字左右) | |
```
e.g.
{
    "story":"我的室友是个电脑迷，他的生活几乎都离不开电脑。有一天，他在电脑上发现了一款名叫\"开心农场\"的游戏，从此就迷上了它。每天他都会坐在电脑前，种菜、养鸡、收果子，忙得不亦乐乎。他说，玩这个游戏让他感到了前所未有的开心。我看着他那专注的样子，忍不住也笑了出来。虽然他的行为有些古怪，但我知道，这就是他找到的快乐。所以，我也决定支持他，一起在这个虚拟的农场里开心地玩耍。"
}
```