package actions

import (
	"github.com/urfave/cli"

	"github.com/namely/zipkin-proxy/pkg/destination"
	"github.com/namely/zipkin-proxy/pkg/destination/datadog"
	"github.com/namely/zipkin-proxy/pkg/destination/logs"
	"github.com/namely/zipkin-proxy/pkg/server"
)

// RegisterActions appends all of the actions for this proxy
func RegisterActions(a *cli.App) {
	a.Commands = append(a.Commands, []cli.Command{
		{
			Name:   "server",
			Usage:  "Starts a zipkin proxy server",
			Action: ServerCmd,
			Flags: []cli.Flag{
				cli.StringFlag{Name: "addr", EnvVar: "LISTEN_ON"},
				cli.StringFlag{Name: "dd-agent-url", EnvVar: "DD_AGENT_URL"},
			},
		},
	}...)
}

// ServerCmd is an action that can be passed to a zipkin proxy main application
// for running a server
func ServerCmd(c *cli.Context) error {
	listenOn := c.String("addr")

	var dest destination.Interface
	if ddAgentUrl := c.String("dd-agent-url"); len(ddAgentUrl) > 0 {
		dest = datadog.NewShipper(ddAgentUrl)
	} else {
		dest = &logs.LogShipper{}
	}

	s := server.NewServer(listenOn, dest)
	return s.Start()
}
