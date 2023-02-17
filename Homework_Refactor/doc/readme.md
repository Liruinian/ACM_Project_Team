# 大大作业-介绍文档

电气 2204 李睿年

2.4 技术验证 2.10 重构 2.16 开发完毕

## 项目简介

### 前端

原生 html js css  前后端分离

- Linux + Nginx
- 使用 XHR 与后端进行数据交互

- 使用editor.md编辑文章 支持markdown解析



### 后端

go(gin框架) MySQL

-  https://api.liruinian.top
- 服务器后端在 8880端口建立
- CORS跨域、基于JWT的权限控制和TLS（HTTPS）支持
- 与MySQL数据库对接的账户登录注册
- 文章数据库增删改查操作



部署网址： https://acm.liruinian.top/final

## 主要功能

- 登录 注册 账号管理

- 文章的增加 删除 修改 查询

- 自动给文章编号并排序

- 全页面均含有自适应设计，自动调节元素宽高以适应设备和浏览器尺寸

