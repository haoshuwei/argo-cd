package dyplsapi

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

// BuySecretNo invokes the dyplsapi.BuySecretNo API synchronously
// api document: https://help.aliyun.com/api/dyplsapi/buysecretno.html
func (client *Client) BuySecretNo(request *BuySecretNoRequest) (response *BuySecretNoResponse, err error) {
	response = CreateBuySecretNoResponse()
	err = client.DoAction(request, response)
	return
}

// BuySecretNoWithChan invokes the dyplsapi.BuySecretNo API asynchronously
// api document: https://help.aliyun.com/api/dyplsapi/buysecretno.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BuySecretNoWithChan(request *BuySecretNoRequest) (<-chan *BuySecretNoResponse, <-chan error) {
	responseChan := make(chan *BuySecretNoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BuySecretNo(request)
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

// BuySecretNoWithCallback invokes the dyplsapi.BuySecretNo API asynchronously
// api document: https://help.aliyun.com/api/dyplsapi/buysecretno.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BuySecretNoWithCallback(request *BuySecretNoRequest, callback func(response *BuySecretNoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BuySecretNoResponse
		var err error
		defer close(result)
		response, err = client.BuySecretNo(request)
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

// BuySecretNoRequest is the request struct for api BuySecretNo
type BuySecretNoRequest struct {
	*requests.RpcRequest
	SpecId               requests.Integer `position:"Query" name:"SpecId"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	City                 string           `position:"Query" name:"City"`
	SecretNo             string           `position:"Query" name:"SecretNo"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DisplayPool          requests.Boolean `position:"Query" name:"DisplayPool"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PoolKey              string           `position:"Query" name:"PoolKey"`
}

// BuySecretNoResponse is the response struct for api BuySecretNo
type BuySecretNoResponse struct {
	*responses.BaseResponse
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	Code             string           `json:"Code" xml:"Code"`
	Message          string           `json:"Message" xml:"Message"`
	SecretBuyInfoDTO SecretBuyInfoDTO `json:"SecretBuyInfoDTO" xml:"SecretBuyInfoDTO"`
}

// CreateBuySecretNoRequest creates a request to invoke BuySecretNo API
func CreateBuySecretNoRequest() (request *BuySecretNoRequest) {
	request = &BuySecretNoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dyplsapi", "2017-05-25", "BuySecretNo", "dypls", "openAPI")
	return
}

// CreateBuySecretNoResponse creates a response to parse from BuySecretNo response
func CreateBuySecretNoResponse() (response *BuySecretNoResponse) {
	response = &BuySecretNoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
