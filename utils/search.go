package utils

import (
	"gopkg.in/olivere/elastic.v5"
	"log"
	"fmt"
)

var (
	elasticClient *elastic.Client
)

type ElasticConnect struct {
	Host string
	Port int
}

func NewElasticClient(config *ElasticConnect) *elastic.Client {
	if config == nil ||
		config.Host == "" ||
		config.Port <= 0 {
		panic("elastic search connect config error")
	}

	url := fmt.Sprintf("%s:%d", config.Host, config.Port)

	client, err := elastic.NewClient(elastic.SetURL(url))
	if err != nil {
		panic(err)
	}

	version, err := client.ElasticsearchVersion(url)
	if err != nil {
		panic(err)
	}
	log.Printf("elastic search version %s\n", version)

	return client
}