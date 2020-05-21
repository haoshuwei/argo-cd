package ehpc

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

// SetQueue invokes the ehpc.SetQueue API synchronously
// api document: https://help.aliyun.com/api/ehpc/setqueue.html
func (client *Client) SetQueue(request *SetQueueRequest) (response *SetQueueResponse, err error) {
	response = CreateSetQueueResponse()
	err = client.DoAction(request, response)
	return
}

// SetQueueWithChan invokes the ehpc.SetQueue API asynchronously
// api document: https://help.aliyun.com/api/ehpc/setqueue.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetQueueWithChan(request *SetQueueRequest) (<-chan *SetQueueResponse, <-chan error) {
	responseChan := make(chan *SetQueueResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetQueue(request)
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

// SetQueueWithCallback invokes the ehpc.SetQueue API asynchronously
// api document: https://help.aliyun.com/api/ehpc/setqueue.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetQueueWithCallback(request *SetQueueRequest, callback func(response *SetQueueResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetQueueResponse
		var err error
		defer close(result)
		response, err = client.SetQueue(request)
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

// SetQueueRequest is the request struct for api SetQueue
type SetQueueRequest struct {
	*requests.RpcRequest
	QueueName string          `position:"Query" name:"QueueName"`
	ClusterId string          `position:"Query" name:"ClusterId"`
	Node      *[]SetQueueNode `position:"Query" name:"Node"  type:"Repeated"`
}

// SetQueueNode is a repeated param struct in SetQueueRequest
type SetQueueNode struct {
	Name string `name:"Name"`
}

// SetQueueResponse is the response struct for api SetQueue
type SetQueueResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetQueueRequest creates a request to invoke SetQueue API
func CreateSetQueueRequest() (request *SetQueueRequest) {
	request = &SetQueueRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "SetQueue", "", "")
	return
}

// CreateSetQueueResponse creates a response to parse from SetQueue response
func CreateSetQueueResponse() (response *SetQueueResponse) {
	response = &SetQueueResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
