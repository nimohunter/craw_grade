package engine

import (
	"log"
	"nimohunter.com/fetcher"
	"nimohunter.com/model"
	"nimohunter.com/parser"
)

func doWork(r model.Request) (model.ParseResult, error) {

	body, err := fetcher.Fetch(r.Url)

	if err != nil {
		log.Printf("Fetcher: error fetching url %s %v", r.Url, err)
		return model.ParseResult{}, err
	}

	GetParseFunc(r.ParseMethod)
	return GetParseFunc(r.ParseMethod)(body), nil
}

func GetParseFunc(parseMethod model.ParseType) func(bytes []byte) model.ParseResult {
	switch parseMethod {
	case model.CityListParse:
		return func(bytes []byte) model.ParseResult {
			return parser.ParseCityList(bytes)
		}
	}
	//TODO ADD other method
	return nil
}

func createWorker(in chan model.Request, out chan model.ParseResult) {
	request := <- in

	result, err := doWork(request)
	if err != nil {
		go func() {out <- result}()
	}
}