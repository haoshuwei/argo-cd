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

// GetBatchEvaluateTaskReport invokes the drds.GetBatchEvaluateTaskReport API synchronously
// api document: https://help.aliyun.com/api/drds/getbatchevaluatetaskreport.html
func (client *Client) GetBatchEvaluateTaskReport(request *GetBatchEvaluateTaskReportRequest) (response *GetBatchEvaluateTaskReportResponse, err error) {
	response = CreateGetBatchEvaluateTaskReportResponse()
	err = client.DoAction(request, response)
	return
}

// GetBatchEvaluateTaskReportWithChan invokes the drds.GetBatchEvaluateTaskReport API asynchronously
// api document: https://help.aliyun.com/api/drds/getbatchevaluatetaskreport.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetBatchEvaluateTaskReportWithChan(request *GetBatchEvaluateTaskReportRequest) (<-chan *GetBatchEvaluateTaskReportResponse, <-chan error) {
	responseChan := make(chan *GetBatchEvaluateTaskReportResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetBatchEvaluateTaskReport(request)
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

// GetBatchEvaluateTaskReportWithCallback invokes the drds.GetBatchEvaluateTaskReport API asynchronously
// api document: https://help.aliyun.com/api/drds/getbatchevaluatetaskreport.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetBatchEvaluateTaskReportWithCallback(request *GetBatchEvaluateTaskReportRequest, callback func(response *GetBatchEvaluateTaskReportResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetBatchEvaluateTaskReportResponse
		var err error
		defer close(result)
		response, err = client.GetBatchEvaluateTaskReport(request)
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

// GetBatchEvaluateTaskReportRequest is the request struct for api GetBatchEvaluateTaskReport
type GetBatchEvaluateTaskReportRequest struct {
	*requests.RpcRequest
	BatchEvaluateTaskId requests.Integer `position:"Query" name:"BatchEvaluateTaskId"`
}

// GetBatchEvaluateTaskReportResponse is the response struct for api GetBatchEvaluateTaskReport
type GetBatchEvaluateTaskReportResponse struct {
	*responses.BaseResponse
	RequestId string                           `json:"RequestId" xml:"RequestId"`
	Success   bool                             `json:"Success" xml:"Success"`
	Data      DataInGetBatchEvaluateTaskReport `json:"Data" xml:"Data"`
}

// CreateGetBatchEvaluateTaskReportRequest creates a request to invoke GetBatchEvaluateTaskReport API
func CreateGetBatchEvaluateTaskReportRequest() (request *GetBatchEvaluateTaskReportRequest) {
	request = &GetBatchEvaluateTaskReportRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "GetBatchEvaluateTaskReport", "Drds", "openAPI")
	return
}

// CreateGetBatchEvaluateTaskReportResponse creates a response to parse from GetBatchEvaluateTaskReport response
func CreateGetBatchEvaluateTaskReportResponse() (response *GetBatchEvaluateTaskReportResponse) {
	response = &GetBatchEvaluateTaskReportResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
