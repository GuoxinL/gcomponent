# 系统优化(pprof)
## 如何引用
1. 代码
```go
import (
    _ "github.com/GuoxinL/gcomponent/components/optimize-pprof"
)
```
2. 配置文件（application.yaml）
```yaml
components:
  optimize:
    pprof:
      # 端口号
      port: 19190
      # 是否开启
      enable: true
```
## 如何使用
### 性能观察、调优
https://segmentfault.com/a/119000001641201