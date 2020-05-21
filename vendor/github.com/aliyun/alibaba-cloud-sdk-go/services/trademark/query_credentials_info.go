package trademark

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

// QueryCredentialsInfo invokes the trademark.QueryCredentialsInfo API synchronously
// api document: https://help.aliyun.com/api/trademark/querycredentialsinfo.html
func (client *Client) QueryCredentialsInfo(request *QueryCredentialsInfoRequest) (response *QueryCredentialsInfoResponse, err error) {
	response = CreateQueryCredentialsInfoResponse()
	err = client.DoAction(request, response)
	return
}

// QueryCredentialsInfoWithChan invokes the trademark.QueryCredentialsInfo API asynchronously
// api document: https://help.aliyun.com/api/trademark/querycredentialsinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryCredentialsInfoWithChan(request *QueryCredentialsInfoRequest) (<-chan *QueryCredentialsInfoResponse, <-chan error) {
	responseChan := make(chan *QueryCredentialsInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryCredentialsInfo(request)
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

// QueryCredentialsInfoWithCallback invokes the trademark.QueryCredentialsInfo API asynchronously
// api document: https://help.aliyun.com/api/trademark/querycredentialsinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryCredentialsInfoWithCallback(request *QueryCredentialsInfoRequest, callback func(response *QueryCredentialsInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryCredentialsInfoResponse
		var err error
		defer close(result)
		response, err = client.QueryCredentialsInfo(request)
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

// QueryCredentialsInfoRequest is the request struct for api QueryCredentialsInfo
type QueryCredentialsInfoRequest struct {
	*requests.RpcRequest
	OssKey       string `position:"Body" name:"OssKey"`
	MaterialType string `position:"Body" name:"MaterialType"`
	CompanyName  string `position:"Query" name:"CompanyName"`
}

// QueryCredentialsInfoResponse is the response struct for api QueryCredentialsInfo
type QueryCredentialsInfoResponse struct {
	*responses.BaseResponse
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	CredentialsInfo CredentialsInfo `json:"CredentialsInfo" xml:"CredentialsInfo"`
}

// CreateQueryCredentialsInfoRequest creates a request to invoke QueryCredentialsInfo API
func CreateQueryCredentialsInfoRequest() (request *QueryCredentialsInfoRequest) {
	request = &QueryCredentialsInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Trademark", "2018-07-24", "QueryCredentialsInfo", "trademark", "openAPI")
	return
}

// CreateQueryCredentialsInfoResponse creates a response to parse from QueryCredentialsInfo response
func CreateQueryCredentialsInfoResponse() (response *QueryCredentialsInfoResponse) {
	response = &QueryCredentialsInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
