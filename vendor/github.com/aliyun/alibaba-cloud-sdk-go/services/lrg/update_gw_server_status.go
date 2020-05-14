package lrg

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

// UpdateGwServerStatus invokes the lrg.UpdateGwServerStatus API synchronously
// api document: https://help.aliyun.com/api/lrg/updategwserverstatus.html
func (client *Client) UpdateGwServerStatus(request *UpdateGwServerStatusRequest) (response *UpdateGwServerStatusResponse, err error) {
	response = CreateUpdateGwServerStatusResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateGwServerStatusWithChan invokes the lrg.UpdateGwServerStatus API asynchronously
// api document: https://help.aliyun.com/api/lrg/updategwserverstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateGwServerStatusWithChan(request *UpdateGwServerStatusRequest) (<-chan *UpdateGwServerStatusResponse, <-chan error) {
	responseChan := make(chan *UpdateGwServerStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateGwServerStatus(request)
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

// UpdateGwServerStatusWithCallback invokes the lrg.UpdateGwServerStatus API asynchronously
// api document: https://help.aliyun.com/api/lrg/updategwserverstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateGwServerStatusWithCallback(request *UpdateGwServerStatusRequest, callback func(response *UpdateGwServerStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateGwServerStatusResponse
		var err error
		defer close(result)
		response, err = client.UpdateGwServerStatus(request)
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

// UpdateGwServerStatusRequest is the request struct for api UpdateGwServerStatus
type UpdateGwServerStatusRequest struct {
	*requests.RoaRequest
	BigRegionName string `position:"Body" name:"big_region_name"`
	Status        string `position:"Body" name:"Status"`
}

// UpdateGwServerStatusResponse is the response struct for api UpdateGwServerStatus
type UpdateGwServerStatusResponse struct {
	*responses.BaseResponse
	Code    int                      `json:"code" xml:"code"`
	Message string                   `json:"message" xml:"message"`
	Success bool                     `json:"success" xml:"success"`
	Data    []map[string]interface{} `json:"data" xml:"data"`
}

// CreateUpdateGwServerStatusRequest creates a request to invoke UpdateGwServerStatus API
func CreateUpdateGwServerStatusRequest() (request *UpdateGwServerStatusRequest) {
	request = &UpdateGwServerStatusRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("LRG", "2019-10-10", "UpdateGwServerStatus", "/api/v2/gw/xgw/gw-server?action=updateStatus", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateGwServerStatusResponse creates a response to parse from UpdateGwServerStatus response
func CreateUpdateGwServerStatusResponse() (response *UpdateGwServerStatusResponse) {
	response = &UpdateGwServerStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
