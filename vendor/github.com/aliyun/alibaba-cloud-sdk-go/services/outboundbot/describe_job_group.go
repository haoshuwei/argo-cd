package outboundbot

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

// DescribeJobGroup invokes the outboundbot.DescribeJobGroup API synchronously
// api document: https://help.aliyun.com/api/outboundbot/describejobgroup.html
func (client *Client) DescribeJobGroup(request *DescribeJobGroupRequest) (response *DescribeJobGroupResponse, err error) {
	response = CreateDescribeJobGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeJobGroupWithChan invokes the outboundbot.DescribeJobGroup API asynchronously
// api document: https://help.aliyun.com/api/outboundbot/describejobgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeJobGroupWithChan(request *DescribeJobGroupRequest) (<-chan *DescribeJobGroupResponse, <-chan error) {
	responseChan := make(chan *DescribeJobGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeJobGroup(request)
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

// DescribeJobGroupWithCallback invokes the outboundbot.DescribeJobGroup API asynchronously
// api document: https://help.aliyun.com/api/outboundbot/describejobgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeJobGroupWithCallback(request *DescribeJobGroupRequest, callback func(response *DescribeJobGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeJobGroupResponse
		var err error
		defer close(result)
		response, err = client.DescribeJobGroup(request)
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

// DescribeJobGroupRequest is the request struct for api DescribeJobGroup
type DescribeJobGroupRequest struct {
	*requests.RpcRequest
	BriefTypes *[]string `position:"Query" name:"BriefTypes"  type:"Repeated"`
	InstanceId string    `position:"Query" name:"InstanceId"`
	JobGroupId string    `position:"Query" name:"JobGroupId"`
}

// DescribeJobGroupResponse is the response struct for api DescribeJobGroup
type DescribeJobGroupResponse struct {
	*responses.BaseResponse
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Success        bool     `json:"Success" xml:"Success"`
	Code           string   `json:"Code" xml:"Code"`
	Message        string   `json:"Message" xml:"Message"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	JobGroup       JobGroup `json:"JobGroup" xml:"JobGroup"`
}

// CreateDescribeJobGroupRequest creates a request to invoke DescribeJobGroup API
func CreateDescribeJobGroupRequest() (request *DescribeJobGroupRequest) {
	request = &DescribeJobGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "DescribeJobGroup", "outboundbot", "openAPI")
	return
}

// CreateDescribeJobGroupResponse creates a response to parse from DescribeJobGroup response
func CreateDescribeJobGroupResponse() (response *DescribeJobGroupResponse) {
	response = &DescribeJobGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
