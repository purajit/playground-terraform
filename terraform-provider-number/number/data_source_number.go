package number

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func initNumberSchema() *schema.Resource {
	return &schema.Resource{
		Read: read,

		Schema: map[string]*schema.Schema{
			"number": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func read(d *schema.ResourceData, meta interface{}) (err error) {
	return
}
