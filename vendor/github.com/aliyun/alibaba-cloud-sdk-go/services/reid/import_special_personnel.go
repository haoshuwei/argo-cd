package reid

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

// ImportSpecialPersonnel invokes the reid.ImportSpecialPersonnel API synchronously
// api document: https://help.aliyun.com/api/reid/importspecialpersonnel.html
func (client *Client) ImportSpecialPersonnel(request *ImportSpecialPersonnelRequest) (response *ImportSpecialPersonnelResponse, err error) {
	response = CreateImportSpecialPersonnelResponse()
	err = client.DoAction(request, response)
	return
}

// ImportSpecialPersonnelWithChan invokes the reid.ImportSpecialPersonnel API asynchronously
// api document: https://help.aliyun.com/api/reid/importspecialpersonnel.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ImportSpecialPersonnelWithChan(request *ImportSpecialPersonnelRequest) (<-chan *ImportSpecialPersonnelResponse, <-chan error) {
	responseChan := make(chan *ImportSpecialPersonnelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ImportSpecialPersonnel(request)
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

// ImportSpecialPersonnelWithCallback invokes the reid.ImportSpecialPersonnel API asynchronously
// api document: https://help.aliyun.com/api/reid/importspecialpersonnel.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ImportSpecialPersonnelWithCallback(request *ImportSpecialPersonnelRequest, callback func(response *ImportSpecialPersonnelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ImportSpecialPersonnelResponse
		var err error
		defer close(result)
		response, err = client.ImportSpecialPersonnel(request)
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

// ImportSpecialPersonnelRequest is the request struct for api ImportSpecialPersonnel
type ImportSpecialPersonnelRequest struct {
	*requests.RpcRequest
	UkId        requests.Integer `position:"Body" name:"UkId"`
	Description string           `position:"Body" name:"Description"`
	ExternalId  string           `position:"Body" name:"ExternalId"`
	PersonType  string           `position:"Body" name:"PersonType"`
	Urls        string           `position:"Body" name:"Urls"`
	PersonName  string           `position:"Body" name:"PersonName"`
	StoreIds    string           `position:"Body" name:"StoreIds"`
	Status      string           `position:"Body" name:"Status"`
}

// ImportSpecialPersonnelResponse is the response struct for api ImportSpecialPersonnel
type ImportSpecialPersonnelResponse struct {
	*responses.BaseResponse
	ErrorCode            string               `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage         string               `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	Success              bool                 `json:"Success" xml:"Success"`
	SpecialPersonnelMaps SpecialPersonnelMaps `json:"SpecialPersonnelMaps" xml:"SpecialPersonnelMaps"`
}

// CreateImportSpecialPersonnelRequest creates a request to invoke ImportSpecialPersonnel API
func CreateImportSpecialPersonnelRequest() (request *ImportSpecialPersonnelRequest) {
	request = &ImportSpecialPersonnelRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("reid", "2019-09-28", "ImportSpecialPersonnel", "", "")
	return
}

// CreateImportSpecialPersonnelResponse creates a response to parse from ImportSpecialPersonnel response
func CreateImportSpecialPersonnelResponse() (response *ImportSpecialPersonnelResponse) {
	response = &ImportSpecialPersonnelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
