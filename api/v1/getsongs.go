package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yushengguo557/music-ranking/dao"
)

func GetSongs(c *gin.Context) {
	s := dao.NewDao()
	s.Lock.RLock()
	songs, err := s.QueryMultiSongs()
	s.Lock.RUnlock()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"songs": songs})
}
