package connections

import (
	"os"

	elastic "gopkg.in/olivere/elastic.v6"
)

func ElasticInit() *elastic.Client {
	clientURL := os.Getenv("ELASTIC_CLIENT_URL")
	client, err := elastic.NewClient(elastic.SetURL(clientURL))

	if err != nil {
		panic(err)
	}

	return client
}
