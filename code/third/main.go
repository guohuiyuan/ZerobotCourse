package main

import (
	"fmt"
	"time"

	sql "github.com/FloatTech/sqlite"
)

type article struct {
	ID         int64  `db:"id"`
	Title      string `db:"title"`
	Author     string `db:"author"`
	CreateTime string `db:"createTime"`
	Content    string `db:"content"`
}

var db = &sql.Sqlite{}

// 暂时随机选择一个小作文
func getArticleByKeyword(keyword string) (a article) {
	_ = db.Find("main", &a, "where content LIKE '%"+keyword+"%'")
	return
}

func getRandomArticle() (a article) {
	_ = db.Pick("main", &a)
	return
}

func main() {
	db.DBPath = "code/third/小作文.db"
	err := db.Open(time.Hour * 24)
	if err != nil {
		fmt.Println(err)
	}
	err = db.Create("main", &article{})
	if err != nil {
		fmt.Println(err)
	}
	n, err := db.Count("main")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("一共有", n, "条记录")
	a := getRandomArticle()
	fmt.Printf("%+v\n", a)
}
