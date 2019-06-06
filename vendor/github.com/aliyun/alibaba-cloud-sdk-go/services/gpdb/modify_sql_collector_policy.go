package gpdb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifySQLCollectorPolicy invokes the gpdb.ModifySQLCollectorPolicy API synchronously
// api document: https://help.aliyun.com/api/gpdb/modifysqlcollectorpolicy.html
func (client *Client) ModifySQLCollectorPolicy(request *ModifySQLCollectorPolicyRequest) (response *ModifySQLCollectorPolicyResponse, err error) {
	response = CreateModifySQLCollectorPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySQLCollectorPolicyWithChan invokes the gpdb.ModifySQLCollectorPolicy API asynchronously
// api document: https://help.aliyun.com/api/gpdb/modifysqlcollectorpolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySQLCollectorPolicyWithChan(request *ModifySQLCollectorPolicyRequest) (<-chan *ModifySQLCollectorPolicyResponse, <-chan error) {
	responseChan := make(chan *ModifySQLCollectorPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySQLCollectorPolicy(request)
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

// ModifySQLCollectorPolicyWithCallback invokes the gpdb.ModifySQLCollectorPolicy API asynchronously
// api document: https://help.aliyun.com/api/gpdb/modifysqlcollectorpolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySQLCollectorPolicyWithCallback(request *ModifySQLCollectorPolicyRequest, callback func(response *ModifySQLCollectorPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySQLCollectorPolicyResponse
		var err error
		defer close(result)
		response, err = client.ModifySQLCollectorPolicy(request)
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

// ModifySQLCollectorPolicyRequest is the request struct for api ModifySQLCollectorPolicy
type ModifySQLCollectorPolicyRequest struct {
	*requests.RpcRequest
	SQLCollectorStatus string `position:"Query" name:"SQLCollectorStatus"`
	DBInstanceId       string `position:"Query" name:"DBInstanceId"`
}

// ModifySQLCollectorPolicyResponse is the response struct for api ModifySQLCollectorPolicy
type ModifySQLCollectorPolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifySQLCollectorPolicyRequest creates a request to invoke ModifySQLCollectorPolicy API
func CreateModifySQLCollectorPolicyRequest() (request *ModifySQLCollectorPolicyRequest) {
	request = &ModifySQLCollectorPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "ModifySQLCollectorPolicy", "gpdb", "openAPI")
	return
}

// CreateModifySQLCollectorPolicyResponse creates a response to parse from ModifySQLCollectorPolicy response
func CreateModifySQLCollectorPolicyResponse() (response *ModifySQLCollectorPolicyResponse) {
	response = &ModifySQLCollectorPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}