package crawler

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/yushengguo557/music-ranking/model"
)

func (c *Crawler) CrawlerNetEaseMusic() error {
	// d := dao.NewDao()
	c.OnHTML(".tt", func(e *colly.HTMLElement) {
		name := e.ChildText("span.txt a b")
		singer := e.ChildText("td div.text span a")
		duration := e.ChildText("span.u-dur")
		link_addr := e.ChildAttr("div.ttc span.txt a", "href")
		link := fmt.Sprintf("%s%s%s", "https://", e.Request.URL.Host, link_addr)
		s := model.NewSong(name, singer, duration, link, "")
		fmt.Println("song: ", s)
		// d.Lock.Lock()
		// if err := d.InsertSong(s); err != nil {
		// 	// panic(err)
		// 	log.Println(err)
		// }
		// d.Lock.Unlock()
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://music.163.com/#/discover/toplist")
	return nil
}
