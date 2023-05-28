package crawler

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
	"github.com/yushengguo557/music-ranking-service/dao"
	"github.com/yushengguo557/music-ranking-service/model"
)

// CrawlKuGouMusic crawling kugo music
func (c *Crawler) CrawlKuGouMusic(urls map[string]string) (err error) {
	d := dao.NewDao()
	for tag, url := range urls {
		c.OnHTML("div.pc_temp_songlist > ul > li", func(e *colly.HTMLElement) {
			name, err := e.DOM.Find("a").Html()
			if err != nil {
				panic(err)
			}
			re := regexp.MustCompile("<span[^>]*>.*?</span>")
			name = re.ReplaceAllString(name, "")
			name = strings.TrimSpace(name)

			singer := e.ChildText("a span")
			duration := e.ChildText("span.pc_temp_tips_r > span")
			link_addr := e.ChildAttr("a", "href")
			cover := ""

			s := model.NewSong(name, singer, duration, link_addr, cover, tag, model.KugouMusic)

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
