package cms

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

// DeleteMetricRuleTargets invokes the cms.DeleteMetricRuleTargets API synchronously
// api document: https://help.aliyun.com/api/cms/deletemetricruletargets.html
func (client *Client) DeleteMetricRuleTargets(request *DeleteMetricRuleTargetsRequest) (response *DeleteMetricRuleTargetsResponse, err error) {
	response = CreateDeleteMetricRuleTargetsResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteMetricRuleTargetsWithChan invokes the cms.DeleteMetricRuleTargets API asynchronously
// api document: https://help.aliyun.com/api/cms/deletemetricruletargets.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteMetricRuleTargetsWithChan(request *DeleteMetricRuleTargetsRequest) (<-chan *DeleteMetricRuleTargetsResponse, <-chan error) {
	responseChan := make(chan *DeleteMetricRuleTargetsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteMetricRuleTargets(request)
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

// DeleteMetricRuleTargetsWithCallback invokes the cms.DeleteMetricRuleTargets API asynchronously
// api document: https://help.aliyun.com/api/cms/deletemetricruletargets.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteMetricRuleTargetsWithCallback(request *DeleteMetricRuleTargetsRequest, callback func(response *DeleteMetricRuleTargetsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteMetricRuleTargetsResponse
		var err error
		defer close(result)
		response, err = client.DeleteMetricRuleTargets(request)
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

// DeleteMetricRuleTargetsRequest is the request struct for api DeleteMetricRuleTargets
type DeleteMetricRuleTargetsRequest struct {
	*requests.RpcRequest
	TargetIds *[]string `position:"Query" name:"TargetIds"  type:"Repeated"`
	RuleId    string    `position:"Query" name:"RuleId"`
}

// DeleteMetricRuleTargetsResponse is the response struct for api DeleteMetricRuleTargets
type DeleteMetricRuleTargetsResponse struct {
	*responses.BaseResponse
	Success   bool    `json:"Success" xml:"Success"`
	Code      string  `json:"Code" xml:"Code"`
	Message   string  `json:"Message" xml:"Message"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	FailIds   FailIds `json:"FailIds" xml:"FailIds"`
}

// CreateDeleteMetricRuleTargetsRequest creates a request to invoke DeleteMetricRuleTargets API
func CreateDeleteMetricRuleTargetsRequest() (request *DeleteMetricRuleTargetsRequest) {
	request = &DeleteMetricRuleTargetsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DeleteMetricRuleTargets", "cms", "openAPI")
	return
}

// CreateDeleteMetricRuleTargetsResponse creates a response to parse from DeleteMetricRuleTargets response
func CreateDeleteMetricRuleTargetsResponse() (response *DeleteMetricRuleTargetsResponse) {
	response = &DeleteMetricRuleTargetsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
