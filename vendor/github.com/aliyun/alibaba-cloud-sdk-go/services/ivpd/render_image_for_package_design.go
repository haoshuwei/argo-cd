package ivpd

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

// RenderImageForPackageDesign invokes the ivpd.RenderImageForPackageDesign API synchronously
// api document: https://help.aliyun.com/api/ivpd/renderimageforpackagedesign.html
func (client *Client) RenderImageForPackageDesign(request *RenderImageForPackageDesignRequest) (response *RenderImageForPackageDesignResponse, err error) {
	response = CreateRenderImageForPackageDesignResponse()
	err = client.DoAction(request, response)
	return
}

// RenderImageForPackageDesignWithChan invokes the ivpd.RenderImageForPackageDesign API asynchronously
// api document: https://help.aliyun.com/api/ivpd/renderimageforpackagedesign.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RenderImageForPackageDesignWithChan(request *RenderImageForPackageDesignRequest) (<-chan *RenderImageForPackageDesignResponse, <-chan error) {
	responseChan := make(chan *RenderImageForPackageDesignResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RenderImageForPackageDesign(request)
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

// RenderImageForPackageDesignWithCallback invokes the ivpd.RenderImageForPackageDesign API asynchronously
// api document: https://help.aliyun.com/api/ivpd/renderimageforpackagedesign.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RenderImageForPackageDesignWithCallback(request *RenderImageForPackageDesignRequest, callback func(response *RenderImageForPackageDesignResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RenderImageForPackageDesignResponse
		var err error
		defer close(result)
		response, err = client.RenderImageForPackageDesign(request)
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

// RenderImageForPackageDesignRequest is the request struct for api RenderImageForPackageDesign
type RenderImageForPackageDesignRequest struct {
	*requests.RpcRequest
	DisplayType  string                                    `position:"Body" name:"DisplayType"`
	MaterialName string                                    `position:"Body" name:"MaterialName"`
	JobId        string                                    `position:"Body" name:"JobId"`
	MaterialType string                                    `position:"Body" name:"MaterialType"`
	ModelType    string                                    `position:"Body" name:"ModelType"`
	TargetWidth  requests.Integer                          `position:"Body" name:"TargetWidth"`
	ElementList  *[]RenderImageForPackageDesignElementList `position:"Body" name:"ElementList"  type:"Repeated"`
	Category     string                                    `position:"Body" name:"Category"`
	TargetHeight requests.Integer                          `position:"Body" name:"TargetHeight"`
}

// RenderImageForPackageDesignElementList is a repeated param struct in RenderImageForPackageDesignRequest
type RenderImageForPackageDesignElementList struct {
	ImageUrl string `name:"ImageUrl"`
	SideName string `name:"SideName"`
}

// RenderImageForPackageDesignResponse is the response struct for api RenderImageForPackageDesign
type RenderImageForPackageDesignResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateRenderImageForPackageDesignRequest creates a request to invoke RenderImageForPackageDesign API
func CreateRenderImageForPackageDesignRequest() (request *RenderImageForPackageDesignRequest) {
	request = &RenderImageForPackageDesignRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ivpd", "2019-06-25", "RenderImageForPackageDesign", "", "")
	return
}

// CreateRenderImageForPackageDesignResponse creates a response to parse from RenderImageForPackageDesign response
func CreateRenderImageForPackageDesignResponse() (response *RenderImageForPackageDesignResponse) {
	response = &RenderImageForPackageDesignResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
