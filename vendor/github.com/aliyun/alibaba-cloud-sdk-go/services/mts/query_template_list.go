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

// QueryTemplateList invokes the mts.QueryTemplateList API synchronously
// api document: https://help.aliyun.com/api/mts/querytemplatelist.html
func (client *Client) QueryTemplateList(request *QueryTemplateListRequest) (response *QueryTemplateListResponse, err error) {
	response = CreateQueryTemplateListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryTemplateListWithChan invokes the mts.QueryTemplateList API asynchronously
// api document: https://help.aliyun.com/api/mts/querytemplatelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryTemplateListWithChan(request *QueryTemplateListRequest) (<-chan *QueryTemplateListResponse, <-chan error) {
	responseChan := make(chan *QueryTemplateListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryTemplateList(request)
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

// QueryTemplateListWithCallback invokes the mts.QueryTemplateList API asynchronously
// api document: https://help.aliyun.com/api/mts/querytemplatelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryTemplateListWithCallback(request *QueryTemplateListRequest, callback func(response *QueryTemplateListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryTemplateListResponse
		var err error
		defer close(result)
		response, err = client.QueryTemplateList(request)
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

// QueryTemplateListRequest is the request struct for api QueryTemplateList
type QueryTemplateListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	TemplateIds          string           `position:"Query" name:"TemplateIds"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// QueryTemplateListResponse is the response struct for api QueryTemplateList
type QueryTemplateListResponse struct {
	*responses.BaseResponse
	RequestId    string                          `json:"RequestId" xml:"RequestId"`
	NonExistTids NonExistTidsInQueryTemplateList `json:"NonExistTids" xml:"NonExistTids"`
	TemplateList TemplateListInQueryTemplateList `json:"TemplateList" xml:"TemplateList"`
}

// CreateQueryTemplateListRequest creates a request to invoke QueryTemplateList API
func CreateQueryTemplateListRequest() (request *QueryTemplateListRequest) {
	request = &QueryTemplateListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryTemplateList", "", "")
	return
}

// CreateQueryTemplateListResponse creates a response to parse from QueryTemplateList response
func CreateQueryTemplateListResponse() (response *QueryTemplateListResponse) {
	response = &QueryTemplateListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
