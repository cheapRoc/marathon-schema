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
	build string
	fileProto string = "file://"
	assetName string
)

func DefaultAction(ctx *cli.Context) error {
	var schemaData string

	if ctx.Bool("appdef") {
		assetName = "AppDefinition.json"
	} else if ctx.Bool("group") {
		assetName = "Group.json"
	} else {
		assetName = "AppDefinition.json"
	}

	if ctx.String("tag") != "" {
		fmt.Printf("Implement pulling schema off interwebz for tag %s", ctx.String("tag"))
		schemaData = ""
	} else {
		schemaBinData, err := Asset(assetName)
		if err != nil {
			panic(err.Error())
		}
		schemaData = fmt.Sprintf("%s", schemaBinData)
	}

	userFile := ctx.Args().First()
	if userFile == "" {
		cli.ShowCommandHelp(ctx, "")
		os.Exit(1)
		return nil
	}

	// Deal with user input JSON, `ctx.Args().First()` or `os.Stdin`
	inputName, err := filepath.Abs(filepath.Join(os.Getenv("PWD"), userFile))
	if err != nil {
		panic(err.Error())
	}

	inputPath := strings.Join([]string{fileProto, inputName}, "")
	schemaLoader := schema.NewStringLoader(schemaData)
	fileLoader := schema.NewReferenceLoader(inputPath)

	result, err := schema.Validate(schemaLoader, fileLoader)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		if !ctx.Bool("silent") {
			fmt.Printf("The document is valid\n")
		}

		os.Exit(0)
	} else {
		if !ctx.Bool("silent") {
			fmt.Printf("\n`%s` is not valid and contains the following errors:\n\n", userFile)
		}

		var succ int = 0;
		for _, desc := range result.Errors() {
			if !ctx.Bool("silent") {
				fmt.Printf("- %s\n", desc)
			}
			succ += 1
		}
		if !ctx.Bool("silent") {
			fmt.Println("\n")
		}

		os.Exit(succ)
	}

	return nil
}

func SetupApp() *cli.App {
	app := cli.NewApp()

	app.Name = "marathon-schema"
	app.Version = strings.Join([]string{"1.0.0-", build}, "")
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

	return app
}

func main() {
	app := SetupApp()
	app.Action = DefaultAction
	app.Run(os.Args)
}
