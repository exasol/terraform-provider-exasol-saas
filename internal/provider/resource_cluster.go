package provider

import (
	"context"
	openapi "github.com/exasol/terraform-provider-exasol-saas/internal/client"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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

	resp, err := apiClient.ClustersApi.DeleteCluster(ctx, config.Account, d.Get("database").(string), d.Id()).Execute()
	if err != nil {
		return handleApiError(ctx, err, resp)
	}

	return nil
}

func resourceReadCluster(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(*providerConfig)
	apiClient := config.Client

	account := config.Account
	dbID := d.Get("database").(string)

	tflog.Debug(ctx, "Reading cluster "+dbID, map[string]interface{}{"account": account, "database": dbID, "cluster": d.Id()})

	cluster, resp, err := apiClient.ClustersApi.GetCluster(ctx, account, dbID, d.Id()).Execute()

	// Cluster or Database was removed by hand
	if resp.StatusCode == 404 || resp.StatusCode == 400 {
		d.SetId("")
		return nil
	}

	if err != nil {
		return handleApiError(ctx, err, resp)
	}

	err = d.Set("name", cluster.Name)
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
	return updateCluster(ctx, d, config, "name", d.Get("database").(string), d.Id())
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

	mainCluster, diagnostics := getMainCluster(ctx, d, apiClient, account, dbID)
	if diagnostics != nil {
		return diagnostics
	}

	// Database is not running, start to allow adding worker cluster
	if mainCluster.Status != openapi.RUNNING {
		diagnostics = startDatabase(ctx, d, apiClient, account, dbID)
		if diagnostics != nil {
			return diagnostics
		}
	}

	err := waitForDatabaseStatus(ctx, d, apiClient, account, dbID, []openapi.Status{openapi.RUNNING})
	if err != nil {
		return diag.FromErr(err)
	}

	cluster, resp, err := apiClient.ClustersApi.CreateCluster(ctx, account, dbID).CreateCluster(newCluster).Execute()

	diagnostics = handleApiError(ctx, err, resp)
	if diagnostics != nil {
		return diagnostics
	}
	tflog.Debug(ctx, "Cluster created "+name, map[string]interface{}{"database": dbID, "size": size, "account": account, "id": cluster.Id})

	d.SetId(cluster.Id)

	err = waitForClusterStatus(ctx, d, apiClient, account, dbID, cluster.Id, []openapi.Status{openapi.RUNNING})
	if err != nil {
		return diag.FromErr(err)
	}

	return setClusterConnectionInfo(ctx, d, apiClient, account, dbID, cluster)
}
