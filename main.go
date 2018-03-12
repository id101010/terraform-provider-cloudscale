package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/id101010/terraform-provider-cloudscale/cloudscale"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: cloudscale.Provider,
	})
}
