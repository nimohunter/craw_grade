package parser

import (
	"fmt"
	"nimohunter.com/model"
	"regexp"
)

const cityListRe = `href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`

//解析城市信息
func ParseCityList(contents []byte) model.ParseResult {

	uniqueMap := make(map[string]bool)

	re := regexp.MustCompile(cityListRe)
	all := re.FindAllSubmatch(contents, -1)

	result := model.ParseResult{}
	i := 0
	for _, c := range all {
		if _, ok := uniqueMap[string(c[2])]; !ok {
			item := model.Item{
				// TODO fill the item
				Url:  string(c[1]),
				Type: "cityList",
			}
			//print city name
			fmt.Printf("city: %s  url: %s\n", string(c[2]), string(c[1]))
			result.Items = append(result.Items, item)

			result.Requests = append(result.Requests, model.Request{
				Url:         string(c[1]),
				ParseMethod: model.CityParse,
			})
			//FIXME just run 2 city
			i++
			if i > 1 {
				break
			}
			uniqueMap[string(c[2])] = true
		}

	}

	return result
}
