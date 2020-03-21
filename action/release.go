package action

import (
	"fmt"
	"github/controller"

	"github.com/urfave/cli/v2"
)

func List(c *cli.Context) error {
	token := c.String("token")
	fmt.Printf("token:%s\n", token)
	controller.List(token)
	return nil
}

func Add(c *cli.Context) error {
	tag := c.String("tag")
	body := c.String("desc")
	branch := c.String("branch")
	token := c.String("token")
	fmt.Printf("tag:%s body:%s branch:%s token:%s\n", tag, body, branch, token)
	controller.Add(&controller.AddReq{
		Tag:    tag,
		Body:   body,
		Branch: branch,
	}, token)
	return nil
}

func Edit(c *cli.Context) error {
	fmt.Println("removed task template: ", c.Args().First())
	return nil
}

func Delete(c *cli.Context) error {
	fmt.Println("removed task template: ", c.Args().First())
	return nil
}
