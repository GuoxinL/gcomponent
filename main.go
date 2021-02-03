/*
   Created by guoxin in 2020/4/10 3:26 下午
*/
package main

import (
    ggorm "github.com/GuoxinL/gcomponent/gorm"
    gzap "github.com/GuoxinL/gcomponent/zap"
)

func main() {
    gzap.New()
    ggorm.New()

}
