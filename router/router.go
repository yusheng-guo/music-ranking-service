package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/yushengguo557/music-ranking/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) { // 微信公众号
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// QQ Music
	qq := r.Group("/qq")
	{
		qq.GET("/hotsong", v1.GetQQHotSong) // 热歌榜
		qq.GET("/newsong", v1.GetQQNewSong) // 新歌榜
		qq.GET("/soaring", v1.GetQQSoaring) // 飙升榜
		qq.GET("/popular", v1.GetQQPopular) // 流行指数榜
	}

	// migu Music
	migu := r.Group("/migu")
	{
		migu.GET("/newsong", v1.GetMiguNewSong)   // 新歌榜
		migu.GET("/hotsong", v1.GetMiguHotSong)   // 热歌榜
		migu.GET("/original", v1.GetMiguOriginal) // 原创榜
	}

	// kugou Music
	kugou := r.Group("/kugou")
	{
		kugou.GET("/soaring", v1.GetKugouSoaring)   // 飙升榜
		kugou.GET("/top500", v1.GetKugouTop500)     // 酷狗TOP500
		kugou.GET("/popular", v1.GetKugouPopular)   // 蜂鸟流行音乐榜
		kugou.GET("/douyin", v1.GetKugouDouyin)     // 抖音热歌榜
		kugou.GET("/kuaishou", v1.GetKugouKuaishou) // 快手热歌榜
	}
	// wx := r.Group("wx") // 微信公众号
	return r
}
