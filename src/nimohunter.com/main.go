package main

import (
	"nimohunter.com/engine"
	"nimohunter.com/model"
)

func main() {
	url := "http://www.zhenai.com/zhenghun"

	e := engine.SimpleEngine{
		WorkCount:        100,
		ItemCollectCount: 10,
	}

	e.Run(model.Request{
		Url:         url,
		ParseMethod: model.CityListParse,
	})
}
