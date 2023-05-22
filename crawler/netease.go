package crawler

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/yushengguo557/music-ranking/dao"
	"github.com/yushengguo557/music-ranking/model"
)

// CrawlNetEaseMusic 爬取 网易音乐
func (c *Crawler) CrawlNetEaseMusic(urls map[string]string) (err error) {
	d := dao.NewDao()
	for tag, url := range urls {
		c.OnHTML(".tt", func(e *colly.HTMLElement) {
			name := e.ChildText("span.txt a b")
			singer := e.ChildText("td div.text span a")
			duration := e.ChildText("span.u-dur")
			link_addr := e.ChildAttr("div.ttc span.txt a", "href")
			link := fmt.Sprintf("%s%s%s", "https://", e.Request.URL.Host, link_addr)
			cover := ""
			s := model.NewSong(name, singer, duration, link, cover, tag, model.NeteaseMusic)
			d.Lock.Lock()
			if err = d.InsertSong(s); err != nil {
				panic(err)
			}
			d.Lock.Unlock()
		})

		if err = c.Visit(url); err != nil {
			return fmt.Errorf("CrawlNetEaseMusic, err: %w", err)
		}
	}
	return nil
}
