package main

import "github.com/hashicorp/terraform/helper/schema"

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
	name := d.Get("name").(string)
	d.SetId(name)
	return resourceModelRead(d, m)
}

func resourceModelRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceModelUpdate(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)


	return resourceModelRead(d, m)
}

func resourceModelDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
