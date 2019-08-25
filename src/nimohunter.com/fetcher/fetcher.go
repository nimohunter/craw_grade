package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func Post(url string, params url.Values) ([]byte, error) {
	request, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(params.Encode()))
	resp, err := invoke(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func FetchRaw(url string) ([]byte, error) {
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	resp, err := invoke(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func Fetch(url string) ([]byte, error) {
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	resp, err := invoke(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	//如果页面传来的不是utf8，我们需要转为utf8格式
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func invoke(request *http.Request) (*http.Response, error) {
	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")

	resp, err := http.DefaultClient.Do(request)

	if resp == nil {
		fmt.Println("resp:", resp)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error:status code:%d", resp.StatusCode)
	}
	return resp, nil
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Ftcher error:%v", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
