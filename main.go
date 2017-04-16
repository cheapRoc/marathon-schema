package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli"
	schema "github.com/xeipuuv/gojsonschema"
)

const (
	MarathonVersion string = "1.3.3"
	FileProto string = "file://"
)

var (
	buildCommit string
	assetName string
)

func DefaultAction(ctx *cli.Context) error {
	var schemaData string

	if ctx.Bool("list") {
		for _, name := range AssetNames() {
			fmt.Printf("%s\n", name)
		}
		os.Exit(0)
		return nil
	}

	if ctx.Bool("appdef") {
		assetName = "AppDefinition.json"
	} else if ctx.Bool("group") {
		assetName = "Group.json"
	} else {
		assetName = "AppDefinition.json"
	}

	if ctx.String("tag") != "" {
		schemaData = ""
		assetName = fmt.Sprintf("%s-%s", ctx.String("tag"), assetName)
	}

	schemaBinData, err := Asset(assetName)
	if err != nil {
		fmt.Printf("Could not find schema asset filename: %s\n", assetName)
		panic(err.Error())
	}
	schemaData = fmt.Sprintf("%s", schemaBinData)

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

	inputPath := strings.Join([]string{FileProto, inputName}, "")
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
	app.Version = strings.Join([]string{"1.2.0-", buildCommit}, "")
	app.Usage = "Provides validation of marathon.json files against App Definition schema."
	app.Authors = []cli.Author{
		cli.Author{
			Name: "Justin Reagor",
			Email: "justinwr@gmail.com",
		},
	}

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "tag, t",
			Usage: fmt.Sprintf("version of Marathon to validate, default is %s", MarathonVersion),
		},
		cli.BoolTFlag{
			Name: "appdef, a",
			Usage: "process AppDefinition JSON input",
		},
		cli.BoolFlag{
			Name: "group, g",
			Usage: "process Group JSON input",
		},
		cli.BoolFlag{
			Name: "silent, s",
			Usage: "silent command output",
		},
		cli.BoolFlag{
			Name: "list, l",
			Usage: "list schemas (pre-packaged)",
		},
	}

	return app
}

func main() {
	app := SetupApp()
	app.Action = DefaultAction
	app.Run(os.Args)
}
