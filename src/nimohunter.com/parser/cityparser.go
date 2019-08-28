package parser

import (
	"nimohunter.com/model"
	"regexp"
)

var (
	cityRe    = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func CityParser(contents []byte) model.ParseResult {
	all := cityRe.FindAllSubmatch(contents, -1)
	result := model.ParseResult{}

	for _, c := range all {
		url := string(c[1])
		result.Requests = append(result.Requests, model.Request{
			Url:         url,
			ParseMethod: model.ProfileParse,
		})
	}
	nextPage := cityUrlRe.FindAllSubmatch(contents, -1)
	for _, c := range nextPage {
		result.Requests = append(result.Requests, model.Request{
			Url:         string(c[1]),
			ParseMethod: model.CityParse,
		})
	}
	return result
}
