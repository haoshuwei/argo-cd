package live

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

// DeleteLiveDetectNotifyConfig invokes the live.DeleteLiveDetectNotifyConfig API synchronously
// api document: https://help.aliyun.com/api/live/deletelivedetectnotifyconfig.html
func (client *Client) DeleteLiveDetectNotifyConfig(request *DeleteLiveDetectNotifyConfigRequest) (response *DeleteLiveDetectNotifyConfigResponse, err error) {
	response = CreateDeleteLiveDetectNotifyConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteLiveDetectNotifyConfigWithChan invokes the live.DeleteLiveDetectNotifyConfig API asynchronously
// api document: https://help.aliyun.com/api/live/deletelivedetectnotifyconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteLiveDetectNotifyConfigWithChan(request *DeleteLiveDetectNotifyConfigRequest) (<-chan *DeleteLiveDetectNotifyConfigResponse, <-chan error) {
	responseChan := make(chan *DeleteLiveDetectNotifyConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteLiveDetectNotifyConfig(request)
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

// DeleteLiveDetectNotifyConfigWithCallback invokes the live.DeleteLiveDetectNotifyConfig API asynchronously
// api document: https://help.aliyun.com/api/live/deletelivedetectnotifyconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteLiveDetectNotifyConfigWithCallback(request *DeleteLiveDetectNotifyConfigRequest, callback func(response *DeleteLiveDetectNotifyConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteLiveDetectNotifyConfigResponse
		var err error
		defer close(result)
		response, err = client.DeleteLiveDetectNotifyConfig(request)
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

// DeleteLiveDetectNotifyConfigRequest is the request struct for api DeleteLiveDetectNotifyConfig
type DeleteLiveDetectNotifyConfigRequest struct {
	*requests.RpcRequest
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// DeleteLiveDetectNotifyConfigResponse is the response struct for api DeleteLiveDetectNotifyConfig
type DeleteLiveDetectNotifyConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteLiveDetectNotifyConfigRequest creates a request to invoke DeleteLiveDetectNotifyConfig API
func CreateDeleteLiveDetectNotifyConfigRequest() (request *DeleteLiveDetectNotifyConfigRequest) {
	request = &DeleteLiveDetectNotifyConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DeleteLiveDetectNotifyConfig", "live", "openAPI")
	return
}

// CreateDeleteLiveDetectNotifyConfigResponse creates a response to parse from DeleteLiveDetectNotifyConfig response
func CreateDeleteLiveDetectNotifyConfigResponse() (response *DeleteLiveDetectNotifyConfigResponse) {
	response = &DeleteLiveDetectNotifyConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
