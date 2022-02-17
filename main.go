package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/utils"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func lastSlashOfString(imgSrc string) string {
	lastSlashIndex := strings.LastIndex(imgSrc, "/")
	newName := imgSrc[lastSlashIndex+1:]
	log.Printf("Downloading %s", newName)
	return newName
}

func main() {
	qUrl := flag.String("q", "https://news.hmgjournal.com/Group-Story/%ED%98%84%EB%8C%80%EC%9E%90%EB%8F%99%EC%B0%A8%EA%B7%B8%EB%A3%B9-JD%ED%8C%8C%EC%9B%8C%EB%A1%9C%EB%B6%80%ED%84%B0-%EA%B8%80%EB%A1%9C%EB%B2%8C-%EC%B5%9C%EA%B3%A0-%EC%88%98%EC%A4%80%EC%9D%98-%ED%92%88%EC%A7%88%EA%B3%BC-%EB%82%B4%EA%B5%AC%EC%84%B1%EC%9D%84-%EC%9E%85%EC%A6%9D-%EB%B0%9B%EB%8B%A4", "URL")
	flag.Parse()

	pageName, err := url.QueryUnescape(lastSlashOfString(*qUrl))
	if err != nil {
		log.Fatal(err)
		return
	}

	// page := rod.New().MustConnect().MustPage("https://news.hmgjournal.com/Group-Story/%ED%98%84%EB%8C%80%EC%9E%90%EB%8F%99%EC%B0%A8%EA%B7%B8%EB%A3%B9-JD%ED%8C%8C%EC%9B%8C%EB%A1%9C%EB%B6%80%ED%84%B0-%EA%B8%80%EB%A1%9C%EB%B2%8C-%EC%B5%9C%EA%B3%A0-%EC%88%98%EC%A4%80%EC%9D%98-%ED%92%88%EC%A7%88%EA%B3%BC-%EB%82%B4%EA%B5%AC%EC%84%B1%EC%9D%84-%EC%9E%85%EC%A6%9D-%EB%B0%9B%EB%8B%A4")
	page := rod.New().MustConnect().MustPage(*qUrl)

	page.MustWaitLoad()

	content := page.MustElement(".view-cont")
	images := content.MustElements("img")

	for i, img := range images {
		log.Printf("Index %d", i)
		img_src := *img.MustAttribute("src")
		log.Println(img_src)
		bin, _ := page.GetResource(img_src)
		utils.OutputFile("./download/"+lastSlashOfString(img_src), bin)
	}

	f, err := os.Create("./download/" + pageName + ".html")
	check(err)
	defer f.Close()

	html := content.MustHTML()
	_, err2 := f.WriteString(html)
	check(err2)
	log.Printf("Downloading %s", html[:200])

	time.Sleep(time.Second * 3)

}
