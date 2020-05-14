package qualitycheck

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// UpdateScoreForApi invokes the qualitycheck.UpdateScoreForApi API synchronously
// api document: https://help.aliyun.com/api/qualitycheck/updatescoreforapi.html
func (client *Client) UpdateScoreForApi(request *UpdateScoreForApiRequest) (response *UpdateScoreForApiResponse, err error) {
	response = CreateUpdateScoreForApiResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateScoreForApiWithChan invokes the qualitycheck.UpdateScoreForApi API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/updatescoreforapi.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateScoreForApiWithChan(request *UpdateScoreForApiRequest) (<-chan *UpdateScoreForApiResponse, <-chan error) {
	responseChan := make(chan *UpdateScoreForApiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateScoreForApi(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// UpdateScoreForApiWithCallback invokes the qualitycheck.UpdateScoreForApi API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/updatescoreforapi.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateScoreForApiWithCallback(request *UpdateScoreForApiRequest, callback func(response *UpdateScoreForApiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateScoreForApiResponse
		var err error
		defer close(result)
		response, err = client.UpdateScoreForApi(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// UpdateScoreForApiRequest is the request struct for api UpdateScoreForApi
type UpdateScoreForApiRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// UpdateScoreForApiResponse is the response struct for api UpdateScoreForApi
type UpdateScoreForApiResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateUpdateScoreForApiRequest creates a request to invoke UpdateScoreForApi API
func CreateUpdateScoreForApiRequest() (request *UpdateScoreForApiRequest) {
	request = &UpdateScoreForApiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "UpdateScoreForApi", "", "")
	return
}

// CreateUpdateScoreForApiResponse creates a response to parse from UpdateScoreForApi response
func CreateUpdateScoreForApiResponse() (response *UpdateScoreForApiResponse) {
	response = &UpdateScoreForApiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
