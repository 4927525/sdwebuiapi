package helper

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"io"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"regexp"
	"sdwebuiapi/utils/jwt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var WeekDayMap = map[string]string{
	"Monday":    "星期一",
	"Tuesday":   "星期二",
	"Wednesday": "星期三",
	"Thursday":  "星期四",
	"Friday":    "星期五",
	"Saturday":  "星期六",
	"Sunday":    "星期天",
}

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func TodayRemainTime() int {
	todayLast := time.Now().Format("2006-01-02") + " 23:59:59"
	todayLastTime, _ := time.ParseInLocation("2006-01-02 15:04:05", todayLast, time.Local)

	return int(todayLastTime.Unix() - time.Now().Local().Unix())
}

// RandInt 生成随机数
func RandInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Intn(max-min) + min
}

func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

// JwtUserId 获取UserId
func JwtUserId(c *gin.Context) int {
	token := c.GetHeader("Authorization")
	userId := 0
	if token != "" {
		claims, _ := jwt.ParseToken(token)
		userId = claims.Id
	}

	return userId
}

func RandStr(n int) string {
	result := make([]byte, n/2)
	rand.Read(result)
	return hex.EncodeToString(result)
}

func GetOneStringByRegex(str, rule string) (string, error) {
	reg, err := regexp.Compile(rule)
	if reg == nil || err != nil {
		return "", errors.New("正则Compile错误:" + err.Error())
	}
	//提取关键信息
	result := reg.FindStringSubmatch(str)
	if len(result) < 1 {
		return "", errors.New("没有获取到子字符串")
	}
	return result[1], nil
}

// IsChineseChar 是否包含中文
func IsChineseChar(str string) bool {
	var a = regexp.MustCompile("^[\u4e00-\u9fa5]$")
	//接受正则表达式的范围
	for _, v := range str {
		//golang中string的底层是byte类型，所以单纯的for输出中文会出现乱码，这里选择for-range来输出
		if a.MatchString(string(v)) {
			//判断是否为中文，如果是返回一个true，不是返回false。这俩面MatchString的参数要求是string
			//但是 for-range 返回的 value 是 rune 类型，所以需要做一个 string() 转换
			//fmt.Printf("str 字符串第 %v 个字符是中文。是“%v”字\n", i+1, string(v))
			return true
		}
	}
	return false
	//for _, r := range str {
	//	if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
	//		return true
	//	}
	//}
	//return false
}

func ZuiJin(this int, arr []int) int {
	min := 0
	if this == arr[0] {
		return arr[0]
	} else if this > arr[0] {
		min = this - arr[0]
	} else if this < arr[0] {
		min = arr[0] - this
	}

	for _, v := range arr {
		if v == this {
			return v
		} else if v > this {
			if min > v-this {
				min = v - this
			}
		} else if v < this {
			if min > this-v {
				min = this - v
			}
		}
	}

	for _, v := range arr {
		if this+min == v {
			return v
		} else if this-min == v {
			return v
		}
	}
	return min
}

func ResolveTime(seconds int) (h, m string) {
	var hour, minute int
	var day = seconds / (24 * 3600)
	hour = (seconds - day*3600*24) / 3600
	minute = (seconds - day*24*3600 - hour*3600) / 60
	h = strconv.Itoa(hour)
	m = strconv.Itoa(minute)
	if hour < 10 {
		h = "0" + strconv.Itoa(hour)
	}
	if minute < 10 {
		m = "0" + strconv.Itoa(minute)
	}
	return h, m
}

func TrimLeftChars(s string, n int) string {
	m := 0
	for i := range s {
		if m >= n {
			return s[i:]
		}
		m++
	}
	return s[:0]
}

// GetIpPlace 获取ip归属地
func GetIpPlace(ip string) (string, error) {

	type MyJsonName struct {
		Data []struct {
			ExtendedLocation string `json:"ExtendedLocation"`
			OriginQuery      string `json:"OriginQuery"`
			Appinfo          string `json:"appinfo"`
			DispType         int64  `json:"disp_type"`
			Fetchkey         string `json:"fetchkey"`
			Location         string `json:"location"`
			Origip           string `json:"origip"`
			Origipquery      string `json:"origipquery"`
			Resourceid       string `json:"resourceid"`
			RoleID           int64  `json:"role_id"`
			ShareImage       int64  `json:"shareImage"`
			ShowLikeShare    int64  `json:"showLikeShare"`
			Showlamp         string `json:"showlamp"`
			Titlecont        string `json:"titlecont"`
			Tplt             string `json:"tplt"`
		} `json:"data"`
		SetCacheTime string `json:"set_cache_time"`
		Status       string `json:"status"`
		T            string `json:"t"`
	}

	url := "https://opendata.baidu.com/api.php?query=" + ip + "&co=&resource_id=6006&oe=utf8"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var ipPlace MyJsonName

	err = json.Unmarshal(body, &ipPlace)

	if err != nil {
		return "", err
	}

	if len(ipPlace.Data) == 0 {
		return "未知", nil
	}

	return ipPlace.Data[0].Location, nil
}

func Decimal(num float64) float64 {
	num, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", num), 64)
	return num
}

func LoadImageFromURL(URL string) (image.Image, error) {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("received non 200 response code")
	}

	img, _, err := image.Decode(response.Body)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func GetUrlImgBase64(path string) (baseImg string, err error) {
	//获取网络图片
	client := &http.Client{
		Timeout: time.Second * 5, //超时时间
	}
	var bodyImg io.Reader
	request, err := http.NewRequest("GET", path, bodyImg)
	if err != nil {
		err = errors.New("获取网络图片失败")
		return
	}
	respImg, _ := client.Do(request)
	defer respImg.Body.Close()
	imgByte, _ := ioutil.ReadAll(respImg.Body)

	// 判断文件类型，生成一个前缀，拼接base64后可以直接粘贴到浏览器打开，不需要可以不用下面代码
	//取图片类型
	mimeType := http.DetectContentType(imgByte)
	switch mimeType {
	case "image/jpeg":
		baseImg = "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(imgByte)
	case "image/png":
		baseImg = "data:image/png;base64," + base64.StdEncoding.EncodeToString(imgByte)
	}
	return
}

func GetUrlImgBase64NoPrefix(path string) (baseImg string, err error) {
	//获取网络图片
	client := &http.Client{
		Timeout: time.Second * 5, //超时时间
	}
	var bodyImg io.Reader
	request, err := http.NewRequest("GET", path, bodyImg)
	if err != nil {
		err = errors.New("获取网络图片失败")
		return
	}
	respImg, _ := client.Do(request)
	defer respImg.Body.Close()
	imgByte, _ := ioutil.ReadAll(respImg.Body)

	// 判断文件类型，生成一个前缀，拼接base64后可以直接粘贴到浏览器打开，不需要可以不用下面代码
	//取图片类型
	mimeType := http.DetectContentType(imgByte)
	switch mimeType {
	case "image/jpeg":
		baseImg = base64.StdEncoding.EncodeToString(imgByte)
	case "image/png":
		baseImg = base64.StdEncoding.EncodeToString(imgByte)
	}
	return
}

func Base64ToLocalURL(base64Data string) string {
	if len(base64Data) == 0 {
		return ""
	}
	r, _ := regexp.Compile("^(data:\\s*image\\/(\\w+);base64,)")
	if r.MatchString(base64Data) {
		matchs := r.FindStringSubmatch(base64Data)
		names := strconv.Itoa(RandInt(10000, 99999))
		fileName := fmt.Sprintf("%s.%s", "/resource/upload/"+names, matchs[2])
		filePath := fmt.Sprintf("./%s", fileName)
		decodedData, err := base64.StdEncoding.DecodeString(strings.Replace(base64Data, matchs[1], "", -1))
		if err != nil {
			return ""
		}
		err = ioutil.WriteFile(filePath, decodedData, 0644)
		if err != nil {
			return ""
		}
		return filePath
	}
	return ""
}
