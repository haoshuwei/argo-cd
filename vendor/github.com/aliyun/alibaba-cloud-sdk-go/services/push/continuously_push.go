package push

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

// ContinuouslyPush invokes the push.ContinuouslyPush API synchronously
// api document: https://help.aliyun.com/api/push/continuouslypush.html
func (client *Client) ContinuouslyPush(request *ContinuouslyPushRequest) (response *ContinuouslyPushResponse, err error) {
	response = CreateContinuouslyPushResponse()
	err = client.DoAction(request, response)
	return
}

// ContinuouslyPushWithChan invokes the push.ContinuouslyPush API asynchronously
// api document: https://help.aliyun.com/api/push/continuouslypush.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ContinuouslyPushWithChan(request *ContinuouslyPushRequest) (<-chan *ContinuouslyPushResponse, <-chan error) {
	responseChan := make(chan *ContinuouslyPushResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ContinuouslyPush(request)
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

// ContinuouslyPushWithCallback invokes the push.ContinuouslyPush API asynchronously
// api document: https://help.aliyun.com/api/push/continuouslypush.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ContinuouslyPushWithCallback(request *ContinuouslyPushRequest, callback func(response *ContinuouslyPushResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ContinuouslyPushResponse
		var err error
		defer close(result)
		response, err = client.ContinuouslyPush(request)
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

// ContinuouslyPushRequest is the request struct for api ContinuouslyPush
type ContinuouslyPushRequest struct {
	*requests.RpcRequest
	MessageId   string           `position:"Query" name:"MessageId"`
	Target      string           `position:"Query" name:"Target"`
	AppKey      requests.Integer `position:"Query" name:"AppKey"`
	TargetValue string           `position:"Query" name:"TargetValue"`
}

// ContinuouslyPushResponse is the response struct for api ContinuouslyPush
type ContinuouslyPushResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	MessageId string `json:"MessageId" xml:"MessageId"`
}

// CreateContinuouslyPushRequest creates a request to invoke ContinuouslyPush API
func CreateContinuouslyPushRequest() (request *ContinuouslyPushRequest) {
	request = &ContinuouslyPushRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Push", "2016-08-01", "ContinuouslyPush", "", "")
	return
}

// CreateContinuouslyPushResponse creates a response to parse from ContinuouslyPush response
func CreateContinuouslyPushResponse() (response *ContinuouslyPushResponse) {
	response = &ContinuouslyPushResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
