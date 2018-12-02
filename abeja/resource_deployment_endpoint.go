package abeja

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceDeploymentEndpoint() *schema.Resource {
	return &schema.Resource{
		Create: resourceDeploymentEndpointCreate,
		Read: resourceDeploymentEndpointRead,
		Update: resourceDeploymentEndpointUpdate,
		Delete: resourceDeploymentEndpointDelete,

		Schema: map[string]*schema.Schema{
			"model_id": &schema.Schema{
				Type:	schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:	schema.TypeString,
				Required: true,
			},
			"default_environment": &schema.Schema{
				Type:	schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceDeploymentEndpointCreate(d *schema.ResourceData, m interface{}) error {
	/*
	client := m.(*apiclient.Client)

	param := apiclient.CreateDeploymentEndpointParam{}

	modelId := d.Get("model_id").(string)
	param.Name = d.Get("name").(string)
	param.DefaultEnvironment = d.Get("default_environment").(string)

	res, err := client.CreateDeploymentEndpoint(modelId, &param)
	if err != nil {
		return err
	}

	d.SetId(res.DeploymentEndpointId)
	*/
	return resourceDeploymentEndpointRead(d, m)
}

func resourceDeploymentEndpointRead(d *schema.ResourceData, m interface{}) error {
	/*
	client := m.(*apiclient.Client)

	res, err := client.RetrieveDeploymentEndpoint(d.Id())
	if err != nil {
		return err
	}
	d.SetId(res.DeploymentEndpointId)
	*/
	return nil
}

func resourceDeploymentEndpointUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceDeploymentEndpointRead(d, m)
}

func resourceDeploymentEndpointDelete(d *schema.ResourceData, m interface{}) error {
	/*
	client := m.(*apiclient.Client)
	err := client.DeleteDeploymentEndpoint(d.Id())
	if err != nil {
		return err
	}

	d.SetId("")
	*/

	return nil
}
