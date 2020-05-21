package yundun_dbaudit

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

// DescribeSessionLogs invokes the yundun_dbaudit.DescribeSessionLogs API synchronously
// api document: https://help.aliyun.com/api/yundun-dbaudit/describesessionlogs.html
func (client *Client) DescribeSessionLogs(request *DescribeSessionLogsRequest) (response *DescribeSessionLogsResponse, err error) {
	response = CreateDescribeSessionLogsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSessionLogsWithChan invokes the yundun_dbaudit.DescribeSessionLogs API asynchronously
// api document: https://help.aliyun.com/api/yundun-dbaudit/describesessionlogs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSessionLogsWithChan(request *DescribeSessionLogsRequest) (<-chan *DescribeSessionLogsResponse, <-chan error) {
	responseChan := make(chan *DescribeSessionLogsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSessionLogs(request)
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

// DescribeSessionLogsWithCallback invokes the yundun_dbaudit.DescribeSessionLogs API asynchronously
// api document: https://help.aliyun.com/api/yundun-dbaudit/describesessionlogs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSessionLogsWithCallback(request *DescribeSessionLogsRequest, callback func(response *DescribeSessionLogsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSessionLogsResponse
		var err error
		defer close(result)
		response, err = client.DescribeSessionLogs(request)
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

// DescribeSessionLogsRequest is the request struct for api DescribeSessionLogs
type DescribeSessionLogsRequest struct {
	*requests.RpcRequest
	Dmac          string           `position:"Query" name:"Dmac"`
	ReqFlow       string           `position:"Query" name:"ReqFlow"`
	Smac          string           `position:"Query" name:"Smac"`
	Dip           string           `position:"Query" name:"Dip"`
	StartTime     string           `position:"Query" name:"StartTime"`
	Sessionid     string           `position:"Query" name:"Sessionid"`
	Dir           string           `position:"Query" name:"Dir"`
	Dport         string           `position:"Query" name:"Dport"`
	HostName      string           `position:"Query" name:"HostName"`
	ClientUser    string           `position:"Query" name:"ClientUser"`
	SessionStatus string           `position:"Query" name:"SessionStatus"`
	SourceIp      string           `position:"Query" name:"SourceIp"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	Sip           string           `position:"Query" name:"Sip"`
	LoginUser     string           `position:"Query" name:"LoginUser"`
	RspFlow       string           `position:"Query" name:"RspFlow"`
	EndTime       string           `position:"Query" name:"EndTime"`
	CurrentPage   requests.Integer `position:"Query" name:"CurrentPage"`
	Sort          string           `position:"Query" name:"Sort"`
	InstanceId    string           `position:"Query" name:"InstanceId"`
	DbName        string           `position:"Query" name:"DbName"`
	DbId          string           `position:"Query" name:"DbId"`
	DbType        string           `position:"Query" name:"DbType"`
	ClientPrg     string           `position:"Query" name:"ClientPrg"`
	SqlCount      string           `position:"Query" name:"SqlCount"`
	Sport         string           `position:"Query" name:"Sport"`
}

// DescribeSessionLogsResponse is the response struct for api DescribeSessionLogs
type DescribeSessionLogsResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	TotalCount  int    `json:"TotalCount" xml:"TotalCount"`
	SessionLogs []Item `json:"SessionLogs" xml:"SessionLogs"`
}

// CreateDescribeSessionLogsRequest creates a request to invoke DescribeSessionLogs API
func CreateDescribeSessionLogsRequest() (request *DescribeSessionLogsRequest) {
	request = &DescribeSessionLogsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Yundun-dbaudit", "2018-10-29", "DescribeSessionLogs", "dbaudit", "openAPI")
	return
}

// CreateDescribeSessionLogsResponse creates a response to parse from DescribeSessionLogs response
func CreateDescribeSessionLogsResponse() (response *DescribeSessionLogsResponse) {
	response = &DescribeSessionLogsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
