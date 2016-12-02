package main

import (
	"fmt"
	"os"
	"strings"
	"go/build"
	"path/filepath"
	"github.com/urfave/cli"
	schema "github.com/xeipuuv/gojsonschema"
)

// compile passing -ldflags "-X main.Build <build sha1>"
var Build string
var AppDefPath string
var AppDefFile string
var SamplePath string
var SchemaPath string
var FilePath string
var FileProto string

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
		p, err := build.Default.Import("github.com/user/repo/dummy", "", build.FindOnly)
		if err != nil {
			// something
		}
		fname, _ := filepath.Abs(filepath.Join(p.Dir, c.Args().First()))
		FileProto = "file://"

		FilePath = strings.Join([]string{FileProto, fname}, "")

		SchemaPath = "file:///Users/justin/Development/chariot/ims/marathon-schema/schemas/"

		AppDefFile = strings.Join([]string{c.String("marathon"), "-AppDefinition.json"}, "")

		AppDefPath = strings.Join([]string{SchemaPath, AppDefFile}, "")

		schemaLoader := schema.NewReferenceLoader(AppDefPath)
		fileLoader := schema.NewReferenceLoader(FilePath)

		result, err := schema.Validate(schemaLoader, fileLoader)
		if err != nil {
			panic(err.Error())
		}

		if result.Valid() {
			fmt.Printf("The document is valid")
		} else {
			fmt.Printf("The document is not valid. Errors :\n")
			for _, desc := range result.Errors() {
				fmt.Printf("- %s\n", desc)
			}
		}

		return nil
	}

	app.Run(os.Args)
}
