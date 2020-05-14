package retailcloud

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

// GetRdsBackUp invokes the retailcloud.GetRdsBackUp API synchronously
// api document: https://help.aliyun.com/api/retailcloud/getrdsbackup.html
func (client *Client) GetRdsBackUp(request *GetRdsBackUpRequest) (response *GetRdsBackUpResponse, err error) {
	response = CreateGetRdsBackUpResponse()
	err = client.DoAction(request, response)
	return
}

// GetRdsBackUpWithChan invokes the retailcloud.GetRdsBackUp API asynchronously
// api document: https://help.aliyun.com/api/retailcloud/getrdsbackup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetRdsBackUpWithChan(request *GetRdsBackUpRequest) (<-chan *GetRdsBackUpResponse, <-chan error) {
	responseChan := make(chan *GetRdsBackUpResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRdsBackUp(request)
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

// GetRdsBackUpWithCallback invokes the retailcloud.GetRdsBackUp API asynchronously
// api document: https://help.aliyun.com/api/retailcloud/getrdsbackup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetRdsBackUpWithCallback(request *GetRdsBackUpRequest, callback func(response *GetRdsBackUpResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRdsBackUpResponse
		var err error
		defer close(result)
		response, err = client.GetRdsBackUp(request)
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

// GetRdsBackUpRequest is the request struct for api GetRdsBackUp
type GetRdsBackUpRequest struct {
	*requests.RpcRequest
	BackupId     string           `position:"Body" name:"BackupId"`
	PageSize     requests.Integer `position:"Body" name:"PageSize"`
	DbInstanceId string           `position:"Body" name:"DbInstanceId"`
	BackupType   string           `position:"Body" name:"BackupType"`
	PageNumber   requests.Integer `position:"Body" name:"PageNumber"`
}

// GetRdsBackUpResponse is the response struct for api GetRdsBackUp
type GetRdsBackUpResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	ErrMsg    string `json:"ErrMsg" xml:"ErrMsg"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateGetRdsBackUpRequest creates a request to invoke GetRdsBackUp API
func CreateGetRdsBackUpRequest() (request *GetRdsBackUpRequest) {
	request = &GetRdsBackUpRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "GetRdsBackUp", "retailcloud", "openAPI")
	return
}

// CreateGetRdsBackUpResponse creates a response to parse from GetRdsBackUp response
func CreateGetRdsBackUpResponse() (response *GetRdsBackUpResponse) {
	response = &GetRdsBackUpResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
