package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/higumachan/terraform-abeja-platform/apiclient"
)

func resourceModelVersion() *schema.Resource {
	return &schema.Resource{
		Create: resourceModelVersionCreate,
		Read: resourceModelVersionRead,
		Update: resourceModelVersionUpdate,
		Delete: resourceModelVersionDelete,

		Schema: map[string]*schema.Schema{
			"model_id": &schema.Schema{
				Type:	schema.TypeString,
				Required: true,
			},
			"version": &schema.Schema{
				Type:	schema.TypeString,
				Required: true,
			},
			"image": &schema.Schema{
				Type:	schema.TypeString,
				Required: true,
			},
			"handler": &schema.Schema{
				Type:	schema.TypeString,
				Required: true,
			},
			"upload_file_path": &schema.Schema{
				Type:	schema.TypeString,
				Required: true,
			},
			"content_type": &schema.Schema{
				Type:	schema.TypeString,
				Optional: true,
				Default: "application/octet-stream",
			},
		},
	}
}

func resourceModelVersionCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*apiclient.Client)

	modelId := d.Get("model_id").(string)
	uploadFilePath := d.Get("upload_file_path").(string)

	var param apiclient.CreateModelVersionParam
	param.Version = d.Get("version").(string)
	param.Image = d.Get("image").(string)
	param.Handler = d.Get("handler").(string)
	param.ContentType = d.Get("content_type").(string)

	res, err := client.CreateModelVersion(modelId, uploadFilePath, &param)
	if err != nil {
		return err
	}

	d.SetId(res.VersionId)
	return resourceModelVersionRead(d, m)
}

func resourceModelVersionRead(d *schema.ResourceData, m interface{}) error {
	/*
	client := m.(*apiclient.Client)

	res, err := client.RetrieveModelVersion(d.Id())
	if err != nil {
		return err
	}
	d.SetId(res.ModelVersionId)
	*/
	return nil
}

func resourceModelVersionUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceModelVersionRead(d, m)
}

func resourceModelVersionDelete(d *schema.ResourceData, m interface{}) error {
	/*
	d.SetId("")

	client := m.(*apiclient.Client)
	err := client.DeleteModelVersion(d.Id())
	if err != nil {
		return err
	}
	*/

	return nil
}
