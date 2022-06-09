package number

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

type ProviderClient struct {
}

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		ConfigureFunc: ProviderClient{},

		Schema: map[string]*schema.Schema{},

		DataSourcesMap: map[string]*schema.Resource{
			"number": initNumberSchema(),

		},

		ResourcesMap: map[string]*schema.Resource{},
	}
}
