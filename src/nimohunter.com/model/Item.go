package model

type Item struct {
	Url string //URL
	Type string //存储到ElasticSearch时的type
	Id  string //用户Id
	Payload Profile
}
