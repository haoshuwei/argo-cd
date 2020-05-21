package cloudesl

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

// DeleteStore invokes the cloudesl.DeleteStore API synchronously
// api document: https://help.aliyun.com/api/cloudesl/deletestore.html
func (client *Client) DeleteStore(request *DeleteStoreRequest) (response *DeleteStoreResponse, err error) {
	response = CreateDeleteStoreResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteStoreWithChan invokes the cloudesl.DeleteStore API asynchronously
// api document: https://help.aliyun.com/api/cloudesl/deletestore.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteStoreWithChan(request *DeleteStoreRequest) (<-chan *DeleteStoreResponse, <-chan error) {
	responseChan := make(chan *DeleteStoreResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteStore(request)
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

// DeleteStoreWithCallback invokes the cloudesl.DeleteStore API asynchronously
// api document: https://help.aliyun.com/api/cloudesl/deletestore.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteStoreWithCallback(request *DeleteStoreRequest, callback func(response *DeleteStoreResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteStoreResponse
		var err error
		defer close(result)
		response, err = client.DeleteStore(request)
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

// DeleteStoreRequest is the request struct for api DeleteStore
type DeleteStoreRequest struct {
	*requests.RpcRequest
	StoreId string `position:"Body" name:"StoreId"`
}

// DeleteStoreResponse is the response struct for api DeleteStore
type DeleteStoreResponse struct {
	*responses.BaseResponse
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	Message        string `json:"Message" xml:"Message"`
	DynamicCode    string `json:"DynamicCode" xml:"DynamicCode"`
	Code           string `json:"Code" xml:"Code"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateDeleteStoreRequest creates a request to invoke DeleteStore API
func CreateDeleteStoreRequest() (request *DeleteStoreRequest) {
	request = &DeleteStoreRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudesl", "2020-02-01", "DeleteStore", "cloudesl", "openAPI")
	return
}

// CreateDeleteStoreResponse creates a response to parse from DeleteStore response
func CreateDeleteStoreResponse() (response *DeleteStoreResponse) {
	response = &DeleteStoreResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
