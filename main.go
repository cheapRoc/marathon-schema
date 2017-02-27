package main

import (
	// "bytes"
	// "path/filepath"

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
			// pull schema off internet
			fmt.Printf("Implement pulling schema off interwebz for tag %s", c.String("tag"))
			schemaData = ""
		} else {
			schemaBinData, err := Asset(AssetName)
			if err != nil {
				panic(err.Error())
			}
			schemaData = fmt.Sprintf("%s", schemaBinData)
		}

		fmt.Printf("schemaData:\n%s", schemaData)

		// Deal with user input JSON, `c.Args().First()` or `os.Stdin`
		inputName, err := filepath.Abs(filepath.Join(os.Getenv("PWD"), c.Args().First()))
		if err != nil {
			panic(err.Error())
		}

		inputPath := strings.Join([]string{FileProto, inputName}, "")
		fmt.Println("inputPath: %s\n", inputPath)

		schemaLoader := schema.NewStringLoader(schemaData)
		fileLoader := schema.NewReferenceLoader(inputPath)

		result, err := schema.Validate(schemaLoader, fileLoader)
		if err != nil {
			panic(err.Error())
		}

		if result.Valid() {
			fmt.Printf("The document is valid")
		} else {
			fmt.Printf("%s is not valid and contains the following errors:\n\n", inputName)
			for _, desc := range result.Errors() {
				fmt.Printf("- %s\n", desc)
			}
		}

		return nil
	}

	app.Run(os.Args)
}
