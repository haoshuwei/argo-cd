package vod

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

// UpdateAITemplate invokes the vod.UpdateAITemplate API synchronously
// api document: https://help.aliyun.com/api/vod/updateaitemplate.html
func (client *Client) UpdateAITemplate(request *UpdateAITemplateRequest) (response *UpdateAITemplateResponse, err error) {
	response = CreateUpdateAITemplateResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateAITemplateWithChan invokes the vod.UpdateAITemplate API asynchronously
// api document: https://help.aliyun.com/api/vod/updateaitemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateAITemplateWithChan(request *UpdateAITemplateRequest) (<-chan *UpdateAITemplateResponse, <-chan error) {
	responseChan := make(chan *UpdateAITemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateAITemplate(request)
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

// UpdateAITemplateWithCallback invokes the vod.UpdateAITemplate API asynchronously
// api document: https://help.aliyun.com/api/vod/updateaitemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateAITemplateWithCallback(request *UpdateAITemplateRequest, callback func(response *UpdateAITemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateAITemplateResponse
		var err error
		defer close(result)
		response, err = client.UpdateAITemplate(request)
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

// UpdateAITemplateRequest is the request struct for api UpdateAITemplate
type UpdateAITemplateRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	TemplateConfig       string           `position:"Query" name:"TemplateConfig"`
	TemplateName         string           `position:"Query" name:"TemplateName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	TemplateId           string           `position:"Query" name:"TemplateId"`
}

// UpdateAITemplateResponse is the response struct for api UpdateAITemplate
type UpdateAITemplateResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TemplateId string `json:"TemplateId" xml:"TemplateId"`
}

// CreateUpdateAITemplateRequest creates a request to invoke UpdateAITemplate API
func CreateUpdateAITemplateRequest() (request *UpdateAITemplateRequest) {
	request = &UpdateAITemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "UpdateAITemplate", "vod", "openAPI")
	return
}

// CreateUpdateAITemplateResponse creates a response to parse from UpdateAITemplate response
func CreateUpdateAITemplateResponse() (response *UpdateAITemplateResponse) {
	response = &UpdateAITemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
