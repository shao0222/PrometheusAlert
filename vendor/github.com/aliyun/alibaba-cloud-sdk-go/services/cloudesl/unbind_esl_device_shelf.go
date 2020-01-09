package cloudesl

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

// UnbindEslDeviceShelf invokes the cloudesl.UnbindEslDeviceShelf API synchronously
// api document: https://help.aliyun.com/api/cloudesl/unbindesldeviceshelf.html
func (client *Client) UnbindEslDeviceShelf(request *UnbindEslDeviceShelfRequest) (response *UnbindEslDeviceShelfResponse, err error) {
	response = CreateUnbindEslDeviceShelfResponse()
	err = client.DoAction(request, response)
	return
}

// UnbindEslDeviceShelfWithChan invokes the cloudesl.UnbindEslDeviceShelf API asynchronously
// api document: https://help.aliyun.com/api/cloudesl/unbindesldeviceshelf.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnbindEslDeviceShelfWithChan(request *UnbindEslDeviceShelfRequest) (<-chan *UnbindEslDeviceShelfResponse, <-chan error) {
	responseChan := make(chan *UnbindEslDeviceShelfResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnbindEslDeviceShelf(request)
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

// UnbindEslDeviceShelfWithCallback invokes the cloudesl.UnbindEslDeviceShelf API asynchronously
// api document: https://help.aliyun.com/api/cloudesl/unbindesldeviceshelf.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnbindEslDeviceShelfWithCallback(request *UnbindEslDeviceShelfRequest, callback func(response *UnbindEslDeviceShelfResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnbindEslDeviceShelfResponse
		var err error
		defer close(result)
		response, err = client.UnbindEslDeviceShelf(request)
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

// UnbindEslDeviceShelfRequest is the request struct for api UnbindEslDeviceShelf
type UnbindEslDeviceShelfRequest struct {
	*requests.RpcRequest
	EslBarCode string `position:"Query" name:"EslBarCode"`
	StoreId    string `position:"Query" name:"StoreId"`
}

// UnbindEslDeviceShelfResponse is the response struct for api UnbindEslDeviceShelf
type UnbindEslDeviceShelfResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
}

// CreateUnbindEslDeviceShelfRequest creates a request to invoke UnbindEslDeviceShelf API
func CreateUnbindEslDeviceShelfRequest() (request *UnbindEslDeviceShelfRequest) {
	request = &UnbindEslDeviceShelfRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudesl", "2018-08-01", "UnbindEslDeviceShelf", "", "")
	return
}

// CreateUnbindEslDeviceShelfResponse creates a response to parse from UnbindEslDeviceShelf response
func CreateUnbindEslDeviceShelfResponse() (response *UnbindEslDeviceShelfResponse) {
	response = &UnbindEslDeviceShelfResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}