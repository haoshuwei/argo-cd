package iot

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

// UpdateProductFilterConfig invokes the iot.UpdateProductFilterConfig API synchronously
// api document: https://help.aliyun.com/api/iot/updateproductfilterconfig.html
func (client *Client) UpdateProductFilterConfig(request *UpdateProductFilterConfigRequest) (response *UpdateProductFilterConfigResponse, err error) {
	response = CreateUpdateProductFilterConfigResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateProductFilterConfigWithChan invokes the iot.UpdateProductFilterConfig API asynchronously
// api document: https://help.aliyun.com/api/iot/updateproductfilterconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateProductFilterConfigWithChan(request *UpdateProductFilterConfigRequest) (<-chan *UpdateProductFilterConfigResponse, <-chan error) {
	responseChan := make(chan *UpdateProductFilterConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateProductFilterConfig(request)
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

// UpdateProductFilterConfigWithCallback invokes the iot.UpdateProductFilterConfig API asynchronously
// api document: https://help.aliyun.com/api/iot/updateproductfilterconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateProductFilterConfigWithCallback(request *UpdateProductFilterConfigRequest, callback func(response *UpdateProductFilterConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateProductFilterConfigResponse
		var err error
		defer close(result)
		response, err = client.UpdateProductFilterConfig(request)
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

// UpdateProductFilterConfigRequest is the request struct for api UpdateProductFilterConfig
type UpdateProductFilterConfigRequest struct {
	*requests.RpcRequest
	PropertyTimestampFilter requests.Boolean `position:"Query" name:"PropertyTimestampFilter"`
	ProductKey              string           `position:"Query" name:"ProductKey"`
	ApiProduct              string           `position:"Body" name:"ApiProduct"`
	ApiRevision             string           `position:"Body" name:"ApiRevision"`
	PropertyValueFilter     requests.Boolean `position:"Query" name:"PropertyValueFilter"`
}

// UpdateProductFilterConfigResponse is the response struct for api UpdateProductFilterConfig
type UpdateProductFilterConfigResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateUpdateProductFilterConfigRequest creates a request to invoke UpdateProductFilterConfig API
func CreateUpdateProductFilterConfigRequest() (request *UpdateProductFilterConfigRequest) {
	request = &UpdateProductFilterConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "UpdateProductFilterConfig", "Iot", "openAPI")
	return
}

// CreateUpdateProductFilterConfigResponse creates a response to parse from UpdateProductFilterConfig response
func CreateUpdateProductFilterConfigResponse() (response *UpdateProductFilterConfigResponse) {
	response = &UpdateProductFilterConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
