# Golang 通用组件 
## 组件介绍与配置
### Environment 环境组件
1. 引用该包  
main文件中引用  
```
import _ "github.com/GuoxinL/gcomponent/components/environment"
```
2. 二进制包环境变量配置  
```
./main --profile=dev
```
main 是golang二进制包
profile的类型：
建议使用`dev`,`uat`,`pro`也可自定义环境名字，多环境文件夹中会提到

3. 多环境配置文件夹  
配置文件夹的名称被指定为`resource`，`resource`文件夹下包含环境文件夹，环境文件夹的名字和`profile`的类型一致  
```
- resource
--| dev
--| pro
--| test
--| uat
```
4. application.yaml
`application.yaml`的定义和名字一样是这个应用的配置，简单的配置可以在`application.yaml`中定义


### Logging 环境组件
1. 引用该包  
main文件中引用  
```
import _ "github.com/GuoxinL/gcomponent/components/logging"
```