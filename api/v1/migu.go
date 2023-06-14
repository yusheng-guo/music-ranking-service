package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yushengguo557/music-ranking-service/dao"
	"github.com/yushengguo557/music-ranking-service/model"
	"github.com/yushengguo557/music-ranking-service/model/response"
)

func GetMiguNewSong(c *gin.Context) {
	s := dao.NewDao()
	s.Lock.RLock()
	songs, err := s.QueryMultiSongs("尖叫新歌榜", model.MiguMusic)
	s.Lock.RUnlock()
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	if len(songs) > 0 {
		response.OKWithDataAndMessage(c, songs, "success")
		return
	}
	response.OKWithMessage(c, "len(songs)=0")
}

func GetMiguHotSong(c *gin.Context) {
	s := dao.NewDao()
	s.Lock.RLock()
	songs, err := s.QueryMultiSongs("尖叫热歌榜", model.MiguMusic)
	s.Lock.RUnlock()
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	if len(songs) > 0 {
		response.OKWithDataAndMessage(c, songs, "success")
		return
	}
	response.OKWithMessage(c, "len(songs)=0")
}

func GetMiguOriginal(c *gin.Context) {
	s := dao.NewDao()
	s.Lock.RLock()
	songs, err := s.QueryMultiSongs("尖叫原创榜", model.MiguMusic)
	s.Lock.RUnlock()
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	if len(songs) > 0 {
		response.OKWithDataAndMessage(c, songs, "success")
		return
	}
	response.OKWithMessage(c, "len(songs)=0")
}
