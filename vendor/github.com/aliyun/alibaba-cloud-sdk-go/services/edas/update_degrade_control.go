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

// UpdateDegradeControl invokes the edas.UpdateDegradeControl API synchronously
// api document: https://help.aliyun.com/api/edas/updatedegradecontrol.html
func (client *Client) UpdateDegradeControl(request *UpdateDegradeControlRequest) (response *UpdateDegradeControlResponse, err error) {
	response = CreateUpdateDegradeControlResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateDegradeControlWithChan invokes the edas.UpdateDegradeControl API asynchronously
// api document: https://help.aliyun.com/api/edas/updatedegradecontrol.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateDegradeControlWithChan(request *UpdateDegradeControlRequest) (<-chan *UpdateDegradeControlResponse, <-chan error) {
	responseChan := make(chan *UpdateDegradeControlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateDegradeControl(request)
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

// UpdateDegradeControlWithCallback invokes the edas.UpdateDegradeControl API asynchronously
// api document: https://help.aliyun.com/api/edas/updatedegradecontrol.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateDegradeControlWithCallback(request *UpdateDegradeControlRequest, callback func(response *UpdateDegradeControlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateDegradeControlResponse
		var err error
		defer close(result)
		response, err = client.UpdateDegradeControl(request)
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

// UpdateDegradeControlRequest is the request struct for api UpdateDegradeControl
type UpdateDegradeControlRequest struct {
	*requests.RoaRequest
	Duration    requests.Integer `position:"Query" name:"Duration"`
	RuleType    string           `position:"Query" name:"RuleType"`
	AppId       string           `position:"Query" name:"AppId"`
	UrlVar      string           `position:"Query" name:"UrlVar"`
	RtThreshold requests.Integer `position:"Query" name:"RtThreshold"`
	ServiceName string           `position:"Query" name:"ServiceName"`
	RuleId      string           `position:"Query" name:"RuleId"`
	MethodName  string           `position:"Query" name:"MethodName"`
}

// UpdateDegradeControlResponse is the response struct for api UpdateDegradeControl
type UpdateDegradeControlResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateDegradeControlRequest creates a request to invoke UpdateDegradeControl API
func CreateUpdateDegradeControlRequest() (request *UpdateDegradeControlRequest) {
	request = &UpdateDegradeControlRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "UpdateDegradeControl", "/pop/v5/degradeControl", "Edas", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateUpdateDegradeControlResponse creates a response to parse from UpdateDegradeControl response
func CreateUpdateDegradeControlResponse() (response *UpdateDegradeControlResponse) {
	response = &UpdateDegradeControlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
