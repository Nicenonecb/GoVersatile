# GoVersatile 

GoVersatile  是一个用于学习 Go 语言并应用于多个场景的仓库，例如商城、CMS管理系统以及博客等。本项目是由两位合作者共同开发，旨在通过实践项目来提高对 Go 语言的理解和应用能力。

## 目录

- [项目介绍](#项目介绍)
- [技术栈](#技术栈)
- [如何开始](#如何开始)
- [贡献指南](#贡献指南)
- [联系我们](#联系我们)

## 项目介绍

在这里，您可以添加关于项目的详细描述。例如，您可以描述项目的背景、目标以及已经实现的功能。同时，您也可以介绍项目中所用到的技术、实现原理以及项目的架构。

## 技术栈

- 编程语言：Go
- Web框架：Gin
- 数据库：PostgreSQL
- ORM库：GORM
- ... （其他后续补充）


## 代码提交格式：Type(scope):subjects
1. type（必须） : commit 的类别，只允许使用下面几个标识：
  - feat : 新功能
  - fix : 修复bug
  - docs : 文档改变
  - style : 代码格式改变
  - refactor : 某个已有功能重构
  - perf : 性能优化
  - test : 增加测试
  - build : 改变了build工具 如 grunt换成了 npm
  - revert : 撤销上一次的 commit
  - chore : 构建过程或辅助工具的变动
2.  scope（可选） : 用于说明 commit 影响的范围，比如数据层、控制层、视图层等等，视项目不同而不同。
3. subject（必须） : commit 的简短描述，不超过50个字符。

## 如何开始

### 环境要求

- Go 1.x
- PostgreSQL 
- 其他相关依赖库

### 安装与运行

1. 克隆仓库到本地

```bash
git clone https://github.com/Nicenonecb/GoVersatile.git
