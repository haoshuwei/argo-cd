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

// UpdateImageLib invokes the green.UpdateImageLib API synchronously
// api document: https://help.aliyun.com/api/green/updateimagelib.html
func (client *Client) UpdateImageLib(request *UpdateImageLibRequest) (response *UpdateImageLibResponse, err error) {
	response = CreateUpdateImageLibResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateImageLibWithChan invokes the green.UpdateImageLib API asynchronously
// api document: https://help.aliyun.com/api/green/updateimagelib.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateImageLibWithChan(request *UpdateImageLibRequest) (<-chan *UpdateImageLibResponse, <-chan error) {
	responseChan := make(chan *UpdateImageLibResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateImageLib(request)
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

// UpdateImageLibWithCallback invokes the green.UpdateImageLib API asynchronously
// api document: https://help.aliyun.com/api/green/updateimagelib.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateImageLibWithCallback(request *UpdateImageLibRequest, callback func(response *UpdateImageLibResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateImageLibResponse
		var err error
		defer close(result)
		response, err = client.UpdateImageLib(request)
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

// UpdateImageLibRequest is the request struct for api UpdateImageLib
type UpdateImageLibRequest struct {
	*requests.RpcRequest
	Scene    string           `position:"Query" name:"Scene"`
	SourceIp string           `position:"Query" name:"SourceIp"`
	Enable   requests.Boolean `position:"Query" name:"Enable"`
	Id       requests.Integer `position:"Query" name:"Id"`
	BizTypes string           `position:"Query" name:"BizTypes"`
	Name     string           `position:"Query" name:"Name"`
	Category string           `position:"Query" name:"Category"`
}

// UpdateImageLibResponse is the response struct for api UpdateImageLib
type UpdateImageLibResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateImageLibRequest creates a request to invoke UpdateImageLib API
func CreateUpdateImageLibRequest() (request *UpdateImageLibRequest) {
	request = &UpdateImageLibRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-23", "UpdateImageLib", "green", "openAPI")
	return
}

// CreateUpdateImageLibResponse creates a response to parse from UpdateImageLib response
func CreateUpdateImageLibResponse() (response *UpdateImageLibResponse) {
	response = &UpdateImageLibResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
