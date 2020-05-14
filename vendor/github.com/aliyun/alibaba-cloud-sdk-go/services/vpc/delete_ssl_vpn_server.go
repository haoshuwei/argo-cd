package vpc

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

// DeleteSslVpnServer invokes the vpc.DeleteSslVpnServer API synchronously
// api document: https://help.aliyun.com/api/vpc/deletesslvpnserver.html
func (client *Client) DeleteSslVpnServer(request *DeleteSslVpnServerRequest) (response *DeleteSslVpnServerResponse, err error) {
	response = CreateDeleteSslVpnServerResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteSslVpnServerWithChan invokes the vpc.DeleteSslVpnServer API asynchronously
// api document: https://help.aliyun.com/api/vpc/deletesslvpnserver.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteSslVpnServerWithChan(request *DeleteSslVpnServerRequest) (<-chan *DeleteSslVpnServerResponse, <-chan error) {
	responseChan := make(chan *DeleteSslVpnServerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteSslVpnServer(request)
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

// DeleteSslVpnServerWithCallback invokes the vpc.DeleteSslVpnServer API asynchronously
// api document: https://help.aliyun.com/api/vpc/deletesslvpnserver.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteSslVpnServerWithCallback(request *DeleteSslVpnServerRequest, callback func(response *DeleteSslVpnServerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteSslVpnServerResponse
		var err error
		defer close(result)
		response, err = client.DeleteSslVpnServer(request)
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

// DeleteSslVpnServerRequest is the request struct for api DeleteSslVpnServer
type DeleteSslVpnServerRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	SslVpnServerId       string           `position:"Query" name:"SslVpnServerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteSslVpnServerResponse is the response struct for api DeleteSslVpnServer
type DeleteSslVpnServerResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteSslVpnServerRequest creates a request to invoke DeleteSslVpnServer API
func CreateDeleteSslVpnServerRequest() (request *DeleteSslVpnServerRequest) {
	request = &DeleteSslVpnServerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DeleteSslVpnServer", "Vpc", "openAPI")
	return
}

// CreateDeleteSslVpnServerResponse creates a response to parse from DeleteSslVpnServer response
func CreateDeleteSslVpnServerResponse() (response *DeleteSslVpnServerResponse) {
	response = &DeleteSslVpnServerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
