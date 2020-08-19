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
        # 连接名称在在吗中可以调用GetInstance("indicatorsys")获得实例
      - name: indicatorsys
        url: 192.168.145.151:3306
        database: indicatorsys_test_docker
        username: mysql
        password: mysql
        MaxIdleConns: 100
        MaxOpenConns: 10
        ConnMaxLifetime: 60
      - name: vesta
        url: 192.168.155.178:3306
        database: vesta_test
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
