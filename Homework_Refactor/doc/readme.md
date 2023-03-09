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



其中密码强度分为三档，使用正则判断：

- <font color="#ec6e2d">弱强度：</font>密码为六位及以上且含有数字/字母/符号
- <font color=#70c6fc>中强度：</font>密码为七位及以上且至少含有数字/小写字母/大写字母中的两种
- <font color=#4abf5d>高强度：</font>密码为八位及以上且必须含有数字、小写字母、大写字母、符号四种

## API 接口列表

https://api.liruinian.top

**POST   /user/login**

用户登录

~~~json
{"username":"手机号或邮箱","password":"密码","login_type":1或2}
~~~

其中login_type: 1为手机号登录 2为邮箱登录

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

用户登出

无发送值 返回setCookie header

**POST   /user/register**

用户注册

~~~json
{"phone":"手机号","email":"邮箱","username":"用户名","password":"密码"}
~~~

**POST   /user/info**

用户信息

**以下api在发送请求时必须携带header，格式如下**

~~~code
username:用户名

authorization:eyJhbGciOiJIUzI1NiIs...

adminauth:eyJhbGciOiJIUzI1NiIsInR5...
~~~

~~~json
{"code":200,"msg":"success","AForm":[{"id":1,"phone":"15765505517","email":"2941330150@qq.com","username":"tim_lrn2016","usertype":"admin"}]}
~~~

**POST   /article/list**

文章列表

~~~json
{"code":200,"msg":"success","Articles":[{"id":1,"title":"文章标题","category":"文章类别","content":"文章内容","author":"作者","time":"发送时间","views":"浏览量","href":"链接"}，{ ... }]}
~~~

**POST   /article/create**

创建文章

*必须为admin*

~~~~json
{"edit":false,"title":"标题","category":"类别","content":"内容","author":"作者","time":"时间","views":"浏览量","href":"链接"}
~~~~

**DELETE /article/delete/:id** 

删除文章

*必须为admin*

**POST   /article/edit**

修改文章

*必须为admin*

~~~~json
{"edit":true,"id":对应文章id,"title":"标题","category":"类别","content":"内容","author":"作者","time":"时间","views":"浏览量","href":"链接"}
~~~~

**POST   /article/:id** 

获取单个文章

返回值格式同 /article/list

**POST   /article/comments/:id** 

获取一个文章的所有评论

~~~~json
{"code":200,"msg":"success","Comments":[{"id":文章评论id,"articleId":对应文章id,"user":"用户名","commentText":"评论文本","time":"评论时间","thumbUp":点赞数},{...}]}
~~~~

**POST   /article/create-comment/:id**

对某一个文章创建评论

~~~~json
{"commentText":"评论文本","time":"评论时间"}
~~~~

**POST   /article/like-comment/:id** 

对一个评论点赞

**POST   /article/remove-comment/:id** 

*会判断是否为留言者或者为管理员*

删除一个评论

