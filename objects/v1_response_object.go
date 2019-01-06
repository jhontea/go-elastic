package objects

import (
	elastic "gopkg.in/olivere/elastic.v6"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type ResponseElasticGetResultObject struct {
	Response
	Data *elastic.GetResult `json:"data"`
}
