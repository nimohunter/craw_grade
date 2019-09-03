package service

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"nimohunter.com/model"
	"reflect"
)

type ElasticService struct {
	ctx    context.Context
	client *elastic.Client
}

func (service ElasticService) AdvanceSearch(age int, gender string, beauty string) model.SearchResult {

	result, e := service.client.Search("datint_profile").
		Query(elastic.NewMatchQuery("gender", true)).
		Query(elastic.NewRangeQuery("age").Lt(age)).Do(service.ctx)

	if e != nil {
		return model.SearchResult{
			Hits: 0,
		}
	}

	return genSearchResultFromElasticResult(result)
}

func genSearchResultFromElasticResult(res *elastic.SearchResult) model.SearchResult {
	var typ model.Item
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(model.Item)
		fmt.Printf("%#v\n", t)
	}
	return model.SearchResult{
		Hits: 0,
	}
}

func NewElasticService() *ElasticService {
	theClient, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	return &ElasticService{
		ctx:    context.Background(),
		client: theClient,
	}
}
