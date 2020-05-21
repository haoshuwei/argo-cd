package drds

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

// CreateEvaluateDataImportTask invokes the drds.CreateEvaluateDataImportTask API synchronously
// api document: https://help.aliyun.com/api/drds/createevaluatedataimporttask.html
func (client *Client) CreateEvaluateDataImportTask(request *CreateEvaluateDataImportTaskRequest) (response *CreateEvaluateDataImportTaskResponse, err error) {
	response = CreateCreateEvaluateDataImportTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateEvaluateDataImportTaskWithChan invokes the drds.CreateEvaluateDataImportTask API asynchronously
// api document: https://help.aliyun.com/api/drds/createevaluatedataimporttask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateEvaluateDataImportTaskWithChan(request *CreateEvaluateDataImportTaskRequest) (<-chan *CreateEvaluateDataImportTaskResponse, <-chan error) {
	responseChan := make(chan *CreateEvaluateDataImportTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateEvaluateDataImportTask(request)
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

// CreateEvaluateDataImportTaskWithCallback invokes the drds.CreateEvaluateDataImportTask API asynchronously
// api document: https://help.aliyun.com/api/drds/createevaluatedataimporttask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateEvaluateDataImportTaskWithCallback(request *CreateEvaluateDataImportTaskRequest, callback func(response *CreateEvaluateDataImportTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateEvaluateDataImportTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateEvaluateDataImportTask(request)
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

// CreateEvaluateDataImportTaskRequest is the request struct for api CreateEvaluateDataImportTask
type CreateEvaluateDataImportTaskRequest struct {
	*requests.RpcRequest
	ImportParam string `position:"Query" name:"ImportParam"`
}

// CreateEvaluateDataImportTaskResponse is the response struct for api CreateEvaluateDataImportTask
type CreateEvaluateDataImportTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	TaskId    int64  `json:"TaskId" xml:"TaskId"`
}

// CreateCreateEvaluateDataImportTaskRequest creates a request to invoke CreateEvaluateDataImportTask API
func CreateCreateEvaluateDataImportTaskRequest() (request *CreateEvaluateDataImportTaskRequest) {
	request = &CreateEvaluateDataImportTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "CreateEvaluateDataImportTask", "Drds", "openAPI")
	return
}

// CreateCreateEvaluateDataImportTaskResponse creates a response to parse from CreateEvaluateDataImportTask response
func CreateCreateEvaluateDataImportTaskResponse() (response *CreateEvaluateDataImportTaskResponse) {
	response = &CreateEvaluateDataImportTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
