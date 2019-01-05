package services

import (
	"net/http"

	"github.com/jhontea/go-elastic/objects"
	repositories "github.com/jhontea/go-elastic/repositories/elastics"
	elastic "gopkg.in/olivere/elastic.v6"
)

type V1ElasticService struct {
	ElasticRepository repositories.V1ElasticRepository
}

func V1ElasticServiceHandler(client elastic.Client) V1ElasticService {
	service := V1ElasticService{
		ElasticRepository: repositories.V1ElasticRepositoryHandler(client),
	}
	return service
}

func (service *V1ElasticService) GetElasticVersion() (objects.Response, error) {
	var response objects.Response

	result, err := service.ElasticRepository.GetElasticVersion()
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		response.Status = "failed"

		return response, err
	}

	response.Code = http.StatusOK
	response.Message = result
	response.Status = "success"

	return response, err
}
