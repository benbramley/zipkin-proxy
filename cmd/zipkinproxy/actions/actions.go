package actions

import (
	"github.com/urfave/cli"

	"github.com/namely/zipkin-proxy/pkg/destination/logs"
	"github.com/namely/zipkin-proxy/pkg/server"
)

//
func RegisterActions(a *cli.App) {
	a.Commands = append(a.Commands, []cli.Command{
		{
			Name:   "server",
			Usage:  "Starts a zipkin proxy server",
			Action: ServerCmd,
			Flags: []cli.Flag{
				cli.StringFlag{Name: "addr", EnvVar: "LISTEN_ON"},
			},
		},
	}...)
}

// ServerCmd is an action that can be passed to a zipkin proxy main application
// for running a server
func ServerCmd(c *cli.Context) error {
	listenOn := c.String("addr")
	s := server.NewServer(listenOn, &logs.LogShipper{})

	return s.Start()
}
