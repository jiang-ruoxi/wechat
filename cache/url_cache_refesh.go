package cache

import (
	"github.com/chenyahui/gin-cache/persist"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

//redisCacheForcedRefresh redis中的路由缓存key强制更新
func redisCacheForcedRefresh(cacheKey string,cfg *Config, cacheDuration time.Duration, cacheStore persist.CacheStore, respCache *ResponseCache) {
	//log.Printf("%+v \n", "进入的协程")
	url := "http://localhost:8089" + cacheKey + "&no_cache_ext=" + strconv.FormatInt(time.Now().Unix(), 10)
	httpGetContent, httpCode := sendHttpGet(url)
	//log.Printf("url:%+v  code: %+v httpGet:%+v\n", url, httpCode, httpGetContent)
	//log.Printf("url:%+v\n", url)
	//log.Printf("code:%+v\n", httpCode)
	//log.Printf("httpGetContent:%+v\n", httpGetContent)
	//log.Printf("len(httpGetContent):%+v\n", len(httpGetContent))
	//log.Printf("code:%+v \n", httpCode)
	// only cache 2xx response and httpGetContent length gt 0
	if httpCode == http.StatusOK && len(httpGetContent) > 0 {
		var respCacheNew = respCache
		respCacheNew.Data = httpGetContent
		if err := cacheStore.Set(cacheKey, respCacheNew, cacheDuration); err != nil {
			cfg.logger.Errorf("set cache key error: %s, cache key: %s", err, cacheKey)
		}
	}
}

func sendHttpGet(url string) (bodyContent []byte, httpStatusCode int) {
	// 发起GET请求
	response, err := http.Get(url)
	if err != nil {
		return
	}
	defer response.Body.Close() // 确保在函数退出前关闭响应体

	httpStatusCode = response.StatusCode
	// 读取响应体内容
	bodyContent, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	return
}
