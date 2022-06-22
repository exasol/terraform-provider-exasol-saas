package provider

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/exasol/exasol-driver-go"
	openapi "github.com/exasol/terraform-provider-exasol-saas/internal/client"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceItemDatabase() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The name of the resource",
				ValidateDiagFunc: validateName,
			},
			"region": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				Description:      "The region of the resource",
				ValidateDiagFunc: validateRegions,
			},
			"size": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The size of the database",
				ValidateDiagFunc: validateSize,
			},
			"main_cluster_name": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "The name of the main cluster",
				ValidateDiagFunc: validateName,
				Default:          "Main",
			},
			"main_cluster_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Id of the main cluster",
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
			"initial_sql": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "Sql which will be run after database creation",
			},
		},
		CreateContext: resourceCreateDatabase,
		ReadContext:   resourceReadDatabase,
		UpdateContext: resourceUpdateDatabase,
		DeleteContext: resourceDeleteDatabase,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(45 * time.Minute),
		},
	}
}

func resourceDeleteDatabase(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(*providerConfig)
	apiClient := config.Client
	account := config.Account
	dbID := d.Id()
	_, err := apiClient.DatabasesApi.DeleteDatabase(ctx, account, dbID).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceReadDatabase(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(*providerConfig)
	apiClient := config.Client
	account := config.Account
	dbID := d.Id()
	tflog.Debug(ctx, "Reading database "+dbID, map[string]interface{}{"account": account})

	database, resp, err := apiClient.DatabasesApi.GetDatabase(ctx, account, dbID).Execute()
	if resp != nil && resp.StatusCode == 404 {
		d.SetId("")
		return nil
	}

	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("name", database.Name)
	if err != nil {
		return diag.FromErr(err)
	}

	mainCluster, diagnostics := getMainCluster(ctx, d, apiClient, account, database.Id)
	if diagnostics != nil {
		return diagnostics
	}

	connection, resp, err := apiClient.ClustersApi.GetClusterConnection(ctx, account, database.Id, mainCluster.Id).Execute()

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
	err = d.Set("main_cluster_id", mainCluster.Id)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("main_cluster_name", mainCluster.Name)
	if err != nil {
		return diag.FromErr(err)
	}
	if mainCluster.AutoStop != nil && mainCluster.AutoStop.Enabled {
		err = d.Set("autostop", int(mainCluster.AutoStop.IdleTime))
		if err != nil {
			return diag.FromErr(err)
		}
	} else {
		err = d.Set("autostop", nil)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return nil
}

func resourceUpdateDatabase(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(*providerConfig)
	apiClient := config.Client
	account := config.Account
	region := d.Get("region").(string)
	dbID := d.Id()
	mainClusterId := d.Get("main_cluster_id").(string)

	if d.HasChange("name") {
		oldName, newName := d.GetChange("name")
		tflog.Debug(ctx, fmt.Sprintf("Renaming database %s -> %s", oldName.(string), newName.(string)), map[string]interface{}{"region": region, "account": account})

		resp, err := apiClient.DatabasesApi.UpdateDatabase(ctx, account, dbID).UpdateDatabase(openapi.UpdateDatabase{
			Name: newName.(string),
		}).Execute()

		if err != nil {
			return handleApiError(ctx, err, resp)
		}
	}

	if d.HasChange("size") {
		mainCluster, diagnostics := getMainCluster(ctx, d, apiClient, account, dbID)
		if diagnostics != nil {
			return diagnostics
		}

		if mainCluster.Status != openapi.RUNNING {
			diagnostics = startDatabase(ctx, d, apiClient, account, dbID)
			if diagnostics != nil {
				return diagnostics
			}
		}
	}

	return updateCluster(ctx, d, config, "main_cluster_name", dbID, mainClusterId)
}

func resourceCreateDatabase(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(*providerConfig)
	apiClient := config.Client
	account := config.Account
	name := d.Get("name").(string)
	mainClusterName := d.Get("main_cluster_name").(string)
	region := d.Get("region").(string)
	size := d.Get("size").(string)

	tflog.Debug(ctx, "Creating database"+name, map[string]interface{}{"region": region, "size": size, "account": account})

	initialCluster := openapi.CreateCluster{
		Name: mainClusterName,
		Size: size,
	}

	if d.Get("autostop") != nil {
		initialCluster.AutoStop = &openapi.AutoStop{
			Enabled:  true,
			IdleTime: int32(d.Get("autostop").(int)),
		}
	}

	database, resp, err := apiClient.DatabasesApi.CreateDatabase(ctx, account).CreateDatabase(openapi.CreateDatabase{
		Name:           name,
		InitialCluster: initialCluster,
		Provider:       "aws",
		Region:         region,
	}).Execute()

	diagnostics := handleApiError(ctx, err, resp)
	if diagnostics != nil {
		return diagnostics
	}
	tflog.Debug(ctx, "Database created "+name, map[string]interface{}{"region": region, "size": size, "account": account, "id": database.Id})

	d.SetId(database.Id)
	err = resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate)-time.Minute, func() *resource.RetryError {
		currentDBStatus, _, err := apiClient.DatabasesApi.GetDatabase(ctx, account, database.Id).Execute()
		tflog.Debug(ctx, "Check database state for "+name, map[string]interface{}{"region": region, "size": size, "account": account, "id": database.Id})

		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("error describing database: %s", err))
		}

		if currentDBStatus.Status != openapi.RUNNING {
			return resource.RetryableError(fmt.Errorf("expected instance to be running but was in state %s", currentDBStatus.Status))
		}

		return nil
	})
	if err != nil {
		return diag.FromErr(err)
	}

	clusters, resp, err := apiClient.ClustersApi.ListClusters(ctx, account, database.Id).Execute()
	diagnostics = handleApiError(ctx, err, resp)
	if diagnostics != nil {
		return diagnostics
	}

	connection, resp, err := apiClient.ClustersApi.GetClusterConnection(ctx, account, database.Id, clusters[0].Id).Execute()
	diagnostics = handleApiError(ctx, err, resp)
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
	err = d.Set("main_cluster_id", clusters[0].Id)
	if err != nil {
		return diag.FromErr(err)
	}

	initialSqlRaw := d.Get("initial_sql")
	if initialSqlRaw != nil {
		// TODO add retry, can happen that the DB is running but the user ist not there yet.
		databaseClient, err := sql.Open("exasol", exasol.NewConfigWithRefreshToken(config.Token).
			Port(int(connection.Port)).
			Host(connection.Dns).
			String())

		if err != nil {
			return diag.FromErr(err)
		}

		var diags diag.Diagnostics
		for _, sql := range initialSqlRaw.([]interface{}) {
			if _, err := databaseClient.ExecContext(ctx, sql.(string)); err != nil {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Warning,
					Summary:  "Could not execute initial sql, please make sure the database is reachable",
					Detail:   err.Error(),
				})
			}
		}
		if len(diags) > 0 {
			return diags
		}

	}

	return nil
}
