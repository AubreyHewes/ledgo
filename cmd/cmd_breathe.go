package cmd

import (
	//"bufio"
	"fmt"
	//"log"
	//"os"
	//"strings"

	"github.com/urfave/cli/v2"
)

func createBreathe() cli.Command {
	return cli.Command{
		Name:  "breathe",
		Usage: "Single color breathing",
		Before: func(ctx *cli.Context) error {
			return nil
		},
		Action: runBreath,
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

func runBreath(ctx *cli.Context) error {
	fmt.Println("!!!! HEADS UP !!!!")
	return nil
}
