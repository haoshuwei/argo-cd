package live

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

// DescribeLiveCertificateList invokes the live.DescribeLiveCertificateList API synchronously
// api document: https://help.aliyun.com/api/live/describelivecertificatelist.html
func (client *Client) DescribeLiveCertificateList(request *DescribeLiveCertificateListRequest) (response *DescribeLiveCertificateListResponse, err error) {
	response = CreateDescribeLiveCertificateListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveCertificateListWithChan invokes the live.DescribeLiveCertificateList API asynchronously
// api document: https://help.aliyun.com/api/live/describelivecertificatelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveCertificateListWithChan(request *DescribeLiveCertificateListRequest) (<-chan *DescribeLiveCertificateListResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveCertificateListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveCertificateList(request)
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

// DescribeLiveCertificateListWithCallback invokes the live.DescribeLiveCertificateList API asynchronously
// api document: https://help.aliyun.com/api/live/describelivecertificatelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveCertificateListWithCallback(request *DescribeLiveCertificateListRequest, callback func(response *DescribeLiveCertificateListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveCertificateListResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveCertificateList(request)
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

// DescribeLiveCertificateListRequest is the request struct for api DescribeLiveCertificateList
type DescribeLiveCertificateListRequest struct {
	*requests.RpcRequest
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// DescribeLiveCertificateListResponse is the response struct for api DescribeLiveCertificateList
type DescribeLiveCertificateListResponse struct {
	*responses.BaseResponse
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	CertificateListModel CertificateListModel `json:"CertificateListModel" xml:"CertificateListModel"`
}

// CreateDescribeLiveCertificateListRequest creates a request to invoke DescribeLiveCertificateList API
func CreateDescribeLiveCertificateListRequest() (request *DescribeLiveCertificateListRequest) {
	request = &DescribeLiveCertificateListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveCertificateList", "live", "openAPI")
	return
}

// CreateDescribeLiveCertificateListResponse creates a response to parse from DescribeLiveCertificateList response
func CreateDescribeLiveCertificateListResponse() (response *DescribeLiveCertificateListResponse) {
	response = &DescribeLiveCertificateListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
