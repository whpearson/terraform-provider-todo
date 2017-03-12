package main

import (
 "github.com/hashicorp/terraform/plugin"
 "github.com/whpearson/terraform-provider-todo/todo"
)

func main () {
  plugin.Serve (&plugin.ServeOpts{
    ProviderFunc: todo.Provider,
  })
}
