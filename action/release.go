package action

import (
	"fmt"
	"github/controller"

	"github.com/urfave/cli/v2"
)

func List(c *cli.Context) error {
	token := c.String("token")
	fmt.Printf("token:%s\n", token)
	return controller.List(token)
}

func Add(c *cli.Context) error {
	tag := c.String("tag")
	body := c.String("desc")
	branch := c.String("branch")
	token := c.String("token")
	fmt.Printf("tag:%s body:%s branch:%s token:%s\n", tag, body, branch, token)
	return controller.Add(&controller.AddReq{
		Tag:    tag,
		Body:   body,
		Branch: branch,
	}, token)
}

func Edit(c *cli.Context) error {
	id := c.String("id")
	tag := c.String("tag")
	body := c.String("desc")
	branch := c.String("branch")
	token := c.String("token")
	fmt.Printf("tag:%s body:%s branch:%s token:%s\n", tag, body, branch, token)
	return controller.Edit(&controller.EditReq{
		Id:     id,
		Tag:    tag,
		Body:   body,
		Branch: branch,
	}, token)
}

func Delete(c *cli.Context) error {
	id := c.String("id")
	token := c.String("token")
	return controller.Delete(&controller.DeleteReq{
		Id: id,
	}, token)
}
