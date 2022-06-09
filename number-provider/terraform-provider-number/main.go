package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/purajit/terraform-playground/number-provider/number"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: number.Provider})
}
