package main

import (
"github.com/hashicorp/terraform/helper/schema"
	"github.com/higumachan/terraform-abeja-platform/apiclient"
)


func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"user_id": {
				Type:	schema.TypeString,
				Optional: true,
				Default: "",
				Description: "abeja platform user id",
			},
			"personal_access_token": {
				Type:	schema.TypeString,
				Optional: true,
				Default: "",
				Description: "abeja platform personal access token",
			},
			"organization_id": {
				Type:	schema.TypeString,
				Optional: true,
				Default: "",
				Description: "abeja platform organization id",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"abeja_model": resourceModel(),
			"abeja_model_version": resourceModelVersion(),
			"abeja_deployment": resourceDeployment(),
			"abeja_deployment_service": resourceDeploymentService(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	client := apiclient.NewClient(d.Get("user_id").(string), d.Get("personal_access_token").(string),d.Get("organization_id").(string))
	return client, nil
}
