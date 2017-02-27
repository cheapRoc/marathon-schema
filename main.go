package main

import (
	"fmt"
	"os"
	// "bytes"
	// "path/filepath"
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
		// 1. If appdefSchema then load AppDefinition.json
		// - error: app def could not be in binary asset

		// 2. If groupSchema then load Group.json
		// - error: group schema could not be in binary asset

		// 3. If localFileSchema then load schema file on the file system
		// - error: schema file is not on local FS

		// 4. Pass asset binary information into jsonSchema
		// - error:

		// 5. Process schema result
		// - error: JSON error processing schema or file

		// 6. Optionally colorize output, for friendlier display...

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
			schemaData = fmt.Sprintf("schema: %s", schemaBinData)
		}

		// Deal with user input JSON, file or stdin

		fmt.Printf("schemaData:\n%s", schemaData)

		fileName := c.Args().First()

		// fmt.Printf("fileName: %s", fileName)
		// os.Stdin

		userFile, err := os.Open(fileName)
		if err != nil {
			if os.IsNotExist(err) {
				panic(fmt.Sprintf("Missing file %s\n%s", fileName, err.Error()))
			}

			panic(err.Error())
		}

		fileInfo, err := os.Stat(userFile)
		if err != nil {
			panic(err.Error())
		}

		fileData, err := userFile.Read()
		if err != nil {
			panic(err.Error())
		}

		// ----

		// userfile := strings.Join([]string{FileProto, fileName}, "")

		// schemaTag := c.String("tag")
		// schemaName := strings.Join([]string{schemaTag, "AppDefinition.json"}, "")
		// schemaFile := strings.Join([]string{SchemaPath, schemaName}, "")

		// schemaLoader := schema.NewReferenceLoader(schemaFile)

		// ----

		schemaLoader := schema.NewStringLoader(schemaData)
		fileLoader := schema.NewReferenceLoader(fileData)

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
