package cdn

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

// DescribeDomainSrcTopUrlVisit invokes the cdn.DescribeDomainSrcTopUrlVisit API synchronously
// api document: https://help.aliyun.com/api/cdn/describedomainsrctopurlvisit.html
func (client *Client) DescribeDomainSrcTopUrlVisit(request *DescribeDomainSrcTopUrlVisitRequest) (response *DescribeDomainSrcTopUrlVisitResponse, err error) {
	response = CreateDescribeDomainSrcTopUrlVisitResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainSrcTopUrlVisitWithChan invokes the cdn.DescribeDomainSrcTopUrlVisit API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomainsrctopurlvisit.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainSrcTopUrlVisitWithChan(request *DescribeDomainSrcTopUrlVisitRequest) (<-chan *DescribeDomainSrcTopUrlVisitResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainSrcTopUrlVisitResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainSrcTopUrlVisit(request)
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

// DescribeDomainSrcTopUrlVisitWithCallback invokes the cdn.DescribeDomainSrcTopUrlVisit API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomainsrctopurlvisit.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainSrcTopUrlVisitWithCallback(request *DescribeDomainSrcTopUrlVisitRequest, callback func(response *DescribeDomainSrcTopUrlVisitResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainSrcTopUrlVisitResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainSrcTopUrlVisit(request)
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

// DescribeDomainSrcTopUrlVisitRequest is the request struct for api DescribeDomainSrcTopUrlVisit
type DescribeDomainSrcTopUrlVisitRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	DomainName string           `position:"Query" name:"DomainName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
	SortBy     string           `position:"Query" name:"SortBy"`
}

// DescribeDomainSrcTopUrlVisitResponse is the response struct for api DescribeDomainSrcTopUrlVisit
type DescribeDomainSrcTopUrlVisitResponse struct {
	*responses.BaseResponse
	RequestId  string                                   `json:"RequestId" xml:"RequestId"`
	DomainName string                                   `json:"DomainName" xml:"DomainName"`
	StartTime  string                                   `json:"StartTime" xml:"StartTime"`
	AllUrlList AllUrlListInDescribeDomainSrcTopUrlVisit `json:"AllUrlList" xml:"AllUrlList"`
	Url200List Url200ListInDescribeDomainSrcTopUrlVisit `json:"Url200List" xml:"Url200List"`
	Url300List Url300ListInDescribeDomainSrcTopUrlVisit `json:"Url300List" xml:"Url300List"`
	Url400List Url400ListInDescribeDomainSrcTopUrlVisit `json:"Url400List" xml:"Url400List"`
	Url500List Url500ListInDescribeDomainSrcTopUrlVisit `json:"Url500List" xml:"Url500List"`
}

// CreateDescribeDomainSrcTopUrlVisitRequest creates a request to invoke DescribeDomainSrcTopUrlVisit API
func CreateDescribeDomainSrcTopUrlVisitRequest() (request *DescribeDomainSrcTopUrlVisitRequest) {
	request = &DescribeDomainSrcTopUrlVisitRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeDomainSrcTopUrlVisit", "", "")
	return
}

// CreateDescribeDomainSrcTopUrlVisitResponse creates a response to parse from DescribeDomainSrcTopUrlVisit response
func CreateDescribeDomainSrcTopUrlVisitResponse() (response *DescribeDomainSrcTopUrlVisitResponse) {
	response = &DescribeDomainSrcTopUrlVisitResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
