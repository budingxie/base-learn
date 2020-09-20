package matchers

import (
	"base-learn/search"
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

type (
	item struct {
		XMLName     xml.Name `xml:"item"`
		PubData     string   `xml:"pubData"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}

	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}

	channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubData        string   `xml:"pubData"`
		LastBuildData  string   `xml:"lastBuildData"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          image    `xml:"image"`
		Item           []item   `xml:"item"`
	}
	// xml中对应的文档对象
	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)

type rssMatcher struct {
}

func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

func (m rssMatcher) retrieve(feed *search.Feed) (*rssDocument, error) {
	// 判断数据源的url是否为空
	if feed.URL == "" {
		return nil, errors.New("No rss feed URL provided")
	}
	// 获取到url对的xml数据
	resp, err := http.Get(feed.URL)
	if err != nil {
		return nil, err
	}
	// 关闭资源
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Http Response Error %d\n", resp.StatusCode)
	}
	// 创建xml文档对象
	var document rssDocument
	//解析数据，赋值到document对象
	err = xml.NewDecoder(resp.Body).Decode(&document)

	return &document, err
}

// rss类型的数据源，实现搜索接口
func (m rssMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
	// 创建返回值
	var results []*search.Result
	// 打印正在解析的数据源
	log.Printf("Search Feed Type[%s] Site[%s] For Uri[%s]\n", feed.Type, feed.Name, feed.URL)
	// 通过解析器，解析数据源对象，并获取到xml文档对象
	document, err := m.retrieve(feed)
	if err != nil {
		return nil, err
	}
	// 搜索rssDocument文档中
	// channel标签
	// 下的item标签
	// 里面的title标签
	// 的内容和searchTerm匹配的
	// 然后封装成Result对象，并添加到results中
	for _, channelItem := range document.Channel.Item {
		matched, err := regexp.MatchString(searchTerm, channelItem.Title)
		if err != nil {
			return nil, err
		}
		if matched {
			results = append(results, &search.Result{
				Field:   "Title",
				Content: channelItem.Title,
			})
		}

		matched, err = regexp.MatchString(searchTerm, channelItem.Description)
		if err != nil {
			return nil, err
		}
		if matched {
			results = append(results, &search.Result{
				Field:   "Description",
				Content: channelItem.Description,
			})
		}
	}
	return results, nil
}
