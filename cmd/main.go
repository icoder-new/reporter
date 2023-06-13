package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/icoder-new/reporter"
)

func main() {
	srv := new(reporter.Server)
	go func() {
		if err := srv.Run("1234", nil); err != nil {
			// TODO: make logger and write it into log file
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		// TODO: make logger and write it into log file
		return
	}
}
