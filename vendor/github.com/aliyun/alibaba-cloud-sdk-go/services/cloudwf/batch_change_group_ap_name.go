package cloudwf

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

// BatchChangeGroupApName invokes the cloudwf.BatchChangeGroupApName API synchronously
// api document: https://help.aliyun.com/api/cloudwf/batchchangegroupapname.html
func (client *Client) BatchChangeGroupApName(request *BatchChangeGroupApNameRequest) (response *BatchChangeGroupApNameResponse, err error) {
	response = CreateBatchChangeGroupApNameResponse()
	err = client.DoAction(request, response)
	return
}

// BatchChangeGroupApNameWithChan invokes the cloudwf.BatchChangeGroupApName API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/batchchangegroupapname.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchChangeGroupApNameWithChan(request *BatchChangeGroupApNameRequest) (<-chan *BatchChangeGroupApNameResponse, <-chan error) {
	responseChan := make(chan *BatchChangeGroupApNameResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchChangeGroupApName(request)
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

// BatchChangeGroupApNameWithCallback invokes the cloudwf.BatchChangeGroupApName API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/batchchangegroupapname.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchChangeGroupApNameWithCallback(request *BatchChangeGroupApNameRequest, callback func(response *BatchChangeGroupApNameResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchChangeGroupApNameResponse
		var err error
		defer close(result)
		response, err = client.BatchChangeGroupApName(request)
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

// BatchChangeGroupApNameRequest is the request struct for api BatchChangeGroupApName
type BatchChangeGroupApNameRequest struct {
	*requests.RpcRequest
	JsonData string `position:"Query" name:"JsonData"`
}

// BatchChangeGroupApNameResponse is the response struct for api BatchChangeGroupApName
type BatchChangeGroupApNameResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateBatchChangeGroupApNameRequest creates a request to invoke BatchChangeGroupApName API
func CreateBatchChangeGroupApNameRequest() (request *BatchChangeGroupApNameRequest) {
	request = &BatchChangeGroupApNameRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "BatchChangeGroupApName", "cloudwf", "openAPI")
	return
}

// CreateBatchChangeGroupApNameResponse creates a response to parse from BatchChangeGroupApName response
func CreateBatchChangeGroupApNameResponse() (response *BatchChangeGroupApNameResponse) {
	response = &BatchChangeGroupApNameResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
