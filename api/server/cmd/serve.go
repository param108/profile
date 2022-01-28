package cmd

import (
	"log"

	srv "github.com/param108/profile/api/server/instance"
	"github.com/soellman/pidfile"
	"github.com/urfave/cli/v2"
)

var (
	servePort int

	// The Actual command
	serve = &cli.Command{
		Name:   "serve",
		Usage:  "run the server",
		Action: serveCmd,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "port",
				Usage:       "port to listen on for http requests",
				Required:    true,
				Value:       8080,
				EnvVars:     []string{"PORT"},
				Destination: &servePort,
			}},
	}
)

func serveCmd(c *cli.Context) error {

	if err := pidfile.Write("PID"); err != nil {
		log.Fatal("failed PID file:" + err.Error())
	}

	s, err := srv.NewServer(servePort)
	if err != nil {
		log.Fatalf("Failed to start server:%s", err.Error())
	}

	// never returns
	s.Serve()

	return nil
}

func init() {
	register(serve)
}
