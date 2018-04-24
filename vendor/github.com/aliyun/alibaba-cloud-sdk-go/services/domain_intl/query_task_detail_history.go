package domain_intl

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

// QueryTaskDetailHistory invokes the domain_intl.QueryTaskDetailHistory API synchronously
// api document: https://help.aliyun.com/api/domain-intl/querytaskdetailhistory.html
func (client *Client) QueryTaskDetailHistory(request *QueryTaskDetailHistoryRequest) (response *QueryTaskDetailHistoryResponse, err error) {
	response = CreateQueryTaskDetailHistoryResponse()
	err = client.DoAction(request, response)
	return
}

// QueryTaskDetailHistoryWithChan invokes the domain_intl.QueryTaskDetailHistory API asynchronously
// api document: https://help.aliyun.com/api/domain-intl/querytaskdetailhistory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryTaskDetailHistoryWithChan(request *QueryTaskDetailHistoryRequest) (<-chan *QueryTaskDetailHistoryResponse, <-chan error) {
	responseChan := make(chan *QueryTaskDetailHistoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryTaskDetailHistory(request)
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

// QueryTaskDetailHistoryWithCallback invokes the domain_intl.QueryTaskDetailHistory API asynchronously
// api document: https://help.aliyun.com/api/domain-intl/querytaskdetailhistory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryTaskDetailHistoryWithCallback(request *QueryTaskDetailHistoryRequest, callback func(response *QueryTaskDetailHistoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryTaskDetailHistoryResponse
		var err error
		defer close(result)
		response, err = client.QueryTaskDetailHistory(request)
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

// QueryTaskDetailHistoryRequest is the request struct for api QueryTaskDetailHistory
type QueryTaskDetailHistoryRequest struct {
	*requests.RpcRequest
	Lang               string           `position:"Query" name:"Lang"`
	UserClientIp       string           `position:"Query" name:"UserClientIp"`
	TaskNo             string           `position:"Query" name:"TaskNo"`
	DomainName         string           `position:"Query" name:"DomainName"`
	DomainNameCursor   string           `position:"Query" name:"DomainNameCursor"`
	TaskStatus         requests.Integer `position:"Query" name:"TaskStatus"`
	PageSize           requests.Integer `position:"Query" name:"PageSize"`
	TaskDetailNoCursor string           `position:"Query" name:"TaskDetailNoCursor"`
}

// QueryTaskDetailHistoryResponse is the response struct for api QueryTaskDetailHistory
type QueryTaskDetailHistoryResponse struct {
	*responses.BaseResponse
	RequestId         string              `json:"RequestId" xml:"RequestId"`
	PageSize          int                 `json:"PageSize" xml:"PageSize"`
	CurrentPageCursor CurrentPageCursor   `json:"CurrentPageCursor" xml:"CurrentPageCursor"`
	NextPageCursor    NextPageCursor      `json:"NextPageCursor" xml:"NextPageCursor"`
	PrePageCursor     PrePageCursor       `json:"PrePageCursor" xml:"PrePageCursor"`
	Objects           []TaskDetailHistory `json:"Objects" xml:"Objects"`
}

// CreateQueryTaskDetailHistoryRequest creates a request to invoke QueryTaskDetailHistory API
func CreateQueryTaskDetailHistoryRequest() (request *QueryTaskDetailHistoryRequest) {
	request = &QueryTaskDetailHistoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain-intl", "2017-12-18", "QueryTaskDetailHistory", "", "")
	return
}

// CreateQueryTaskDetailHistoryResponse creates a response to parse from QueryTaskDetailHistory response
func CreateQueryTaskDetailHistoryResponse() (response *QueryTaskDetailHistoryResponse) {
	response = &QueryTaskDetailHistoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
