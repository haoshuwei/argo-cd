package ddosbgp

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

// DescribeExcpetionCount invokes the ddosbgp.DescribeExcpetionCount API synchronously
// api document: https://help.aliyun.com/api/ddosbgp/describeexcpetioncount.html
func (client *Client) DescribeExcpetionCount(request *DescribeExcpetionCountRequest) (response *DescribeExcpetionCountResponse, err error) {
	response = CreateDescribeExcpetionCountResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeExcpetionCountWithChan invokes the ddosbgp.DescribeExcpetionCount API asynchronously
// api document: https://help.aliyun.com/api/ddosbgp/describeexcpetioncount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeExcpetionCountWithChan(request *DescribeExcpetionCountRequest) (<-chan *DescribeExcpetionCountResponse, <-chan error) {
	responseChan := make(chan *DescribeExcpetionCountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeExcpetionCount(request)
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

// DescribeExcpetionCountWithCallback invokes the ddosbgp.DescribeExcpetionCount API asynchronously
// api document: https://help.aliyun.com/api/ddosbgp/describeexcpetioncount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeExcpetionCountWithCallback(request *DescribeExcpetionCountRequest, callback func(response *DescribeExcpetionCountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeExcpetionCountResponse
		var err error
		defer close(result)
		response, err = client.DescribeExcpetionCount(request)
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

// DescribeExcpetionCountRequest is the request struct for api DescribeExcpetionCount
type DescribeExcpetionCountRequest struct {
	*requests.RpcRequest
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	DdosRegionId    string `position:"Query" name:"DdosRegionId"`
}

// DescribeExcpetionCountResponse is the response struct for api DescribeExcpetionCount
type DescribeExcpetionCountResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	ExceptionIpCount int    `json:"ExceptionIpCount" xml:"ExceptionIpCount"`
	ExpireTimeCount  int    `json:"ExpireTimeCount" xml:"ExpireTimeCount"`
}

// CreateDescribeExcpetionCountRequest creates a request to invoke DescribeExcpetionCount API
func CreateDescribeExcpetionCountRequest() (request *DescribeExcpetionCountRequest) {
	request = &DescribeExcpetionCountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddosbgp", "2018-07-20", "DescribeExcpetionCount", "ddosbgp", "openAPI")
	return
}

// CreateDescribeExcpetionCountResponse creates a response to parse from DescribeExcpetionCount response
func CreateDescribeExcpetionCountResponse() (response *DescribeExcpetionCountResponse) {
	response = &DescribeExcpetionCountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
