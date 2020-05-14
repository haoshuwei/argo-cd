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

// SubmitMediaFpDeleteJob invokes the mts.SubmitMediaFpDeleteJob API synchronously
// api document: https://help.aliyun.com/api/mts/submitmediafpdeletejob.html
func (client *Client) SubmitMediaFpDeleteJob(request *SubmitMediaFpDeleteJobRequest) (response *SubmitMediaFpDeleteJobResponse, err error) {
	response = CreateSubmitMediaFpDeleteJobResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitMediaFpDeleteJobWithChan invokes the mts.SubmitMediaFpDeleteJob API asynchronously
// api document: https://help.aliyun.com/api/mts/submitmediafpdeletejob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitMediaFpDeleteJobWithChan(request *SubmitMediaFpDeleteJobRequest) (<-chan *SubmitMediaFpDeleteJobResponse, <-chan error) {
	responseChan := make(chan *SubmitMediaFpDeleteJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitMediaFpDeleteJob(request)
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

// SubmitMediaFpDeleteJobWithCallback invokes the mts.SubmitMediaFpDeleteJob API asynchronously
// api document: https://help.aliyun.com/api/mts/submitmediafpdeletejob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitMediaFpDeleteJobWithCallback(request *SubmitMediaFpDeleteJobRequest, callback func(response *SubmitMediaFpDeleteJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitMediaFpDeleteJobResponse
		var err error
		defer close(result)
		response, err = client.SubmitMediaFpDeleteJob(request)
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

// SubmitMediaFpDeleteJobRequest is the request struct for api SubmitMediaFpDeleteJob
type SubmitMediaFpDeleteJobRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	FpDBId               string           `position:"Query" name:"FpDBId"`
	UserData             string           `position:"Query" name:"UserData"`
	PrimaryKey           string           `position:"Query" name:"PrimaryKey"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PipelineId           string           `position:"Query" name:"PipelineId"`
}

// SubmitMediaFpDeleteJobResponse is the response struct for api SubmitMediaFpDeleteJob
type SubmitMediaFpDeleteJobResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	JobId     string `json:"JobId" xml:"JobId"`
}

// CreateSubmitMediaFpDeleteJobRequest creates a request to invoke SubmitMediaFpDeleteJob API
func CreateSubmitMediaFpDeleteJobRequest() (request *SubmitMediaFpDeleteJobRequest) {
	request = &SubmitMediaFpDeleteJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "SubmitMediaFpDeleteJob", "", "")
	return
}

// CreateSubmitMediaFpDeleteJobResponse creates a response to parse from SubmitMediaFpDeleteJob response
func CreateSubmitMediaFpDeleteJobResponse() (response *SubmitMediaFpDeleteJobResponse) {
	response = &SubmitMediaFpDeleteJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
