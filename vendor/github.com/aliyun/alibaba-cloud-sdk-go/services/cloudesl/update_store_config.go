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

// UpdateStoreConfig invokes the cloudesl.UpdateStoreConfig API synchronously
// api document: https://help.aliyun.com/api/cloudesl/updatestoreconfig.html
func (client *Client) UpdateStoreConfig(request *UpdateStoreConfigRequest) (response *UpdateStoreConfigResponse, err error) {
	response = CreateUpdateStoreConfigResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateStoreConfigWithChan invokes the cloudesl.UpdateStoreConfig API asynchronously
// api document: https://help.aliyun.com/api/cloudesl/updatestoreconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateStoreConfigWithChan(request *UpdateStoreConfigRequest) (<-chan *UpdateStoreConfigResponse, <-chan error) {
	responseChan := make(chan *UpdateStoreConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateStoreConfig(request)
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

// UpdateStoreConfigWithCallback invokes the cloudesl.UpdateStoreConfig API asynchronously
// api document: https://help.aliyun.com/api/cloudesl/updatestoreconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateStoreConfigWithCallback(request *UpdateStoreConfigRequest, callback func(response *UpdateStoreConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateStoreConfigResponse
		var err error
		defer close(result)
		response, err = client.UpdateStoreConfig(request)
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

// UpdateStoreConfigRequest is the request struct for api UpdateStoreConfig
type UpdateStoreConfigRequest struct {
	*requests.RpcRequest
	EnableNotification      requests.Boolean `position:"Body" name:"EnableNotification"`
	NotificationWebHook     string           `position:"Body" name:"NotificationWebHook"`
	StoreId                 string           `position:"Body" name:"StoreId"`
	NotificationSilentTimes string           `position:"Body" name:"NotificationSilentTimes"`
}

// UpdateStoreConfigResponse is the response struct for api UpdateStoreConfig
type UpdateStoreConfigResponse struct {
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

// CreateUpdateStoreConfigRequest creates a request to invoke UpdateStoreConfig API
func CreateUpdateStoreConfigRequest() (request *UpdateStoreConfigRequest) {
	request = &UpdateStoreConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudesl", "2020-02-01", "UpdateStoreConfig", "cloudesl", "openAPI")
	return
}

// CreateUpdateStoreConfigResponse creates a response to parse from UpdateStoreConfig response
func CreateUpdateStoreConfigResponse() (response *UpdateStoreConfigResponse) {
	response = &UpdateStoreConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
