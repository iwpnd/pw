package main

import (
	"testing"

	"github.com/urfave/cli/v2"
)

func TestCreate(t *testing.T) {
	app := cli.NewApp()
	app.Name = "pw"
	app.Commands = cli.Commands{&createCommand}

	type tcase struct {
		cmd []string
	}

	fn := func(tc tcase) func(t *testing.T) {
		return func(t *testing.T) {
			err := app.Run(tc.cmd)
			if err != nil {
				t.Errorf("create should not return an error")
			}
		}
	}

	tests := map[string]tcase{
		"vanilla": {
			cmd: []string{"pw", "create"},
		},
		"no uppercase": {
			cmd: []string{"pw", "create", "--no-upper"},
		},
		"no lowercase": {
			cmd: []string{"pw", "create", "--no-lower"},
		},
		"no numbers": {
			cmd: []string{"pw", "create", "--no-numbers"},
		},
		"no special": {
			cmd: []string{"pw", "create", "--no-special"},
		},
		"no all == default": {
			cmd: []string{"pw", "create", "--no-special", "--no-numbers", "--no-upper", "--no-lower"},
		},
	}

	for name, tc := range tests {
		t.Run(name, fn(tc))
	}

}
