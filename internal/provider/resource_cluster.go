package provider

import (
	"context"
	"fmt"
	openapi "github.com/exasol/terraform-provider-exasol-saas/internal/client"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceItemCluster() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCreateCluster,
		ReadContext:   resourceReadCluster,
		UpdateContext: resourceUpdateCluster,
		DeleteContext: resourceDeleteCluster,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(45 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"database": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the database to which the cluster should be added",
			},
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The name of the cluster",
				ValidateDiagFunc: validateName,
			},
			"size": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The size of the cluster",
				ValidateDiagFunc: validateSize,
			},
			"autostop": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Auto stop cluster after x minutes",
			},
			"dns": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "DNS of the cluster",
			},
			"port": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "DNS of the cluster",
			},
		},
	}
}

func resourceDeleteCluster(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(*providerConfig)
	apiClient := config.Client
	account := config.Account
	dbID := d.Get("database").(string)
	clusterID := d.Id()
	_, err := apiClient.ClustersApi.DeleteCluster(ctx, account, dbID, clusterID).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceReadCluster(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(*providerConfig)
	apiClient := config.Client
	account := config.Account
	dbID := d.Get("database").(string)
	clusterID := d.Id()
	tflog.Debug(ctx, "Reading cluster "+dbID, map[string]interface{}{"account": account, "database": dbID})

	cluster, resp, err := apiClient.ClustersApi.GetCluster(ctx, account, dbID, clusterID).Execute()
	if resp.StatusCode == 404 {
		d.SetId("")
		return nil
	}
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("name", cluster.Name)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("size", cluster.Size)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("size", cluster.Size)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceUpdateCluster(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(*providerConfig)
	apiClient := config.Client
	account := config.Account
	dbID := d.Get("database").(string)
	clusterID := d.Id()
	//size := d.Get("size").(string)

	diagnostics := updateCluster(ctx, d, "name", account, dbID, clusterID, apiClient)
	if diagnostics != nil {
		return diagnostics
	}

	return nil
}

func resourceCreateCluster(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(*providerConfig)
	apiClient := config.Client
	account := config.Account
	dbID := d.Get("database").(string)
	name := d.Get("name").(string)
	size := d.Get("size").(string)

	tflog.Debug(ctx, "Creating cluster "+name, map[string]interface{}{"database": dbID, "account": account})

	newCluster := openapi.CreateCluster{
		Name: name,
		Size: size,
	}

	if d.Get("autostop") != nil {
		newCluster.AutoStop = &openapi.AutoStop{
			Enabled:  true,
			IdleTime: int32(d.Get("autostop").(int)),
		}
	}

	cluster, resp, err := apiClient.ClustersApi.CreateCluster(ctx, account, dbID).CreateCluster(newCluster).Execute()

	diagnostics := handleApiError(ctx, err, resp)
	if diagnostics != nil {
		return diagnostics
	}
	tflog.Debug(ctx, "Cluster created "+name, map[string]interface{}{"database": dbID, "size": size, "account": account, "id": cluster.Id})

	d.SetId(cluster.Id)

	err = resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate)-time.Minute, func() *resource.RetryError {
		currentStatus, _, err := apiClient.ClustersApi.GetCluster(ctx, account, dbID, cluster.Id).Execute()
		tflog.Debug(ctx, "Check cluster state for "+name, map[string]interface{}{"database": dbID, "size": size, "account": account, "id": cluster.Id})

		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("Error describing database: %s", err))
		}

		if currentStatus.Status != openapi.RUNNING {
			return resource.RetryableError(fmt.Errorf("Expected instance to be running but was in state %s", currentStatus.Status))
		}

		return nil
	})
	if err != nil {
		return diag.FromErr(err)
	}

	return setClusterConnectionInfo(ctx, d, apiClient, account, dbID, cluster)
}

func setClusterConnectionInfo(ctx context.Context, d *schema.ResourceData, apiClient *openapi.APIClient, account string, dbID string, cluster *openapi.Cluster) diag.Diagnostics {
	connection, resp, err := apiClient.ClustersApi.GetClusterConnection(ctx, account, dbID, cluster.Id).Execute()
	diagnostics := handleApiError(ctx, err, resp)
	if diagnostics != nil {
		return diagnostics
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
