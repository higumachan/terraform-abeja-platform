package abeja

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/higumachan/terraform-abeja-platform/apiclient"
)

func resourceModel() *schema.Resource {
	return &schema.Resource{
		Create: resourceModelCreate,
		Read: resourceModelRead,
		Update: resourceModelUpdate,
		Delete: resourceModelDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:	schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:	schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceModelCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*apiclient.Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)

	res, err := client.CreateModel(&apiclient.CreateModelParam{Name: name, Description: description})
	if err != nil {
		return err
	}

	d.SetId(res.ModelId)
	return resourceModelRead(d, m)
}

func resourceModelRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*apiclient.Client)

	res, err := client.RetrieveModel(d.Id())
	if err != nil {
		return err
	}
	d.SetId(res.ModelId)
	return nil
}

func resourceModelUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceModelRead(d, m)
}

func resourceModelDelete(d *schema.ResourceData, m interface{}) error {

	client := m.(*apiclient.Client)
	err := client.DeleteModel(d.Id())
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
