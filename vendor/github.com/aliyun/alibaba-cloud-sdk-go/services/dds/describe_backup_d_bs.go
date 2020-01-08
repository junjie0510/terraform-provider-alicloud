package dds

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

// DescribeBackupDBs invokes the dds.DescribeBackupDBs API synchronously
// api document: https://help.aliyun.com/api/dds/describebackupdbs.html
func (client *Client) DescribeBackupDBs(request *DescribeBackupDBsRequest) (response *DescribeBackupDBsResponse, err error) {
	response = CreateDescribeBackupDBsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeBackupDBsWithChan invokes the dds.DescribeBackupDBs API asynchronously
// api document: https://help.aliyun.com/api/dds/describebackupdbs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeBackupDBsWithChan(request *DescribeBackupDBsRequest) (<-chan *DescribeBackupDBsResponse, <-chan error) {
	responseChan := make(chan *DescribeBackupDBsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBackupDBs(request)
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

// DescribeBackupDBsWithCallback invokes the dds.DescribeBackupDBs API asynchronously
// api document: https://help.aliyun.com/api/dds/describebackupdbs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeBackupDBsWithCallback(request *DescribeBackupDBsRequest, callback func(response *DescribeBackupDBsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBackupDBsResponse
		var err error
		defer close(result)
		response, err = client.DescribeBackupDBs(request)
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

// DescribeBackupDBsRequest is the request struct for api DescribeBackupDBs
type DescribeBackupDBsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	RestoreTime          string           `position:"Query" name:"RestoreTime"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	BackupId             string           `position:"Query" name:"BackupId"`
	SourceDBInstance     string           `position:"Query" name:"SourceDBInstance"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeBackupDBsResponse is the response struct for api DescribeBackupDBs
type DescribeBackupDBsResponse struct {
	*responses.BaseResponse
	RequestId  string    `json:"RequestId" xml:"RequestId"`
	PageNumber int       `json:"PageNumber" xml:"PageNumber"`
	PageSize   int       `json:"PageSize" xml:"PageSize"`
	TotalCount int       `json:"TotalCount" xml:"TotalCount"`
	Databases  Databases `json:"Databases" xml:"Databases"`
}

// CreateDescribeBackupDBsRequest creates a request to invoke DescribeBackupDBs API
func CreateDescribeBackupDBsRequest() (request *DescribeBackupDBsRequest) {
	request = &DescribeBackupDBsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "DescribeBackupDBs", "Dds", "openAPI")
	return
}

// CreateDescribeBackupDBsResponse creates a response to parse from DescribeBackupDBs response
func CreateDescribeBackupDBsResponse() (response *DescribeBackupDBsResponse) {
	response = &DescribeBackupDBsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
