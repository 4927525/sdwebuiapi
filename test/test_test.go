package test

import (
	"errors"
	"regexp"
	"strconv"
	"testing"
)

func TestTest(t *testing.T) {

	str := "npf71b963c1b7b6d119://auth#session_state=a5e4b12fc3190996c53b1f0377522ee2df9ff6d8f690c4d2b9224d69b33ddebe&session_token_code=eyJhbGciOiJIUzI1NiJ9.eyJleHAiOjE2NjM3NTM0NjUsInR5cCI6InNlc3Npb25fdG9rZW5fY29kZSIsInN1YiI6IjQ4YmRlOTQwNzM0MGQ2MjgiLCJhdWQiOiI3MWI5NjNjMWI3YjZkMTE5Iiwic3RjOmMiOiJZMk5tTVRFMVlUVXhaR0U0WlRrMU5ESXpPVFF5WVRBMlpXTTJZV1l4TkROalltRmkiLCJpc3MiOiJodHRwczovL2FjY291bnRzLm5pbnRlbmRvLmNvbSIsInN0YzpzY3AiOlswLDgsOSwxNywyM10sImlhdCI6MTY2Mzc1Mjg2NSwic3RjOm0iOiJTMjU2IiwianRpIjoiNjE4NDk3MjMzMjcifQ.4v6oZYemAn1BdRK97Cfkafc6tUT7pBbxpxPfo5BqpEM&state=YmI0MTAzNTYzMWY3ODViYzBmMjdlMmY3ZmYyMmEwODJlYzk4"

	url, _ := GetOneStringByRegex(str, "session_token_code=(.*)&state")
	println(url)
	println(resolveTime(5200))
}
func resolveTime(seconds int) (h, m string) {
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
