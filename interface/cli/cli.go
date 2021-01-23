package cli

import (
	"cacing/interface/socket"

	"github.com/urfave/cli/v2"
)

// NewCliApp func
func NewCliApp(args []string) error {
	c := &cli.App{
		Name:  "cacing",
		Usage: "Quite simple in memory cache storage",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "username",
				Usage: "Username to access the storage server",
			},
			&cli.StringFlag{
				Name:  "password",
				Usage: "Secret password for username",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "run",
				Usage: "Start cacing server",
				Action: func(ctx *cli.Context) error {
					return socket.RunServer()
				},
			},
			{
				Name:  "connect",
				Usage: "Connect to a cacing server as client",
				Action: func(ctx *cli.Context) error {
					return socket.ConnectTo()
				},
			},
		},
	}
	return c.Run(args)
}
