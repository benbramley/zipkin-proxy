package main

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/namely/zipkin-proxy/cmd/zipkinproxy/actions"
)

func main() {
	app := cli.NewApp()
	app.Name = "zipkin-proxy"
	actions.RegisterActions(app)

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err.Error())
		os.Exit(1)
	}
}

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}
