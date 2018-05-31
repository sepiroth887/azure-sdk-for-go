package hdinsight

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// ConfigurationsClient is the hDInsight Management Client
type ConfigurationsClient struct {
	BaseClient
}

// NewConfigurationsClient creates an instance of the ConfigurationsClient client.
func NewConfigurationsClient(subscriptionID string) ConfigurationsClient {
	return NewConfigurationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewConfigurationsClientWithBaseURI creates an instance of the ConfigurationsClient client.
func NewConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) ConfigurationsClient {
	return ConfigurationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get the configuration object for the specified cluster.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the cluster.
// configurationName - the constant for configuration type of gateway.
func (client ConfigurationsClient) Get(ctx context.Context, resourceGroupName string, clusterName string, configurationName string) (result SetString, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, clusterName, configurationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ConfigurationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hdinsight.ConfigurationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ConfigurationsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ConfigurationsClient) GetPreparer(ctx context.Context, resourceGroupName string, clusterName string, configurationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"configurationName": autorest.Encode("path", configurationName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/configurations/{configurationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ConfigurationsClient) GetResponder(resp *http.Response) (result SetString, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UpdateHTTPSettings configures the HTTP settings on the specified cluster.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the cluster.
// parameters - the name of the resource group.
func (client ConfigurationsClient) UpdateHTTPSettings(ctx context.Context, resourceGroupName string, clusterName string, parameters HTTPConnectivitySettings) (result ConfigurationsUpdateHTTPSettingsFuture, err error) {
	req, err := client.UpdateHTTPSettingsPreparer(ctx, resourceGroupName, clusterName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ConfigurationsClient", "UpdateHTTPSettings", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateHTTPSettingsSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ConfigurationsClient", "UpdateHTTPSettings", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdateHTTPSettingsPreparer prepares the UpdateHTTPSettings request.
func (client ConfigurationsClient) UpdateHTTPSettingsPreparer(ctx context.Context, resourceGroupName string, clusterName string, parameters HTTPConnectivitySettings) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"configurationName": autorest.Encode("path", "gateway"),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/configurations/{configurationName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateHTTPSettingsSender sends the UpdateHTTPSettings request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationsClient) UpdateHTTPSettingsSender(req *http.Request) (future ConfigurationsUpdateHTTPSettingsFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	err = autorest.Respond(resp, azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// UpdateHTTPSettingsResponder handles the response to the UpdateHTTPSettings request. The method always
// closes the http.Response Body.
func (client ConfigurationsClient) UpdateHTTPSettingsResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
