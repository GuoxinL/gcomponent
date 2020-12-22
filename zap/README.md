# 缓存组件(redis)
## 如何引用
### 代码配置
```
package main

import (
    "github.com/GuoxinL/gcomponent/gzap"
    "go.uber.org/zap"
)

func main() {
    gzap.New()
    zap.L().Info("zap.L().Info", zap.Int("balabala", 1))
    zap.S().Infof("zap.S().Infof %v", "balabala")
}
```
### 配置文件（application.yaml）
```yaml
components:
  zap:
    addCallerSkip: 0
    console:
      enable: true
      level: debug
      encoder:
        messageKey: msg
        levelKey: level
        timeKey: time
        nameKey: logger
        callerKey: file
        stacktraceKey: stacktrace
        lineEnding: \n
        timeFormat: 2006-01-02T15:04:05.000
    files:
      - enable: true
        level: debug
        encoder:
          messageKey: msg
          levelKey: level
          timeKey: time
          nameKey: logger
          callerKey: file
          stacktraceKey: stacktrace
          lineEnding: \n
          timeFormat: 2006-01-02T15:04:05.000
        logger:
          filename: gcomponent_zap.log
          maxsize: 128
          maxage: 7
          maxbackups: 30
          localtime:
          compress: false
```
