package main

import (
	"fmt"
	"go/build"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli"
	schema "github.com/xeipuuv/gojsonschema"
)

var (
	Build string
	SchemaPath string = "file:///Users/justin/Development/chariot/ims/marathon-schema/schemas/"
	FileProto string = "file://"
)

func main() {
	app := cli.NewApp()

	app.Name = "marathon-schema"
	app.Version = strings.Join([]string{"1.0.0-", Build}, "")
	app.Usage = "Provides validation of marathon.json files against App Definition schema"
	app.Authors = []cli.Author{
		cli.Author{
			Name: "Justin Reagor",
			Email: "jreagor@chariotsolutions.com",
		},
	}

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "marathon, m",
			Value: "1.3.3",
			Usage: "version of Marathon to target",
		},
		cli.BoolTFlag{
			Name: "appdef, a",
			Usage: "work on AppDefinition JSON input",
		},
		cli.BoolFlag{
			Name: "group, g",
			Usage: "work on Group JSON input",
		},
	}

	app.Action = func(c *cli.Context) error {
		p, err := build.Default.Import("github.com/not/found", "", build.FindOnly)
		if err != nil {
			// panic(err.Error())
		}

		fname, _ := filepath.Abs(filepath.Join(p.Dir, c.Args().First()))
		userfile := strings.Join([]string{FileProto, fname}, "")
		schemaBuild := c.String("marathon")
		schemaName := strings.Join([]string{schemaBuild, "-AppDefinition.json"}, "")
		schemaFile := strings.Join([]string{SchemaPath, schemaName}, "")
		schemaLoader := schema.NewReferenceLoader(schemaFile)
		fileLoader := schema.NewReferenceLoader(userfile)

		result, err := schema.Validate(schemaLoader, fileLoader)
		if err != nil {
			panic(err.Error())
		}

		if result.Valid() {
			fmt.Printf("The document is valid")
		} else {
			fmt.Printf("%s is not valid and contains the following errors:\n\n", fname)
			for _, desc := range result.Errors() {
				fmt.Printf("- %s\n", desc)
			}
		}

		return nil
	}

	app.Run(os.Args)
}
