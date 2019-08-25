package facerank

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"net/url"
	"nimohunter.com/fetcher"
	"sort"
	"strconv"
	"strings"
	"time"
)

const AppID = "2121104776"
const AppKey = "Sn3VRV4281fIV6WZ"
const url_preffix = "https://api.ai.qq.com/fcgi-bin/face/face_detectface"

func GetFaceRank(photoUrl string) int {

	photoBase64 := getPhotoBase64(photoUrl)
	if len(photoBase64) == 0 {
		return 0
	}
	params := getReqParams(photoBase64)
	resultBytes, e := fetcher.Post(url_preffix, params)
	if e != nil {
		return 0
	}

	fmt.Println(string(resultBytes))
	return 0
}

func getReqParams(photoBase64 string) url.Values {
	params := url.Values{}
	params.Set("app_id", AppID)
	params.Set("app_key", AppKey)
	params.Set("mode", "0")
	timeStamp := time.Now().Unix()
	timeStampStr := strconv.FormatInt(timeStamp, 10)
	params.Set("time_stamp", timeStampStr)
	params.Set("nonce_str", timeStampStr)
	params.Set("image", photoBase64)
	signStr := genSignString(params)
	params.Set("sign", signStr)
	return params
}

func genSignString(params url.Values) string {
	signString := ""
	keys := getSortedKeys(params)
	for _, key := range keys {
		if key == "app_key" {
			continue
		}
		value := params.Get(key)
		signString += key + "=" + url.QueryEscape(value) + "&"
	}
	signString += "app_key=" + AppKey
	has := md5.Sum([]byte(signString))
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return strings.ToUpper(md5str)
}

func getSortedKeys(values url.Values) []string {
	var keys []string
	for k := range values {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func getPhotoBase64(photoUrl string) string {
	photo, err := fetcher.FetchRaw(photoUrl)
	if err != nil {
		return ""
	}
	photoBase64 := base64.StdEncoding.EncodeToString(photo)
	return photoBase64
}
