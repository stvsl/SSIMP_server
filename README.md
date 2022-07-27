# SSIMP_server 服务器设计定义

预期要实现的功能：所有业务的最终逻辑实现

### 文件目录：

db：mysql（mariadb）数据库相关组件实现以及相关业务的封装实现

security：安全加密以及安全验证中间层实现

service：各类实际请求业务的处理以及逻辑实现

web：静态资源存储

utils：上述部分需要的组件以及封装

### 技术框架以及相关模块：

服务器框架：gin

redis数据库：github.com/go-redis/redis/v8

mysql数据库框架：gorm





