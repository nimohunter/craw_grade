package engine

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"log"
	"nimohunter.com/model"
)

const mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"profile":{
			"properties":{
				"Id":{
					"type":"keyword"
				},
				"name":{
					"type":"string"
				},
				"marriage":{
					"type":"string"
				},
				"age":{
					"type":"string"
				},
				"gender":{
					"type":"string"
				},
				"height":{
					"type":"string"
				},
				"weight":{
					"type":"string"
				},
				"income":{
					"type":"string"
				},
				"photo_url":{
					"type":"url"
				},
				"update_time":{
					"type":"date"
					"index" : "not_analyzed",
					"doc_values" : true,
					"format" : "dd/MM/YYYY:HH:mm:ss Z"
				},
			}
		}
	}
}`

func createItemCollectWorker(startSignal chan int, itemChan chan model.Item) {
	<-startSignal

	ctx := context.Background()
	client, err := elastic.NewClient(elastic.SetSniff(false))
	_, _, err = client.Ping("http://127.0.0.1:9200").Do(ctx)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case item := <-itemChan:
			SaveItem(item, ctx, client)
		}
	}
}

func SaveItem(item model.Item, ctx context.Context, client *elastic.Client) {
	fmt.Println("SaveItem" + item.Name)

	resp, err := client.Index(). //存储数据，可以添加或者修改，要看id是否存在
					Index("datint_profile").
					BodyJson(item).
					Do(ctx)

	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%+v", resp) //格式化输出结构体对象的时候包含了字段名称
}
