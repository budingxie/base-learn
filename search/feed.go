package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

// 封装数据源信息
type Feed struct {
	Name string `json:"site"`
	URL  string `json:"link"`
	Type string `json:"type"`
}

// 读取文件并反序列化
func RetrieveFeeds() ([]*Feed, error) {
	// 打开文件
	file, err := os.Open(dataFile)

	if err != nil {
		return nil, err
	}

	// 当打开文件正常时，关闭资源
	defer file.Close()

	var feeds []*Feed
	// 将json解码到feeds数组中
	err = json.NewDecoder(file).Decode(&feeds)
	return feeds, err
}
