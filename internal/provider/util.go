package provider

import (
	"context"
	"encoding/json"
	"fmt"
	openapi "github.com/exasol/terraform-provider-exasol-saas/internal/client"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"net/http"
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

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func updateCluster(ctx context.Context, d *schema.ResourceData, nameKey string, account string, dbID string, clusterID string, apiClient *openapi.APIClient) diag.Diagnostics {
	updateCluster := openapi.UpdateCluster{}
	hasChange := false

	if d.HasChange(nameKey) {
		oldName, newName := d.GetChange(nameKey)
		tflog.Debug(ctx, fmt.Sprintf("Renaming cluster %s -> %s", oldName.(string), newName.(string)), map[string]interface{}{"account": account, "database": dbID, "cluster": clusterID})
		updateCluster.Name = newName.(string)
		hasChange = true
	}

	if d.HasChange("autostop") {
		oldStop, newStop := d.GetChange("autostop")
		tflog.Debug(ctx, fmt.Sprintf("Upating autostop %s -> %s", oldStop.(string), newStop.(string)), map[string]interface{}{"account": account, "database": dbID, "cluster": clusterID})
		if newStop != nil {
			updateCluster.AutoStop = &openapi.AutoStop{
				Enabled:  true,
				IdleTime: int32(newStop.(int)),
			}
		}
		if newStop == nil {
			updateCluster.AutoStop = &openapi.AutoStop{
				Enabled:  false,
				IdleTime: int32(oldStop.(int)),
			}
		}
		hasChange = true
	}

	if hasChange {
		resp, err := apiClient.ClustersApi.UpdateCluster(ctx, account, dbID, clusterID).UpdateCluster(updateCluster).Execute()
		diagnostics := handleApiError(ctx, err, resp)
		if diagnostics != nil {
			return diagnostics
		}
	}
	return nil
}