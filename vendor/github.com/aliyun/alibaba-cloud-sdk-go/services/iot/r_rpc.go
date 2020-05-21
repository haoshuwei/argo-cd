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

// RRpc invokes the iot.RRpc API synchronously
// api document: https://help.aliyun.com/api/iot/rrpc.html
func (client *Client) RRpc(request *RRpcRequest) (response *RRpcResponse, err error) {
	response = CreateRRpcResponse()
	err = client.DoAction(request, response)
	return
}

// RRpcWithChan invokes the iot.RRpc API asynchronously
// api document: https://help.aliyun.com/api/iot/rrpc.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RRpcWithChan(request *RRpcRequest) (<-chan *RRpcResponse, <-chan error) {
	responseChan := make(chan *RRpcResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RRpc(request)
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

// RRpcWithCallback invokes the iot.RRpc API asynchronously
// api document: https://help.aliyun.com/api/iot/rrpc.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RRpcWithCallback(request *RRpcRequest, callback func(response *RRpcResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RRpcResponse
		var err error
		defer close(result)
		response, err = client.RRpc(request)
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

// RRpcRequest is the request struct for api RRpc
type RRpcRequest struct {
	*requests.RpcRequest
	RequestBase64Byte string           `position:"Query" name:"RequestBase64Byte"`
	ProductKey        string           `position:"Query" name:"ProductKey"`
	Timeout           requests.Integer `position:"Query" name:"Timeout"`
	IotInstanceId     string           `position:"Query" name:"IotInstanceId"`
	Topic             string           `position:"Query" name:"Topic"`
	DeviceName        string           `position:"Query" name:"DeviceName"`
}

// RRpcResponse is the response struct for api RRpc
type RRpcResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	Success           bool   `json:"Success" xml:"Success"`
	Code              string `json:"Code" xml:"Code"`
	ErrorMessage      string `json:"ErrorMessage" xml:"ErrorMessage"`
	RrpcCode          string `json:"RrpcCode" xml:"RrpcCode"`
	PayloadBase64Byte string `json:"PayloadBase64Byte" xml:"PayloadBase64Byte"`
	MessageId         int64  `json:"MessageId" xml:"MessageId"`
}

// CreateRRpcRequest creates a request to invoke RRpc API
func CreateRRpcRequest() (request *RRpcRequest) {
	request = &RRpcRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "RRpc", "Iot", "openAPI")
	return
}

// CreateRRpcResponse creates a response to parse from RRpc response
func CreateRRpcResponse() (response *RRpcResponse) {
	response = &RRpcResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
