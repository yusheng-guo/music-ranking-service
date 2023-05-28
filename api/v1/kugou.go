package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yushengguo557/music-ranking-service/dao"
	"github.com/yushengguo557/music-ranking-service/model"
)

func GetKugouSoaring(c *gin.Context) {
	s := dao.NewDao()
	s.Lock.RLock()
	songs, err := s.QueryMultiSongs("酷狗飙升榜", model.KugouMusic)
	s.Lock.RUnlock()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"songs": songs})
}

func GetKugouTop500(c *gin.Context) {
	s := dao.NewDao()
	s.Lock.RLock()
	songs, err := s.QueryMultiSongs("酷狗TOP500", model.KugouMusic)
	s.Lock.RUnlock()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"songs": songs})
}

func GetKugouPopular(c *gin.Context) {
	s := dao.NewDao()
	s.Lock.RLock()
	songs, err := s.QueryMultiSongs("蜂鸟流行音乐榜", model.KugouMusic)
	s.Lock.RUnlock()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"songs": songs})
}

func GetKugouDouyin(c *gin.Context) {
	s := dao.NewDao()
	s.Lock.RLock()
	songs, err := s.QueryMultiSongs("抖音热歌榜", model.KugouMusic)
	s.Lock.RUnlock()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"songs": songs})
}

func GetKugouKuaishou(c *gin.Context) {
	s := dao.NewDao()
	s.Lock.RLock()
	songs, err := s.QueryMultiSongs("快手热歌榜", model.KugouMusic)
	s.Lock.RUnlock()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		panic(err)
		// return
	}
	c.JSON(http.StatusOK, gin.H{"songs": songs})
}
