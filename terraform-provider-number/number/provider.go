package number

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

type Dummy struct {
}

func configureFunc(d *schema.ResourceData) (i interface{}, err error) {
	i = Dummy{}
	return
}

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		ConfigureFunc: configureFunc,

		Schema: map[string]*schema.Schema{},

		DataSourcesMap: map[string]*schema.Resource{
			"number": initNumberSchema(),

		},

		ResourcesMap: map[string]*schema.Resource{},
	}
}
