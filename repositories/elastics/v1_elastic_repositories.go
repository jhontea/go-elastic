package repositories

import (
	"context"
	"fmt"
	"os"

	elastic "gopkg.in/olivere/elastic.v6"
)

type V1ElasticRepository struct {
	Client elastic.Client
}

func V1ElasticRepositoryHandler(client elastic.Client) V1ElasticRepository {
	repository := V1ElasticRepository{Client: client}
	return repository
}

func (repository *V1ElasticRepository) GetElasticVersion() (string, error) {
	var result string
	clientURL := os.Getenv("ELASTIC_CLIENT_URL")
	esversion, err := repository.Client.ElasticsearchVersion(clientURL)

	if err != nil {
		return result, err
	}

	result = "Elasticsearch version " + esversion

	return result, nil
}

func (repository *V1ElasticRepository) GetByField(field string, value string, index string) (*elastic.SearchResult, string, error) {
	termQuery := elastic.NewTermQuery(field, value)
	searchResult, err := repository.Client.Search().
		Index(index).
		Query(termQuery).
		From(0).Size(10).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		return nil, err.Error(), err
	}

	message := fmt.Sprintf("Query took %d milliseconds\n", searchResult.TookInMillis)

	return searchResult, message, err
}

func (repository *V1ElasticRepository) GetById(uuid string, index string) (*elastic.GetResult, string, error) {
	var message string

	data, err := repository.Client.Get().
		Index(index).
		Type("doc").
		Id(uuid).
		Do(context.Background())
	if err != nil {
		switch {
		case elastic.IsNotFound(err):
			message = fmt.Sprintf("Document not found: %v", err)
		case elastic.IsTimeout(err):
			message = fmt.Sprintf("Timeout retrieving document: %v", err)
		case elastic.IsConnErr(err):
			message = fmt.Sprintf("Connection problem: %v", err)
		default:
		}
	} else {
		message = fmt.Sprintf("Got document %s in version %d from index %s, type %s\n", data.Id, data.Version, data.Index, data.Type)
	}

	return data, message, err
}
