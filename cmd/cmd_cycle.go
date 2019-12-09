package cmd

import (
	//"bufio"
	"fmt"

	//"log"
	//"os"
	//"strings"

	"github.com/urfave/cli/v2"
)

func createCycle() cli.Command {
	return cli.Command{
		Name:  "cycle",
		Usage: "Cycle through all colors",
		Before: func(ctx *cli.Context) error {
			return nil
		},
		Action: cycle,
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

func cycle(ctx *cli.Context) error {
	fmt.Println("!!!! HEADS UP !!!!")
	return nil
}
