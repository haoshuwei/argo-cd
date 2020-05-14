package cassandra

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

// CreateBackupPlan invokes the cassandra.CreateBackupPlan API synchronously
// api document: https://help.aliyun.com/api/cassandra/createbackupplan.html
func (client *Client) CreateBackupPlan(request *CreateBackupPlanRequest) (response *CreateBackupPlanResponse, err error) {
	response = CreateCreateBackupPlanResponse()
	err = client.DoAction(request, response)
	return
}

// CreateBackupPlanWithChan invokes the cassandra.CreateBackupPlan API asynchronously
// api document: https://help.aliyun.com/api/cassandra/createbackupplan.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateBackupPlanWithChan(request *CreateBackupPlanRequest) (<-chan *CreateBackupPlanResponse, <-chan error) {
	responseChan := make(chan *CreateBackupPlanResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateBackupPlan(request)
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

// CreateBackupPlanWithCallback invokes the cassandra.CreateBackupPlan API asynchronously
// api document: https://help.aliyun.com/api/cassandra/createbackupplan.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateBackupPlanWithCallback(request *CreateBackupPlanRequest, callback func(response *CreateBackupPlanResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateBackupPlanResponse
		var err error
		defer close(result)
		response, err = client.CreateBackupPlan(request)
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

// CreateBackupPlanRequest is the request struct for api CreateBackupPlan
type CreateBackupPlanRequest struct {
	*requests.RpcRequest
	RetentionPeriod requests.Integer `position:"Query" name:"RetentionPeriod"`
	DataCenterId    string           `position:"Query" name:"DataCenterId"`
	Active          requests.Boolean `position:"Query" name:"Active"`
	ClusterId       string           `position:"Query" name:"ClusterId"`
	BackupTime      string           `position:"Query" name:"BackupTime"`
	BackupPeriod    string           `position:"Query" name:"BackupPeriod"`
}

// CreateBackupPlanResponse is the response struct for api CreateBackupPlan
type CreateBackupPlanResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateBackupPlanRequest creates a request to invoke CreateBackupPlan API
func CreateCreateBackupPlanRequest() (request *CreateBackupPlanRequest) {
	request = &CreateBackupPlanRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cassandra", "2019-01-01", "CreateBackupPlan", "Cassandra", "openAPI")
	return
}

// CreateCreateBackupPlanResponse creates a response to parse from CreateBackupPlan response
func CreateCreateBackupPlanResponse() (response *CreateBackupPlanResponse) {
	response = &CreateBackupPlanResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
