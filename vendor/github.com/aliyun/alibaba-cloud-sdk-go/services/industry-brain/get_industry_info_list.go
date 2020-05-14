package industry_brain

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

// GetIndustryInfoList invokes the industry_brain.GetIndustryInfoList API synchronously
// api document: https://help.aliyun.com/api/industry-brain/getindustryinfolist.html
func (client *Client) GetIndustryInfoList(request *GetIndustryInfoListRequest) (response *GetIndustryInfoListResponse, err error) {
	response = CreateGetIndustryInfoListResponse()
	err = client.DoAction(request, response)
	return
}

// GetIndustryInfoListWithChan invokes the industry_brain.GetIndustryInfoList API asynchronously
// api document: https://help.aliyun.com/api/industry-brain/getindustryinfolist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetIndustryInfoListWithChan(request *GetIndustryInfoListRequest) (<-chan *GetIndustryInfoListResponse, <-chan error) {
	responseChan := make(chan *GetIndustryInfoListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetIndustryInfoList(request)
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

// GetIndustryInfoListWithCallback invokes the industry_brain.GetIndustryInfoList API asynchronously
// api document: https://help.aliyun.com/api/industry-brain/getindustryinfolist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetIndustryInfoListWithCallback(request *GetIndustryInfoListRequest, callback func(response *GetIndustryInfoListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetIndustryInfoListResponse
		var err error
		defer close(result)
		response, err = client.GetIndustryInfoList(request)
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

// GetIndustryInfoListRequest is the request struct for api GetIndustryInfoList
type GetIndustryInfoListRequest struct {
	*requests.RpcRequest
}

// GetIndustryInfoListResponse is the response struct for api GetIndustryInfoList
type GetIndustryInfoListResponse struct {
	*responses.BaseResponse
	RequestId        string                 `json:"RequestId" xml:"RequestId"`
	IndustryInfoList []IndustryInfoListItem `json:"IndustryInfoList" xml:"IndustryInfoList"`
}

// CreateGetIndustryInfoListRequest creates a request to invoke GetIndustryInfoList API
func CreateGetIndustryInfoListRequest() (request *GetIndustryInfoListRequest) {
	request = &GetIndustryInfoListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("industry-brain", "2018-07-12", "GetIndustryInfoList", "", "")
	return
}

// CreateGetIndustryInfoListResponse creates a response to parse from GetIndustryInfoList response
func CreateGetIndustryInfoListResponse() (response *GetIndustryInfoListResponse) {
	response = &GetIndustryInfoListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
