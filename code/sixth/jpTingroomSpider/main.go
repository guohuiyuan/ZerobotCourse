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
	re     = regexp.MustCompile(`http://jp\.tingroom\.com/(tingli|gequ)/.*\.html`)
	tre    = regexp.MustCompile(`http://jp\.tingroom\.com/(tingli|gequ)`)
	timeRe = regexp.MustCompile(`(\d{4}-\d{2}-\d{2})`)
)

func init() {
	_ = os.MkdirAll(dbpath, 0666)
	idb = initialize(dbfile)
}
func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("jp.tingroom.com"),
		colly.URLFilters(tre),
		colly.MaxDepth(4),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36"),
		colly.CacheDir("./cache"),
	)

	detailCollector := c.Clone()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if re.MatchString(link) {
			detailCollector.Visit(link)
		}
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	detailCollector.OnHTML("div[class=viewbox]", func(e *colly.HTMLElement) {
		fmt.Println("Visiting", e.Request.URL)
		title := e.ChildText("div[class=title] > h2")
		timeStr := e.ChildText("div[class=info]:nth-child(2)")
		content := strings.ReplaceAll(e.ChildText("#article"), "\u3000", "")
		var category string
		var datetime time.Time
		var audioURL string
		if re.MatchString(e.Request.URL.String()) {
			category = re.FindStringSubmatch(e.Request.URL.String())[1]
		}
		if timeRe.MatchString(timeStr) {
			datetime, _ = time.ParseInLocation("2006-01-02", timeRe.FindStringSubmatch(timeStr)[1], time.Local)
		}
		intro := e.ChildText("div[class=intro]")
		t := colly.NewCollector()
		t.OnHTML("body", func(h *colly.HTMLElement) {
			audioURL = h.ChildAttr("#tbc_01 > audio", "src")
			if audioURL == "" {
				audioURL = h.ChildAttr("#tbc_02 > audio", "src")
			}
		})
		playURL := e.ChildAttr("IFRAME", "src")
		t.Visit(playURL)
		it := item{Datetime: datetime, Title: title, Intro: intro, AudioURL: audioURL, Content: content, PageURL: e.Request.URL.String(), Category: category}
		idb.Debug().Model(&item{}).Create(&it)
	})

	// detailCollector.Visit("http://jp.tingroom.com/gequ/dmgequ/8233.html")
	// detailCollector.Visit("http://jp.tingroom.com/gequ/lxgequ/66580.html")
	c.Visit("http://jp.tingroom.com/gequ/")
	c.Visit("http://jp.tingroom.com/tingli/")
}
