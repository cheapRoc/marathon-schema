package main

import (
	"bytes"
	"os"
	"testing"

	assert "github.com/stretchr/testify/assert"
	cli "github.com/urfave/cli"
)

// func TestRunDefault(t *testing.T) {
// 	os.Args = []string{"marathon-schema"}

// 	output := new(bytes.Buffer)
// 	app := SetupApp()
// 	app.Writer = output
// 	app.Action = func(ctx *cli.Context) error {
// 		DefaultAction(ctx)
// 		return nil
// 	}

// 	app.Run(os.Args)

// 	assert := assert.New(t)
// 	assert.Contains(output.String(), "AUTHOR(S):", "should contain 'AUTHOR(S):'")
// }

// func TestRunHelp(t *testing.T) {
// 	os.Args = []string{"marathon-schema", "-h"}

// 	output := new(bytes.Buffer)
// 	app := SetupApp()
// 	app.Writer = output
// 	app.Action = DefaultAction
// 	app.Run(os.Args)

// 	assert := assert.New(t)
// 	assert.Contains(output.String(), "AUTHOR(S):", "should contain 'AUTHOR(S):'")
// }

// marathon-schema ./marathon.json
// app.Run([]string{"marathon-schema", "samples/marathon.json"})

// func TestMainFile(t *testing.T) {
// 	output := new(bytes.Buffer)
// 	app := cli.NewApp()
// 	app.Writer = output
// }

// cat ./marathon.json | marathon-schema

// func TestMainStdin(t *testing.T) {

// }

// marathon-schema -a ./marathon.json
// app.Run([]string{"marathon-schema", "-a", "samples/marathon.json"})

// func TestMainAppDefFlag(t *testing.T) {

// }

// marathon-schema -g ./group.json
// app.Run([]string{"marathon-schema", "-g", "samples/group.json"})

// func TestMainGroupFlag(t *testing.T) {

// }

// marathon-schema -g -t 1.3.3 ./marathon.json
// app.Run([]string{"marathon-schema", "-g", "-t", "1.3.3", "samples/group.json"})

// marathon-schema -a -t 1.3.3 ./marathon.json
// app.Run([]string{"marathon-schema", "-a", "-t", "1.3.3", "samples/marathon.json"})

// func TestMainRemoteTag(t *testing.T) {

// }


