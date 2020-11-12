# 环境组件(environment)
## 如何引用
1. 代码
```go
import (
    _ "github.com/GuoxinL/gcomponent/environment"
)
```
2. 多环境配置文件夹  
配置文件夹的名称被指定为`resource`，`resource`文件夹下包含环境文件夹（用于区分不同环境），环境文件夹的名字和`profile`的类型一致  
```
-| resource
--| dev
---| application.yaml
--| pro
---| application.yaml
--| test
---| application.yaml
--| uat
---| application.yaml
```

3. application.yaml
`application.yaml`的定义和名字一样是这个应用的配置，简单的配置可以在`application.yaml`中定义，不同环境的`application.yaml`会在不同
环境生效  

4. 环境变量参数
```
./main --profile=dev
```
main 是golang二进制包
profile的类型：
建议使用`dev`,`uat`,`pro`也可自定义环境名字，和环境名称一致即可

## 如何使用
1. GetApplicationFile() viper.Viper  
>注释：获得当前环境目录下application.yaml的Viper对象
>参数：
>返回值：获得配置文件对象实例

2. GetConfig(prefix string, config interface{}) error  
>注释：获得当前环境目录下application.yaml中的配置
>参数：
>    prefix: yaml配置前缀
>    config: yaml配置对象映射
>返回值：异常

3. GetEnvironmentDirectory() string  
>注释：获得当前环境目录
>参数：
>返回值：当前环境目录