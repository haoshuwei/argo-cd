package baas

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

// UpdateAntChainContractContent invokes the baas.UpdateAntChainContractContent API synchronously
// api document: https://help.aliyun.com/api/baas/updateantchaincontractcontent.html
func (client *Client) UpdateAntChainContractContent(request *UpdateAntChainContractContentRequest) (response *UpdateAntChainContractContentResponse, err error) {
	response = CreateUpdateAntChainContractContentResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateAntChainContractContentWithChan invokes the baas.UpdateAntChainContractContent API asynchronously
// api document: https://help.aliyun.com/api/baas/updateantchaincontractcontent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateAntChainContractContentWithChan(request *UpdateAntChainContractContentRequest) (<-chan *UpdateAntChainContractContentResponse, <-chan error) {
	responseChan := make(chan *UpdateAntChainContractContentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateAntChainContractContent(request)
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

// UpdateAntChainContractContentWithCallback invokes the baas.UpdateAntChainContractContent API asynchronously
// api document: https://help.aliyun.com/api/baas/updateantchaincontractcontent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateAntChainContractContentWithCallback(request *UpdateAntChainContractContentRequest, callback func(response *UpdateAntChainContractContentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateAntChainContractContentResponse
		var err error
		defer close(result)
		response, err = client.UpdateAntChainContractContent(request)
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

// UpdateAntChainContractContentRequest is the request struct for api UpdateAntChainContractContent
type UpdateAntChainContractContentRequest struct {
	*requests.RpcRequest
	Content         string `position:"Body" name:"Content"`
	ParentContentId string `position:"Body" name:"ParentContentId"`
	ContentName     string `position:"Body" name:"ContentName"`
	ContentId       string `position:"Body" name:"ContentId"`
}

// UpdateAntChainContractContentResponse is the response struct for api UpdateAntChainContractContent
type UpdateAntChainContractContentResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    string `json:"Result" xml:"Result"`
}

// CreateUpdateAntChainContractContentRequest creates a request to invoke UpdateAntChainContractContent API
func CreateUpdateAntChainContractContentRequest() (request *UpdateAntChainContractContentRequest) {
	request = &UpdateAntChainContractContentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-12-21", "UpdateAntChainContractContent", "baas", "openAPI")
	return
}

// CreateUpdateAntChainContractContentResponse creates a response to parse from UpdateAntChainContractContent response
func CreateUpdateAntChainContractContentResponse() (response *UpdateAntChainContractContentResponse) {
	response = &UpdateAntChainContractContentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
