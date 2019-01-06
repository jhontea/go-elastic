package services

import (
	"net/http"

	elastic "gopkg.in/olivere/elastic.v6"

	"github.com/jhontea/go-elastic/constants"
	"github.com/jhontea/go-elastic/objects"
	repositories "github.com/jhontea/go-elastic/repositories/elastics"
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
		response.Status = constants.RESPONSE_FAILED

		return response, err
	}

	response.Code = http.StatusOK
	response.Message = result
	response.Status = constants.RESPONSE_SUCCESS

	return response, err
}

func (service *V1ElasticService) GetById(uuid string, index string) objects.ResponseElasticGetResultObject {
	var response objects.ResponseElasticGetResultObject

	response.Data = &elastic.GetResult{}
	response.Code = http.StatusUnprocessableEntity
	response.Status = constants.RESPONSE_FAILED

	if uuid == "" {
		response.Message = "Uuid required"
		return response
	} else if index == "" {
		response.Message = "Index required"
		return response
	}

	result, message, err := service.ElasticRepository.GetById(uuid, index)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()

		return response
	}

	response.Data = result
	response.Code = http.StatusOK
	response.Message = message
	response.Status = constants.RESPONSE_SUCCESS

	return response
}
