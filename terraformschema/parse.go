package terraformschema

import (
	"github.com/hashicorp/terraform/helper/schema"
	"reflect"
)

func Parse(d *schema.ResourceData, v interface{}) error {
	rv := reflect.ValueOf(v).Elem()
	rt := reflect.TypeOf(v).Elem()

	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		key := getTagOrName(f)
		val := d.Get(key)
		rv.Field(i).Set(reflect.ValueOf(val))
	}

	return nil
}

func getTagOrName(f reflect.StructField) string {
	res := f.Tag.Get("tf-schema")
	if res == "" {
		res = f.Name
	}

	return res
}
