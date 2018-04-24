package push

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

// UnbindPhone invokes the push.UnbindPhone API synchronously
// api document: https://help.aliyun.com/api/push/unbindphone.html
func (client *Client) UnbindPhone(request *UnbindPhoneRequest) (response *UnbindPhoneResponse, err error) {
	response = CreateUnbindPhoneResponse()
	err = client.DoAction(request, response)
	return
}

// UnbindPhoneWithChan invokes the push.UnbindPhone API asynchronously
// api document: https://help.aliyun.com/api/push/unbindphone.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnbindPhoneWithChan(request *UnbindPhoneRequest) (<-chan *UnbindPhoneResponse, <-chan error) {
	responseChan := make(chan *UnbindPhoneResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnbindPhone(request)
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

// UnbindPhoneWithCallback invokes the push.UnbindPhone API asynchronously
// api document: https://help.aliyun.com/api/push/unbindphone.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnbindPhoneWithCallback(request *UnbindPhoneRequest, callback func(response *UnbindPhoneResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnbindPhoneResponse
		var err error
		defer close(result)
		response, err = client.UnbindPhone(request)
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

// UnbindPhoneRequest is the request struct for api UnbindPhone
type UnbindPhoneRequest struct {
	*requests.RpcRequest
	AppKey   requests.Integer `position:"Query" name:"AppKey"`
	DeviceId string           `position:"Query" name:"DeviceId"`
}

// UnbindPhoneResponse is the response struct for api UnbindPhone
type UnbindPhoneResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUnbindPhoneRequest creates a request to invoke UnbindPhone API
func CreateUnbindPhoneRequest() (request *UnbindPhoneRequest) {
	request = &UnbindPhoneRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Push", "2016-08-01", "UnbindPhone", "", "")
	return
}

// CreateUnbindPhoneResponse creates a response to parse from UnbindPhone response
func CreateUnbindPhoneResponse() (response *UnbindPhoneResponse) {
	response = &UnbindPhoneResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
