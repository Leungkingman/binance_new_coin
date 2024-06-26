package binance

import (
	"binanceNewCoin/server/global"
	"binanceNewCoin/server/logger"
	"binanceNewCoin/server/model"
	"binanceNewCoin/server/tool"
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

var article model.Article
var news model.News

const announcementUrl = "https://www.bGetCurrentNewTitleinancezh.com/en/support/announcement"
const mobileAnnouncementUrl = "https://www.binancezh.sh/en/support/announcement"
const listingUrl = "https://www.binancezh.com/en/support/announcement/c-48?navId=48"
const newListingUrl = "https://www.binancezh.top/en/support/announcement/c-48?navId=48"
const apiUrl = "/bapi/composite/v1/public/cms/article/list/query?type=1&pageNo=1&pageSize=4"
const newApiUrl = "/bapi/composite/v1/public/marketing/app/findAllProAppDynamicMenuConfig?page=1&recommendCurrency=CNY&recommendType=1&simpleStyle=0&rows=3&type=2&userCurrency=USD"

var scriptReg = regexp.MustCompile(`<script id="__APP_DATA" type="application/json">(?s:(.*?))</script>`)
var coinReg = regexp.MustCompile(`\((.*?)\)`)
var tags = [5]string{"65uLk7Nkp0pRea2ps4Aj", "Z73vs0d3eu67rDqMCbSH", "tfph2mpTPAuwxbiMHoQc", "pch5D9lsORjgObhyjdSK", "tvXLzOPgJFiMa8Omltoo"}
var numberUpperChars = [36]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
var numberChars = [36]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
var allChars = [62]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func GetCurrentNewTitle() string {
	global.DB_ENGINE.Table("bm_articles").First(&article)
	return article.Content
}

func GetCurrentNews() string {
	global.DB_ENGINE.Table("bm_news_title").First(&news)
	return news.Title
}

func UpdateNewTitleFromNewApi(title, title_type, publish_time string) bool {
	isAddNew := false
	var articles []model.Notice
	global.DB_ENGINE.Table("bm_articles").Where("content = ?", title).Find(&articles)
	if len(articles) == 0 {
		create_time := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
		db_article := model.Notice{
			Content:      title,
			Create_time:  create_time,
			Title_type:   title_type,
			Publish_time: publish_time,
		}
		global.DB_ENGINE.Table("bm_articles").Create(&db_article)
		isAddNew = true
	}
	return isAddNew
}

func UpdateNewTitle(title string) bool {
	isAddNew := false
	var articles []model.Article
	global.DB_ENGINE.Table("bm_articles").Where("content = ?", title).Find(&articles)
	if len(articles) == 0 {
		create_time := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
		db_article := model.Article{
			Content:     title,
			Create_time: create_time,
		}
		global.DB_ENGINE.Table("bm_articles").Create(&db_article)
		isAddNew = true
	}
	return isAddNew
}

func UpdcateNews(title string) {
	global.DB_ENGINE.Table("bm_news_title").First(&news)
	news.Title = title
	global.DB_ENGINE.Table("bm_news_title").Save(&news)
}

func FetchAnnouncement() (string, error) {
	firstTitle := ""
	resp, err := http.Get(announcementUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("币安爬虫失败 status code: " + string(resp.StatusCode))
	}

	utf8Reader := transform.NewReader(resp.Body, determinEncoding(resp.Body).NewDecoder())
	webSite, readErr := ioutil.ReadAll(utf8Reader)
	if readErr != nil {
		return "", readErr
	}
	websiteJson := scriptReg.FindAllStringSubmatch(string(webSite), -1)
	//articles := gjson.Get(websiteJson[0][1], "routeProps.*.navDataResource.0.articles")
	articles := gjson.Get(websiteJson[0][1], "routeProps.*.catalogs.0.articles")
	//fmt.Printf("articles = %v \n", articles)

	if len(articles.Array()) == 0 {
		return "", errors.New("announce json 格式出错")
	}

	firstTitle = articles.Array()[0].Get("title").String()
	return firstTitle, nil
}

func FetchNewListing() (bool, int, string, error) {
	firstTitle := ""
	completeUrl := newListingUrl + "&t=" + tool.Int64ToString(time.Now().Unix())
	resp, err := http.Get(completeUrl)
	if err != nil {
		return false, resp.StatusCode, "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return false, resp.StatusCode, "", errors.New("币安爬虫失败 status code: " + string(resp.StatusCode))
	}
	utf8Reader := transform.NewReader(resp.Body, determinEncoding(resp.Body).NewDecoder())
	webSite, readErr := ioutil.ReadAll(utf8Reader)
	if readErr != nil {
		return false, resp.StatusCode, "", readErr
	}
	websiteJson := scriptReg.FindAllStringSubmatch(string(webSite), -1)
	//articles := gjson.Get(websiteJson[0][1], "routeProps.*.navDataResource.0.articles")
	articles := gjson.Get(websiteJson[0][1], "routeProps.*.catalogs.0.articles")
	if len(articles.Array()) == 0 {
		return false, resp.StatusCode, "", errors.New("newListing json 格式出错")
	}
	firstTitle = articles.Array()[0].Get("title").String()
	return true, resp.StatusCode, firstTitle, nil
}

func FetchFromApi(index int, domain string) (bool, int, string, int64, error) {
	success := false
	var title string
	var create_time int64
	req := BuildNewApiReq(domain)
	client := &http.Client{}
	resp, httpErr := client.Do(req)
	//fmt.Printf("resp = %V \n", resp)
	//fmt.Printf("================================== \n")
	if httpErr != nil {
		return success, -1, "", 0, httpErr
	}
	//logger.Addlog("请求前时间：" + strconv.FormatInt(time.Now().UnixNano() / 1000000, 10), "normal", "mobileApi")
	defer resp.Body.Close()
	//fmt.Printf("RemoteAddr = %v \n", req.RemoteAddr)
	//fmt.Printf("================================== \n")
	httpResult, readHttpErr := ioutil.ReadAll(resp.Body)
	if readHttpErr != nil {
		return success, -2, "", 0, readHttpErr
	}
	var announceApiRes model.AnnounceApiResponse
	_ = json.Unmarshal(httpResult, &announceApiRes)
	if announceApiRes.Success {
		success = true
		articles := gjson.Get(string(httpResult), "data.catalogs.0.articles")
		title = articles.Array()[index].Get("title").String()
		create_time = articles.Array()[index].Get("releaseDate").Int()
	} else {
		var err error
		create_time = 0
		if resp.StatusCode == http.StatusOK { // 查看 200 状态码返回的是什么
			err = errors.New(string(httpResult))
		} else {
			title = "statusCode = " + strconv.Itoa(resp.StatusCode) + "；Code = " + strconv.FormatInt(announceApiRes.Code, 10) + "；message = " + announceApiRes.Message + "；messageDetail = " + announceApiRes.MessageDetail
			err = errors.New(title)
		}
		return success, resp.StatusCode, title, create_time, err
	}
	return success, resp.StatusCode, title, create_time, nil
}

func FetchFromNewApi(domain string) (bool, int, string, string, string, error) {
	success := false
	var title, noticeType, publishTime string
	req := BuildNewApiReq(domain)
	client := &http.Client{}
	resp, httpErr := client.Do(req)
	if httpErr != nil {
		return success, -1, "", "", "", httpErr
	}
	defer resp.Body.Close()
	httpResult, readHttpErr := ioutil.ReadAll(resp.Body)
	if readHttpErr != nil {
		return success, -2, "", "", "", readHttpErr
	}
	var announceApiRes model.AnnounceApiResponse
	_ = json.Unmarshal(httpResult, &announceApiRes)
	if announceApiRes.Success {
		notices := gjson.Get(string(httpResult), "data.notices")
		noticesData := notices.Array()
		if len(noticesData) > 0 {
			success = true
			title = notices.Array()[0].Get("title").String()
			noticeType = notices.Array()[0].Get("type").String()
			publishTime = notices.Array()[0].Get("time").String()
		} else {
			logger.Addlog("空数据", "normal", "operate")
			logger.Addlog(string(httpResult), "normal", "operate")
		}
	} else {
		title = "statusCode = " + strconv.Itoa(resp.StatusCode) + "；Code = " + strconv.FormatInt(announceApiRes.Code, 10) + "；message = " + announceApiRes.Message + "；messageDetail = " + announceApiRes.MessageDetail
		err := errors.New(title)
		return success, resp.StatusCode, title, "", "", err
	}
	return success, resp.StatusCode, title, noticeType, publishTime, nil
}

func BuildNewApiReq(domain string) (*http.Request) {
	url := "https://" + domain + apiUrl
	completeApi := url + "&timestamp=" + tool.Int64ToString(time.Now().UnixNano()/1e6)
	//logger.Addlog("completeApi = " + completeApi, "normal", "operate")
	req, _ := http.NewRequest("GET", completeApi, nil)
	rand.Seed(time.Now().UnixNano() / 1e6)
	index := rand.Intn(5)
	tag := tags[index]
	req.Header.Set("mclient-x-tag", tag)
	req.Header.Set("lang", "en")
	req.Header.Set("user-agent", "BNC/2.40.1 (build 1; iOS 14.6.0) Alamofire/4.9.0")
	req.Header.Set("versioncode", "2.40.1")
	req.Header.Set("bnc-app-channel", "appstore")
	req.Header.Set("clienttype", "ios")
	req.Header.Set("bnc-time-zone", "Asia/Shanghai")
	req.Header.Set("isnight", "false")
	req.Header.Set("bnc-app-mode", "pro")
	deviceInfo := GetRandomStr(648, 3)
	req.Header.Set("device-info", deviceInfo)
	uuid1 := GetRandomStr(8, 1)
	uuid2 := GetRandomStr(4, 1)
	uuid3 := GetRandomStr(4, 1)
	uuid4 := GetRandomStr(4, 1)
	uuid5 := GetRandomStr(12, 1)
	completeUUID := uuid1 + "-" + uuid2 + "-" + uuid3 + "-" + uuid4 + "-" + uuid5
	req.Header.Set("bnc-uuid", completeUUID)
	videoId := GetRandomStr(40, 2)
	req.Header.Set("fvideo-id", videoId)
	traceId1 := GetRandomStr(8, 1)
	traceId2 := GetRandomStr(4, 1)
	traceId3 := GetRandomStr(4, 1)
	traceId4 := GetRandomStr(4, 1)
	traceId5 := GetRandomStr(12, 1)
	completeTraceId := traceId1 + "-" + traceId2 + "-" + traceId3 + "-" + traceId4 + "-" + traceId5
	req.Header.Set("x-trace-id", completeTraceId)
	cookie := "cid=" + GetRandomStr(8, 3)
	req.Header.Set("cookie", cookie)
	return req
}

func GetRandomStr(length int, charType int) string {
	str := ""
	rand.Seed(time.Now().UnixNano() / 1e6)
	if charType == 1 {
		for i := 0; i < length; i++ {
			index := rand.Intn(len(numberUpperChars))
			char := numberUpperChars[index]
			str = str + char
		}
	} else if charType == 2 {
		for i := 0; i < length; i++ {
			index := rand.Intn(len(numberChars))
			char := numberChars[index]
			str = str + char
		}
	} else if charType == 3 {
		for i := 0; i < length; i++ {
			index := rand.Intn(len(allChars))
			char := allChars[index]
			str = str + char
		}
	}
	return str
}

func TestFetchNewApi() {
	req, _ := http.NewRequest("GET", newApiUrl, nil)
	//req.Header.Set("lang", "en")
	req.Header.Set("lang", "en")
	// pch5D9lsORjgObhyjdSK、tfph2mpTPAuwxbiMHoQc、tvXLzOPgJFiMa8Omltoo
	req.Header.Set("mclient-x-tag", "pch5D9lsORjgObhyjdSK")
	req.Header.Set("user-agent", "Binance/2.30.3 (com.czzhao.binance; build:0; iOS 14.6.0) Alamofire/2.30.3")
	req.Header.Set("versioncode", "2.30.3")
	client := &http.Client{}
	resp, httpErr := client.Do(req)
	if httpErr != nil {
		fmt.Printf("httpErr = %V \n", httpErr.Error())
	} else {
		defer resp.Body.Close()
		httpResult, _ := ioutil.ReadAll(resp.Body)
		logger.Addlog(string(httpResult), "normal", "operate")
		fmt.Printf("================================== \n")
	}
}

func TestFetchFromApi(index int) (bool, int, string, error) {
	success := false
	var title string
	logger.Addlog("请求前时间："+strconv.FormatInt(time.Now().UnixNano()/1000000, 10), "normal", "mobileApi")
	req, _ := http.NewRequest("GET", apiUrl, nil)
	client := &http.Client{}
	resp, httpErr := client.Do(req)
	logger.Addlog("请求完成时间："+strconv.FormatInt(time.Now().UnixNano()/1000000, 10), "normal", "mobileApi")
	if httpErr != nil {
		return success, -1, "", httpErr
	}
	defer resp.Body.Close()
	httpResult, readHttpErr := ioutil.ReadAll(resp.Body)
	if readHttpErr != nil {
		return success, -2, "", readHttpErr
	}
	logger.Addlog("读取请求时间："+strconv.FormatInt(time.Now().UnixNano()/1000000, 10), "normal", "mobileApi")
	var announceApiRes model.AnnounceApiResponse
	_ = json.Unmarshal(httpResult, &announceApiRes)
	if announceApiRes.Success {
		success = true
		articles := gjson.Get(string(httpResult), "data.catalogs.0.articles")
		title = articles.Array()[index].Get("title").String()
	} else {
		title = "statusCode = " + strconv.Itoa(resp.StatusCode) + "；Code = " + strconv.FormatInt(announceApiRes.Code, 10) + "；message = " + announceApiRes.Message + "；messageDetail = " + announceApiRes.MessageDetail
		err := errors.New(title)
		return success, resp.StatusCode, title, err
	}
	return success, resp.StatusCode, title, nil
}

func determinEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
