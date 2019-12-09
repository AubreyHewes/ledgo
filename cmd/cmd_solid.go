package cmd

import (
	//"bufio"
	"fmt"
	//"log"
	//"os"
	//"strings"

	"github.com/urfave/cli/v2"

	"github.com/AubreyHewes/ledgo/v1/ledgo"
)

func createSolid() cli.Command {
	return cli.Command{
		Name:  "solid",
		Usage: "Set a solid color",
		Before: func(ctx *cli.Context) error {
			return nil
		},
		Action: run,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "only-logo",
				Usage: "Only set the logo",
			},
			&cli.BoolFlag{
				Name:  "only-wheel",
				Usage: "Only set the wheel",
			},
		},
	}
}

func run(ctx *cli.Context) error {
	fmt.Println("!!!! HEADS UP !!!!")
	ledgo.NewClient()
	return nil
}
