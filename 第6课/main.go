package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

var (
	re = regexp.MustCompile(`https://news.cctv.com/(\d{4}/\d{2}/\d{2})/[A-Za-z0-9]+.shtml`)
)

func init() {
	_ = os.RemoveAll(dbpath)
	_ = os.MkdirAll(dbpath, 0666)
	adb = initialize(dbfile)
}
func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("news.cctv.com"),
		colly.MaxDepth(4),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36"),
		colly.CacheDir("./cache"),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if re.MatchString(link) {
			t, err := time.ParseInLocation("2006/01/02", re.FindStringSubmatch(link)[1], time.Local)
			if err != nil {
				fmt.Println(err)
			} else {
				a := article{Datetime: t, ArticleURL: link, Description: strings.TrimSpace(e.Text)}
				adb.Debug().Model(&article{}).Create(&a)
			}
		}
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://news.cctv.com/")
}
