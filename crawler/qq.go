package crawler

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
	"github.com/yushengguo557/music-ranking/dao"
	"github.com/yushengguo557/music-ranking/model"
)

// CrawlerQQMusic 爬取 QQ 音乐
func (c *Crawler) CrawlerQQMusic() error {
	d := dao.NewDao()

	c.OnHTML(".songlist__item", func(e *colly.HTMLElement) {
		name := e.ChildText(".songlist__songname_txt")
		singer := e.ChildText(".songlist__artist a")
		duration := e.ChildText(".songlist__time")
		link_addr := e.ChildAttr("div.songlist__songname > span > a:nth-child(2)", "href")
		link := fmt.Sprintf("%s%s%s", "https://", e.Request.URL.Host, link_addr)
		s := model.NewSong(name, singer, duration, link, "")
		fmt.Println(s)
		d.Lock.Lock()
		if err := d.InsertSong(s); err != nil {
			// panic(err)
			log.Println(err)
		}
		d.Lock.Unlock()
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://y.qq.com/n/ryqq/toplist/4")
	c.Visit("https://y.qq.com/n/ryqq/toplist/62")
	c.Visit("https://y.qq.com/n/ryqq/toplist/26")
	c.Visit("https://y.qq.com/n/ryqq/toplist/27")
	c.Visit("https://y.qq.com/n/ryqq/toplist/52")
	c.Visit("https://y.qq.com/n/ryqq/toplist/67")
	return nil
}
