/**
  Create by guoxin 2020.12.22
*/
package ggorm

import (
    "github.com/GuoxinL/gcomponent/core"
    "gorm.io/gorm"
    "testing"
    "time"
)

type User struct {
    gorm.Model
    Name     string
    Age      int
    Birthday time.Time
}

func TestCreate(t *testing.T) {
    if err := core.SetWorkDirectory(); err != nil {
        t.Error(err)
    }
    New()
    db := GetInstance("test")

    // 创建表
    if err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}); err != nil {
        t.Error("Create table ", err)
    }

    user := User{Name: "Liuguoxin", Age: 18, Birthday: time.Now()}

    if err := db.Create(&user).Error; err != nil {
        t.Error("Add ", err)
    }

    if err := db.Model(&user).Update("name", "hello").Error; err != nil {
        t.Error("Update ", err)
    }

    user = User{}
    if err := db.Where("name = ?", "hello").Find(&user).Error; err != nil {
        t.Error("Select ", err)
    }

    if err := db.Delete(user).Error; err != nil {
        t.Error("Soft delete ", err)
    }

    user = User{}
    if row := db.Where("age = 18").Find(&user).RowsAffected; row != 0 {
        t.Error("Select skip soft delete ", row)
    }

    if err := db.Unscoped().Where("age = 18").Find(&user).Error; err != nil {
        t.Error("Unscoped Select ", err)
    }

    if err := db.Unscoped().Where("name = ?", "hello").Delete(&user).Error; err != nil {
        t.Error("Unscoped delete ", err)
    }

    // 如果存在表则删除（删除时会忽略、删除外键约束)
    if err := db.Migrator().DropTable(&User{}); err != nil {
        t.Error("Drop table ", err)
    }
    t.Log("test success")
}
