package rds

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

// CreateDedicatedHostUser invokes the rds.CreateDedicatedHostUser API synchronously
// api document: https://help.aliyun.com/api/rds/creatededicatedhostuser.html
func (client *Client) CreateDedicatedHostUser(request *CreateDedicatedHostUserRequest) (response *CreateDedicatedHostUserResponse, err error) {
	response = CreateCreateDedicatedHostUserResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDedicatedHostUserWithChan invokes the rds.CreateDedicatedHostUser API asynchronously
// api document: https://help.aliyun.com/api/rds/creatededicatedhostuser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDedicatedHostUserWithChan(request *CreateDedicatedHostUserRequest) (<-chan *CreateDedicatedHostUserResponse, <-chan error) {
	responseChan := make(chan *CreateDedicatedHostUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDedicatedHostUser(request)
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

// CreateDedicatedHostUserWithCallback invokes the rds.CreateDedicatedHostUser API asynchronously
// api document: https://help.aliyun.com/api/rds/creatededicatedhostuser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDedicatedHostUserWithCallback(request *CreateDedicatedHostUserRequest, callback func(response *CreateDedicatedHostUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDedicatedHostUserResponse
		var err error
		defer close(result)
		response, err = client.CreateDedicatedHostUser(request)
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

// CreateDedicatedHostUserRequest is the request struct for api CreateDedicatedHostUser
type CreateDedicatedHostUserRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	UserPassword         string           `position:"Query" name:"UserPassword"`
	DedicatedHostName    string           `position:"Query" name:"DedicatedHostName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	UserName             string           `position:"Query" name:"UserName"`
}

// CreateDedicatedHostUserResponse is the response struct for api CreateDedicatedHostUser
type CreateDedicatedHostUserResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateDedicatedHostUserRequest creates a request to invoke CreateDedicatedHostUser API
func CreateCreateDedicatedHostUserRequest() (request *CreateDedicatedHostUserRequest) {
	request = &CreateDedicatedHostUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "CreateDedicatedHostUser", "rds", "openAPI")
	return
}

// CreateCreateDedicatedHostUserResponse creates a response to parse from CreateDedicatedHostUser response
func CreateCreateDedicatedHostUserResponse() (response *CreateDedicatedHostUserResponse) {
	response = &CreateDedicatedHostUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
