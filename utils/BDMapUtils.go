package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// ConvertToBaiduCoords 将给定的JSON格式的字符串数据转换为百度地图坐标，并返回相应的JSON格式字符串。
func ConvertToBaiduCoords(jsonData string) (string, error) {
	// 定义请求的URL和参数
	apiUrl := "http://api.map.baidu.com/geoconv/v1/"
	params := url.Values{}
	params.Set("ak", "oRqbxhmwbE30ZgrbXzzUM19O6kyBjKQD") // 请替换为您的百度地图开发者AK

	// 解析JSON数据
	var coordinates []map[string]float64
	if err := json.Unmarshal([]byte(jsonData), &coordinates); err != nil {
		return "", fmt.Errorf("解析JSON出错：%v", err)
	}

	// 构造coords参数
	var coords string
	for _, coordinate := range coordinates {
		coords += fmt.Sprintf("%.4f,%.4f;", coordinate["lng"], coordinate["lat"])
	}
	coords = coords[:len(coords)-1] // 去掉最后一个分号

	// 添加coords参数
	params.Set("coords", coords)

	// 发送HTTP请求
	resp, err := http.Get(apiUrl + "?" + params.Encode())
	if err != nil {
		return "", fmt.Errorf("请求出错：%v", err)
	}
	defer resp.Body.Close()

	// 解析响应结果
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应出错：%v", err)
	}
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("解析JSON出错：%v", err)
	}
	fmt.Println(result)

	// 构造返回的JSON格式字符串
	var convertedCoords []map[string]float64
	for _, coordinate := range result["result"].([]interface{}) {
		convertedCoordinate := map[string]float64{
			"lng": coordinate.(map[string]interface{})["x"].(float64),
			"lat": coordinate.(map[string]interface{})["y"].(float64),
		}
		convertedCoords = append(convertedCoords, convertedCoordinate)
	}
	convertedData, err := json.Marshal(convertedCoords)
	if err != nil {
		return "", fmt.Errorf("JSON编码出错：%v", err)
	}
	fmt.Println(string(convertedData))
	return string(convertedData), nil
}
