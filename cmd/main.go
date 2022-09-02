package main

import (
	"fmt"
	"log"
	"os"

	"github.com/iwpnd/pw"
	"github.com/urfave/cli/v2"
)

func create(ctx *cli.Context) error {
	var options []pw.Option

	if !ctx.Bool("no-upper") {
		options = append(options, pw.WithUpperCase())
	}
	if !ctx.Bool("no-lower") {
		options = append(options, pw.WithLowerCase())
	}
	if !ctx.Bool("no-numbers") {
		options = append(options, pw.WithNumbers())
	}
	if !ctx.Bool("no-special") {
		options = append(options, pw.WithSpecial())
	}

	fmt.Println(pw.NewPassword(ctx.Int("length"), options...))
	return nil
}

var createCommand cli.Command
var noUpperFlag, noLowerFlag, noNumbersFlag, noSpecialFlag cli.BoolFlag
var lengthFlag cli.IntFlag

func init() {
	lengthFlag = cli.IntFlag{
		Name:    "length",
		Aliases: []string{"l"},
		Value:   50,
		Usage:   "length of password",
	}

	noUpperFlag = cli.BoolFlag{
		Name:    "no-upper",
		Aliases: []string{"nu"},
		Value:   false,
		Usage:   "password contains NO uppercase characters",
	}

	noLowerFlag = cli.BoolFlag{
		Name:    "no-lower",
		Aliases: []string{"nl"},
		Value:   false,
		Usage:   "password contains NO lowercase characters",
	}

	noNumbersFlag = cli.BoolFlag{
		Name:    "no-numbers",
		Aliases: []string{"nn"},
		Value:   false,
		Usage:   "password contains NO number characters",
	}

	noSpecialFlag = cli.BoolFlag{
		Name:    "no-special",
		Aliases: []string{"ns"},
		Value:   false,
		Usage:   "password contains NO special characters",
	}

	createCommand = cli.Command{
		Name:   "create",
		Usage:  "create a new password",
		Action: create,
		Flags: []cli.Flag{
			&lengthFlag,
			&noUpperFlag,
			&noLowerFlag,
			&noSpecialFlag,
			&noNumbersFlag,
		},
	}
}

func main() {
	app := &cli.App{
		Name:     "pw",
		Usage:    "create a secure password from the command line",
		Commands: []*cli.Command{&createCommand},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
