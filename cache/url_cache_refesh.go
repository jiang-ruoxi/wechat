package cache

import (
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
	"log"
	"time"
)

//redisCacheForcedRefresh redis中的路由缓存key强制更新
func redisCacheForcedRefresh(cacheKey string, c *gin.Context, cfg *Config, cacheDuration time.Duration, cacheStore persist.CacheStore, sfGroup singleflight.Group, respCache *ResponseCache) {
	log.Println("进入到了协程")
	//cfg.missCacheCallback(c)
	//replyWithCacheNext(c, cfg, respCache)
	//log.Printf("respCache:%+v \n", respCache)

	//cacheWriter := &responseCacheWriter{
	//	ResponseWriter: c.Writer,
	//}
	//c.Writer = cacheWriter



	//log.Printf("c.Writer:%+v \n", c.Writer)
	//
	//
	//
	//log.Printf("里面面c.cacheWriter:%+v \n", cacheWriter)
}
