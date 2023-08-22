package router

import (
	"fmt"
	"github.com/chenyahui/gin-cache"
	"github.com/gin-gonic/gin"
	"log"
	"time"
	"wechat/global"
)

func routerCache(sec int64) (res gin.HandlerFunc) {
	fmt.Printf("sec:%v", sec)
	return cache.CacheByRequestURI(global.GVA_HTTP_CACHE, time.Duration(sec)*time.Second)
}

func routerCacheWithOption(sec int64) (res gin.HandlerFunc) {
	fmt.Printf("sec:%v", sec)
	return cache.CacheByRequestURI(
		global.GVA_HTTP_CACHE,
		time.Duration(sec)*time.Second,
		//当生成缓存键时，IgnoreQueryOrder将忽略url中的查询顺序。该选项仅在CacheByRequestURI函数中生效
		cache.IgnoreQueryOrder(),
		//WithPrefixKey将为键添加前缀
		cache.WithPrefixKey("jiang"),
		//当缓存命中时，将调用WithOnHitCache。
		cache.WithOnHitCache(func(c *gin.Context) {
			log.Printf("%+v \n","这里是命中的缓存")
		}),
		//当缓存丢失时将调用WithOnMissCache
		cache.WithOnMissCache(func(c *gin.Context) {
			log.Printf("%+v \n","没有命中的缓存")
		}),
		//withbeforeereplywithcache将在回复缓存之前被调用。
		cache.WithBeforeReplyWithCache(func(c *gin.Context, cache *cache.ResponseCache) {
			log.Printf("%+v \n","WithBeforeReplyWithCache")
		}),
		//WithOnShareSingleFlight将在共享singleflight结果时调用
		cache.WithOnShareSingleFlight(func(c *gin.Context) {
			log.Printf("%+v \n","WithOnShareSingleFlight")
		}),
		//WithCacheStrategyByRequest根据每个请求设置自定义策略
		cache.WithCacheStrategyByRequest(func(c *gin.Context) (bool, cache.Strategy) {
			// 在这里根据请求的内容来判断是否要使用缓存
			shouldCache := true // 根据实际情况设置
			var cacheDuration time.Duration = time.Minute * 5 // 根据实际情况设置
			cacheKey := "custom_key" // 根据实际情况设置

			// 构造缓存策略对象
			strategy := cache.Strategy{
				CacheKey:      cacheKey,
				CacheDuration: cacheDuration,
			}

			return shouldCache, strategy
		}),
	)
}
