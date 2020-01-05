# User RPC Service

A gRPC service provides user-related methods.

## 运行方式

- cmd/startup.sh 启动应用
- cmd/test.sh 执行测试

## 环境依赖

- .env 环境变量文件，参见 .env.example 文件
- MySQL

## 开发规范

### 数据库迁移
1. 数据库表或字段变动，不可以直接在数据库修改，而是应该通过迁移功能进行。由于迁移记录的存在，通过 Git 可以查出新增的迁移，在服务上线之前先复现迁移，这样可保证不同环境下的数据库表和字段的统一性。
2. 由于 gorm 的 AutoMigrate 方法只能迁移新增的表和字段，而不可以迁移对表和字段所做的修改和删除，所以新增表和字段只需新增模型结构体和模型结构体字段即可，**但对于修改或删除操作，不仅要修改模型结构体或删除模型，还要写迁移记录，这是为了保持模型和数据表的统一。**
3. 一条迁移记录就是一个 SQL 文件，命名应以 `datetime_description.sql` 的形式，如 `20200102112312_user_table_modify_salt_field.sql`。