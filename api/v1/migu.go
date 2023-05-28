package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yushengguo557/music-ranking-service/dao"
	"github.com/yushengguo557/music-ranking-service/model"
)

func GetMiguNewSong(c *gin.Context) {
	s := dao.NewDao()
	s.Lock.RLock()
	songs, err := s.QueryMultiSongs("尖叫新歌榜", model.MiguMusic)
	s.Lock.RUnlock()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"songs": songs})
}

func GetMiguHotSong(c *gin.Context) {
	s := dao.NewDao()
	s.Lock.RLock()
	songs, err := s.QueryMultiSongs("尖叫热歌榜", model.MiguMusic)
	s.Lock.RUnlock()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"songs": songs})
}

func GetMiguOriginal(c *gin.Context) {
	s := dao.NewDao()
	s.Lock.RLock()
	songs, err := s.QueryMultiSongs("尖叫原创榜", model.MiguMusic)
	s.Lock.RUnlock()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"songs": songs})
}
