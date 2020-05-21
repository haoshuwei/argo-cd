package oos

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

// GenerateExecutionPolicy invokes the oos.GenerateExecutionPolicy API synchronously
// api document: https://help.aliyun.com/api/oos/generateexecutionpolicy.html
func (client *Client) GenerateExecutionPolicy(request *GenerateExecutionPolicyRequest) (response *GenerateExecutionPolicyResponse, err error) {
	response = CreateGenerateExecutionPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// GenerateExecutionPolicyWithChan invokes the oos.GenerateExecutionPolicy API asynchronously
// api document: https://help.aliyun.com/api/oos/generateexecutionpolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GenerateExecutionPolicyWithChan(request *GenerateExecutionPolicyRequest) (<-chan *GenerateExecutionPolicyResponse, <-chan error) {
	responseChan := make(chan *GenerateExecutionPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GenerateExecutionPolicy(request)
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

// GenerateExecutionPolicyWithCallback invokes the oos.GenerateExecutionPolicy API asynchronously
// api document: https://help.aliyun.com/api/oos/generateexecutionpolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GenerateExecutionPolicyWithCallback(request *GenerateExecutionPolicyRequest, callback func(response *GenerateExecutionPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GenerateExecutionPolicyResponse
		var err error
		defer close(result)
		response, err = client.GenerateExecutionPolicy(request)
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

// GenerateExecutionPolicyRequest is the request struct for api GenerateExecutionPolicy
type GenerateExecutionPolicyRequest struct {
	*requests.RpcRequest
	TemplateVersion string `position:"Query" name:"TemplateVersion"`
	TemplateName    string `position:"Query" name:"TemplateName"`
}

// GenerateExecutionPolicyResponse is the response struct for api GenerateExecutionPolicy
type GenerateExecutionPolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Policy    string `json:"Policy" xml:"Policy"`
}

// CreateGenerateExecutionPolicyRequest creates a request to invoke GenerateExecutionPolicy API
func CreateGenerateExecutionPolicyRequest() (request *GenerateExecutionPolicyRequest) {
	request = &GenerateExecutionPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("oos", "2019-06-01", "GenerateExecutionPolicy", "oos", "openAPI")
	return
}

// CreateGenerateExecutionPolicyResponse creates a response to parse from GenerateExecutionPolicy response
func CreateGenerateExecutionPolicyResponse() (response *GenerateExecutionPolicyResponse) {
	response = &GenerateExecutionPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
