package global

import "github.com/yushengguo557/music-ranking-service/log"

var MRS_LOG *log.Logger

func InitLog() {
	var err error
	MRS_LOG, err = log.NewLogger()
	if err != nil {
		panic(err)
	}
}
