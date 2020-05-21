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

// SearchWorkflowStatus invokes the lrg.SearchWorkflowStatus API synchronously
// api document: https://help.aliyun.com/api/lrg/searchworkflowstatus.html
func (client *Client) SearchWorkflowStatus(request *SearchWorkflowStatusRequest) (response *SearchWorkflowStatusResponse, err error) {
	response = CreateSearchWorkflowStatusResponse()
	err = client.DoAction(request, response)
	return
}

// SearchWorkflowStatusWithChan invokes the lrg.SearchWorkflowStatus API asynchronously
// api document: https://help.aliyun.com/api/lrg/searchworkflowstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SearchWorkflowStatusWithChan(request *SearchWorkflowStatusRequest) (<-chan *SearchWorkflowStatusResponse, <-chan error) {
	responseChan := make(chan *SearchWorkflowStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SearchWorkflowStatus(request)
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

// SearchWorkflowStatusWithCallback invokes the lrg.SearchWorkflowStatus API asynchronously
// api document: https://help.aliyun.com/api/lrg/searchworkflowstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SearchWorkflowStatusWithCallback(request *SearchWorkflowStatusRequest, callback func(response *SearchWorkflowStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SearchWorkflowStatusResponse
		var err error
		defer close(result)
		response, err = client.SearchWorkflowStatus(request)
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

// SearchWorkflowStatusRequest is the request struct for api SearchWorkflowStatus
type SearchWorkflowStatusRequest struct {
	*requests.RoaRequest
	Id requests.Integer `position:"Path" name:"id"`
}

// SearchWorkflowStatusResponse is the response struct for api SearchWorkflowStatus
type SearchWorkflowStatusResponse struct {
	*responses.BaseResponse
	Code    int                    `json:"code" xml:"code"`
	Data    map[string]interface{} `json:"data" xml:"data"`
	Message string                 `json:"message" xml:"message"`
	Success bool                   `json:"success" xml:"success"`
}

// CreateSearchWorkflowStatusRequest creates a request to invoke SearchWorkflowStatus API
func CreateSearchWorkflowStatusRequest() (request *SearchWorkflowStatusRequest) {
	request = &SearchWorkflowStatusRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("LRG", "2019-10-10", "SearchWorkflowStatus", "/api/v2/tianji/process/[id]?action=queryStatus", "", "")
	request.Method = requests.GET
	return
}

// CreateSearchWorkflowStatusResponse creates a response to parse from SearchWorkflowStatus response
func CreateSearchWorkflowStatusResponse() (response *SearchWorkflowStatusResponse) {
	response = &SearchWorkflowStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
