package edas

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

// ListConvertableEcu invokes the edas.ListConvertableEcu API synchronously
// api document: https://help.aliyun.com/api/edas/listconvertableecu.html
func (client *Client) ListConvertableEcu(request *ListConvertableEcuRequest) (response *ListConvertableEcuResponse, err error) {
	response = CreateListConvertableEcuResponse()
	err = client.DoAction(request, response)
	return
}

// ListConvertableEcuWithChan invokes the edas.ListConvertableEcu API asynchronously
// api document: https://help.aliyun.com/api/edas/listconvertableecu.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListConvertableEcuWithChan(request *ListConvertableEcuRequest) (<-chan *ListConvertableEcuResponse, <-chan error) {
	responseChan := make(chan *ListConvertableEcuResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListConvertableEcu(request)
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

// ListConvertableEcuWithCallback invokes the edas.ListConvertableEcu API asynchronously
// api document: https://help.aliyun.com/api/edas/listconvertableecu.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListConvertableEcuWithCallback(request *ListConvertableEcuRequest, callback func(response *ListConvertableEcuResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListConvertableEcuResponse
		var err error
		defer close(result)
		response, err = client.ListConvertableEcu(request)
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

// ListConvertableEcuRequest is the request struct for api ListConvertableEcu
type ListConvertableEcuRequest struct {
	*requests.RoaRequest
	ClusterId string `position:"Query" name:"clusterId"`
}

// ListConvertableEcuResponse is the response struct for api ListConvertableEcu
type ListConvertableEcuResponse struct {
	*responses.BaseResponse
	Code         int          `json:"Code" xml:"Code"`
	Message      string       `json:"Message" xml:"Message"`
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	InstanceList InstanceList `json:"InstanceList" xml:"InstanceList"`
}

// CreateListConvertableEcuRequest creates a request to invoke ListConvertableEcu API
func CreateListConvertableEcuRequest() (request *ListConvertableEcuRequest) {
	request = &ListConvertableEcuRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "ListConvertableEcu", "/pop/v5/resource/convertable_ecu_list", "Edas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListConvertableEcuResponse creates a response to parse from ListConvertableEcu response
func CreateListConvertableEcuResponse() (response *ListConvertableEcuResponse) {
	response = &ListConvertableEcuResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
