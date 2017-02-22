package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
)

func Test_ShowAppHelp(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(123, 123, "should be equal")

	output := new(bytes.Buffer)
	app := cli.NewApp()
	app.Writer = output

	// c := cli.NewContext(app, nil, nil)
	if bytes.Index(output.Bytes(), []byte("AUTHOR(S):")) != -1 {
		t.Errorf("expected\n%s not to include %s", output.String(), "AUTHOR(S):")
	}

}
