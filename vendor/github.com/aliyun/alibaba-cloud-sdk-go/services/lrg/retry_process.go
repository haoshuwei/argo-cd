package lrg

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

// RetryProcess invokes the lrg.RetryProcess API synchronously
// api document: https://help.aliyun.com/api/lrg/retryprocess.html
func (client *Client) RetryProcess(request *RetryProcessRequest) (response *RetryProcessResponse, err error) {
	response = CreateRetryProcessResponse()
	err = client.DoAction(request, response)
	return
}

// RetryProcessWithChan invokes the lrg.RetryProcess API asynchronously
// api document: https://help.aliyun.com/api/lrg/retryprocess.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RetryProcessWithChan(request *RetryProcessRequest) (<-chan *RetryProcessResponse, <-chan error) {
	responseChan := make(chan *RetryProcessResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RetryProcess(request)
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

// RetryProcessWithCallback invokes the lrg.RetryProcess API asynchronously
// api document: https://help.aliyun.com/api/lrg/retryprocess.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RetryProcessWithCallback(request *RetryProcessRequest, callback func(response *RetryProcessResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RetryProcessResponse
		var err error
		defer close(result)
		response, err = client.RetryProcess(request)
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

// RetryProcessRequest is the request struct for api RetryProcess
type RetryProcessRequest struct {
	*requests.RoaRequest
	Id requests.Integer `position:"Path" name:"id"`
}

// RetryProcessResponse is the response struct for api RetryProcess
type RetryProcessResponse struct {
	*responses.BaseResponse
	Code    int                    `json:"code" xml:"code"`
	Data    map[string]interface{} `json:"data" xml:"data"`
	Message string                 `json:"message" xml:"message"`
	Success bool                   `json:"success" xml:"success"`
}

// CreateRetryProcessRequest creates a request to invoke RetryProcess API
func CreateRetryProcessRequest() (request *RetryProcessRequest) {
	request = &RetryProcessRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("LRG", "2019-10-10", "RetryProcess", "/api/v2/tianji/process/[id]/retry", "", "")
	request.Method = requests.POST
	return
}

// CreateRetryProcessResponse creates a response to parse from RetryProcess response
func CreateRetryProcessResponse() (response *RetryProcessResponse) {
	response = &RetryProcessResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
