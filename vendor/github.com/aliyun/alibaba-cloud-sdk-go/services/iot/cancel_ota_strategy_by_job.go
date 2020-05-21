package iot

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

// CancelOTAStrategyByJob invokes the iot.CancelOTAStrategyByJob API synchronously
// api document: https://help.aliyun.com/api/iot/cancelotastrategybyjob.html
func (client *Client) CancelOTAStrategyByJob(request *CancelOTAStrategyByJobRequest) (response *CancelOTAStrategyByJobResponse, err error) {
	response = CreateCancelOTAStrategyByJobResponse()
	err = client.DoAction(request, response)
	return
}

// CancelOTAStrategyByJobWithChan invokes the iot.CancelOTAStrategyByJob API asynchronously
// api document: https://help.aliyun.com/api/iot/cancelotastrategybyjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CancelOTAStrategyByJobWithChan(request *CancelOTAStrategyByJobRequest) (<-chan *CancelOTAStrategyByJobResponse, <-chan error) {
	responseChan := make(chan *CancelOTAStrategyByJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CancelOTAStrategyByJob(request)
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

// CancelOTAStrategyByJobWithCallback invokes the iot.CancelOTAStrategyByJob API asynchronously
// api document: https://help.aliyun.com/api/iot/cancelotastrategybyjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CancelOTAStrategyByJobWithCallback(request *CancelOTAStrategyByJobRequest, callback func(response *CancelOTAStrategyByJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CancelOTAStrategyByJobResponse
		var err error
		defer close(result)
		response, err = client.CancelOTAStrategyByJob(request)
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

// CancelOTAStrategyByJobRequest is the request struct for api CancelOTAStrategyByJob
type CancelOTAStrategyByJobRequest struct {
	*requests.RpcRequest
	JobId         string `position:"Query" name:"JobId"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// CancelOTAStrategyByJobResponse is the response struct for api CancelOTAStrategyByJob
type CancelOTAStrategyByJobResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateCancelOTAStrategyByJobRequest creates a request to invoke CancelOTAStrategyByJob API
func CreateCancelOTAStrategyByJobRequest() (request *CancelOTAStrategyByJobRequest) {
	request = &CancelOTAStrategyByJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "CancelOTAStrategyByJob", "Iot", "openAPI")
	return
}

// CreateCancelOTAStrategyByJobResponse creates a response to parse from CancelOTAStrategyByJob response
func CreateCancelOTAStrategyByJobResponse() (response *CancelOTAStrategyByJobResponse) {
	response = &CancelOTAStrategyByJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
