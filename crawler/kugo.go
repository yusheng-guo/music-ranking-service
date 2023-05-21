package crawler

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
	"github.com/yushengguo557/music-ranking/model"
)

func (c *Crawler) CrawlerKuGouMusic() (err error) {
	// d := dao.NewDao()
	// document.querySelector("#rankWrap > div.pc_temp_songlist > ul > li:nth-child(1) > a")
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
		s := model.NewSong(name, singer, duration, link_addr, "")
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

	urls := []string{
		"https://www.kugou.com/yy/rank/home/1-6666.html",
		"https://www.kugou.com/yy/rank/home/1-8888.html",
		"https://www.kugou.com/yy/rank/home/1-59703.html",
		"https://www.kugou.com/yy/rank/home/1-52144.html",
	}

	for _, url := range urls {
		if err = c.Visit(url); err != nil {
			return fmt.Errorf("visit %s, err: %w", url, err)
		}
	}

	return nil
}
