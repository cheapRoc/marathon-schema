package main

import (
	"fmt"
	"os"
	"github.com/urfave/cli"
)

// compile passing -ldflags "-X main.Build <build sha1>"
var Build string

func main() {
	fmt.Printf("Using build: %s\n", Build)

	app := cli.NewApp()
	app.Name = "marathon-schema"
	app.Usage = "Provides validation of marathon.json files against App Definition schema"
	app.Action = func(c *cli.Context) error {
		fmt.Println("boom! I say!")
		return nil
	}

	app.Run(os.Args)
}
