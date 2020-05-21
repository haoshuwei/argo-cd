package dds

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

// DescribeAvailableTimeRange invokes the dds.DescribeAvailableTimeRange API synchronously
// api document: https://help.aliyun.com/api/dds/describeavailabletimerange.html
func (client *Client) DescribeAvailableTimeRange(request *DescribeAvailableTimeRangeRequest) (response *DescribeAvailableTimeRangeResponse, err error) {
	response = CreateDescribeAvailableTimeRangeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAvailableTimeRangeWithChan invokes the dds.DescribeAvailableTimeRange API asynchronously
// api document: https://help.aliyun.com/api/dds/describeavailabletimerange.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAvailableTimeRangeWithChan(request *DescribeAvailableTimeRangeRequest) (<-chan *DescribeAvailableTimeRangeResponse, <-chan error) {
	responseChan := make(chan *DescribeAvailableTimeRangeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAvailableTimeRange(request)
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

// DescribeAvailableTimeRangeWithCallback invokes the dds.DescribeAvailableTimeRange API asynchronously
// api document: https://help.aliyun.com/api/dds/describeavailabletimerange.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAvailableTimeRangeWithCallback(request *DescribeAvailableTimeRangeRequest, callback func(response *DescribeAvailableTimeRangeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAvailableTimeRangeResponse
		var err error
		defer close(result)
		response, err = client.DescribeAvailableTimeRange(request)
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

// DescribeAvailableTimeRangeRequest is the request struct for api DescribeAvailableTimeRange
type DescribeAvailableTimeRangeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	NodeId               string           `position:"Query" name:"NodeId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// DescribeAvailableTimeRangeResponse is the response struct for api DescribeAvailableTimeRange
type DescribeAvailableTimeRangeResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	TimeRange TimeRange `json:"TimeRange" xml:"TimeRange"`
}

// CreateDescribeAvailableTimeRangeRequest creates a request to invoke DescribeAvailableTimeRange API
func CreateDescribeAvailableTimeRangeRequest() (request *DescribeAvailableTimeRangeRequest) {
	request = &DescribeAvailableTimeRangeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "DescribeAvailableTimeRange", "dds", "openAPI")
	return
}

// CreateDescribeAvailableTimeRangeResponse creates a response to parse from DescribeAvailableTimeRange response
func CreateDescribeAvailableTimeRangeResponse() (response *DescribeAvailableTimeRangeResponse) {
	response = &DescribeAvailableTimeRangeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
