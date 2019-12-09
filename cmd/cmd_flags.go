package cmd

import (
	//"github.com/AubreyHewes/ledgo/v1/ledgo"
	"github.com/urfave/cli/v2"
)

func CreateFlags(defaultPath string) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "path",
			Usage: "Directory to use for storing the data.",
			Value: defaultPath,
		},
	}
}
