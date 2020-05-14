package trademark

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

// DeleteTmMonitorRule invokes the trademark.DeleteTmMonitorRule API synchronously
// api document: https://help.aliyun.com/api/trademark/deletetmmonitorrule.html
func (client *Client) DeleteTmMonitorRule(request *DeleteTmMonitorRuleRequest) (response *DeleteTmMonitorRuleResponse, err error) {
	response = CreateDeleteTmMonitorRuleResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteTmMonitorRuleWithChan invokes the trademark.DeleteTmMonitorRule API asynchronously
// api document: https://help.aliyun.com/api/trademark/deletetmmonitorrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteTmMonitorRuleWithChan(request *DeleteTmMonitorRuleRequest) (<-chan *DeleteTmMonitorRuleResponse, <-chan error) {
	responseChan := make(chan *DeleteTmMonitorRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteTmMonitorRule(request)
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

// DeleteTmMonitorRuleWithCallback invokes the trademark.DeleteTmMonitorRule API asynchronously
// api document: https://help.aliyun.com/api/trademark/deletetmmonitorrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteTmMonitorRuleWithCallback(request *DeleteTmMonitorRuleRequest, callback func(response *DeleteTmMonitorRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteTmMonitorRuleResponse
		var err error
		defer close(result)
		response, err = client.DeleteTmMonitorRule(request)
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

// DeleteTmMonitorRuleRequest is the request struct for api DeleteTmMonitorRule
type DeleteTmMonitorRuleRequest struct {
	*requests.RpcRequest
	Id requests.Integer `position:"Query" name:"Id"`
}

// DeleteTmMonitorRuleResponse is the response struct for api DeleteTmMonitorRule
type DeleteTmMonitorRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
}

// CreateDeleteTmMonitorRuleRequest creates a request to invoke DeleteTmMonitorRule API
func CreateDeleteTmMonitorRuleRequest() (request *DeleteTmMonitorRuleRequest) {
	request = &DeleteTmMonitorRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Trademark", "2018-07-24", "DeleteTmMonitorRule", "trademark", "openAPI")
	return
}

// CreateDeleteTmMonitorRuleResponse creates a response to parse from DeleteTmMonitorRule response
func CreateDeleteTmMonitorRuleResponse() (response *DeleteTmMonitorRuleResponse) {
	response = &DeleteTmMonitorRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
