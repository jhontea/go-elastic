package repositories

import (
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
	esversion, err := repository.Client.ElasticsearchVersion("http://127.0.0.1:9200")

	if err != nil {
		return result, err
	}

	result = "Elasticsearch version " + esversion

	return result, nil
}
