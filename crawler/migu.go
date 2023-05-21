package crawler

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/yushengguo557/music-ranking/dao"
	"github.com/yushengguo557/music-ranking/model"
)

// CrawlMiguMusic 爬取 咪咕音乐
func (c *Crawler) CrawlMiguMusic(urls map[string]string) (err error) {
	d := dao.NewDao()
	for tag, url := range urls {
		c.OnHTML(".songlist-item", func(e *colly.HTMLElement) {
			name := e.ChildText(".song-name-text > a")
			if name == "" {
				return
			}
			singer := ""
			duration := e.ChildText(".song-length")
			link := e.ChildAttr(".song-name-text > a", "href")
			link = fmt.Sprintf("%s%s", "https://music.migu.cn/v3/", link)
			cover := e.ChildAttr("div.song-name > div > img", "src")
			cover = fmt.Sprintf("%s%s", "https:", cover)

			s := model.NewSong(name, singer, duration, link, cover, tag, model.MiguMusic)

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
