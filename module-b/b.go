package module_b

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var B = "module b"

func init() {
	_, _ = gorm.Open(mysql.Open("foo"))
	_, _ = gorm.Open(sqlite.Open("foo"))
}
