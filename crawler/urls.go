package crawler

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

// CrawlKugouRankUrls -- find all kugou rank links
func (c *Crawler) CrawlKugouRankUrls() (map[string]string, error) {
	var err error
	urls := make(map[string]string)
	c.OnHTML("div.pc_temp_side > div > ul > li > a", func(e *colly.HTMLElement) {
		url := e.Attr("href")
		if strings.HasPrefix(url, "https://www.kugou.com/yy/rank/home/") {
			title := e.Attr("title")
			urls[title] = url
		}
	})

	if err = c.Visit("https://www.kugou.com/yy/html/rank.html"); err != nil {
		return nil, fmt.Errorf("visit %s, err: %w", "https://www.kugou.com/yy/html/rank.html", err)
	}
	return urls, nil
}

// CrawlMiguRankUrls -- find all migu rank links
func (c *Crawler) CrawlMiguRankUrls() (map[string]string, error) {
	var err error
	urls := make(map[string]string)
	c.OnHTML("div.board-sord > div > ul > li > a", func(e *colly.HTMLElement) {
		url := e.Attr("href")
		title := e.Text
		urls[title] = fmt.Sprintf("%s%s", "https://music.migu.cn", url)
	})

	if err = c.Visit("https://music.migu.cn/v3/music/top/"); err != nil {
		return urls, fmt.Errorf("visit %s, err: %w", "https://music.migu.cn/v3/music/top/", err)
	}

	return urls, nil
}

// CrawlQQRankUrls -- find all qq rank links
func (c *Crawler) CrawlQQRankUrls() (map[string]string, error) {
	urls := make(map[string]string)
	urls["飙升榜"] = "https://y.qq.com/n/ryqq/toplist/62"
	urls["热歌榜"] = "https://y.qq.com/n/ryqq/toplist/26"
	urls["新歌榜"] = "https://y.qq.com/n/ryqq/toplist/27"
	urls["流行指数榜"] = "https://y.qq.com/n/ryqq/toplist/4"
	urls["内地榜"] = "https://y.qq.com/n/ryqq/toplist/5"
	urls["香港地区榜"] = "https://y.qq.com/n/ryqq/toplist/59"
	urls["欧美榜"] = "https://y.qq.com/n/ryqq/toplist/3"
	return urls, nil
}

// CrawlNeteaseRankUrls -- find all netease rank links
func (c *Crawler) CrawlNeteaseRankUrls() (map[string]string, error) {
	var err error
	urls := make(map[string]string)
	c.OnHTML("ul > li > div", func(e *colly.HTMLElement) {
		url := e.ChildAttr("p.name > a", "href")
		title := e.ChildText("p.s-fc4")
		urls[title] = fmt.Sprintf("%s%s", "https://music.163.com/#", url)
	})

	if err = c.Visit("https://music.163.com/#/discover/toplist"); err != nil {
		return urls, fmt.Errorf("visit %s, err: %w", "https://music.163.com/#/discover/toplist", err)
	}

	return urls, nil
}
