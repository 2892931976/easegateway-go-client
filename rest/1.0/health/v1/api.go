package v1

import (
	"net/url"
	"strings"

	"github.com/hexdecteam/easegateway-go-client/rest/1.0/common/v1"
)

type HealthApi struct {
	Configuration *v1.Configuration
}

func NewHealthApi() *HealthApi {
	configuration := v1.NewConfiguration("http://localhost:9090/health/v1")
	return &HealthApi{
		Configuration: configuration,
	}
}

func NewHealthApiWithBasePath(basePath string) *HealthApi {
	configuration := v1.NewConfiguration(basePath)
	return &HealthApi{
		Configuration: configuration,
	}
}

func (a HealthApi) Check() (*v1.APIResponse, error) {
	method := strings.ToUpper("Get")
	path := a.Configuration.BasePath + "/check"
	headers := make(map[string]string)
	queryParams := url.Values{}

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headers[key] = a.Configuration.DefaultHeader[key]
	}

	// set Content-Type header
	contentTypes := a.Configuration.APIClient.SelectHeaderContentType([]string{})
	if contentTypes != "" {
		headers["Content-Type"] = contentTypes
	}

	// set Accept header
	accepts := a.Configuration.APIClient.SelectHeaderAccept([]string{"application/json"})
	if accepts != "" {
		headers["Accept"] = accepts
	}
	response, err := a.Configuration.APIClient.CallAPI(path, method, nil, headers, queryParams)

	u, _ := url.Parse(path)
	u.RawQuery = queryParams.Encode()
	ret := &v1.APIResponse{Operation: "Check", Method: method, RequestURL: u.String()}
	if response != nil {
		ret.Response = response.RawResponse
		ret.Payload = response.Body()
	}

	if err != nil {
		return ret, err
	}
	return ret, err
}