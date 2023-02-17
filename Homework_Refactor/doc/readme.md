# 大大作业-介绍文档

电气 2204 李睿年

2.4 技术验证 2.10 重构 2.16 开发完毕

## 项目简介

### 前端

原生 html js css  前后端分离

https://acm.liruinian.top/articles

- Linux + Nginx
- 使用 XHR 与后端进行数据交互
- 使用editor.md编辑文章 支持markdown解析
- 其他特性请查阅前端大作业文档



### 后端

go(gin框架) MySQL

https://api.liruinian.top

- 服务器后端在 8880端口建立
- CORS跨域、基于JWT的权限控制和TLS（HTTPS）支持
- 与MySQL数据库对接的账户登录注册
- 文章数据库增删改查操作



部署网址： https://acm.liruinian.top/final

## 主要功能

- 登录 注册 登录态token鉴权
- 文章的增加 删除 修改 查询
- 对文章的评论 对评论内容的点赞
- 自动给文章编号并倒序显示最新文章
- 全页面均含有自适应设计，自动调节元素宽高以适应设备和浏览器尺寸
- 对手机等竖屏设备有针对性优化



## API 接口列表

https://api.liruinian.top

**POST   /user/login**

发送：(raw)

~~~json
{"username":"手机号或邮箱","password":"密码","login_type":1或2}
~~~

其中login_type: 1为手机号登录 2为邮箱登录

响应：(json)

~~~json
{"code":事件码,"msg":"消息文本"}
~~~

^ 之后的此类响应格式将不再赘述，如有特殊格式将会说明

2001：登录失败：输入格式不正确

2002：登录失败：用户名或密码不能为空

2003：登录失败：请检查用户名或密码是否正确

2004：登录失败：存在多个对应账号，请尝试使用其他方式登录

成功响应：

~~~json
{
    "code": 2000,
    "data": {
        "admin_token": "eyJhbGciOiJIUzI1N..." 或 "NOT ADMIN",
        "login_token": "eyJhbGciOiJIUzI1N...",
        "username": "用户名"
    },
    "msg": "success"
}
~~~

ps: admin_token ≠ login_token

**POST   /user/logout**

发送：无

接收：(json code msg响应)

2021：登出失败：可能登录已过期

2020：success

**POST   /user/register**

2011：注册失败：输入格式不正确

2012：注册失败：注册部分不能有空值

2013：注册失败：用户已存在 请尝试更改邮箱、手机号或用户名

2010：success

**POST   /user/info**

~~~json
{"code":200,"msg":"success","AForm":[{"id":1,"phone":"15765505517","email":"2941330150@qq.com","username":"tim_lrn2016","usertype":"admin"}]}
~~~



**POST   /user/edit-info**

**POST   /article/list**

**POST   /article/create**

**DELETE /article/delete/:id** 

**POST   /article/edit**

**POST   /article/:id** 

**POST   /article/comments/:id** 

**POST   /article/create-comment/:id**

**POST   /article/like-comment/:id** 

