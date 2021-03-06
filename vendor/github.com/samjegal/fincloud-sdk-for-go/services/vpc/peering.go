package vpc

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// PeeringClient is the VPC Client
type PeeringClient struct {
	BaseClient
}

// NewPeeringClient creates an instance of the PeeringClient client.
func NewPeeringClient() PeeringClient {
	return NewPeeringClientWithBaseURI(DefaultBaseURI)
}

// NewPeeringClientWithBaseURI creates an instance of the PeeringClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewPeeringClientWithBaseURI(baseURI string) PeeringClient {
	return PeeringClient{NewWithBaseURI(baseURI)}
}

// AcceptOrReject VPC Peering 요청을 수락/거절
// Parameters:
// responseFormatType - 반환 데이터 포맷 타입
// vpcPeeringInstanceNo - VPC Peering 인스턴스 번호
// isAccept - 수락 여부
// regionCode - REGION 코드
func (client PeeringClient) AcceptOrReject(ctx context.Context, responseFormatType string, vpcPeeringInstanceNo string, isAccept string, regionCode string) (result PeeringInstanceResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PeeringClient.AcceptOrReject")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.AcceptOrRejectPreparer(ctx, responseFormatType, vpcPeeringInstanceNo, isAccept, regionCode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.PeeringClient", "AcceptOrReject", nil, "Failure preparing request")
		return
	}

	resp, err := client.AcceptOrRejectSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vpc.PeeringClient", "AcceptOrReject", resp, "Failure sending request")
		return
	}

	result, err = client.AcceptOrRejectResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.PeeringClient", "AcceptOrReject", resp, "Failure responding to request")
	}

	return
}

// AcceptOrRejectPreparer prepares the AcceptOrReject request.
func (client PeeringClient) AcceptOrRejectPreparer(ctx context.Context, responseFormatType string, vpcPeeringInstanceNo string, isAccept string, regionCode string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"isAccept":             autorest.Encode("query", isAccept),
		"responseFormatType":   autorest.Encode("query", responseFormatType),
		"vpcPeeringInstanceNo": autorest.Encode("query", vpcPeeringInstanceNo),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc/v2/acceptOrRejectVpcPeering"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AcceptOrRejectSender sends the AcceptOrReject request. The method will close the
// http.Response Body if it receives an error.
func (client PeeringClient) AcceptOrRejectSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AcceptOrRejectResponder handles the response to the AcceptOrReject request. The method always
// closes the http.Response Body.
func (client PeeringClient) AcceptOrRejectResponder(resp *http.Response) (result PeeringInstanceResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Create VPC Peering 인스턴스를 생성
// Parameters:
// responseFormatType - 반환 데이터 포맷 타입
// sourceVpcNo - 요청 VPC 번호
// targetVpcNo - 수락 VPC 번호
// regionCode - REGION 코드
// vpcPeeringName - VPC Peering 이름
// targetVpcName - 수락 VPC 이름
// targetVpcLoginID - 수락 VPC 소유자 ID
// vpcPeeringDescription - VPC Peering 설명
func (client PeeringClient) Create(ctx context.Context, responseFormatType string, sourceVpcNo string, targetVpcNo string, regionCode string, vpcPeeringName string, targetVpcName string, targetVpcLoginID string, vpcPeeringDescription string) (result PeeringInstanceResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PeeringClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, responseFormatType, sourceVpcNo, targetVpcNo, regionCode, vpcPeeringName, targetVpcName, targetVpcLoginID, vpcPeeringDescription)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.PeeringClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vpc.PeeringClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.PeeringClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client PeeringClient) CreatePreparer(ctx context.Context, responseFormatType string, sourceVpcNo string, targetVpcNo string, regionCode string, vpcPeeringName string, targetVpcName string, targetVpcLoginID string, vpcPeeringDescription string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", responseFormatType),
		"sourceVpcNo":        autorest.Encode("query", sourceVpcNo),
		"targetVpcNo":        autorest.Encode("query", targetVpcNo),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}
	if len(vpcPeeringName) > 0 {
		queryParameters["vpcPeeringName"] = autorest.Encode("query", vpcPeeringName)
	}
	if len(targetVpcName) > 0 {
		queryParameters["targetVpcName"] = autorest.Encode("query", targetVpcName)
	}
	if len(targetVpcLoginID) > 0 {
		queryParameters["targetVpcLoginId"] = autorest.Encode("query", targetVpcLoginID)
	}
	if len(vpcPeeringDescription) > 0 {
		queryParameters["vpcPeeringDescription"] = autorest.Encode("query", vpcPeeringDescription)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc/v2/createVpcPeeringInstance"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client PeeringClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client PeeringClient) CreateResponder(resp *http.Response) (result PeeringInstanceResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete VPC Peering 인스턴스를 삭제
// Parameters:
// responseFormatType - 반환 데이터 포맷 타입
// vpcPeeringInstanceNo - VPC Peering 인스턴스 번호
// regionCode - REGION 코드
func (client PeeringClient) Delete(ctx context.Context, responseFormatType string, vpcPeeringInstanceNo string, regionCode string) (result PeeringInstanceResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PeeringClient.Delete")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, responseFormatType, vpcPeeringInstanceNo, regionCode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.PeeringClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vpc.PeeringClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.PeeringClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client PeeringClient) DeletePreparer(ctx context.Context, responseFormatType string, vpcPeeringInstanceNo string, regionCode string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType":   autorest.Encode("query", responseFormatType),
		"vpcPeeringInstanceNo": autorest.Encode("query", vpcPeeringInstanceNo),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc/v2/deleteVpcPeeringInstance"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client PeeringClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client PeeringClient) DeleteResponder(resp *http.Response) (result PeeringInstanceResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetDetail VPC Peering 인스턴스 상세 정보를 조회
// Parameters:
// responseFormatType - 반환 데이터 포맷 타입
// vpcPeeringInstanceNo - VPC Peering 인스턴스 번호
// regionCode - REGION 코드
func (client PeeringClient) GetDetail(ctx context.Context, responseFormatType string, vpcPeeringInstanceNo string, regionCode string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PeeringClient.GetDetail")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetDetailPreparer(ctx, responseFormatType, vpcPeeringInstanceNo, regionCode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.PeeringClient", "GetDetail", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDetailSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "vpc.PeeringClient", "GetDetail", resp, "Failure sending request")
		return
	}

	result, err = client.GetDetailResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.PeeringClient", "GetDetail", resp, "Failure responding to request")
	}

	return
}

// GetDetailPreparer prepares the GetDetail request.
func (client PeeringClient) GetDetailPreparer(ctx context.Context, responseFormatType string, vpcPeeringInstanceNo string, regionCode string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType":   autorest.Encode("query", responseFormatType),
		"vpcPeeringInstanceNo": autorest.Encode("query", vpcPeeringInstanceNo),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc/v2/getVpcPeeringInstanceDetail"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetDetailSender sends the GetDetail request. The method will close the
// http.Response Body if it receives an error.
func (client PeeringClient) GetDetailSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetDetailResponder handles the response to the GetDetail request. The method always
// closes the http.Response Body.
func (client PeeringClient) GetDetailResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetList VPC Peering 인스턴스 리스트를 조회
// Parameters:
// responseFormatType - 반환 데이터 포맷 타입
// regionCode - REGION 코드
// vpcPeeringInstanceNoListN - VPC Peering 인스턴스 번호 리스트
// sourceVpcName - 요청 VPC 이름
// vpcPeeringName - VPC Peering 이름
// targetVpcName - 수락 VPC 이름
// vpcPeeringInstanceStatusCode - VPC Peering 인스턴스 상태 코드
// pageNo - 페이지 번호
// pageSize - 페이지 사이즈
// sortedBy - 정렬 대상
// sortingOrder - 정렬 순서
func (client PeeringClient) GetList(ctx context.Context, responseFormatType string, regionCode string, vpcPeeringInstanceNoListN string, sourceVpcName string, vpcPeeringName string, targetVpcName string, vpcPeeringInstanceStatusCode PeeringInstanceStatusCode, pageNo string, pageSize string, sortedBy SortedBy, sortingOrder SortingOrder) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PeeringClient.GetList")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetListPreparer(ctx, responseFormatType, regionCode, vpcPeeringInstanceNoListN, sourceVpcName, vpcPeeringName, targetVpcName, vpcPeeringInstanceStatusCode, pageNo, pageSize, sortedBy, sortingOrder)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.PeeringClient", "GetList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetListSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "vpc.PeeringClient", "GetList", resp, "Failure sending request")
		return
	}

	result, err = client.GetListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.PeeringClient", "GetList", resp, "Failure responding to request")
	}

	return
}

// GetListPreparer prepares the GetList request.
func (client PeeringClient) GetListPreparer(ctx context.Context, responseFormatType string, regionCode string, vpcPeeringInstanceNoListN string, sourceVpcName string, vpcPeeringName string, targetVpcName string, vpcPeeringInstanceStatusCode PeeringInstanceStatusCode, pageNo string, pageSize string, sortedBy SortedBy, sortingOrder SortingOrder) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", responseFormatType),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}
	if len(vpcPeeringInstanceNoListN) > 0 {
		queryParameters["vpcPeeringInstanceNoList.N"] = autorest.Encode("query", vpcPeeringInstanceNoListN)
	}
	if len(sourceVpcName) > 0 {
		queryParameters["sourceVpcName"] = autorest.Encode("query", sourceVpcName)
	}
	if len(vpcPeeringName) > 0 {
		queryParameters["vpcPeeringName"] = autorest.Encode("query", vpcPeeringName)
	}
	if len(targetVpcName) > 0 {
		queryParameters["targetVpcName"] = autorest.Encode("query", targetVpcName)
	}
	if len(string(vpcPeeringInstanceStatusCode)) > 0 {
		queryParameters["vpcPeeringInstanceStatusCode"] = autorest.Encode("query", vpcPeeringInstanceStatusCode)
	}
	if len(pageNo) > 0 {
		queryParameters["pageNo"] = autorest.Encode("query", pageNo)
	}
	if len(pageSize) > 0 {
		queryParameters["pageSize"] = autorest.Encode("query", pageSize)
	}
	if len(string(sortedBy)) > 0 {
		queryParameters["sortedBy"] = autorest.Encode("query", sortedBy)
	}
	if len(string(sortingOrder)) > 0 {
		queryParameters["sortingOrder"] = autorest.Encode("query", sortingOrder)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc/v2/getVpcPeeringInstanceList"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetListSender sends the GetList request. The method will close the
// http.Response Body if it receives an error.
func (client PeeringClient) GetListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetListResponder handles the response to the GetList request. The method always
// closes the http.Response Body.
func (client PeeringClient) GetListResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
