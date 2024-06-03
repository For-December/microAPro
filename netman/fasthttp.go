package netman

import (
	"github.com/valyala/fasthttp"
	"microAPro/constant/config"
	"microAPro/utils/logger"
	"time"
)

var reqClient = &fasthttp.Client{
	// 读超时时间,不设置read超时,可能会造成连接复用失效
	ReadTimeout: time.Second * 5,
	// 写超时时间
	WriteTimeout: time.Second * 5,
	// 5秒后，关闭空闲的活动连接
	MaxIdleConnDuration: time.Second * 5,
	// 当true时,从请求中去掉User-Agent标头
	NoDefaultUserAgentHeader: true,
	// 当true时，header中的key按照原样传输，默认会根据标准化转化
	DisableHeaderNamesNormalizing: true,
	//当true时,路径按原样传输，默认会根据标准化转化
	DisablePathNormalizing: true,
	// 配置http代理
	//Dial: fasthttpproxy.FasthttpHTTPDialer("localhost:7891"),
}

func getFastReqClient() *fasthttp.Client {
	return reqClient
}

func FastGet(url string) []byte {
	// 获取客户端
	client := getFastReqClient()

	_, body, err := client.Get(nil, url)
	if err != nil {
		panic(err)
		return nil
	}

	// 读取结果
	return body
}

func FastPost(url string, data string) []byte {
	// 获取客户端
	client := getFastReqClient()

	// 从请求池中分别获取一个request、response实例
	req, resp := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	// 回收到请求池
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	req.SetRequestURI(url)

	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentType("application/json")

	req.Header.Set("Authorization", "Bearer "+config.EnvCfg.DouBaoApiKey)
	req.SetBodyString(data)

	if err := client.Do(req, resp); err != nil {
		logger.Error("req err ", err)
		return nil
	}

	return resp.Body()
}
