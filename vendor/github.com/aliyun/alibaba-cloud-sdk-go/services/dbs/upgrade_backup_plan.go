package dbs

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

// UpgradeBackupPlan invokes the dbs.UpgradeBackupPlan API synchronously
// api document: https://help.aliyun.com/api/dbs/upgradebackupplan.html
func (client *Client) UpgradeBackupPlan(request *UpgradeBackupPlanRequest) (response *UpgradeBackupPlanResponse, err error) {
	response = CreateUpgradeBackupPlanResponse()
	err = client.DoAction(request, response)
	return
}

// UpgradeBackupPlanWithChan invokes the dbs.UpgradeBackupPlan API asynchronously
// api document: https://help.aliyun.com/api/dbs/upgradebackupplan.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpgradeBackupPlanWithChan(request *UpgradeBackupPlanRequest) (<-chan *UpgradeBackupPlanResponse, <-chan error) {
	responseChan := make(chan *UpgradeBackupPlanResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpgradeBackupPlan(request)
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

// UpgradeBackupPlanWithCallback invokes the dbs.UpgradeBackupPlan API asynchronously
// api document: https://help.aliyun.com/api/dbs/upgradebackupplan.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpgradeBackupPlanWithCallback(request *UpgradeBackupPlanRequest, callback func(response *UpgradeBackupPlanResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpgradeBackupPlanResponse
		var err error
		defer close(result)
		response, err = client.UpgradeBackupPlan(request)
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

// UpgradeBackupPlanRequest is the request struct for api UpgradeBackupPlan
type UpgradeBackupPlanRequest struct {
	*requests.RpcRequest
	ClientToken   string `position:"Query" name:"ClientToken"`
	BackupPlanId  string `position:"Query" name:"BackupPlanId"`
	OwnerId       string `position:"Query" name:"OwnerId"`
	InstanceClass string `position:"Query" name:"InstanceClass"`
}

// UpgradeBackupPlanResponse is the response struct for api UpgradeBackupPlan
type UpgradeBackupPlanResponse struct {
	*responses.BaseResponse
	Success        bool   `json:"Success" xml:"Success"`
	ErrCode        string `json:"ErrCode" xml:"ErrCode"`
	ErrMessage     string `json:"ErrMessage" xml:"ErrMessage"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	BackupPlanId   string `json:"BackupPlanId" xml:"BackupPlanId"`
	OrderId        string `json:"OrderId" xml:"OrderId"`
}

// CreateUpgradeBackupPlanRequest creates a request to invoke UpgradeBackupPlan API
func CreateUpgradeBackupPlanRequest() (request *UpgradeBackupPlanRequest) {
	request = &UpgradeBackupPlanRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dbs", "2019-03-06", "UpgradeBackupPlan", "cbs", "openAPI")
	return
}

// CreateUpgradeBackupPlanResponse creates a response to parse from UpgradeBackupPlan response
func CreateUpgradeBackupPlanResponse() (response *UpgradeBackupPlanResponse) {
	response = &UpgradeBackupPlanResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
