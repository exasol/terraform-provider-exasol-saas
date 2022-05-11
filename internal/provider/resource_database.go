package provider

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/exasol/exasol-driver-go"
	openapi "github.com/exasol/terraform-provider-exasol-saas/internal/client"
	"time"

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
	if resp.StatusCode == 404 {
		d.SetId("")
		return nil
	}
	if err != nil {

		return diag.FromErr(err)
	}

	err = d.Set("database", database)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceUpdateDatabase(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(*providerConfig)
	apiClient := config.Client
	account := config.Account
	region := d.Get("region").(string)
	//size := d.Get("size").(string)
	dbID := d.Id()

	if d.HasChange("name") {
		oldName, newName := d.GetChange("name")
		tflog.Debug(ctx, fmt.Sprintf("Renaming database %s -> %s", oldName.(string), newName.(string)), map[string]interface{}{"region": region, "account": account})
		resp, err := apiClient.DatabasesApi.UpdateDatabase(ctx, account, dbID).UpdateDatabase(openapi.UpdateDatabase{
			Name: newName.(string),
		}).Execute()
		diagnostics := handleApiError(ctx, err, resp)
		if diagnostics != nil {
			return diagnostics
		}
	}

	/*	if d.HasChange("size") {
		oldSize, newSize := d.GetChange("size")
		tflog.Debug(ctx, fmt.Sprintf("Resize database %s -> %s", oldSize.(string), newSize.(string)), map[string]interface{}{"region": region, "account": account})
		resp, err := apiClient.DatabasesApi.(ctx, account, dbID).UpdateDatabase(openapi.UpdateDatabase{
			Name: newName.(string),
		}).Execute()
		diagnostics := handleApiError(ctx, err, resp)
		if diagnostics != nil {
			return diagnostics
		}
	}*/

	return nil
}

func resourceCreateDatabase(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(*providerConfig)
	apiClient := config.Client
	account := config.Account
	name := d.Get("name").(string)
	region := d.Get("region").(string)
	size := d.Get("size").(string)

	tflog.Debug(ctx, "Creating database"+name, map[string]interface{}{"region": region, "size": size, "account": account})

	database, resp, err := apiClient.DatabasesApi.CreateDatabase(ctx, account).CreateDatabase(openapi.CreateDatabase{
		Name: name,
		InitialCluster: openapi.CreateCluster{
			Name: "Main",
			Size: size,
		},
		Provider: "aws",
		Region:   region,
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
			return resource.NonRetryableError(fmt.Errorf("Error describing database: %s", err))
		}

		if currentDBStatus.Status != openapi.RUNNING {
			return resource.RetryableError(fmt.Errorf("Expected instance to be running but was in state %s", currentDBStatus.Status))
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

	initialSqlRaw := d.Get("initial_sql")
	err = d.Set("dns", connection.Dns)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("port", connection.Port)
	if err != nil {
		return diag.FromErr(err)
	}

	if initialSqlRaw != nil {
		databaseClient, err := sql.Open("exasol", exasol.NewConfigWithRefreshToken(config.Token).
			Port(int(connection.Port)).
			Host(connection.Dns).
			String())

		if err != nil {
			return diag.FromErr(err)
		}

		for _, sql := range initialSqlRaw.([]interface{}) {
			if _, err := databaseClient.ExecContext(ctx, sql.(string)); err != nil {
				return diag.FromErr(err)
			}
		}

	}

	return nil
}
