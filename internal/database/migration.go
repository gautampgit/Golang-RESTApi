package database

import (
	"github.com/gautampgit/Golang-RESTApi/internal/comment"
	"github.com/jinzhu/gorm"
)

//MigrateDB - Migrates the Database and creates a comment table
func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(comment.Comment{}); result.Error != nil {
		return result.Error
	}
	return nil
}
