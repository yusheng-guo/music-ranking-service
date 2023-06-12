package v1

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yushengguo557/music-ranking-service/dao"
	"github.com/yushengguo557/music-ranking-service/model"
	"github.com/yushengguo557/music-ranking-service/model/response"
)

func GetQQHotSong(c *gin.Context) {
	s := dao.NewDao()
	s.Lock.RLock()
	songs, err := s.QueryMultiSongs("热歌榜", model.QQMusic)
	s.Lock.RUnlock()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
	// c.JSON(http.StatusOK, gin.H{"songs": songs})
	// response.OKWithData(c, songs)
	if len(songs) > 0 {
		response.OKWithDataAndMessage(c, songs, "success")
		return
	}
	response.OKWithMessage(c, "len(songs)=0")
}

func GetQQNewSong(c *gin.Context) {
	s := dao.NewDao()
	s.Lock.RLock()
	songs, err := s.QueryMultiSongs("新歌榜", model.QQMusic)
	s.Lock.RUnlock()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"songs": songs})
}

func GetQQSoaring(c *gin.Context) {
	s := dao.NewDao()
	s.Lock.RLock()
	songs, err := s.QueryMultiSongs("飙升榜", model.QQMusic)
	s.Lock.RUnlock()
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"songs": songs})
}

func GetQQPopular(c *gin.Context) {
	s := dao.NewDao()
	s.Lock.RLock()
	songs, err := s.QueryMultiSongs("流行指数榜", model.QQMusic)
	s.Lock.RUnlock()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"songs": songs})
}
