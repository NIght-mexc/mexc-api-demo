package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"demo/config"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
)

//公共get请求
func PublicGet(urlStr string, jsonParams string) interface{} {
	var path string
	if jsonParams == "" {
		path = urlStr
	} else {
		strParams := JsonToParamStr(jsonParams)
		path = urlStr + "?" + strParams
		fmt.Println("路径:", path)
	}
	//创建请求
	client := resty.New()
	//发送请求
	resp, err := client.R().Get(path)

	if err != nil {
		log.Fatal("请求报错：", err)
	}

	// fmt.Println("Response Info:", resp)
	return resp
}

//私有get请求
func PrivateGet(urlStr string, jsonParams string) interface{} {
	var path string
	timestamp := time.Now().UnixNano() / 1e6
	fmt.Println(timestamp)
	if jsonParams == "" {
		message := fmt.Sprintf("timestamp=%d", timestamp)
		sign := ComputeHmac256(message, config.SEC_KEY)
		path = fmt.Sprintf("%s?timestamp=%d&signature=%s", urlStr, timestamp, sign)
		fmt.Println("message:", message)
		fmt.Println("sign:", sign)
		fmt.Println("path:", path)
	} else {
		strParams := JsonToParamStr(jsonParams)
		message := fmt.Sprintf("%s&timestamp=%d", strParams, timestamp)
		sign := ComputeHmac256(message, config.SEC_KEY)
		path = fmt.Sprintf("%s?%s&timestamp=%d&signature=%s", urlStr, strParams, timestamp, sign)
		fmt.Println("message:", ParamsEncode(message))
		fmt.Println("sign:", sign)
		fmt.Println("path:", path)
	}
	//创建请求
	client := resty.New()
	//发送请求
	resp, err := client.R().SetHeaders(map[string]string{
		"X-MEXC-APIKEY": config.API_KEY,
		"Content-Type":  "application/json",
	}).Get(path)

	if err != nil {
		log.Fatal("请求报错：", err)
	}

	// fmt.Println("Response Info:", resp)
	return resp
}

//私有post请求
func PrivatePost(urlStr string, jsonParams string) interface{} {
	var path string
	timestamp := time.Now().UnixNano() / 1e6
	fmt.Println(timestamp)
	if jsonParams == "" {
		message := fmt.Sprintf("timestamp=%d", timestamp)
		sign := ComputeHmac256(message, config.SEC_KEY)
		path = fmt.Sprintf("%s?timestamp=%d&signature=%s", urlStr, timestamp, sign)
		fmt.Println("message:", message)
		fmt.Println("sign:", sign)
		fmt.Println("path:", path)
	} else {
		strParams := JsonToParamStr(jsonParams)
		message := fmt.Sprintf("%s&timestamp=%d", strParams, timestamp)
		sign := ComputeHmac256(message, config.SEC_KEY)
		path = fmt.Sprintf("%s?%s&timestamp=%d&signature=%s", urlStr, strParams, timestamp, sign)
		fmt.Println("message:", ParamsEncode(message))
		fmt.Println("sign:", sign)
		fmt.Println("path:", path)
	}
	//创建请求
	client := resty.New()
	//发送请求
	resp, err := client.R().SetHeaders(map[string]string{
		"X-MEXC-APIKEY": config.API_KEY,
		"Content-Type":  "application/json",
	}).Post(path)

	if err != nil {
		log.Fatal("请求报错：", err)
	}

	// fmt.Println("Response Info:", resp)
	return resp
}

//私有delete请求
func PrivateDelete(urlStr string, jsonParams string) interface{} {
	var path string
	timestamp := time.Now().UnixNano() / 1e6
	fmt.Println(timestamp)
	if jsonParams == "" {
		message := fmt.Sprintf("timestamp=%d", timestamp)
		sign := ComputeHmac256(message, config.SEC_KEY)
		path = fmt.Sprintf("%s?timestamp=%d&signature=%s", urlStr, timestamp, sign)
		fmt.Println("message:", message)
		fmt.Println("sign:", sign)
		fmt.Println("path:", path)
	} else {
		strParams := JsonToParamStr(jsonParams)
		message := fmt.Sprintf("%s&timestamp=%d", strParams, timestamp)
		sign := ComputeHmac256(message, config.SEC_KEY)
		path = fmt.Sprintf("%s?%s&timestamp=%d&signature=%s", urlStr, strParams, timestamp, sign)
		fmt.Println("message:", ParamsEncode(message))
		fmt.Println("sign:", sign)
		fmt.Println("path:", path)
	}
	//创建请求
	client := resty.New()
	//发送请求
	resp, err := client.R().SetHeaders(map[string]string{
		"X-MEXC-APIKEY": config.API_KEY,
		"Content-Type":  "application/json",
	}).Delete(path)

	if err != nil {
		log.Fatal("请求报错：", err)
	}

	// fmt.Println("Response Info:", resp)
	return resp
}

//私有put请求
func PrivatePut(urlStr string, jsonParams string) interface{} {
	var path string
	timestamp := time.Now().UnixNano() / 1e6
	fmt.Println(timestamp)
	if jsonParams == "" {
		message := fmt.Sprintf("timestamp=%d", timestamp)
		sign := ComputeHmac256(message, config.SEC_KEY)
		path = fmt.Sprintf("%s?timestamp=%d&signature=%s", urlStr, timestamp, sign)
		fmt.Println("message:", message)
		fmt.Println("sign:", sign)
		fmt.Println("path:", path)
	} else {
		strParams := JsonToParamStr(jsonParams)
		message := fmt.Sprintf("%s&timestamp=%d", strParams, timestamp)
		sign := ComputeHmac256(message, config.SEC_KEY)
		path = fmt.Sprintf("%s?%s&timestamp=%d&signature=%s", urlStr, strParams, timestamp, sign)
		fmt.Println("message:", ParamsEncode(message))
		fmt.Println("sign:", sign)
		fmt.Println("path:", path)
	}
	//创建请求
	client := resty.New()
	//发送请求
	resp, err := client.R().SetHeaders(map[string]string{
		"X-MEXC-APIKEY": config.API_KEY,
		"Content-Type":  "application/json",
	}).Put(path)

	if err != nil {
		log.Fatal("请求报错：", err)
	}

	// fmt.Println("Response Info:", resp)
	return resp
}

//格式化参数字符串
func JsonToParamStr(jsonParams string) string {
	//转化json参数->参数字符串
	var paramsarr []string
	var arritem string
	m := make(map[string]string)
	err := json.Unmarshal([]byte(jsonParams), &m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("map:%v\n", m)
	i := 0
	for key, value := range m {

		arritem = fmt.Sprintf("%s=%s", key, value)
		paramsarr = append(paramsarr, arritem)
		i++
		fmt.Println("遍历：", i, "总共", len(m))
		if i > len(m) {
			break
		}
	}
	paramsstr := strings.Join(paramsarr, "&")
	fmt.Println("参数字符串：", paramsstr)
	return paramsstr
}

//urlencode
func ParamsEncode(paramStr string) string {
	return url.QueryEscape(paramStr)
}

//加密
func ComputeHmac256(Message string, sec_key string) string {
	key := []byte(sec_key)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(Message))
	return hex.EncodeToString(h.Sum(nil))
}

func jsonToMap(jsonParams string) map[string]string {
	m := make(map[string]string)
	if jsonParams == "" {
		return m
	}
	if err := json.Unmarshal([]byte(jsonParams), &m); err != nil {
		fmt.Println(err)
	}
	return m
}

func substitutePathParams(pathTemplate string, jsonParams string) (string, string) {
	m := jsonToMap(jsonParams)
	result := pathTemplate
	re := regexp.MustCompile(`\{(\w+)\}`)
	for _, match := range re.FindAllStringSubmatch(pathTemplate, -1) {
		key := match[1]
		if val, ok := m[key]; ok {
			result = strings.Replace(result, "{"+key+"}", val, 1)
			delete(m, key)
		}
	}
	remaining, _ := json.Marshal(m)
	return result, string(remaining)
}

func sortedQueryString(jsonParams string) string {
	m := jsonToMap(jsonParams)
	if len(m) == 0 {
		return ""
	}
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	parts := make([]string, 0, len(keys))
	for _, key := range keys {
		parts = append(parts, fmt.Sprintf("%s=%s", key, m[key]))
	}
	return strings.Join(parts, "&")
}

func futuresHeaders(timestamp string, signTarget string) map[string]string {
	sign := ComputeHmac256(signTarget, config.SEC_KEY)
	return map[string]string{
		"ApiKey":        config.API_KEY,
		"Request-Time":  timestamp,
		"Signature":     sign,
		"Content-Type":  "application/json",
	}
}

// FuturesPublicGet public futures GET request
func FuturesPublicGet(pathTemplate string, jsonParams string) interface{} {
	path, remainParams := substitutePathParams(pathTemplate, jsonParams)
	urlStr := config.FUTURES_BASE_URL + path
	query := sortedQueryString(remainParams)
	if query != "" {
		urlStr = urlStr + "?" + query
	}
	fmt.Println("path:", urlStr)
	client := resty.New()
	resp, err := client.R().Get(urlStr)
	if err != nil {
		log.Fatal("请求报错：", err)
	}
	return resp
}

// FuturesPrivateGet signed futures GET request
func FuturesPrivateGet(pathTemplate string, jsonParams string) interface{} {
	path, remainParams := substitutePathParams(pathTemplate, jsonParams)
	urlStr := config.FUTURES_BASE_URL + path
	timestamp := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)
	paramStr := sortedQueryString(remainParams)
	signTarget := config.API_KEY + timestamp + paramStr
	headers := futuresHeaders(timestamp, signTarget)
	if paramStr != "" {
		urlStr = urlStr + "?" + paramStr
	}
	fmt.Println("path:", urlStr)
	client := resty.New()
	resp, err := client.R().SetHeaders(headers).Get(urlStr)
	if err != nil {
		log.Fatal("请求报错：", err)
	}
	return resp
}

// FuturesPrivatePost signed futures POST request
func FuturesPrivatePost(path string, jsonParams string) interface{} {
	urlStr := config.FUTURES_BASE_URL + path
	body := jsonParams
	timestamp := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)
	signTarget := config.API_KEY + timestamp + body
	headers := futuresHeaders(timestamp, signTarget)
	fmt.Println("path:", urlStr)
	client := resty.New()
	req := client.R().SetHeaders(headers)
	if body != "" {
		req = req.SetBody(body)
	}
	resp, err := req.Post(urlStr)
	if err != nil {
		log.Fatal("请求报错：", err)
	}
	return resp
}

// FuturesPrivateDelete signed futures DELETE request
func FuturesPrivateDelete(pathTemplate string, jsonParams string) interface{} {
	path, remainParams := substitutePathParams(pathTemplate, jsonParams)
	urlStr := config.FUTURES_BASE_URL + path
	timestamp := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)
	paramStr := sortedQueryString(remainParams)
	signTarget := config.API_KEY + timestamp + paramStr
	headers := futuresHeaders(timestamp, signTarget)
	if paramStr != "" {
		urlStr = urlStr + "?" + paramStr
	}
	fmt.Println("path:", urlStr)
	client := resty.New()
	resp, err := client.R().SetHeaders(headers).Delete(urlStr)
	if err != nil {
		log.Fatal("请求报错：", err)
	}
	return resp
}
