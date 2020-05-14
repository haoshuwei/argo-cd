package cs

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

// DescirbeWorkflow invokes the cs.DescirbeWorkflow API synchronously
// api document: https://help.aliyun.com/api/cs/descirbeworkflow.html
func (client *Client) DescirbeWorkflow(request *DescirbeWorkflowRequest) (response *DescirbeWorkflowResponse, err error) {
	response = CreateDescirbeWorkflowResponse()
	err = client.DoAction(request, response)
	return
}

// DescirbeWorkflowWithChan invokes the cs.DescirbeWorkflow API asynchronously
// api document: https://help.aliyun.com/api/cs/descirbeworkflow.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescirbeWorkflowWithChan(request *DescirbeWorkflowRequest) (<-chan *DescirbeWorkflowResponse, <-chan error) {
	responseChan := make(chan *DescirbeWorkflowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescirbeWorkflow(request)
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

// DescirbeWorkflowWithCallback invokes the cs.DescirbeWorkflow API asynchronously
// api document: https://help.aliyun.com/api/cs/descirbeworkflow.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescirbeWorkflowWithCallback(request *DescirbeWorkflowRequest, callback func(response *DescirbeWorkflowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescirbeWorkflowResponse
		var err error
		defer close(result)
		response, err = client.DescirbeWorkflow(request)
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

// DescirbeWorkflowRequest is the request struct for api DescirbeWorkflow
type DescirbeWorkflowRequest struct {
	*requests.RoaRequest
	WorkflowName string `position:"Path" name:"workflowName"`
}

// DescirbeWorkflowResponse is the response struct for api DescirbeWorkflow
type DescirbeWorkflowResponse struct {
	*responses.BaseResponse
}

// CreateDescirbeWorkflowRequest creates a request to invoke DescirbeWorkflow API
func CreateDescirbeWorkflowRequest() (request *DescirbeWorkflowRequest) {
	request = &DescirbeWorkflowRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "DescirbeWorkflow", "/gs/workflow/[workflowName]", "csk", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescirbeWorkflowResponse creates a response to parse from DescirbeWorkflow response
func CreateDescirbeWorkflowResponse() (response *DescirbeWorkflowResponse) {
	response = &DescirbeWorkflowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
