package cassandra

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

// ResizeNodeCount invokes the cassandra.ResizeNodeCount API synchronously
// api document: https://help.aliyun.com/api/cassandra/resizenodecount.html
func (client *Client) ResizeNodeCount(request *ResizeNodeCountRequest) (response *ResizeNodeCountResponse, err error) {
	response = CreateResizeNodeCountResponse()
	err = client.DoAction(request, response)
	return
}

// ResizeNodeCountWithChan invokes the cassandra.ResizeNodeCount API asynchronously
// api document: https://help.aliyun.com/api/cassandra/resizenodecount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResizeNodeCountWithChan(request *ResizeNodeCountRequest) (<-chan *ResizeNodeCountResponse, <-chan error) {
	responseChan := make(chan *ResizeNodeCountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResizeNodeCount(request)
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

// ResizeNodeCountWithCallback invokes the cassandra.ResizeNodeCount API asynchronously
// api document: https://help.aliyun.com/api/cassandra/resizenodecount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResizeNodeCountWithCallback(request *ResizeNodeCountRequest, callback func(response *ResizeNodeCountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResizeNodeCountResponse
		var err error
		defer close(result)
		response, err = client.ResizeNodeCount(request)
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

// ResizeNodeCountRequest is the request struct for api ResizeNodeCount
type ResizeNodeCountRequest struct {
	*requests.RpcRequest
	DataCenterId string           `position:"Query" name:"DataCenterId"`
	ClusterId    string           `position:"Query" name:"ClusterId"`
	NodeCount    requests.Integer `position:"Query" name:"NodeCount"`
}

// ResizeNodeCountResponse is the response struct for api ResizeNodeCount
type ResizeNodeCountResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateResizeNodeCountRequest creates a request to invoke ResizeNodeCount API
func CreateResizeNodeCountRequest() (request *ResizeNodeCountRequest) {
	request = &ResizeNodeCountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cassandra", "2019-01-01", "ResizeNodeCount", "Cassandra", "openAPI")
	return
}

// CreateResizeNodeCountResponse creates a response to parse from ResizeNodeCount response
func CreateResizeNodeCountResponse() (response *ResizeNodeCountResponse) {
	response = &ResizeNodeCountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
