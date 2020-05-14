package aliyuncvc

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

// ActiveMeeting invokes the aliyuncvc.ActiveMeeting API synchronously
// api document: https://help.aliyun.com/api/aliyuncvc/activemeeting.html
func (client *Client) ActiveMeeting(request *ActiveMeetingRequest) (response *ActiveMeetingResponse, err error) {
	response = CreateActiveMeetingResponse()
	err = client.DoAction(request, response)
	return
}

// ActiveMeetingWithChan invokes the aliyuncvc.ActiveMeeting API asynchronously
// api document: https://help.aliyun.com/api/aliyuncvc/activemeeting.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ActiveMeetingWithChan(request *ActiveMeetingRequest) (<-chan *ActiveMeetingResponse, <-chan error) {
	responseChan := make(chan *ActiveMeetingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ActiveMeeting(request)
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

// ActiveMeetingWithCallback invokes the aliyuncvc.ActiveMeeting API asynchronously
// api document: https://help.aliyun.com/api/aliyuncvc/activemeeting.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ActiveMeetingWithCallback(request *ActiveMeetingRequest, callback func(response *ActiveMeetingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ActiveMeetingResponse
		var err error
		defer close(result)
		response, err = client.ActiveMeeting(request)
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

// ActiveMeetingRequest is the request struct for api ActiveMeeting
type ActiveMeetingRequest struct {
	*requests.RpcRequest
	MeetingUUID string `position:"Query" name:"MeetingUUID"`
	MeetingCode string `position:"Query" name:"MeetingCode"`
}

// ActiveMeetingResponse is the response struct for api ActiveMeeting
type ActiveMeetingResponse struct {
	*responses.BaseResponse
	ErrorCode   int         `json:"ErrorCode" xml:"ErrorCode"`
	Message     string      `json:"Message" xml:"Message"`
	Success     bool        `json:"Success" xml:"Success"`
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	MeetingInfo MeetingInfo `json:"MeetingInfo" xml:"MeetingInfo"`
}

// CreateActiveMeetingRequest creates a request to invoke ActiveMeeting API
func CreateActiveMeetingRequest() (request *ActiveMeetingRequest) {
	request = &ActiveMeetingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aliyuncvc", "2019-10-30", "ActiveMeeting", "aliyuncvc", "openAPI")
	return
}

// CreateActiveMeetingResponse creates a response to parse from ActiveMeeting response
func CreateActiveMeetingResponse() (response *ActiveMeetingResponse) {
	response = &ActiveMeetingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
