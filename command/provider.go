package command

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func NewProvider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"command_command": resourceCommand(),
		},
	}
}
