package main

import (
	"os"
	"time"

	_ "github.com/fumiama/sqlite3" // import sql
	"github.com/jinzhu/gorm"
)

var (
	dbpath = "./data/"
	dbfile = dbpath + "article.db"
)

type article struct {
	ID          int64
	Datetime    time.Time `gorm:"column:datetime"`
	ArticleURL  string    `gorm:"column:article_url"`
	Description string    `gorm:"column:description"`
}

func (article) TableName() string {
	return "article"
}

// sdb 得分数据库
var adb *gorm.DB

// initialize ...
func initialize(dbpath string) *gorm.DB {
	var err error
	if _, err = os.Stat(dbpath); err != nil || os.IsNotExist(err) {
		// 生成文件
		f, err := os.Create(dbpath)
		if err != nil {
			return nil
		}
		defer f.Close()
	}
	gdb, err := gorm.Open("sqlite3", dbpath)
	if err != nil {
		panic(err)
	}
	gdb.AutoMigrate(article{})
	return gdb
}
