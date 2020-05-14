package cr

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

// GetNamespaceAuthorizationList invokes the cr.GetNamespaceAuthorizationList API synchronously
// api document: https://help.aliyun.com/api/cr/getnamespaceauthorizationlist.html
func (client *Client) GetNamespaceAuthorizationList(request *GetNamespaceAuthorizationListRequest) (response *GetNamespaceAuthorizationListResponse, err error) {
	response = CreateGetNamespaceAuthorizationListResponse()
	err = client.DoAction(request, response)
	return
}

// GetNamespaceAuthorizationListWithChan invokes the cr.GetNamespaceAuthorizationList API asynchronously
// api document: https://help.aliyun.com/api/cr/getnamespaceauthorizationlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetNamespaceAuthorizationListWithChan(request *GetNamespaceAuthorizationListRequest) (<-chan *GetNamespaceAuthorizationListResponse, <-chan error) {
	responseChan := make(chan *GetNamespaceAuthorizationListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetNamespaceAuthorizationList(request)
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

// GetNamespaceAuthorizationListWithCallback invokes the cr.GetNamespaceAuthorizationList API asynchronously
// api document: https://help.aliyun.com/api/cr/getnamespaceauthorizationlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetNamespaceAuthorizationListWithCallback(request *GetNamespaceAuthorizationListRequest, callback func(response *GetNamespaceAuthorizationListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetNamespaceAuthorizationListResponse
		var err error
		defer close(result)
		response, err = client.GetNamespaceAuthorizationList(request)
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

// GetNamespaceAuthorizationListRequest is the request struct for api GetNamespaceAuthorizationList
type GetNamespaceAuthorizationListRequest struct {
	*requests.RoaRequest
	Namespace string `position:"Path" name:"Namespace"`
	Authorize string `position:"Query" name:"Authorize"`
}

// GetNamespaceAuthorizationListResponse is the response struct for api GetNamespaceAuthorizationList
type GetNamespaceAuthorizationListResponse struct {
	*responses.BaseResponse
}

// CreateGetNamespaceAuthorizationListRequest creates a request to invoke GetNamespaceAuthorizationList API
func CreateGetNamespaceAuthorizationListRequest() (request *GetNamespaceAuthorizationListRequest) {
	request = &GetNamespaceAuthorizationListRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("cr", "2016-06-07", "GetNamespaceAuthorizationList", "/namespace/[Namespace]/authorizations", "cr", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetNamespaceAuthorizationListResponse creates a response to parse from GetNamespaceAuthorizationList response
func CreateGetNamespaceAuthorizationListResponse() (response *GetNamespaceAuthorizationListResponse) {
	response = &GetNamespaceAuthorizationListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
