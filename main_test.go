package main

import (
	"bytes"
	"testing"

	// "github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
)

// assert := assert.New(t)
// assert.Equal(123, 123, "should be equal")
// c := cli.NewContext(app, nil, nil)

func TestMainAuthor(t *testing.T) {
	output := new(bytes.Buffer)
	app := cli.NewApp()
	app.Writer = output

	if bytes.Index(output.Bytes(), []byte("AUTHOR(S):")) != -1 {
		t.Errorf("expected\n%s not to include %s", output.String(), "AUTHOR(S):")
	}
}

func TestMainArguments(t *testing.T) {
	output := new(bytes.Buffer)
	app := cli.NewApp()
	app.Writer = output
	app.Run([]string{"-h"})

	if bytes.Contains(output.Bytes(), []byte("Schema not found")) != -1 {
		t.Errorf("expected\n%s not to include %s", output.String(), "Schema not found")
	}
}

// marathon-schema ./marathon.json
// app.Run([]string{"samples/marathon.json"})


// func TestMainFile(t *testing.T) {
// 	output := new(bytes.Buffer)
// 	app := cli.NewApp()
// 	app.Writer = output
// }

// cat ./marathon.json | marathon-schema

// func TestMainStdin(t *testing.T) {

// }

// marathon-schema -a ./marathon.json

// func TestMainAppDefFlag(t *testing.T) {

// }

// marathon-schema -g ./marathon.json

// func TestMainGroupFlag(t *testing.T) {

// }

// marathon-schema -g -t 1.3.3 ./marathon.json
// marathon-schema -a -t 1.3.3 ./marathon.json

// func TestMainRemoteTag(t *testing.T) {

// }


