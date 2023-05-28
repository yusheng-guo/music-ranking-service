package crawler

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/yushengguo557/music-ranking-service/dao"
	"github.com/yushengguo557/music-ranking-service/model"
)

// CrawlQQMusic 爬取 QQ 音乐
func (c *Crawler) CrawlQQMusic(urls map[string]string) (err error) {
	d := dao.NewDao()

	for tag, url := range urls {
		c.OnHTML(".songlist__item", func(e *colly.HTMLElement) {
			name := e.ChildText(".songlist__songname_txt")
			singer := e.ChildText(".songlist__artist a")
			duration := e.ChildText(".songlist__time")
			link := e.ChildAttr("div.songlist__songname > span > a", "href")
			link = fmt.Sprintf("%s%s%s", "https://", e.Request.URL.Host, link)
			cover := ""

			s := model.NewSong(name, singer, duration, link, cover, tag, model.QQMusic)
			d.Lock.Lock()
			if err := d.InsertSong(s); err != nil {
				panic(err)
			}
			d.Lock.Unlock()
		})

		if err = c.Visit(url); err != nil {
			return fmt.Errorf("visit %s, err: %w", url, err)
		}
	}
	return nil
}
