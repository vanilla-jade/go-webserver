package database

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var SqlDB *gorm.DB

func init() {
  var err error
  SqlDB, err = gorm.Open("sqlite3", "test.db")
  if err != nil {
    panic("连接数据库失败")
  }
}
