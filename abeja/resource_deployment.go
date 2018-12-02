package abeja

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/higumachan/terraform-abeja-platform/apiclient"
)

func resourceDeployment() *schema.Resource {
	return &schema.Resource{
		Create: resourceDeploymentCreate,
		Read: resourceDeploymentRead,
		Update: resourceDeploymentUpdate,
		Delete: resourceDeploymentDelete,

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

func resourceDeploymentCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*apiclient.Client)

	param := apiclient.CreateDeploymentParam{}

	modelId := d.Get("model_id").(string)
	param.Name = d.Get("name").(string)
	param.DefaultEnvironment = d.Get("default_environment").(string)

	res, err := client.CreateDeployment(modelId, &param)
	if err != nil {
		return err
	}

	d.SetId(res.DeploymentId)
	return resourceDeploymentRead(d, m)
}

func resourceDeploymentRead(d *schema.ResourceData, m interface{}) error {
	/*
	client := m.(*apiclient.Client)

	res, err := client.RetrieveDeployment(d.Id())
	if err != nil {
		return err
	}
	d.SetId(res.DeploymentId)
	*/
	return nil
}

func resourceDeploymentUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceDeploymentRead(d, m)
}

func resourceDeploymentDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*apiclient.Client)
	err := client.DeleteDeployment(d.Id())
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
