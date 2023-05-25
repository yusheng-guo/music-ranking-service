package main

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/yushengguo557/music-ranking/api/v1"
	"github.com/yushengguo557/music-ranking/crawler"
	"github.com/yushengguo557/music-ranking/global"
)

func init() {
	// 初始化 mysql 数据库
	if err := global.InitDB(); err != nil {
		panic(err)
	}
}

// data := "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAoAAAAKCAQAAAAnOwc2AAAAD0lEQVR4AWM4bYwJh7QgANLSYzmZH3ISAAAAAElFTkSuQmCC" // 这是一个示例数据，实际上应该从爬取结果中获取
// b64data := data[strings.IndexByte(data, ',')+1:]                                                                                 // 去掉"data:image/png;base64,"前缀
// img, err := base64.StdEncoding.DecodeString(b64data)                                                                             // 解码base64字符串
// if err != nil {
// 	panic(err)
// }
// err = ioutil.WriteFile("image.png", img, 0644) // 将图片保存为本地文件
// if err != nil {
// 	panic(err)
// }

func main() {
	// 关闭数据库连接
	defer global.DB.Close()
	go func() {
		// 爬虫
		if err := crawler.Run(); err != nil {
			panic(err)
		}
	}()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) { // 微信公众号
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// QQ Music
	qq := r.Group("/qq")
	{
		qq.GET("/hotsong") // 热歌榜
		qq.GET("/newsong") // 新歌榜
		qq.GET("/soaring") // 飙升榜
		qq.GET("/popular") // 流行指数榜
	}

	// migu Music
	migu := r.Group("/migu")
	{
		migu.GET("/newsong")  // 新歌榜
		migu.GET("/hotsong")  // 热歌榜
		migu.GET("/original") // 原创榜
	}

	// kugou Music
	kugou := r.Group("/kugou")
	{
		kugou.GET("/soaring")  // 飙升榜
		kugou.GET("/top500")   // 酷狗TOP500
		kugou.GET("/popular")  // "蜂鸟流行音乐榜
		kugou.GET("/douyin")   // 抖音热歌榜
		kugou.GET("/kuaishou") // 快手热歌榜
	}
	wx := r.Group("wx") // 微信公众号
	{
		wx.GET("/songs", v1.GetSongs)
	}
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
