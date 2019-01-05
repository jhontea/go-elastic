package connections

import (
	elastic "gopkg.in/olivere/elastic.v6"
)

func ElasticInit() *elastic.Client {
	clientURL := "http://localhost:9200"
	client, err := elastic.NewClient(elastic.SetURL(clientURL))

	if err != nil {
		panic(err)
	}

	return client
}
