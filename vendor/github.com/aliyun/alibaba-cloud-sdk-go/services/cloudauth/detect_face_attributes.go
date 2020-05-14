package cloudauth

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

// DetectFaceAttributes invokes the cloudauth.DetectFaceAttributes API synchronously
// api document: https://help.aliyun.com/api/cloudauth/detectfaceattributes.html
func (client *Client) DetectFaceAttributes(request *DetectFaceAttributesRequest) (response *DetectFaceAttributesResponse, err error) {
	response = CreateDetectFaceAttributesResponse()
	err = client.DoAction(request, response)
	return
}

// DetectFaceAttributesWithChan invokes the cloudauth.DetectFaceAttributes API asynchronously
// api document: https://help.aliyun.com/api/cloudauth/detectfaceattributes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DetectFaceAttributesWithChan(request *DetectFaceAttributesRequest) (<-chan *DetectFaceAttributesResponse, <-chan error) {
	responseChan := make(chan *DetectFaceAttributesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetectFaceAttributes(request)
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

// DetectFaceAttributesWithCallback invokes the cloudauth.DetectFaceAttributes API asynchronously
// api document: https://help.aliyun.com/api/cloudauth/detectfaceattributes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DetectFaceAttributesWithCallback(request *DetectFaceAttributesRequest, callback func(response *DetectFaceAttributesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetectFaceAttributesResponse
		var err error
		defer close(result)
		response, err = client.DetectFaceAttributes(request)
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

// DetectFaceAttributesRequest is the request struct for api DetectFaceAttributes
type DetectFaceAttributesRequest struct {
	*requests.RpcRequest
	MaxNumPhotosPerCategory requests.Integer `position:"Body" name:"MaxNumPhotosPerCategory"`
	MaxFaceNum              requests.Integer `position:"Body" name:"MaxFaceNum"`
	ResourceOwnerId         requests.Integer `position:"Query" name:"ResourceOwnerId"`
	RetAttributes           string           `position:"Body" name:"RetAttributes"`
	ClientTag               string           `position:"Body" name:"ClientTag"`
	SourceIp                string           `position:"Query" name:"SourceIp"`
	MaterialValue           string           `position:"Body" name:"MaterialValue"`
	DontSaveDB              requests.Boolean `position:"Body" name:"DontSaveDB"`
}

// DetectFaceAttributesResponse is the response struct for api DetectFaceAttributes
type DetectFaceAttributesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDetectFaceAttributesRequest creates a request to invoke DetectFaceAttributes API
func CreateDetectFaceAttributesRequest() (request *DetectFaceAttributesRequest) {
	request = &DetectFaceAttributesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2018-09-16", "DetectFaceAttributes", "cloudauth", "openAPI")
	return
}

// CreateDetectFaceAttributesResponse creates a response to parse from DetectFaceAttributes response
func CreateDetectFaceAttributesResponse() (response *DetectFaceAttributesResponse) {
	response = &DetectFaceAttributesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
