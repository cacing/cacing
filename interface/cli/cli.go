package cli

import (
	"github.com/hadihammurabi/cacing/interface/socket"
	"github.com/hadihammurabi/cacing/utils"

	"github.com/urfave/cli/v2"
)

// NewCliApp func
func NewCliApp(args []string) error {
	c := &cli.App{
		Name:  "cacing",
		Usage: "Quite simple in memory cache storage",
		Commands: []*cli.Command{
			{
				Name:  "run",
				Usage: "Start cacing server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "port",
						Usage: "Cacing server's host",
					},
					&cli.StringFlag{
						Name:  "username",
						Usage: "Username to access the storage server",
					},
					&cli.StringFlag{
						Name:  "password",
						Usage: "Secret password for username",
					},
				},
				Action: func(ctx *cli.Context) error {
					config := &socket.Config{
						Port:     ctx.String("port"),
						Username: ctx.String("username"),
						Password: ctx.String("password"),
					}
					return socket.RunServer(config)
				},
			},
			{
				Name:  "connect",
				Usage: "Connect to a cacing server as client",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "dsn",
						Usage: "Cacing connection string",
					},
				},
				Action: func(ctx *cli.Context) error {
					url, err := utils.ParseURL(ctx.String("dsn"))
					if err != nil {
						return err
					}
					return socket.ConnectTo(url)
				},
			},
		},
	}
	return c.Run(args)
}
