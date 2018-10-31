package types

import (
    "time"
    "github.com/jinzhu/gorm"
)

type Product struct {
  gorm.Model
  Code string
  Price uint
}

type Record struct {
  gorm.Model
  Msg string
  Time time.Time
}
