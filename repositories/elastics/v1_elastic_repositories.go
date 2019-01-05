package repositories

import (
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
