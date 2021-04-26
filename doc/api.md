# GPL 接口文档

- [GPL 接口文档](#gpl-接口文档)
  - [请求地址](#请求地址)
  - [健康检测](#健康检测)
  - [登录](#登录)
  - [注销](#注销)
  - [域名管理](#域名管理)
    - [添加域名](#添加域名)
    - [获取域名列表](#获取域名列表)

## 请求地址

* 生产环境 ``http://``
* 测试环境 ``http://``
* 开发环境 ``http://127.0.0.1:8080``

## 健康检测

* URL ``/ok``
* Method ``HEAD``
* Response ``ok``

## 登录

* URL ``/login``
* Method ``POST``
* Params

| Filed    | Type   | Require | Comment |
| -------- | ------ | ------- | ------- |
| username | string | Y       | 用户名 |
| password | string | Y       | 密码 |

* Response

```json
{
    "data": null,
    "meta": {
        "msg": "登录成功",
        "status": 200,
        "token": "4881c9cc2915caa59c1921513acec601"
    }
}
```

## 注销

* URL ``/logout``
* Method ``POST``
* Headers ``Authorization: token_string``
* Response

```json
{
    "data": null,
    "meta": {
        "msg": "注销成功",
        "status": 200
    }
}
```

## 域名管理

### 添加域名

* URL ``/domains``
* Method ``POST``
* Headers ``Authorization: token_string``
* Params

| Filed    | Type   | Require | Comment |
| -------- | ------ | ------- | ------- |
| name     | string | Y       | 域名 |

* Response

```json
{
    "data": null,
    "meta": {
        "msg": "添加域名成功",
        "status": 201
    }
}
```

### 获取域名列表

* URL ``/domains``
* Method ``GET``
* Headers ``Authorization: token_string``
* Params

| Filed    | Type   | Require | Comment |
| -------- | ------ | ------- | ------- |
| name     | string | N       | 域名 |
| page     | int    | N       | 页码，default 1 |
| pagesize | int    | N       | 限制，default 10|
| order    | string | N       | 排序，default updated_at desc |

* Response

```json
{
    "data": {
        "total": 0,
        "domain_list": []
    },
    "meta": {
        "msg": "获取域名列表成功",
        "status": 200
    }
}
```