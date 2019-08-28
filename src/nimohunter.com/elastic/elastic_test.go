package elastic

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"nimohunter.com/engine"
	"nimohunter.com/model"
	"testing"
)

func TestSaveToElastic(t *testing.T) {
	// Starting with elastic.v5, you must pass a context to execute each service
	ctx := context.Background()

	// Obtain a client and connect to the default Elasticsearch installation
	// on 127.0.0.1:9200. Of course you can configure your client to connect
	// to other hosts and configure it in various other ways.
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		// Handle error
		panic(err)
	}

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping("http://127.0.0.1:9200").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	// Getting the ES version number is quite common, so there's a shortcut
	esversion, err := client.ElasticsearchVersion("http://127.0.0.1:9200")
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

	//// Use the IndexExists service to check if a specified index exists.
	//exists, err := client.IndexExists("twitter").Do(ctx)
	//if err != nil {
	//	// Handle error
	//	panic(err)
	//}
	//if !exists {
	//	// Create a new index.
	//	createIndex, err := client.CreateIndex("datint_profile").BodyString(mapping).Do(ctx)
	//	if err != nil {
	//		// Handle error
	//		panic(err)
	//	}
	//	if !createIndex.Acknowledged {
	//		// Not acknowledged
	//	}
	//}
	item := model.Item{
		Name:     "hu",
		Marriage: "ok",
	}
	engine.SaveItem(item, ctx, client)
}
