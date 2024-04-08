package databse

import "gorm.io/gorm"

type Database interface {
	getDB() *gorm.DB
}
