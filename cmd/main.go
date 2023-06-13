package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/icoder-new/reporter"
	"github.com/icoder-new/reporter/db"
	"github.com/icoder-new/reporter/logger"
	"github.com/icoder-new/reporter/pkg/handler"
	"github.com/icoder-new/reporter/pkg/repository"
	"github.com/icoder-new/reporter/pkg/service"
	"github.com/icoder-new/reporter/utils"
)

func main() {
	utils.ReadSettings()
	utils.PutAdditionalSettings()

	logger.Init()

	db.StartDbConnection()

	_db := db.GetDBConn()
	repository := repository.NewRepository(_db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	srv := new(reporter.Server)
	go func() {
		if err := srv.Run(utils.AppSettings.AppParams.PortRun, handler.InitRoutes()); err != nil {
			logger.Error.Fatal("Error occured while running http server. Error is: ", err.Error())
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	db.DisconnectDB(_db)
	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Error.Fatal("Error occured on server shutting down. Error is: ", err.Error())
		return
	}
}
