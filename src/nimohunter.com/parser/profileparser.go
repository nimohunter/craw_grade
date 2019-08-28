package parser

import (
	"errors"
	"github.com/bitly/go-simplejson"
	"nimohunter.com/model"
	"regexp"
	"strings"
)

var profileRe = regexp.MustCompile(`<script>window.__INITIAL_STATE__=(.+);\(function`)

func ProfileParser(context []byte) model.ParseResult {

	match := profileRe.FindSubmatch(context)
	result := model.ParseResult{}

	if len(match) >= 2 {
		json := match[1]
		item, e := parseJson(json)
		if e != nil {
			return result
		}
		result.Items = append(result.Items, item)
	}
	return result

}

//解析json数据
func parseJson(json []byte) (model.Item, error) {
	res, err := simplejson.NewJson(json)
	if err != nil {
		return model.Item{}, errors.New("Json parse error")
	}
	var profile model.Item

	fillBasicInfo(&profile, res)
	fillDetailInfo(&profile, res)
	fillOtherInfo(&profile, res)
	return profile, nil
}

func fillOtherInfo(profile *model.Item, res *simplejson.Json) {
	name, err := res.Get("objectInfo").Get("nickname").String()
	if err == nil {
		profile.Name = name
	}

	gender, err := res.Get("objectInfo").Get("genderString").String()
	if err == nil {
		profile.Gender = gender
	}

	id, err := res.Get("objectInfo").Get("memberID").String()
	if err == nil {
		profile.Id = id
	}

	photoList, err := res.Get("objectInfo").Get("photos").Array()
	if err == nil {
		for i := range photoList {
			photoInfo := res.Get("objectInfo").Get("photos").GetIndex(i)
			profile.PhotoUrl = append(profile.PhotoUrl, photoInfo.Get("photoURL").MustString())
		}
	}

}

func fillDetailInfo(profile *model.Item, res *simplejson.Json) {
	infos2, err := res.Get("objectInfo").Get("detailInfo").Array()
	if err != nil {
		return
	}
	for _, v := range infos2 {
		if e, ok := v.(string); ok {
			if strings.Contains(e, "族") {
				profile.Hukou = e
			} else if strings.Contains(e, "房") {
				profile.House = e
			} else if strings.Contains(e, "车") {
				profile.Car = e
			}
		}
	}
}

func fillBasicInfo(profile *model.Item, res *simplejson.Json) {
	infos, err := res.Get("objectInfo").Get("basicInfo").Array()
	if err != nil {
		return
	}
	length := len(infos)
	for k, v := range infos {
		if e, ok := v.(string); ok {
			if strings.Contains(e, "未婚") || strings.Contains(e, "离异") || strings.Contains(e, "丧偶") {
				profile.Marriage = e
			} else if strings.Contains(e, "岁") {
				profile.Age = e
			} else if strings.Contains(e, "座") {
				profile.Xingzuo = e
			} else if strings.Contains(e, "cm") {
				profile.Height = e
			} else if strings.Contains(e, "kg") {
				profile.Weight = e
			} else if strings.Contains(e, "月收入") {
				profile.Income = e
			} else {
				//剩下的两个内部不太好写，我们可以按照下标来解析
				switch k {
				case length - 2:
					profile.Occupation = e
				case length - 1:
					profile.Education = e
				}
			}
		}
	}
}
