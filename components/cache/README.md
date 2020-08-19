# 缓存组件(redis)
## 如何引用
1. 代码
```go
import (
    _ "github.com/GuoxinL/gcomponent/components/cache"
)
```
2. 配置文件（application.yaml）
```yaml
components:
  redis:
    rds:
        # redis名称
      - name: risk
        # 集群节点
        nodes:
          - 192.168.155.81:7378
          - 192.168.155.82:7378
          - 192.168.155.235:7378
          - 192.168.155.81:7379
          - 192.168.155.82:7379
          - 192.168.155.235:7379
        # 密码
        password: risk
        # 连接池配置
        pool:
          maxIdle: 10
          idleTimeout: 240
          maxConnLifetime: 300
        # 连接超时配置
        connection:
          readTimeout: 50
          writeTimeout: 50
          connectTimeout: 50
```
## 如何使用
1. func (this *clusterWrapper) Get(key string) (string, error)
>注释：get
>参数：key
>返回值：value, err
2. func (this *clusterWrapper) Set(key string, value interface{}) (string, error)
>注释：set
>参数：key, value
>返回值：value, err
