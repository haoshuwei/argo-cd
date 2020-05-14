package imm

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

// GetContentKey invokes the imm.GetContentKey API synchronously
// api document: https://help.aliyun.com/api/imm/getcontentkey.html
func (client *Client) GetContentKey(request *GetContentKeyRequest) (response *GetContentKeyResponse, err error) {
	response = CreateGetContentKeyResponse()
	err = client.DoAction(request, response)
	return
}

// GetContentKeyWithChan invokes the imm.GetContentKey API asynchronously
// api document: https://help.aliyun.com/api/imm/getcontentkey.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetContentKeyWithChan(request *GetContentKeyRequest) (<-chan *GetContentKeyResponse, <-chan error) {
	responseChan := make(chan *GetContentKeyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetContentKey(request)
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

// GetContentKeyWithCallback invokes the imm.GetContentKey API asynchronously
// api document: https://help.aliyun.com/api/imm/getcontentkey.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetContentKeyWithCallback(request *GetContentKeyRequest, callback func(response *GetContentKeyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetContentKeyResponse
		var err error
		defer close(result)
		response, err = client.GetContentKey(request)
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

// GetContentKeyRequest is the request struct for api GetContentKey
type GetContentKeyRequest struct {
	*requests.RpcRequest
	Project     string `position:"Query" name:"Project"`
	VersionId   string `position:"Query" name:"VersionId"`
	DRMServerId string `position:"Query" name:"DRMServerId"`
	KeyIds      string `position:"Query" name:"KeyIds"`
}

// GetContentKeyResponse is the response struct for api GetContentKey
type GetContentKeyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	VersionId string `json:"VersionId" xml:"VersionId"`
	KeyInfos  string `json:"KeyInfos" xml:"KeyInfos"`
}

// CreateGetContentKeyRequest creates a request to invoke GetContentKey API
func CreateGetContentKeyRequest() (request *GetContentKeyRequest) {
	request = &GetContentKeyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "GetContentKey", "imm", "openAPI")
	return
}

// CreateGetContentKeyResponse creates a response to parse from GetContentKey response
func CreateGetContentKeyResponse() (response *GetContentKeyResponse) {
	response = &GetContentKeyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
