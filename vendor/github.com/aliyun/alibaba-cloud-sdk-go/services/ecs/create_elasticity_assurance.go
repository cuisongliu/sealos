package ecs

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

// CreateElasticityAssurance invokes the ecs.CreateElasticityAssurance API synchronously
func (client *Client) CreateElasticityAssurance(request *CreateElasticityAssuranceRequest) (response *CreateElasticityAssuranceResponse, err error) {
	response = CreateCreateElasticityAssuranceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateElasticityAssuranceWithChan invokes the ecs.CreateElasticityAssurance API asynchronously
func (client *Client) CreateElasticityAssuranceWithChan(request *CreateElasticityAssuranceRequest) (<-chan *CreateElasticityAssuranceResponse, <-chan error) {
	responseChan := make(chan *CreateElasticityAssuranceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateElasticityAssurance(request)
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

// CreateElasticityAssuranceWithCallback invokes the ecs.CreateElasticityAssurance API asynchronously
func (client *Client) CreateElasticityAssuranceWithCallback(request *CreateElasticityAssuranceRequest, callback func(response *CreateElasticityAssuranceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateElasticityAssuranceResponse
		var err error
		defer close(result)
		response, err = client.CreateElasticityAssurance(request)
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

// CreateElasticityAssuranceRequest is the request struct for api CreateElasticityAssurance
type CreateElasticityAssuranceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId                 requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken                     string           `position:"Query" name:"ClientToken"`
	Description                     string           `position:"Query" name:"Description"`
	StartTime                       string           `position:"Query" name:"StartTime"`
	Platform                        string           `position:"Query" name:"Platform"`
	PrivatePoolOptionsMatchCriteria string           `position:"Query" name:"PrivatePoolOptions.MatchCriteria"`
	InstanceType                    *[]string        `position:"Query" name:"InstanceType"  type:"Repeated"`
	InstanceChargeType              string           `position:"Query" name:"InstanceChargeType"`
	Period                          requests.Integer `position:"Query" name:"Period"`
	ResourceOwnerAccount            string           `position:"Query" name:"ResourceOwnerAccount"`
	PrivatePoolOptionsName          string           `position:"Query" name:"PrivatePoolOptions.Name"`
	OwnerAccount                    string           `position:"Query" name:"OwnerAccount"`
	AssuranceTimes                  string           `position:"Query" name:"AssuranceTimes"`
	OwnerId                         requests.Integer `position:"Query" name:"OwnerId"`
	ResourceType                    string           `position:"Query" name:"ResourceType"`
	InstanceCpuCoreCount            requests.Integer `position:"Query" name:"InstanceCpuCoreCount"`
	PeriodUnit                      string           `position:"Query" name:"PeriodUnit"`
	ZoneId                          *[]string        `position:"Query" name:"ZoneId"  type:"Repeated"`
	ChargeType                      string           `position:"Query" name:"ChargeType"`
	PackageType                     string           `position:"Query" name:"PackageType"`
	InstanceAmount                  requests.Integer `position:"Query" name:"InstanceAmount"`
}

// CreateElasticityAssuranceResponse is the response struct for api CreateElasticityAssurance
type CreateElasticityAssuranceResponse struct {
	*responses.BaseResponse
	RequestId            string `json:"RequestId" xml:"RequestId"`
	PrivatePoolOptionsId string `json:"PrivatePoolOptionsId" xml:"PrivatePoolOptionsId"`
	OrderId              string `json:"OrderId" xml:"OrderId"`
}

// CreateCreateElasticityAssuranceRequest creates a request to invoke CreateElasticityAssurance API
func CreateCreateElasticityAssuranceRequest() (request *CreateElasticityAssuranceRequest) {
	request = &CreateElasticityAssuranceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "CreateElasticityAssurance", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateElasticityAssuranceResponse creates a response to parse from CreateElasticityAssurance response
func CreateCreateElasticityAssuranceResponse() (response *CreateElasticityAssuranceResponse) {
	response = &CreateElasticityAssuranceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
