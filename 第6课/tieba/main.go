package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var (
	API_BASE_URL = "https://tieba.baidu.com/p/6105274031" // 帖子链接
	maxPage      = 6                                      // 最大页码数
	wg           sync.WaitGroup
)

func main() {
	var baseDir string
	tempBaseDir := getBasictDirectory()
	// fmt.Printf("默认本地路径：%s \n", tempBaseDir)
	fmt.Println("请输入存储路径：")

	_, err := fmt.Scanf("%s", &baseDir)
	if err != nil {
		return
	}

	if baseDir == "" {
		baseDir = tempBaseDir
		log.Println("当前路径为空，已为你使用默认路径")
	}

	for i := 1; i <= maxPage; i++ {
		getTieBaPic(i, API_BASE_URL+"?pn="+strconv.Itoa(i), baseDir)
	}
	log.Println("图片下载完毕")
}

// 获取贴吧图片
func getTieBaPic(page int, url string, basePath string) {
	log.Printf("正在获取第 %d 页图片 \n", page)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("net errer:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Printf("status code error: %d %s \n", response.StatusCode, response.Status)
		return
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return
	}

	var images []string
	doc.Find(".BDE_Image").Each(func(i int, selection *goquery.Selection) {
		src, _ := selection.Attr("src")
		images = append(images, src)
	})
	length := len(images)

	if length <= 0 {
		return
	}

	tempPath := basePath
	createDir(tempPath)
	for _, _item := range images {
		tempIndex := strconv.FormatInt(int64(time.Now().Nanosecond()), 10)
		fileName := tempPath + "/" + tempIndex + ".jpg"
		log.Printf("正在下载图片 -> %s \n", fileName)
		_item := _item
		wg.Add(1)
		go func() {
			defer wg.Done()
			downloadImage(_item, fileName)

		}()
	}
	wg.Wait()
}

// 下载图片
func downloadImage(url string, path string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("net errer:", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	_ = ioutil.WriteFile(path, body, 0755)
}

// 检查文件夹是否存在
func createDir(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		err := os.MkdirAll(path, 0766)
		if err != nil {
			return false
		}
		return true
	}
	if os.IsNotExist(err) {
		err := os.MkdirAll(path, 0766)
		if err != nil {
			return false
		}
		return true
	}
	return true
}

// 获取基本路径
func getBasictDirectory() string {
	// 返回绝对路径
	// filepath.Dir(os.Args[0]) 去除最后一个元素的路径
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Println(err)
	}

	// 将 \\ 替换为 /
	return strings.Replace(dir, "\\", "/", -1)
}
