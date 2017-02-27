package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli"
	schema "github.com/xeipuuv/gojsonschema"
)

var (
	Build string
	FileProto string = "file://"
	AssetName string
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
			Name: "tag, t",
			Usage: "validate against remote tagged version of Marathon schema",
		},
		cli.BoolTFlag{
			Name: "appdef, a",
			Usage: "work on AppDefinition JSON input",
		},
		cli.BoolFlag{
			Name: "group, g",
			Usage: "work on Group JSON input",
		},
		cli.BoolFlag{
			Name: "silent, s",
			Usage: "silent output",
		},
	}

	app.Action = func(c *cli.Context) error {
		var schemaData string

		if c.Bool("appdef") {
			AssetName = "AppDefinition.json"
		} else if c.Bool("group") {
			AssetName = "Group.json"
		} else {
			AssetName = "AppDefinition.json"
		}

		if c.String("tag") != "" {
			fmt.Printf("Implement pulling schema off interwebz for tag %s", c.String("tag"))
			schemaData = ""
		} else {
			schemaBinData, err := Asset(AssetName)
			if err != nil {
				panic(err.Error())
			}
			schemaData = fmt.Sprintf("%s", schemaBinData)
		}

		// fmt.Printf("schemaData:\n%s", schemaData)

		userFile := c.Args().First()

		// Deal with user input JSON, `c.Args().First()` or `os.Stdin`
		inputName, err := filepath.Abs(filepath.Join(os.Getenv("PWD"), userFile))
		if err != nil {
			panic(err.Error())
		}

		inputPath := strings.Join([]string{FileProto, inputName}, "")

		// fmt.Println("inputPath: %s\n", inputPath)

		schemaLoader := schema.NewStringLoader(schemaData)
		fileLoader := schema.NewReferenceLoader(inputPath)

		result, err := schema.Validate(schemaLoader, fileLoader)
		if err != nil {
			panic(err.Error())
		}

		if result.Valid() {
			if !c.Bool("silent") {
				fmt.Printf("The document is valid\n")
			}

			os.Exit(0)
		} else {
			if !c.Bool("silent") {
				fmt.Printf("\n`%s` is not valid and contains the following errors:\n\n", userFile)
			}

			var succ int = 0;
			for _, desc := range result.Errors() {
				if !c.Bool("silent") {
					fmt.Printf("- %s\n", desc)
				}
				succ += 1
			}
			if !c.Bool("silent") {
				fmt.Println("\n")
			}

			os.Exit(succ)
		}

		return nil
	}

	app.Run(os.Args)
}
