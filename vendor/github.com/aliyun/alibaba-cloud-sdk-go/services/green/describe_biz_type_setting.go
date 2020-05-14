package green

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

// DescribeBizTypeSetting invokes the green.DescribeBizTypeSetting API synchronously
// api document: https://help.aliyun.com/api/green/describebiztypesetting.html
func (client *Client) DescribeBizTypeSetting(request *DescribeBizTypeSettingRequest) (response *DescribeBizTypeSettingResponse, err error) {
	response = CreateDescribeBizTypeSettingResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeBizTypeSettingWithChan invokes the green.DescribeBizTypeSetting API asynchronously
// api document: https://help.aliyun.com/api/green/describebiztypesetting.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeBizTypeSettingWithChan(request *DescribeBizTypeSettingRequest) (<-chan *DescribeBizTypeSettingResponse, <-chan error) {
	responseChan := make(chan *DescribeBizTypeSettingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBizTypeSetting(request)
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

// DescribeBizTypeSettingWithCallback invokes the green.DescribeBizTypeSetting API asynchronously
// api document: https://help.aliyun.com/api/green/describebiztypesetting.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeBizTypeSettingWithCallback(request *DescribeBizTypeSettingRequest, callback func(response *DescribeBizTypeSettingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBizTypeSettingResponse
		var err error
		defer close(result)
		response, err = client.DescribeBizTypeSetting(request)
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

// DescribeBizTypeSettingRequest is the request struct for api DescribeBizTypeSetting
type DescribeBizTypeSettingRequest struct {
	*requests.RpcRequest
	ResourceType string `position:"Query" name:"ResourceType"`
	SourceIp     string `position:"Query" name:"SourceIp"`
	BizTypeName  string `position:"Query" name:"BizTypeName"`
}

// DescribeBizTypeSettingResponse is the response struct for api DescribeBizTypeSetting
type DescribeBizTypeSettingResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	Terrorism Terrorism `json:"Terrorism" xml:"Terrorism"`
	Porn      Porn      `json:"Porn" xml:"Porn"`
	Antispam  Antispam  `json:"Antispam" xml:"Antispam"`
}

// CreateDescribeBizTypeSettingRequest creates a request to invoke DescribeBizTypeSetting API
func CreateDescribeBizTypeSettingRequest() (request *DescribeBizTypeSettingRequest) {
	request = &DescribeBizTypeSettingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-23", "DescribeBizTypeSetting", "green", "openAPI")
	return
}

// CreateDescribeBizTypeSettingResponse creates a response to parse from DescribeBizTypeSetting response
func CreateDescribeBizTypeSettingResponse() (response *DescribeBizTypeSettingResponse) {
	response = &DescribeBizTypeSettingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
