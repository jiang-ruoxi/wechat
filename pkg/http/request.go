package http

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
	"reflect"
	"time"
)

const (
	// SuccessServiceReq 成功
	SuccessServiceReq int = 100
	// ErrorServiceReq 请求异常
	ErrorServiceReq       int = 101
	ErrorServiceUnmarshal int = 102
)

func Request(url, method string, headers map[string]string, args map[string]string, js string) ([]byte, bool) {
	params := map[string]interface{}{
		"params": args,
		"msg":    "request请求信息",
		"url":    url,
		// "headers": headers,
		"js":     js,
		"method": method,
	}
	paramsJSON, _ := json.Marshal(params)
	log.Printf(fmt.Sprintf("request start params:%v", string(paramsJSON)))
	if method == "POST" {
		return p(url, headers, args, js)
	} else if method == "POST1" {
		return p1(url, headers, args, js)
	}
	return g(url, headers, args, js)
}

func p1(url string, headers map[string]string, args map[string]string, js string) ([]byte, bool) {
	params := map[string]interface{}{
		"params": args,
		"msg":    "request请求信息",
		"url":    url,
		// "headers": headers,
		"js":     js,
		"method": "POST",
	}
	var paramsJSON []byte

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req) // 用完需要释放资源
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentTypeBytes([]byte("application/json"))

	// 请求地址
	req.SetRequestURI(url)
	// 请求参数
	if args != nil && len(args) > 0 {
		for k, v := range args {
			req.PostArgs().Add(k, v)
		}
	}
	// header
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	// body
	req.SetBodyString(js)

	statusCode, body, err := fasthttp.Post([]byte{}, string(req.URI().FullURI()), req.PostArgs())
	if err == nil {
		if statusCode == http.StatusOK {
			params["msg"] = "request请求信息-返回"
			params["res"] = string(body)
			params["statusCode"] = statusCode
			paramsJSON, _ = json.Marshal(params)
			log.Printf(fmt.Sprintf("request end params:%v", string(paramsJSON)))
			return body, true
		} else {
			params["msg"] = "请求接口错误"
			params["err"] = err
			params["info"] = fmt.Sprintf("ERR invalid HTTP response code: %d\n", statusCode)
			paramsJSON, _ = json.Marshal(params)
			log.Printf(fmt.Sprintf("request end params:%v", string(paramsJSON)))
			return nil, false
		}

	} else {
		errName, known := httpConnError(err)
		info := ""
		if known {
			info = fmt.Sprintf("WARN conn error: %v\n", errName)
		} else {
			info = fmt.Sprintf("ERR conn failure: %v %v\n", errName, err)
		}
		params["msg"] = "请求接口错误"
		params["err"] = err
		params["info"] = info
		paramsJSON, _ = json.Marshal(params)
		log.Printf(fmt.Sprintf("request end params:%v", string(paramsJSON)))
		return []byte{}, false
	}
}

func p(url string, headers map[string]string, args map[string]string, js string) ([]byte, bool) {
	params := map[string]interface{}{
		"params": args,
		"msg":    "request请求信息",
		"url":    url,
		// "headers": headers,
		"js":     js,
		"method": "POST",
	}
	var paramsJSON []byte

	reqTimeout := time.Duration(10000) * time.Millisecond

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req) // 用完需要释放资源
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentTypeBytes([]byte("application/json"))

	// 请求地址
	req.SetRequestURI(url)
	// 请求参数
	if args != nil && len(args) > 0 {
		for k, v := range args {
			req.PostArgs().Add(k, v)
		}
	}
	// header
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	// body
	req.SetBodyString(js)

	c := &fasthttp.Client{}
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp) // 用完需要释放资源

	err := c.DoTimeout(req, resp, reqTimeout)
	if err == nil {
		statusCode := resp.StatusCode()
		body := resp.Body()
		if statusCode == http.StatusOK {
			params["msg"] = "request请求信息-返回"
			params["res"] = string(body)
			params["statusCode"] = statusCode
			paramsJSON, _ = json.Marshal(params)
			log.Printf(fmt.Sprintf("request end params:%v", string(paramsJSON)))
			return body, true
		} else {
			params["msg"] = "请求接口错误"
			params["err"] = err
			params["info"] = fmt.Sprintf("ERR invalid HTTP response code: %d\n", statusCode)
			paramsJSON, _ = json.Marshal(params)
			log.Printf(fmt.Sprintf("request end params:%v", string(paramsJSON)))
			return nil, false
		}

	} else {
		errName, known := httpConnError(err)
		info := ""
		if known {
			info = fmt.Sprintf("WARN conn error: %v\n", errName)
		} else {
			info = fmt.Sprintf("ERR conn failure: %v %v\n", errName, err)
		}
		params["msg"] = "请求接口错误"
		params["err"] = err
		params["info"] = info
		paramsJSON, _ = json.Marshal(params)
		log.Printf(fmt.Sprintf("request end params:%v", string(paramsJSON)))
		return []byte{}, false
	}
}

func g(url string, headers map[string]string, args map[string]string, js string) ([]byte, bool) {
	params := map[string]interface{}{
		"params": args,
		"msg":    "request请求信息",
		"url":    url,
		// "headers": headers,
		"js":     js,
		"method": "GET",
	}
	var paramsJSON []byte

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req) // 用完需要释放资源
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodGet)

	if args != nil && len(args) > 0 {
		for k, v := range args {
			req.URI().QueryArgs().Add(k, v)
		}
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	requestBody := []byte(js)
	req.SetBody(requestBody)

	c := &fasthttp.Client{}
	resp := fasthttp.AcquireResponse()
	if err := c.Do(req, resp); err != nil {
		params["msg"] = "请求接口错误"
		params["err"] = err
		paramsJSON, _ = json.Marshal(params)
		log.Printf(fmt.Sprintf("request end params:%v", string(paramsJSON)))
		return []byte{}, false
	}
	body := resp.Body()
	defer fasthttp.ReleaseResponse(resp) // 用完需要释放资源
	params["msg"] = "request请求信息-返回"
	params["res"] = string(body)
	paramsJSON, _ = json.Marshal(params)
	log.Printf(fmt.Sprintf("request end params:%v", string(paramsJSON)))
	return body, true

}

func httpConnError(err error) (string, bool) {
	errName := ""
	known := false
	if err == fasthttp.ErrTimeout {
		errName = "timeout"
		known = true
	} else if err == fasthttp.ErrNoFreeConns {
		errName = "conn_limit"
		known = true
	} else if err == fasthttp.ErrConnectionClosed {
		errName = "conn_close"
		known = true
	} else {
		errName = reflect.TypeOf(err).String()
		if errName == "*net.OpError" {
			// Write and Read errors are not so often and in fact they just mean timeout problems
			errName = "timeout"
			known = true
		}
	}
	return errName, known
}
