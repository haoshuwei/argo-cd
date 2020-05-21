package mts

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

// UpdateMCTemplate invokes the mts.UpdateMCTemplate API synchronously
// api document: https://help.aliyun.com/api/mts/updatemctemplate.html
func (client *Client) UpdateMCTemplate(request *UpdateMCTemplateRequest) (response *UpdateMCTemplateResponse, err error) {
	response = CreateUpdateMCTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateMCTemplateWithChan invokes the mts.UpdateMCTemplate API asynchronously
// api document: https://help.aliyun.com/api/mts/updatemctemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateMCTemplateWithChan(request *UpdateMCTemplateRequest) (<-chan *UpdateMCTemplateResponse, <-chan error) {
	responseChan := make(chan *UpdateMCTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateMCTemplate(request)
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

// UpdateMCTemplateWithCallback invokes the mts.UpdateMCTemplate API asynchronously
// api document: https://help.aliyun.com/api/mts/updatemctemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateMCTemplateWithCallback(request *UpdateMCTemplateRequest, callback func(response *UpdateMCTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateMCTemplateResponse
		var err error
		defer close(result)
		response, err = client.UpdateMCTemplate(request)
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

// UpdateMCTemplateRequest is the request struct for api UpdateMCTemplate
type UpdateMCTemplateRequest struct {
	*requests.RpcRequest
	Politics             string           `position:"Query" name:"Politics"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Abuse                string           `position:"Query" name:"Abuse"`
	Qrcode               string           `position:"Query" name:"Qrcode"`
	Porn                 string           `position:"Query" name:"Porn"`
	Terrorism            string           `position:"Query" name:"Terrorism"`
	Logo                 string           `position:"Query" name:"Logo"`
	Live                 string           `position:"Query" name:"Live"`
	Contraband           string           `position:"Query" name:"Contraband"`
	Ad                   string           `position:"Query" name:"Ad"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	TemplateId           string           `position:"Query" name:"TemplateId"`
	Name                 string           `position:"Query" name:"Name"`
	Spam                 string           `position:"Query" name:"spam"`
}

// UpdateMCTemplateResponse is the response struct for api UpdateMCTemplate
type UpdateMCTemplateResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Template  Template `json:"Template" xml:"Template"`
}

// CreateUpdateMCTemplateRequest creates a request to invoke UpdateMCTemplate API
func CreateUpdateMCTemplateRequest() (request *UpdateMCTemplateRequest) {
	request = &UpdateMCTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "UpdateMCTemplate", "", "")
	return
}

// CreateUpdateMCTemplateResponse creates a response to parse from UpdateMCTemplate response
func CreateUpdateMCTemplateResponse() (response *UpdateMCTemplateResponse) {
	response = &UpdateMCTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
