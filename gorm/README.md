# 数据库组件(mysql)
## 如何引用
1. 代码
```go
import (
    _ "github.com/GuoxinL/gcomponent/components/mysql-gorm"
)
```
2. 配置文件（application.yaml）
```yaml
components:
  mysql:
    dataSources:
      - name: test1
        url: 127.0.0.1:3306
        database: test1
        username: mysql
        password: mysql
        MaxIdleConns: 100
        MaxOpenConns: 10
        ConnMaxLifetime: 60
      - name: test
        url: 127.0.0.1:3306
        database: test
        username: root
        password: mysql
        MaxIdleConns: 100
        MaxOpenConns: 10
        ConnMaxLifetime: 60
```
## 如何使用
1. func GetInstance(name string) *gorm.DB
>注释：通过该方法获得`gorm.DB`对象
>参数：name与配置文件中`components.mysql.dataSources.name`字段对应
>返回值：*gorm.DB
