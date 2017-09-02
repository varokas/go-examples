package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	fmt.Println("Hello")

	app := cli.NewApp()
	app.Name = "hello"
	app.Usage = "hello"
	app.Action = func(c *cli.Context) error {
		fmt.Println("hello!")
		return nil
	}

	app.Run(os.Args)
}
