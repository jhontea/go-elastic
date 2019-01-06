package objects

import (
	elastic "gopkg.in/olivere/elastic.v6"
)

type ResponseElasticGetResultObject struct {
	Response
	Data *elastic.GetResult `json:"data"`
}

type ResponseElasticSearchResultObject struct {
	Response
	Data *elastic.SearchResult `json:"data"`
}
