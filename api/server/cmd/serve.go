package cmd

import (
	"github.com/urfave/cli/v2"
	srv "github.com/param108/profile/api/server/serve"
)

var (
	servePort int

	// The Actual command
	serve = &cli.Command{
	Name: "serve",
	Usage: "run the server",
	Action: serveCmd,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name: "port",
			Usage: "port to listen on for http requests",
			Required: true,
			Value: 8080,
			EnvVars: []string{"PORT"},
			Destination: &servePort,
		}},
	}
)

func serveCmd(c *cli.Context) error {
	srv.Serve(servePort)
	return nil
}

func init() {
	register(serve)
}
