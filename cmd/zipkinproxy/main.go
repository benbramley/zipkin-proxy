package main

import (
	"fmt"
	"os"

	"github.com/namely/zipkin-proxy/cmd/zipkinproxy/actions"
	"github.com/urfave/cli"
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
