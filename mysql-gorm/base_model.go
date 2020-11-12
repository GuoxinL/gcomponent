/*
   Created by guoxin in 2020/6/15 10:14 上午
*/
package mysql

/**
主键Id 字符串
*/
type BaseIdS struct {
	Id string `gorm:"id"`
}

/**
主键Id 整型
*/
type BaseIdI struct {
	Id int `gorm:"id"`
}

/**
创建时间
*/
type BaseCreateTime struct {
	CreateTime string `gorm:"create_time"`
}

/**
更新时间
*/
type BaseUpdateTime struct {
	UpdateTime string `gorm:"update_time"`
}

/**
是否使用
*/
type BaseUsed struct {
	Used string `gorm:"used"`
}
