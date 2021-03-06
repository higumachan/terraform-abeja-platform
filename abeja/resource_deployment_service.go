package abeja

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/higumachan/terraform-abeja-platform/apiclient"
	"github.com/higumachan/terraform-abeja-platform/terraformschema"
)

func resourceDeploymentService() *schema.Resource {
	return &schema.Resource{
		Create: resourceDeploymentServiceCreate,
		Read: resourceDeploymentServiceRead,
		Update: resourceDeploymentServiceUpdate,
		Delete: resourceDeploymentServiceDelete,

		Schema: map[string]*schema.Schema{
			"deployment_id": &schema.Schema{
				Type:	schema.TypeString,
				Required: true,
			},
			"instance_number": &schema.Schema{
				Type:	schema.TypeInt,
				Optional: true,
			},
			"instance_type": &schema.Schema{
				Type:	schema.TypeString,
				Optional: true,
			},
			"environment": &schema.Schema{
				Type:	schema.TypeString,
				Optional: true,
			},
			"version_id": &schema.Schema{
				Type:	schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceDeploymentServiceCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*apiclient.Client)


	deploymentId := d.Get("deployment_id").(string)
	param := apiclient.CreateDeploymentServiceParam{}
	terraformschema.Parse(d, &param)

	res, err := client.CreateDeploymentService(deploymentId, &param)
	if err != nil {
		return err
	}

	d.SetId(res.ServiceId)
	d.Set("deployment_id", deploymentId)
	return resourceDeploymentServiceRead(d, m)
}

func resourceDeploymentServiceRead(d *schema.ResourceData, m interface{}) error {
	/*
	client := m.(*apiclient.Client)

	res, err := client.RetrieveDeploymentService(d.Id())
	if err != nil {
		return err
	}
	d.SetId(res.DeploymentServiceId)
	*/
	return nil
}

func resourceDeploymentServiceUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceDeploymentServiceRead(d, m)
}

func resourceDeploymentServiceDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*apiclient.Client)
	err := client.DeleteDeploymentService(d.Get("deployment_id").(string), d.Id())
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
