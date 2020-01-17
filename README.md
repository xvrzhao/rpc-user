# micro-stacks/rpc-user

[![Build Status](https://travis-ci.com/micro-stacks/rpc-user.svg?branch=master)](https://travis-ci.com/micro-stacks/rpc-user)

通用的用户 gRPC 服务，用于二次开发。

## 项目架构

- cache *缓存客户端*
- cmd *命令脚本*
- db *数据库客户端*
    - migrations *迁移记录*
    - models *表模型*
- proto *RPC 方法定义*
- server *RPC 方法实现*
- main.go *编译入口*

## 安装方式

1. `git clone` 代码到本地
2. 根据 `.env.example` 文件设置环境变量文件 `.env`
3. 根据环境变量中的 `DB_NAME` 建立一个数据库，首次运行会自动执行数据表迁移，之后的数据库结果变动请遵循下文 *开发规范-数据库迁移* 一节

## 运行方式

- cmd/startup.sh 启动应用
- cmd/test.sh 执行测试
- cmd/proto.sh 重新生成 protobuf Go 代码

## 环境依赖 / 技术栈

- .env 环境变量文件，参见 .env.example 文件
- MySQL（jinzhu/gorm）
- Redis（go-redis/redis/v7）
- gRPC
- Protocol Buffers

## 开发规范

### 数据库迁移
1. 数据库表或字段变动，不可以直接在数据库修改，而是应该通过迁移功能进行。由于迁移记录的存在，通过 Git 可以查出新增的迁移，在服务上线之前先复现迁移，这样可保证不同环境下的数据库表和字段的统一性。
2. 由于 gorm 的 AutoMigrate 方法只能迁移新增的表和字段，而不可以迁移对表和字段所做的修改和删除，所以新增表和字段只需新增模型结构体和模型结构体字段即可，**但对于修改或删除操作，不仅要修改模型结构体或删除模型，还要写迁移记录，这是为了保持模型和数据表的统一。**
3. 一条迁移记录就是一个 SQL 文件，命名应以 `datetime_description.sql` 的形式，如 `20200102112312_user_table_modify_salt_field.sql`。
4. db/models 目录中每个模型结构体的方法必须定义在值上，而非指针上，并且在调用这些方法时需要模型实例为空对象，这样不会出现意外的查询条件的问题。