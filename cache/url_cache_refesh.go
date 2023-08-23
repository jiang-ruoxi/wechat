package cache

import (
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
	"time"
)

//redisCacheForcedRefresh redis中的路由缓存key强制更新
func redisCacheForcedRefresh(cacheKey string, c *gin.Context, cfg *Config, cacheDuration time.Duration, cacheStore persist.CacheStore) {
	// use responseCacheWriter in order to record the response
	cacheWriter := &responseCacheWriter{
		ResponseWriter: c.Writer,
	}
	c.Writer = cacheWriter

	respCache := &ResponseCache{}
	respCache.fillWithCacheWriter(cacheWriter, cfg.withoutHeader)

	// only cache 2xx response
	if cacheWriter.Status() < 300 && cacheWriter.Status() >= 200 {
		if err := cacheStore.Set(cacheKey, respCache, cacheDuration); err != nil {
			cfg.logger.Errorf("set cache key error: %s, cache key: %s", err, cacheKey)
		}
	}
	c.Abort()
}
