package rds

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

// DescribeEvents invokes the rds.DescribeEvents API synchronously
// api document: https://help.aliyun.com/api/rds/describeevents.html
func (client *Client) DescribeEvents(request *DescribeEventsRequest) (response *DescribeEventsResponse, err error) {
	response = CreateDescribeEventsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeEventsWithChan invokes the rds.DescribeEvents API asynchronously
// api document: https://help.aliyun.com/api/rds/describeevents.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeEventsWithChan(request *DescribeEventsRequest) (<-chan *DescribeEventsResponse, <-chan error) {
	responseChan := make(chan *DescribeEventsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeEvents(request)
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

// DescribeEventsWithCallback invokes the rds.DescribeEvents API asynchronously
// api document: https://help.aliyun.com/api/rds/describeevents.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeEventsWithCallback(request *DescribeEventsRequest, callback func(response *DescribeEventsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeEventsResponse
		var err error
		defer close(result)
		response, err = client.DescribeEvents(request)
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

// DescribeEventsRequest is the request struct for api DescribeEvents
type DescribeEventsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeEventsResponse is the response struct for api DescribeEvents
type DescribeEventsResponse struct {
	*responses.BaseResponse
	RequestId        string                     `json:"RequestId" xml:"RequestId"`
	TotalRecordCount int                        `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageSize         int                        `json:"PageSize" xml:"PageSize"`
	PageNumber       int                        `json:"PageNumber" xml:"PageNumber"`
	EventItems       EventItemsInDescribeEvents `json:"EventItems" xml:"EventItems"`
}

// CreateDescribeEventsRequest creates a request to invoke DescribeEvents API
func CreateDescribeEventsRequest() (request *DescribeEventsRequest) {
	request = &DescribeEventsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeEvents", "rds", "openAPI")
	return
}

// CreateDescribeEventsResponse creates a response to parse from DescribeEvents response
func CreateDescribeEventsResponse() (response *DescribeEventsResponse) {
	response = &DescribeEventsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
