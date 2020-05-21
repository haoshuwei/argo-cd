package ddoscoo

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

// DescribeDomains invokes the ddoscoo.DescribeDomains API synchronously
// api document: https://help.aliyun.com/api/ddoscoo/describedomains.html
func (client *Client) DescribeDomains(request *DescribeDomainsRequest) (response *DescribeDomainsResponse, err error) {
	response = CreateDescribeDomainsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainsWithChan invokes the ddoscoo.DescribeDomains API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/describedomains.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainsWithChan(request *DescribeDomainsRequest) (<-chan *DescribeDomainsResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomains(request)
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

// DescribeDomainsWithCallback invokes the ddoscoo.DescribeDomains API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/describedomains.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainsWithCallback(request *DescribeDomainsRequest, callback func(response *DescribeDomainsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomains(request)
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

// DescribeDomainsRequest is the request struct for api DescribeDomains
type DescribeDomainsRequest struct {
	*requests.RpcRequest
	ResourceGroupId string    `position:"Query" name:"ResourceGroupId"`
	SourceIp        string    `position:"Query" name:"SourceIp"`
	InstanceIds     *[]string `position:"Query" name:"InstanceIds"  type:"Repeated"`
	Lang            string    `position:"Query" name:"Lang"`
}

// DescribeDomainsResponse is the response struct for api DescribeDomains
type DescribeDomainsResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Domains   []string `json:"Domains" xml:"Domains"`
}

// CreateDescribeDomainsRequest creates a request to invoke DescribeDomains API
func CreateDescribeDomainsRequest() (request *DescribeDomainsRequest) {
	request = &DescribeDomainsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DescribeDomains", "ddoscoo", "openAPI")
	return
}

// CreateDescribeDomainsResponse creates a response to parse from DescribeDomains response
func CreateDescribeDomainsResponse() (response *DescribeDomainsResponse) {
	response = &DescribeDomainsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
