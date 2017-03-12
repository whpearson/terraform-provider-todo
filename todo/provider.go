package todo

import (
  "log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
  "github.com/mitchellh/mapstructure"
)

func Provider () terraform.ResourceProvider {
  return &schema.Provider{
    Schema: map[string]*schema.Schema {
       "host" : &schema.Schema{
				  Type:     schema.TypeString,
				  Default: "localhost:9001",
				  Optional: false,
			  },
    },
	  ResourcesMap: map[string]*schema.Resource{
			"todo_item": resourceTodoItem(),
		},
		ConfigureFunc: providerConfigure,
  }
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
  var config Config
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &config); err != nil {
		return nil, err
	}

	log.Println("[INFO] Initializing Todo client")
	return config.Client()
}
