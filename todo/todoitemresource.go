package todo

import (


	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"


	client "github.com/whpearson/todo-client/client"
	todos "github.com/whpearson/todo-client/client/todos"
  models "github.com/whpearson/todo-client/models"
)


func resourceTodoItem() *schema.Resource {
	return &schema.Resource{
		Create: resourceTodoItemCreate,
		Read:   resourceTodoItemRead,
		Update: resourceTodoItemUpdate,
		Delete: resourceTodoItemDelete,

		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
		},
	}
}

type commonItemParams struct {
	Description              string
}

func itemForResource(d *schema.ResourceData) (*models.Item, error) {
	itemParams := commonItemParams{}

	// required
	if v, ok := d.GetOk("description"); ok {
		itemParams.Description = v.(string)
	}


	return &models.Item{
			Description:              &itemParams.Description,
	}, nil
}

func resourceTodoItemCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*client.SimpleToDoList)

	item, err := itemForResource(d)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Check create configuration: %#v ", d.Get("description") )
  params :=  todos.NewAddOneParams()
  params.SetBody(item )
	new_item , err := client.Todos.AddOne( params, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(int(new_item.Payload.ID)))

	return nil
}

func resourceTodoItemRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*client.SimpleToDoList)

	id, err := strconv.ParseInt(d.Id(),10,32)
	if err != nil {
		return fmt.Errorf("Error retrieving id for resource: %s", err)
	}
	find, err := client.Todos.Find(todos.NewFindParams().WithTags([]int32{int32(id)}).WithLimit(1), nil)
	if err != nil {
		return fmt.Errorf("Error retrieving list of items: %s", err)
	}
	exists := false
  if len(find.Payload) > 0 {
			exists = true
  }
	if !exists {
		d.SetId("")
		return nil
	}
	item := find.Payload[0]
	if err != nil {
		return fmt.Errorf("Error retrieving check: %s", err)
	}

	d.Set("description", item.Description)

	return nil
}

func resourceTodoItemUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*client.SimpleToDoList)

	item, err := itemForResource(d)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Check update configuration: %#v ", d.Get("description") )
  params := todos.NewUpdateOneParams().WithID(d.Id())
  params.SetBody(item)
	item_ok,err := client.Todos.UpdateOne( params, nil)
	if err != nil {
		return fmt.Errorf("Error updating check: %s", err)
	}

  d.SetId(strconv.Itoa(int(item_ok.Payload.ID)))

	return nil

}

func resourceTodoItemDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*client.SimpleToDoList)

	log.Printf("[INFO] Deleting Check: %v", d.Id())

	_, err := client.Todos.DestroyOne( todos.NewDestroyOneParams().WithID(d.Id() ), nil)
	if err != nil {
		return fmt.Errorf("Error deleting check: %s", err)
	}

	return nil
}

