# 环境组件(environment)
## 如何引用
### 代码配置
```
package main

import (
	"github.com/GuoxinL/gcomponent/environment"
)

func main() {
    environment.New()
    name := environment.GetName()
    fmt.Println(name)
    file := environment.GetApplicationFile()
    fmt.Println(file)
    profile := environment.GetProfile()
    fmt.Println(profile)
}
```
### 配置文件

配置文件夹的名称被指定为`conf`

1. 单环境配置
适合可以不区分环境直接在`conf`文件夹下直接创建配置文件
application.yaml
```
-| conf
--| application.yaml
```
2. 多环境配置  
`conf`文件夹下包含环境文件夹（用于区分不同环境），环境文件夹的名字和`profile`的类型一致  
```
-| conf
--| dev
---| application.yaml
--| pro
---| application.yaml
--| test
---| application.yaml
--| uat
---| application.yaml
```

3. 环境变量参数
```
./main --profile=dev
```
main 是golang二进制包
profile的类型：可以定义任何环境名称只要和配置文件的文件夹名称一致即可
