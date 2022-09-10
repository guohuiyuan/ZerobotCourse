package compare

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

var (
	re    = regexp.MustCompile(`<audio src="(.*?)" controls="controls"></div>`)
	reStr string
	doc   *html.Node
	doc2  *goquery.Document
)

func init() {
	data, err := os.ReadFile("test.html")
	if err != nil {
		fmt.Println("err:", err)
	}
	reStr = string(data)
	doc, err = html.Parse(strings.NewReader(reStr))
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Printf("xpath doc:%#v\n", htmlquery.OutputHTML(doc, true))
	doc2, err = goquery.NewDocumentFromReader(strings.NewReader(reStr))
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Print("goquery doc: ")
	fmt.Println(doc2.Html())
}
func Re() string {
	// fmt.Println(re.FindAllStringSubmatch(reStr, -1))
	return re.FindAllStringSubmatch(reStr, -1)[0][1]
}

func Xpath() string {
	list, err := htmlquery.QueryAll(doc, "//audio")
	if err != nil {
		fmt.Println("err:", err)
	}
	// for _, v := range list {
	// 	fmt.Printf("%#v\n", v)
	// }
	return list[0].Attr[0].Val
}

func Goquery() string {
	// var res []string
	// doc2.Find("audio").Each(func(_ int, s *goquery.Selection) {
	// 	if attr, ok := s.Attr("src"); ok {
	// 		res = append(res, strings.TrimSpace(attr))
	// 	}
	// })
	if attr, ok := doc2.Find("audio").Attr("src"); ok {
		return strings.TrimSpace(attr)
	}
	return ""
}

func TestRe(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(Re())
		})
	}
}

func TestXpath(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(Xpath())
		})
	}
}

func TestGoquery(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(Goquery())
		})
	}
}

func BenchmarkRe(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Re()
	}
}

func BenchmarkXpath(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Xpath()
	}
}

func BenchmarkGoquery(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Goquery()
	}
}

func BenchmarkXpathWithLoad(b *testing.B) {
	var err error
	for n := 0; n < b.N; n++ {
		doc, err = html.Parse(strings.NewReader(reStr))
		if err != nil {
			fmt.Println("err:", err)
		}
		Xpath()
	}
}

func BenchmarkGoqueryWithLoad(b *testing.B) {
	var err error
	for n := 0; n < b.N; n++ {
		doc2, err = goquery.NewDocumentFromReader(strings.NewReader(reStr))
		if err != nil {
			fmt.Println("err:", err)
		}
		Goquery()
	}
}
