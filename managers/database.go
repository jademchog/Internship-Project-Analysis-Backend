package managers

import (
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Init 初始化 SQLite 数据库连接并初始化 Manager
func Init(dbPath string) (*gorm.DB, error) {
	dir := filepath.Dir(dbPath)
	if dir != "." && dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, err
		}
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	initManagers(db)

	return db, nil
}

var (
	GraduateEmploymentManager *graduateEmploymentManager
)

func initManagers(db *gorm.DB) {
	GraduateEmploymentManager = &graduateEmploymentManager{db: db}
}
