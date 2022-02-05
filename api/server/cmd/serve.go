package cmd

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	srv "github.com/param108/profile/api/server/instance"
	"github.com/soellman/pidfile"
	"github.com/urfave/cli/v2"
)

var (
	servePort int

	// The Actual command
	serveCommand = &cli.Command{
		Name:   "serve",
		Usage:  "run the server",
		Action: serveCmd,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "port",
				Usage:       "port to listen on for http requests",
				Required:    false,
				Value:       8080,
				EnvVars:     []string{"PORT"},
				Destination: &servePort,
			}},
	}
)

func serveCmd(c *cli.Context) error {

	if err := pidfile.Write("PID"); err != nil {
		if err != pidfile.ErrProcessRunning {
			// lets remove and try again
			os.Remove("PID")
			if err = pidfile.Write("PID"); err != nil {
				log.Fatal("failed PID file:" + err.Error())
			}
		} else {
			log.Fatal("failed PID file:" + err.Error())
		}
	}

	s, err := srv.NewServer(servePort)
	if err != nil {
		log.Fatalf("Failed to start server:%s", err.Error())
	}

	cancelChan := make(chan os.Signal, 1)
	// catch SIGTERM or SIGINTERRUPT
	signal.Notify(cancelChan, syscall.SIGTERM, syscall.SIGINT)

	done := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.Serve()
		close(done)
	}()

	wg.Add(1)

	// This goroutine listens to cancelChan for SIG_TERM signals
	// In case the Serve goroutine exits, it also listens to the done channel.
	go func() {
		defer wg.Done()
		select {
		case <-cancelChan:
			s.Quit()
			return
		case <-done:
			return
		}
	}()

	wg.Wait()

	return nil
}

func init() {
	register(serveCommand)
}
