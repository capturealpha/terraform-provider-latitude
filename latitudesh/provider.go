package latitude

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api "github.com/latitudesh/latitudesh-go"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"auth_token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("LATITUDE_AUTH_TOKEN", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"latitude_project": resourceProject(),
			"latitude_server":  resourceServer(),
			"latitude_ssh_key": resourceSSHKey(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"latitude_plan":   dataSourcePlan(),
			"latitude_region": dataSourceRegion(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	authToken := d.Get("auth_token").(string)

	var diags diag.Diagnostics

	if authToken != "" {
		c := api.NewClientWithAuth("latitude", authToken, nil)

		return c, diags
	}

	return nil, diags
}
