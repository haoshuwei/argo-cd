package mts

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

// ListCensorPipeline invokes the mts.ListCensorPipeline API synchronously
// api document: https://help.aliyun.com/api/mts/listcensorpipeline.html
func (client *Client) ListCensorPipeline(request *ListCensorPipelineRequest) (response *ListCensorPipelineResponse, err error) {
	response = CreateListCensorPipelineResponse()
	err = client.DoAction(request, response)
	return
}

// ListCensorPipelineWithChan invokes the mts.ListCensorPipeline API asynchronously
// api document: https://help.aliyun.com/api/mts/listcensorpipeline.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListCensorPipelineWithChan(request *ListCensorPipelineRequest) (<-chan *ListCensorPipelineResponse, <-chan error) {
	responseChan := make(chan *ListCensorPipelineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListCensorPipeline(request)
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

// ListCensorPipelineWithCallback invokes the mts.ListCensorPipeline API asynchronously
// api document: https://help.aliyun.com/api/mts/listcensorpipeline.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListCensorPipelineWithCallback(request *ListCensorPipelineRequest, callback func(response *ListCensorPipelineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListCensorPipelineResponse
		var err error
		defer close(result)
		response, err = client.ListCensorPipeline(request)
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

// ListCensorPipelineRequest is the request struct for api ListCensorPipeline
type ListCensorPipelineRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	State                string           `position:"Query" name:"State"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ListCensorPipelineResponse is the response struct for api ListCensorPipeline
type ListCensorPipelineResponse struct {
	*responses.BaseResponse
	RequestId    string                           `json:"RequestId" xml:"RequestId"`
	TotalCount   int64                            `json:"TotalCount" xml:"TotalCount"`
	PageNumber   int64                            `json:"PageNumber" xml:"PageNumber"`
	PageSize     int64                            `json:"PageSize" xml:"PageSize"`
	PipelineList PipelineListInListCensorPipeline `json:"PipelineList" xml:"PipelineList"`
}

// CreateListCensorPipelineRequest creates a request to invoke ListCensorPipeline API
func CreateListCensorPipelineRequest() (request *ListCensorPipelineRequest) {
	request = &ListCensorPipelineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "ListCensorPipeline", "", "")
	return
}

// CreateListCensorPipelineResponse creates a response to parse from ListCensorPipeline response
func CreateListCensorPipelineResponse() (response *ListCensorPipelineResponse) {
	response = &ListCensorPipelineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
