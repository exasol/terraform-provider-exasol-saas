package provider

import (
	"context"
	"fmt"

	openapi "github.com/exasol/terraform-provider-exasol-saas/internal/client"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceItemNetworkAllowList() *schema.Resource {
	return &schema.Resource{
		Description: "Whitelist cidr block",

		CreateContext: resourceCreateNetwork,
		ReadContext:   resourceReadNetwork,
		UpdateContext: resourceUpdateNetwork,
		DeleteContext: resourceDeleteNetwork,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the resource",
			},
			"cidr_block": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The cidr block to allow access",
				ValidateDiagFunc: validation.ToDiagFunc(validation.IsCIDR),
			},
		},
	}
}

func resourceDeleteNetwork(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(*providerConfig)
	apiClient := config.Client
	account := config.Account
	id := d.Id()

	_, err := apiClient.SecurityApi.DeleteAllowedIP(ctx, account, id).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceReadNetwork(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(*providerConfig)
	apiClient := config.Client
	account := config.Account
	id := d.Id()
	tflog.Debug(ctx, "Reading security ip "+id, map[string]interface{}{"account": account})

	ip, resp, err := apiClient.SecurityApi.GetAllowedIP(ctx, account, id).Execute()
	if resp.StatusCode == 404 {
		d.SetId("")
		return nil

	}
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("name", ip.Name)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("cidr_block", ip.CidrIp)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceUpdateNetwork(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(*providerConfig)
	apiClient := config.Client
	account := config.Account
	dbID := d.Id()

	if d.HasChange("name") {
		oldName, newName := d.GetChange("name")
		tflog.Debug(ctx, fmt.Sprintf("Change name %s -> %s", oldName.(string), newName.(string)), map[string]interface{}{"account": account})
	}

	if d.HasChange("cidr_block") {
		oldCidr, newCidr := d.GetChange("cidr_block")
		tflog.Debug(ctx, fmt.Sprintf("Change cidr %s -> %s", oldCidr.(string), newCidr.(string)), map[string]interface{}{"account": account})
	}

	if d.HasChange("name") || d.HasChange("cidr_block") {
		_, newName := d.GetChange("name")
		_, newCidr := d.GetChange("cidr_block")
		resp, err := apiClient.SecurityApi.UpdateAllowedIP(ctx, account, dbID).UpdateAllowedIP(openapi.UpdateAllowedIP{
			Name:   newName.(string),
			CidrIp: newCidr.(string),
		}).Execute()

		if err != nil {
			return handleApiError(ctx, err, resp)
		}
	}
	return nil
}

func resourceCreateNetwork(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(*providerConfig)
	apiClient := config.Client
	account := config.Account
	name := d.Get("name").(string)
	cidr := d.Get("cidr_block").(string)

	tflog.Debug(ctx, "Add cidr to allow list "+name, map[string]interface{}{"cidr": cidr, "account": account})

	ip, resp, err := apiClient.SecurityApi.AddAllowedIP(ctx, account).CreateAllowedIP(openapi.CreateAllowedIP{
		Name:   name,
		CidrIp: cidr,
	}).Execute()

	if err != nil {
		return handleApiError(ctx, err, resp)
	}

	tflog.Debug(ctx, "Cidr added "+name, map[string]interface{}{"cidr": cidr, "account": account, "id": ip.Id})

	d.SetId(ip.Id)
	return nil
}
