package cmd

import (
	//"bufio"
	"fmt"
	//"log"
	//"os"
	//"strings"

	"github.com/urfave/cli/v2"
)

func createIntro() cli.Command {
	return cli.Command{
		Name:  "intro",
		Usage: "Enable/disable startup effect",
		Before: func(ctx *cli.Context) error {
			return nil
		},
		Action: runIntro,
		Flags:  []cli.Flag{},
	}
}

func runIntro(ctx *cli.Context) error {
	fmt.Println("!!!! HEADS UP !!!!")
	return nil
}
