package provider

import (
	"context"
	"encoding/json"
	"fmt"
	openapi "github.com/exasol/terraform-provider-exasol-saas/internal/client"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"net/http"
	"strings"
	"time"
)

func handleApiError(ctx context.Context, err error, resp *http.Response) diag.Diagnostics {
	if err != nil {
		var apiError openapi.APIError
		if resp != nil {
			err := json.NewDecoder(resp.Body).Decode(&apiError)
			if err == nil {
				tflog.Error(ctx, apiError.Message)
				return diag.Errorf(apiError.Message)
			}

		} else {
			tflog.Error(ctx, fmt.Sprintf("Api error: %+v", err))
		}

		return diag.FromErr(err)
	}
	return nil
}

func contains(values []string, expected string) bool {
	for _, value := range values {
		if value == expected {
			return true
		}
	}
	return false
}

func containsStatus(statuses []openapi.Status, expected openapi.Status) bool {
	for _, status := range statuses {
		if status == expected {
			return true
		}
	}
	return false
}

func statusesToString(statuses []openapi.Status) string {
	var values []string
	for _, status := range statuses {
		values = append(values, string(status))
	}
	return strings.Join(values, ", ")
}

func updateCluster(ctx context.Context, d *schema.ResourceData, config *providerConfig, nameKey string, dbID string, clusterID string) diag.Diagnostics {
	apiClient := config.Client
	account := config.Account
	updateClusterBody := openapi.UpdateCluster{}
	hasChange := false

	if d.HasChange(nameKey) {
		oldName, newName := d.GetChange(nameKey)
		tflog.Debug(ctx, fmt.Sprintf("Renaming cluster %s -> %s", oldName.(string), newName.(string)), map[string]interface{}{"account": account, "database": dbID, "cluster": clusterID})
		updateClusterBody.Name = newName.(string)
		hasChange = true
	}

	if d.HasChange("autostop") {
		oldStop, newStop := d.GetChange("autostop")
		tflog.Debug(ctx, fmt.Sprintf("Upating autostop %s -> %s", oldStop.(string), newStop.(string)), map[string]interface{}{"account": account, "database": dbID, "cluster": clusterID})
		if newStop != nil {
			updateClusterBody.AutoStop = &openapi.AutoStop{
				Enabled:  true,
				IdleTime: int32(newStop.(int)),
			}
		}
		if newStop == nil {
			updateClusterBody.AutoStop = &openapi.AutoStop{
				Enabled:  false,
				IdleTime: int32(oldStop.(int)),
			}
		}
		hasChange = true
	}

	if hasChange {
		resp, err := apiClient.ClustersApi.UpdateCluster(ctx, account, dbID, clusterID).UpdateCluster(updateClusterBody).Execute()
		if err != nil {
			return handleApiError(ctx, err, resp)
		}
	}

	//Scale

	if d.HasChange("size") {
		oldSize, newSize := d.GetChange("size")

		tflog.Debug(ctx, fmt.Sprintf("Upating size %s -> %s", oldSize.(string), newSize.(string)), map[string]interface{}{"account": account, "database": dbID, "cluster": clusterID})

		err := waitForClusterStatus(ctx, d, apiClient, account, dbID, clusterID, []openapi.Status{openapi.RUNNING, openapi.STOPPED})
		if err != nil {
			return diag.FromErr(err)
		}

		resp, err := apiClient.ClustersApi.ScaleCluster(ctx, account, dbID, clusterID).ScaleCluster(openapi.ScaleCluster{Size: newSize.(string)}).Execute()
		if err != nil {
			return handleApiError(ctx, err, resp)
		}

		err = waitForClusterStatus(ctx, d, apiClient, account, dbID, clusterID, []openapi.Status{openapi.RUNNING, openapi.STOPPED})
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func waitForClusterStatus(ctx context.Context, d *schema.ResourceData, apiClient *openapi.APIClient, account string, dbID string, clusterID string, status []openapi.Status) error {
	return resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate)-time.Minute, func() *resource.RetryError {
		currentStatus, _, err := apiClient.ClustersApi.GetCluster(ctx, account, dbID, clusterID).Execute()
		tflog.Debug(ctx, "Check cluster state for "+clusterID, map[string]interface{}{"database": dbID, "account": account, "id": clusterID})

		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("error describing cluster: %s", err))
		}

		if !containsStatus(status, currentStatus.Status) {
			return resource.RetryableError(fmt.Errorf("expected instance to be in state %s but was in state %s", statusesToString(status), currentStatus.Status))
		}

		return nil
	})
}

func waitForDatabaseStatus(ctx context.Context, d *schema.ResourceData, apiClient *openapi.APIClient, account string, dbID string, status []openapi.Status) error {
	return resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate)-time.Minute, func() *resource.RetryError {
		currentStatus, _, err := apiClient.DatabasesApi.GetDatabase(ctx, account, dbID).Execute()
		tflog.Debug(ctx, "Check database state for "+dbID, map[string]interface{}{"database": dbID, "account": account})

		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("error describing cluster: %s", err))
		}

		if !containsStatus(status, currentStatus.Status) {
			return resource.RetryableError(fmt.Errorf("expected database to be in state %s but was in state %s", statusesToString(status), currentStatus.Status))
		}

		if currentStatus.Clusters.Running != currentStatus.Clusters.Total {
			return resource.RetryableError(fmt.Errorf("expected database clusters to be running %d/%d", currentStatus.Clusters.Running, currentStatus.Clusters.Total))
		}

		return nil
	})
}

func getMainCluster(ctx context.Context, d *schema.ResourceData, apiClient *openapi.APIClient, account string, databaseId string) (*openapi.Cluster, diag.Diagnostics) {
	tflog.Debug(ctx, fmt.Sprintf("######## %v", d.Get("main_cluster_id")))
	if d.Get("main_cluster_id") != nil {
		mainCluster, resp, err := apiClient.ClustersApi.GetCluster(ctx, account, databaseId, d.Get("main_cluster_id").(string)).Execute()
		diagnostics := handleApiError(ctx, err, resp)
		if diagnostics != nil {
			return nil, diagnostics
		}
		return mainCluster, nil
	}

	clusters, resp, err := apiClient.ClustersApi.ListClusters(ctx, account, databaseId).Execute()
	diagnostics := handleApiError(ctx, err, resp)
	if diagnostics != nil {
		return nil, diagnostics
	}

	for _, cluster := range clusters {
		if cluster.MainCluster {
			return &cluster, nil
		}
	}
	return nil, diag.Errorf("Main cluster not found")
}

func startDatabase(ctx context.Context, d *schema.ResourceData, apiClient *openapi.APIClient, account string, databaseId string) diag.Diagnostics {

	_, err := apiClient.DatabasesApi.StartDatabase(ctx, account, databaseId).Execute()
	diagnostics := handleApiError(ctx, err, nil)
	if diagnostics != nil {
		return diagnostics
	}

	err = waitForDatabaseStatus(ctx, d, apiClient, account, databaseId, []openapi.Status{openapi.RUNNING})
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func setClusterConnectionInfo(ctx context.Context, d *schema.ResourceData, apiClient *openapi.APIClient, account string, dbID string, cluster *openapi.Cluster) diag.Diagnostics {
	connection, resp, err := apiClient.ClustersApi.GetClusterConnection(ctx, account, dbID, cluster.Id).Execute()
	if err != nil {
		return handleApiError(ctx, err, resp)
	}

	err = d.Set("dns", connection.Dns)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("port", connection.Port)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
