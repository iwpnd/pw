package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/iwpnd/pw"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "pw",
		Usage: "create a secure password from the command line",
		Action: func(ctx *cli.Context) error {
			l := ctx.Args().Get(0)
			v, err := strconv.Atoi(l)
			if err != nil {
				panic("AAH!")
			}
			fmt.Println(pw.NewPassword(v))
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
