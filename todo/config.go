package todo

import (
	"log"
  strfmt "github.com/go-openapi/strfmt"

	todoclient "github.com/whpearson/todo-client/client"
)

type Config struct {
   Host string `mapstructure:"host"`
}

// Client() returns a new client for accessing todo.
//
func (c *Config) Client() (*todoclient.SimpleToDoList , error) {

  httptransportconfig := todoclient.DefaultTransportConfig().WithHost(c.Host).WithSchemes([]string{"http"})
  client := todoclient.NewHTTPClientWithConfig(strfmt.Default, httptransportconfig )



	log.Printf("[INFO] Todo Client configured for host: %s", c.Host)

	return client, nil
}
