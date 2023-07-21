package main

import (
	"os"
	"time"

	_ "github.com/fumiama/sqlite3" // import sql
	"github.com/jinzhu/gorm"
)

var (
	dbpath = "./data/"
	dbfile = dbpath + "item.db"
)

type item struct {
	ID       int64
	Title    string    `gorm:"column:title"`
	PageURL  string    `gorm:"column:page_url"`
	Category string    `gorm:"column:category"`
	Intro    string    `gorm:"column:intro"`
	AudioURL string    `gorm:"column:audio_url"`
	Content  string    `gorm:"column:content"`
	Datetime time.Time `gorm:"column:datetime"`
}

func (item) TableName() string {
	return "item"
}

// idb item数据库
var idb *gorm.DB

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
	gdb.AutoMigrate(&item{})
	return gdb
}
