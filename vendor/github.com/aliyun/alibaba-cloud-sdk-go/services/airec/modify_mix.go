package airec

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

// ModifyMix invokes the airec.ModifyMix API synchronously
// api document: https://help.aliyun.com/api/airec/modifymix.html
func (client *Client) ModifyMix(request *ModifyMixRequest) (response *ModifyMixResponse, err error) {
	response = CreateModifyMixResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyMixWithChan invokes the airec.ModifyMix API asynchronously
// api document: https://help.aliyun.com/api/airec/modifymix.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyMixWithChan(request *ModifyMixRequest) (<-chan *ModifyMixResponse, <-chan error) {
	responseChan := make(chan *ModifyMixResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyMix(request)
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

// ModifyMixWithCallback invokes the airec.ModifyMix API asynchronously
// api document: https://help.aliyun.com/api/airec/modifymix.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyMixWithCallback(request *ModifyMixRequest, callback func(response *ModifyMixResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyMixResponse
		var err error
		defer close(result)
		response, err = client.ModifyMix(request)
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

// ModifyMixRequest is the request struct for api ModifyMix
type ModifyMixRequest struct {
	*requests.RoaRequest
	InstanceId string `position:"Path" name:"InstanceId"`
	Name       string `position:"Path" name:"Name"`
}

// ModifyMixResponse is the response struct for api ModifyMix
type ModifyMixResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateModifyMixRequest creates a request to invoke ModifyMix API
func CreateModifyMixRequest() (request *ModifyMixRequest) {
	request = &ModifyMixRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Airec", "2018-10-12", "ModifyMix", "/openapi/instances/[InstanceId]/mixes/[Name]", "airec", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateModifyMixResponse creates a response to parse from ModifyMix response
func CreateModifyMixResponse() (response *ModifyMixResponse) {
	response = &ModifyMixResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
