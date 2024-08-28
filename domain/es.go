package domain

import (
	"bytes"
	"crypto/tls"
	"errors"
	"log"
	"net/http"
	"platform-data/config"
)
import "github.com/elastic/go-elasticsearch/v8"

type IEsMapper interface {
	Insert(documents []byte) error
}

type EsMapper struct {
	es *elasticsearch.Client
}

func NewEsMapper() IEsMapper {
	aConfig := config.Get()
	esClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Username:  aConfig.ElasticSearch.Username,
		Password:  aConfig.ElasticSearch.Password,
		Addresses: aConfig.ElasticSearch.Addr,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
	return &EsMapper{
		es: esClient,
	}
}

func (mapper *EsMapper) Insert(documents []byte) error {

	// 插入文档到指定的索引
	res, err := mapper.es.Bulk(bytes.NewReader(documents))

	if res != nil && res.IsError() {
		log.Println(res.String())
		return errors.New("elastic search 可能存在语法错误")
	}

	if err != nil {
		log.Println(err)
	}

	return err
}
