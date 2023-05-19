package crawler

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/yushengguo557/music-ranking/model"
)

// Find and visit all links
// c.OnHTML("a", func(e *colly.HTMLElement) {
// 	e.Request.Visit(e.Attr("href"))
// })

// c.OnHTML("a", func(e *colly.HTMLElement) {
// 	e.Request.Visit(e.Attr("href"))
// })

// c.OnRequest(func(r *colly.Request) {
// 	fmt.Println("Visiting", r.URL)
// })

// c.Visit("http://go-colly.org/")

func Crawler() {
	c := colly.NewCollector()
	colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")
	colly.MaxDepth(1)
	colly.AllowedDomains("y.qq.com")

	c.OnHTML(".songlist__item", func(e *colly.HTMLElement) {
		name := e.ChildText(".songlist__songname_txt")
		singer := e.ChildText(".songlist__artist a")
		duration := e.ChildText(".songlist__time")
		addr := e.ChildAttr("div.songlist__songname > span > a:nth-child(2)", "href")
		link := fmt.Sprintf("%s%s%s", "https://", e.Request.URL.Host, addr)
		fmt.Println(model.NewSong(name, singer, duration, link))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// c.Visit("https://y.qq.com/n/ryqq/toplist/4")
	c.Visit("https://y.qq.com/n/ryqq/toplist/62")
}
