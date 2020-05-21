package netana

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

// DescribeNetworkQuotas invokes the netana.DescribeNetworkQuotas API synchronously
// api document: https://help.aliyun.com/api/netana/describenetworkquotas.html
func (client *Client) DescribeNetworkQuotas(request *DescribeNetworkQuotasRequest) (response *DescribeNetworkQuotasResponse, err error) {
	response = CreateDescribeNetworkQuotasResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeNetworkQuotasWithChan invokes the netana.DescribeNetworkQuotas API asynchronously
// api document: https://help.aliyun.com/api/netana/describenetworkquotas.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeNetworkQuotasWithChan(request *DescribeNetworkQuotasRequest) (<-chan *DescribeNetworkQuotasResponse, <-chan error) {
	responseChan := make(chan *DescribeNetworkQuotasResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeNetworkQuotas(request)
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

// DescribeNetworkQuotasWithCallback invokes the netana.DescribeNetworkQuotas API asynchronously
// api document: https://help.aliyun.com/api/netana/describenetworkquotas.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeNetworkQuotasWithCallback(request *DescribeNetworkQuotasRequest, callback func(response *DescribeNetworkQuotasResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeNetworkQuotasResponse
		var err error
		defer close(result)
		response, err = client.DescribeNetworkQuotas(request)
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

// DescribeNetworkQuotasRequest is the request struct for api DescribeNetworkQuotas
type DescribeNetworkQuotasRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Product              string           `position:"Query" name:"Product"`
	QuotaPublicityName   string           `position:"Query" name:"QuotaPublicityName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceType         string           `position:"Query" name:"ResourceType"`
}

// DescribeNetworkQuotasResponse is the response struct for api DescribeNetworkQuotas
type DescribeNetworkQuotasResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	TotalCount    int           `json:"TotalCount" xml:"TotalCount"`
	NetworkQuotas NetworkQuotas `json:"NetworkQuotas" xml:"NetworkQuotas"`
}

// CreateDescribeNetworkQuotasRequest creates a request to invoke DescribeNetworkQuotas API
func CreateDescribeNetworkQuotasRequest() (request *DescribeNetworkQuotasRequest) {
	request = &DescribeNetworkQuotasRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Netana", "2018-10-18", "DescribeNetworkQuotas", "Netana", "openAPI")
	return
}

// CreateDescribeNetworkQuotasResponse creates a response to parse from DescribeNetworkQuotas response
func CreateDescribeNetworkQuotasResponse() (response *DescribeNetworkQuotasResponse) {
	response = &DescribeNetworkQuotasResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
