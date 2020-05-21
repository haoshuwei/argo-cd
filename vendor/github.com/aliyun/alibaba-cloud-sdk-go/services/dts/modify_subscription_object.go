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

package dts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifySubscriptionObject invokes the dts.ModifySubscriptionObject API synchronously
// api document: https://help.aliyun.com/api/dts/modifysubscriptionobject.html
func (client *Client) ModifySubscriptionObject(request *ModifySubscriptionObjectRequest) (response *ModifySubscriptionObjectResponse, err error) {
	response = CreateModifySubscriptionObjectResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySubscriptionObjectWithChan invokes the dts.ModifySubscriptionObject API asynchronously
// api document: https://help.aliyun.com/api/dts/modifysubscriptionobject.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySubscriptionObjectWithChan(request *ModifySubscriptionObjectRequest) (<-chan *ModifySubscriptionObjectResponse, <-chan error) {
	responseChan := make(chan *ModifySubscriptionObjectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySubscriptionObject(request)
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

// ModifySubscriptionObjectWithCallback invokes the dts.ModifySubscriptionObject API asynchronously
// api document: https://help.aliyun.com/api/dts/modifysubscriptionobject.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySubscriptionObjectWithCallback(request *ModifySubscriptionObjectRequest, callback func(response *ModifySubscriptionObjectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySubscriptionObjectResponse
		var err error
		defer close(result)
		response, err = client.ModifySubscriptionObject(request)
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

// ModifySubscriptionObjectRequest is the request struct for api ModifySubscriptionObject
type ModifySubscriptionObjectRequest struct {
	*requests.RpcRequest
	SubscriptionInstanceId string `position:"Query" name:"SubscriptionInstanceId"`
	SubscriptionObject     string `position:"Query" name:"SubscriptionObject"`
	OwnerId                string `position:"Query" name:"OwnerId"`
}

// ModifySubscriptionObjectResponse is the response struct for api ModifySubscriptionObject
type ModifySubscriptionObjectResponse struct {
	*responses.BaseResponse
	Success    string `json:"Success" xml:"Success"`
	ErrCode    string `json:"ErrCode" xml:"ErrCode"`
	ErrMessage string `json:"ErrMessage" xml:"ErrMessage"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
}

// CreateModifySubscriptionObjectRequest creates a request to invoke ModifySubscriptionObject API
func CreateModifySubscriptionObjectRequest() (request *ModifySubscriptionObjectRequest) {
	request = &ModifySubscriptionObjectRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2018-08-01", "ModifySubscriptionObject", "dts", "openAPI")
	return
}

// CreateModifySubscriptionObjectResponse creates a response to parse from ModifySubscriptionObject response
func CreateModifySubscriptionObjectResponse() (response *ModifySubscriptionObjectResponse) {
	response = &ModifySubscriptionObjectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
