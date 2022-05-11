package provider

import (
	"context"
	openapi "github.com/exasol/terraform-provider-exasol-saas/internal/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"api": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("EXASOL_API", "https://cloud.exasol.com"),
				},
				"token": {
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("EXASOL_TOKEN", ""),
					Sensitive:   true,
				},
				"account": {
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("EXASOL_ACCOUNT", ""),
				},
			},
			ResourcesMap: map[string]*schema.Resource{
				"exasol_database":           resourceItemDatabase(),
				"exasol_cluster":           resourceItemCluster(),
				"exasol_network_allow_list": resourceItemNetworkAllowList(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

type providerConfig struct {
	Account string
	Token   string
	Client  *openapi.APIClient
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		client := openapi.NewAPIClient(&openapi.Configuration{
			Servers: []openapi.ServerConfiguration{
				{
					URL: d.Get("api").(string),
				},
			},
			DefaultHeader: map[string]string{"Authorization": "Bearer " + d.Get("token").(string)},
		})

		return &providerConfig{Account: d.Get("account").(string), Client: client, Token: d.Get("token").(string)}, nil
	}
}
